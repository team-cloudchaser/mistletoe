#!/bin/bash
# Example runner
# Parse and generate
# Dispatch
tmpFile=".tmp.json"
lineArgs=( $MISTLETOE_LINE )
if [ -f "$(which xray)" ]; then
	rm "$tmpFile" 2>/dev/null
	cp "data/${1}.json" "$tmpFile"
	sed -i "s/__SERVER__/${lineArgs[1]}/g" "$tmpFile"
	xray run -c "$tmpFile"
else
	echo "Required proxy toolchain is not installed."
fi
exit