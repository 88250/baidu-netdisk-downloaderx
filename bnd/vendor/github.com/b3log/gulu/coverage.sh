#!/bin/bash
# see https://gist.github.com/hailiang/0f22736320abe6be71ce for more details
 
set -e
 
# Run test coverage on each subdirectories and merge the coverage profile.
 
echo "mode: count" > profile.cov
 
# Standard go tooling behavior is to ignore dirs with leading underscors
for dir in $(find . -maxdepth 10 -not -path './.git*' -not -path '*/_*' -not -path './vendor*' -type d);
do
if ls $dir/*.go &> /dev/null; then
    go test -covermode=count -coverprofile=$dir/profile.tmp $dir
    if [ -f $dir/profile.tmp ]
    then
        cat $dir/profile.tmp | tail -n +2 >> profile.cov
        rm $dir/profile.tmp
    fi
fi
done
 
go tool cover -func profile.cov

