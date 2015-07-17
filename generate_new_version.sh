#!/bin/bash

VERSION=$(cat VERSION)
NEW_VERSION=${VERSION%.*}.$((${VERSION##*.}+1))
echo $NEW_VERSION > VERSION
sed -i "s:download/v[^\/]*/:download/vX.X.X/:" Dockerfile
git add VERSION Dockerfile
git commit -m "prepare release ${NEW_VERSION}"

