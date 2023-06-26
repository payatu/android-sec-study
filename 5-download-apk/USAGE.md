# Create Digital Ocean functions

1. Create a new namespace
2. Create a Python function with pacakge name `apkpure` and function name of your choice.
3. For the python code, copy it from `do.py`

# Populate namespace-URL.txt

1. After creating the function, copy the URL, strip the function name from the URL such that it ends at `/apkpure`. 

$ echo "https://*.doserverless.co/api/v1/web/*/apkpure" >> namespace-URL.txt

# Populate functions.txt

1. Put all the function names in `functions.txt`

$ functions="a b c" for i in $functions; do echo $i >> functions.txt; done

# Download APKs

1. Make sure aria2c is installed.
2. Put package names in `apk.txt`
3. Run the python script

$ ./dl.py
