#!/bin/bash
args=( "$@" )
# Go - test single
shx gobs "${1}"
if [ "$?" == "0" ]; then
	"./build/${1}" "${args[@]:1}"
fi
exit
