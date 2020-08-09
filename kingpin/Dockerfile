FROM python:3.7

RUN curl -sSL https://raw.githubusercontent.com/python-poetry/poetry/master/get-poetry.py | python

WORKDIR /app
COPY . /app
RUN $HOME/.poetry/bin/poetry install --no-dev -n --no-ansi

EXPOSE 4000
ENTRYPOINT ["/root/.poetry/bin/poetry"]
CMD ["run", "gunicorn", "--bind", "0.0.0.0:4000", "kingpin.wsgi"]