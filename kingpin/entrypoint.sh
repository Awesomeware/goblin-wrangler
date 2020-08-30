#!/bin/sh

$HOME/.poetry/bin/poetry run manage db upgrade
$HOME/.poetry/bin/poetry run gunicorn --bind 0.0.0.0:4000 kingpin.wsgi