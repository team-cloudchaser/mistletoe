#!/bin/bash
# A multi-arch build release script
export CGO_ENABLED=0
rm -r "./build/releases" 2>/dev/null
mkdir -p "./build/releases"
cat conf/gobuild-${1:-release}.txt | while IFS= read -r GOUNION; do
	export GOOS=$(printf $GOUNION | cut -d'/' -f1)
	export GOARCH=$(printf $GOUNION | cut -d'/' -f2)
	buildDir="./build/$GOOS-$GOARCH"
	rm -r "$buildDir" 2>/dev/null
	mkdir -p "$buildDir"
	# Include additional files
	ls -1 includes | while IFS= read -r file; do
		ln -s ../../includes/$file "$buildDir"
	done
	# Build Go executables
	ls -1 go | while IFS= read -r entrypoint; do
		if [ -f "go/${entrypoint}/main.go" ]; then
			if [ "$GOARCH" == "arm64" ]; then
				if [ "$GOOS" == "darwin" ]; then
					GOARM64="v8.4"
				elif [ "$GOOS" != "android" ]; then
					GOARM64="v8.2"
				fi
			fi
			if [ "$GOARCH" == "amd64" ]; then
				GOAMD64="v2"
			fi
			cd go
			echo "Building \"$entrypoint\" for target \"$GOOS-$GOARCH...\""
			go build -o ".$buildDir/$entrypoint" -trimpath -buildvcs=false -ldflags="-s -w -buildid=" "./$entrypoint"
			if [ "$GOOS" == "windows" ]; then
				mv ".$buildDir/$entrypoint" ".$buildDir/$entrypoint.exe" 2>/dev/null
			fi
			cd ..
		fi
	done
	rmdir "$buildDir" 2>/dev/null
	if [ ! -e "$buildDir" ]; then
		echo "Empty build result for \"$GOOS-$GOARCH\"."
		continue
	fi
	# Zip them into bundles
	cd "$buildDir"
	if [ "$GOOS" == "windows" ]; then
		zip -qr9 "../releases/$GOOS-$GOARCH.zip" *
	else
		tar -cf "../releases/$GOOS-$GOARCH.tar" *
		brotli -jv9 "../releases/$GOOS-$GOARCH.tar"
	fi
	cd ../..
	rm -r "$buildDir"
done