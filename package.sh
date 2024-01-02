#!/bin/bash

rm -rf build dist
mkdir -p build/bin

pnpm run build
cp -r dist build/

make -C backend backend 
cp ./backend/out/backend build/bin/backend

cp package.json build/package.json
cp plugin.json build/plugin.json
cp main.py build/main.py
cp README.md build/README.md
cp LICENSE build/LICENSE