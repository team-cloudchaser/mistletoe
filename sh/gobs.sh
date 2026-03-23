#!/bin/bash
# Go - build single
if [ "$1" == "" ]; then
	echo "The following entrypoints are available:"
	ls -1 go | while IFS= read -r folder; do
		if [ -f "./go/${folder}/main.go" ]; then
			echo "- ${folder}"
		fi
	done
	exit 1
fi
if [ ! -f "./go/${1}/main.go" ]; then
	echo "Entrypoint not found."
	exit 1
fi
mkdir -p build
cd go
go build -o "../build/${1}" "./${1}"
exit
