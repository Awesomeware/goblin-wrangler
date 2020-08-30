""" Functions related to manipulating DB objects """

from kingpin.db.context import db
from kingpin.db.models import User


def create_or_update_user(user_info):
    """ Creates or updates a user from an OAuth login """
    existing = User.query.filter_by(email=user_info.email).first()
    if not existing:
        existing = User(email=user_info.email)
        db.session.add(existing)
        db.session.commit()
    return existing
