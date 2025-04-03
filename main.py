from flask import Flask
from elasticapm.contrib.flask import ElasticAPM
import requests
import os

app = Flask(__name__)
apm = ElasticAPM(app)

@app.route('/')
def hello():
    endpoint = os.environ.get('ENDPOINT')
    if endpoint:
        print(requests.get(endpoint).content)
    return "Hello World py!"

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=80)