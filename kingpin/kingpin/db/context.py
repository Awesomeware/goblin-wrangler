""" Provides a SQLAlchemy context """

import os
from flask_sqlalchemy import SQLAlchemy
from kingpin.application import application

db = SQLAlchemy(application)
