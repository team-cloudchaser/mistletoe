#!/bin/bash
# A multi-arch build release script
export CGO_ENABLED=0
rm -r "./build/releases" 2>/dev/null
mkdir -p "./build/releases"
cat conf/goTargets.txt | while IFS= read -r GOUNION; do
	export GOOS=$(printf $GOUNION | cut -d'/' -f1)
	export GOARCH=$(printf $GOUNION | cut -d'/' -f2)
	buildDir="./build/$GOOS-$GOARCH"
	rm -r "$buildDir" 2>/dev/null
	mkdir -p "$buildDir"
	# Include additional files
	# Build Go executables
	ls -1 go | while IFS= read -r entrypoint; do
		if [ -f "go/${entrypoint}/main.go" ]; then
			cd go
			echo "Building \"$entrypoint\" for target \"$GOOS-$GOARCH...\""
			go build -o ".$buildDir/$entrypoint" -trimpath -buildvcs=false -ldflags="-s -w -buildid=" "./$entrypoint"
			if [ "$GOOS" == "windows" ]; then
				mv ".$buildDir/$entrypoint" ".$buildDir/$entrypoint.exe" 2>/dev/null
			fi
			cd ..
		fi
	done
	# Zip them into bundles
	rmdir "$buildDir" 2>/dev/null
done