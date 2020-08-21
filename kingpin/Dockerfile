FROM python:3.7

RUN curl -sSL https://raw.githubusercontent.com/python-poetry/poetry/master/get-poetry.py | python

WORKDIR /app
COPY poetry.lock /app/
COPY *.toml /app/
RUN $HOME/.poetry/bin/poetry install --no-dev -n --no-ansi
COPY . /app

EXPOSE 4000
ENTRYPOINT ["/app/entrypoint.sh"]