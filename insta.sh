#!/bin/sh
VAR="$(python sentance.py)"
echo ${VAR} >> OUTPUT.txt
echo ${VAR} > poem.txt
go run imgGen.go < poem.txt
cat poem.txt
convert test.png -scale x1080 test.png 