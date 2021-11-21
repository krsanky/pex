#!/bin/sh

find . -name "*.go" | xargs grep -l 'go\.d34d\.net\/pex' | \
xargs sed 's#go\.d34d\.net#github.com/krsanky#' 

