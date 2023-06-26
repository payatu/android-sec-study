from bs4 import BeautifulSoup
import urllib.request
import os
import requests
import random
import urllib3
urllib3.disable_warnings(urllib3.exceptions.InsecureRequestWarning)

with open('apk.txt') as f:
       apps = f.read().splitlines()

def get_content(app,apptype):
       with open('namespace-URL.txt', 'r') as namespaceFile:
          namespaceURL = namespaceFile.read().rstrip()

       functionList = []
       f = open("functions.txt", "r")
       for i in f.readlines():
          for value in i.strip().split("\n"):
            functionList.append(value)
       function = random.choice(functionList)
       r = requests.get(namespaceURL + "/" + function + "?app=" +  app + "&type=" + apptype, verify=False)
       return str(r.content.decode("utf-8"))

i = 0
winudf = "winudf"
for app in apps:
    try:
        dl = get_content(app,"APK")
        if not winudf in dl:
          dl = get_content(app,"XAPK")
          if not winudf in dl:
            raise Exception('Error')
        print(app + " : " + dl)
        last_app = open("last.txt", "w")
        last_app.write(app)
        last_app.close()
        dl_file = open("dl.txt", "a")
        dl_file.write(dl+"\n")
        dl_file.close()
        if i == 100:
            os.system('aria2c -d dl/ -i dl.txt -j100 -x16 -c -m5 --retry-wait 5 && rm -rf dl.txt')
            i = 0
        else:
            i = i + 1
    except Exception as e:
        print(app + " : " + str(e))
        err_file = open("err.txt", "a")
        err_file.write(app+"\n")
        err_file.close()
        if i == 100:
            os.system('aria2c -d dl/ -i dl.txt -j100 -x16 -c -m5 --retry-wait 5 && rm -rf dl.txt')
            i = 0
        else:
            i = i + 1

#os.system('aria2c -d dl/ -i dl.txt -j100 -x16 -c -m5 --retry-wait 5')
