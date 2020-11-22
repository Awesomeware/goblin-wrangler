""" Provides a SQLAlchemy context """

from flask_sqlalchemy import SQLAlchemy
from kingpin.application import application

db = SQLAlchemy(application)
