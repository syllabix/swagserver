#!/usr/bin/env bash

############
# This script can be used to build swagger ui into the
############

echo "Fetching swagger-ui-dist from npm"
npm install swagger-ui-dist

mkdir tmp

echo "Copying swagger ui into the target tmp folder"
FILES=(
    'index.html'
    'swagger-ui.css'
    'favicon-32x32.png'
    'favicon-16x16.png'
    'swagger-ui-bundle.js'
    'swagger-ui-standalone-preset.js')


for file in "${FILES[@]}"; do
  cp node_modules/swagger-ui-dist/$file tmp
done

sed -i '' 's%https://petstore.swagger.io/v2/swagger.json%{{ . }}%g' tmp/index.html

echo "Running go generate to create static binary of swagger ui"
go generate

echo "Cleaning up"
rm -rf node_modules
rm -rf tmp

echo "Process Complete"

