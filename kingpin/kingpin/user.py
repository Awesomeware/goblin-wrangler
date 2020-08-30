""" User API """

import uuid
from flask import Blueprint
from kingpin.db.models import User

user_api = Blueprint('user_api', __name__)


@user_api.route('/<user_id>')
def get_user(user_id):
    """ Gets a user's profile info """
    try:
        uuid.UUID(user_id)
    except ValueError:
        return {'error': 'User ID is malformed'}, 400
    user = User.query.get(user_id)
    if not user:
        return {'error': f'User with ID {user_id} not found'}, 404
    return {
        'id': user.id,
        'email': user.email
    }, 200
