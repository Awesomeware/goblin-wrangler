""" The core model objects of our DB """

from flask_login import UserMixin
from sqlalchemy.dialects.postgresql.base import UUID
from sqlalchemy import Table, Column, String, ForeignKey
from sqlalchemy.orm import relationship
from kingpin.db.context import db
import uuid


party_users = Table(
                    'party_users',
                    db.metadata,
                    Column('user_id', UUID(as_uuid=True),
                           ForeignKey('user.id', ondelete="CASCADE")),
                    Column('party_id', UUID(as_uuid=True),
                           ForeignKey('party.id', ondelete="CASCADE"))
                   )


class User(UserMixin, db.Model):
    """ Users represent people who login to our system """
    __tablename__ = 'user'
    id = Column(UUID(as_uuid=True),
                primary_key=True,
                default=uuid.uuid4,
                unique=True,
                nullable=False)
    email = Column(String(120), index=True, unique=True)
    parties = relationship("Party", secondary=party_users,
                           back_populates="users", passive_deletes=True)

    def __repr__(self):
        """ Textual representation for debugging """
        return '<User {}>'.format(self.email)


class Party(db.Model):
    """ Parties represent a group of players who play game together """
    __tablename__ = 'party'
    id = Column(UUID(as_uuid=True),
                primary_key=True,
                default=uuid.uuid4,
                unique=True,
                nullable=False)
    name = Column(String(120), index=True)
    users = relationship("User", secondary=party_users,
                         back_populates="parties", passive_deletes=True)
