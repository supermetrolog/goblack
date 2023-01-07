#!/bin/bash

MOCKS_PATH=tests/mocks

echo "Running mocks manager"

echo "Remove all mocks..."
rm -rf ./${MOCKS_PATH}

echo "Generate PKG mocks..."
pkg_files=(
            http/interfaces/handler/handler.go
            http/interfaces/request/request.go
            http/interfaces/response/response.go
            )

for file in ${pkg_files[*]}
do
    sourcefullpath=pkg/$file
    targetfullpath=pkg/$file
    echo gen $sourcefullpath
    mockgen -source=$sourcefullpath -destination=$MOCKS_PATH/$targetfullpath
done

echo "Generate INTERNAL mocks..."
pkg_files=(
            )

for file in ${pkg_files[*]}
do
    sourcefullpath=internal/$file
    targetfullpath=intern/$file # Нельзя писать INTERNAL т.к GO не позволяет
    echo gen $sourcefullpath
    mockgen -source=$sourcefullpath -destination=$MOCKS_PATH/$targetfullpath
done