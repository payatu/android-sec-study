#!/bin/bash

## Note: make sure wget, curl, gunzip, xmllint are installed.

mkdir -p dump
wget --header="User-Agent: Mozilla/5.0 (X11; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/115.0" -P dump/ $(curl -sSL -H "User-Agent: Mozilla/5.0 (X11; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/115.0" https://apkpure.com/sitemap.xml | grep ".xml.gz" | cut -d ">" -f2 | cut -d "<" -f1)

gunzip dump/*

for name in $(ls dump/)
do
        xmllint --format dump/$name| grep "https://apkpure.com" | cut -d"/" -f5 | cut -d "<" -f1 >> /tmp/package-list.txt
done

cat /tmp/package-list.txt | sort -u > package-list.txt

rm -rf dump/ /tmp/package-list.txt
