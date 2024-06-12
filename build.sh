#!/usr/bin/env bash

############
# This script can be used to embed the latest version of
# swagger ui binary
############

echo "Fetching swagger-ui-dist from npm"
npm install swagger-ui-dist
echo "Fetching css themes"
npm install swagger-ui-themes

mkdir tmp

echo "Copying swagger ui into the target tmp folder"
FILES=(
    'index.html'
    'index.css'
    'swagger-ui.css'
    'favicon-32x32.png'
    'favicon-16x16.png'
    'swagger-ui-bundle.js'
    'swagger-initializer.js'
    'swagger-ui-standalone-preset.js')


for file in "${FILES[@]}"; do
  cp node_modules/swagger-ui-dist/$file tmp
done

echo "Copying swagger ui themes into tmp folder"
THEMES=(
    'theme-feeling-blue.css'
    'theme-flattop.css'
    'theme-material.css'
    'theme-monokai.css'
    'theme-muted.css'
    'theme-newspaper.css'
    'theme-outline.css')

for theme in "${THEMES[@]}"; do
  cp node_modules/swagger-ui-themes/themes/3.x/$theme tmp
done

sed -i '' 's%https://petstore.swagger.io/v2/swagger.json%{{ .SpecURL }}%' tmp/index.html
sed -i '' $'s%</style>%</style>\\\n\\\t\\\t<link rel="stylesheet" type="text/css" href="{{ .BasePath }}/{{ .Theme }}.css" >%' tmp/index.html
sed -i '' 's%\./%{{ .BasePath }}/%' tmp/index.html

echo "Running go generate to create static binary of swagger ui"
go generate

echo "Cleaning up"
rm -rf node_modules
rm -rf tmp
rm package-lock.json

echo "Process Complete"

