FROM python:3.7

RUN curl -sSL https://raw.githubusercontent.com/python-poetry/poetry/master/get-poetry.py | python
RUN poetry config virtualenvs.create false

WORKDIR /app
COPY . /app
RUN poetry install

ENTRYPOINT ["gunicorn", "--bind 0.0.0.0:4000", "kingpin.wsgi"]
