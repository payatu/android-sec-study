import requests

def main(args):
    app = args.get("app", "com.whatsapp")
    apptype = args.get("apptype", "APK")
    headers = {'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; rv:110.0) Gecko/20100101 Firefox/110.2'}
    r = requests.head("https://d.apkpure.com/b/" + apptype + "/" + app + "?version=latest", headers=headers, allow_redirects=False, verify=False)
    return {"body": r.headers['location']}
