""" Sets up the Flask web-app """

import os
from flask import Flask, jsonify
from flask_login import LoginManager, login_user, current_user, login_required
from flask_migrate import Migrate
from authlib.integrations.flask_client import OAuth
from loginpass import create_flask_blueprint
from loginpass import Google

application = Flask(__name__)

application.secret_key = os.environ.get("SECRET_KEY") or os.urandom(24)
cfg = application.config
cfg['GOOGLE_CLIENT_ID'] = os.environ.get('GOOGLE_CLIENT_ID')
cfg['GOOGLE_CLIENT_SECRET'] = os.environ.get('GOOGLE_CLIENT_SECRET')

dbuser = os.environ.get("POSTGRES_USER")
dbpass = os.environ.get("POSTGRES_PASSWORD")
dbs = os.environ.get("POSTGRES_DB")
DBCFG = 'SQLALCHEMY_DATABASE_URI'
cfg[DBCFG] = f"postgresql+psycopg2://{dbuser}:{dbpass}@postgres:5432/{dbs}"
cfg['SQLALCHEMY_TRACK_MODIFICATIONS'] = False

backends = [Google]
login_manager = LoginManager()
login_manager.init_app(application)

# pylint: disable=wrong-import-position
from kingpin.db.context import db  # noqa
from kingpin.db.models import User  # noqa
from kingpin.db import functions  # noqa
from kingpin.healthcheck import healthcheck_api  # noqa
# pylint: enable=wrong-import-position

migrate = Migrate(application, db)


def handle_authorize(_remote, token, user_info):
    """ Handles successful authorization via OAuth """
    if user_info:
        user = functions.create_or_update_user(user_info)
        login_user(user)
    return jsonify({
        'user': user.id,
        'user_info': user_info,
        'token': token
        })


@login_manager.user_loader
def load_user(user_id):
    """Check if user is logged-in on every page load."""
    if user_id is not None:
        return User.query.get(user_id)
    return None


@application.route('/loggedin', methods=['GET'])
@login_required
def logged_in():
    """ Shows whether we are logged in or not """
    return jsonify({
        'id': current_user.id,
        'email': current_user.email
    })


oauth = OAuth(application)
auth_api = create_flask_blueprint(backends, oauth, handle_authorize)

application.register_blueprint(auth_api, url_prefix='')
application.register_blueprint(healthcheck_api, url_prefix='/health')

if __name__ == "__main__":
    application.run(host='0.0.0.0', port=4000)
