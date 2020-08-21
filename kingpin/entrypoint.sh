#!/bin/sh

cd kingpin
$HOME/.poetry/bin/poetry run flask db upgrade
cd ..
$HOME/.poetry/bin/poetry run gunicorn --bind 0.0.0.0:4000 kingpin.wsgi