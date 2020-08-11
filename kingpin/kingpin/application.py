from flask import Flask
from kingpin.healthcheck import healthcheck_api

application = Flask(__name__)
application.register_blueprint(healthcheck_api, url_prefix='/health')

if __name__ == "__main__":
    application.run(host='0.0.0.0', port=4000)
