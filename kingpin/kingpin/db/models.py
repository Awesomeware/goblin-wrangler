""" The core model objects of our DB """

from flask_login import UserMixin
from kingpin.application import db


class User(UserMixin, db.Model):
    """ Users represent people who login to our system """
    id = db.Column(db.Integer, primary_key=True)
    email = db.Column(db.String(120), index=True, unique=True)

    def __repr__(self):
        """ Textual representation for debugging """
        return '<User {}>'.format(self.email)


def create_or_update_user(user_info):
    """ Creates or updates a user from an OAuth login """
    existing = User.query.filter_by(email=user_info.email).first()
    if not existing:
        existing = User(email=user_info.email)
        db.session.add(existing)
        db.session.commit()
    return existing
