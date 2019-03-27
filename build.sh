#!/bin/sh

build_file="wfs_linux_amd64"

gox -os="linux" -arch="amd64"

if [ ! -f "$build_file" ]; then
 echo "编译失败，请检查代码！"
 exit 0
fi

scp wfs_linux_amd64 root@66.15.113.16:/home




