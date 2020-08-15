from flask import Flask, jsonify
from authlib.integrations.flask_client import OAuth
from loginpass import create_flask_blueprint
from loginpass import Google
import os
from kingpin.healthcheck import healthcheck_api

backends = [Google]

def handle_authorize(remote, token, user_info):
    return jsonify(user_info)

application = Flask(__name__)
oauth = OAuth(application)

application.secret_key = os.environ.get("SECRET_KEY") or os.urandom(24)
application.config['GOOGLE_CLIENT_ID'] = os.environ.get('GOOGLE_CLIENT_ID')
application.config['GOOGLE_CLIENT_SECRET'] = os.environ.get('GOOGLE_CLIENT_SECRET')

bp = create_flask_blueprint(backends, oauth, handle_authorize)

application.register_blueprint(bp, url_prefix='')
application.register_blueprint(healthcheck_api, url_prefix='/health')

if __name__ == "__main__":
    application.run(host='0.0.0.0', port=4000)
