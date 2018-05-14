from flask import Flask
import sys
import optparse
import time
import os
import requests
from requests.auth import HTTPBasicAuth

app = Flask(__name__)

start = int(round(time.time()))

@app.route("/colors")
def get_versions():
    awxHostname = os.environ['AWX_API_SERVICE_SERVICE_HOST']
    awxport = os.environ['AWX_API_SERVICE_SERVICE_PORT']
    awxUsername = os.environ['AWX_USERNAME']
    awxPassword = os.environ['AWX_PASSWORD']
    
    url = "http://" + awxHostname + ":" + awxport + "/api/v2/users/";
    print "URL : {0}".format(url)
    
    resp = requests.get(url, auth=HTTPBasicAuth(awxUsername, awxPassword))

    return str(resp.text)

if __name__ == '__main__':
    parser = optparse.OptionParser(usage="python simpleapp.py -p ")
    parser.add_option('-p', '--port', action='store', dest='port', help='The port to listen on.')
    (args, _) = parser.parse_args()
    if args.port == None:
        print "Missing required argument: -p/--port"
        sys.exit(1)
    app.run(host='0.0.0.0', port=int(args.port), debug=False)
