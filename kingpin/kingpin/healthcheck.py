from flask import Blueprint

healthcheck_api = Blueprint('healthcheck_api', __name__)

circuitbreakers = []

@healthcheck_api.route("/check")
def health():
    status = 200
    content = {}
    for func in circuitbreakers:
        check = func()
        if check["healthy"] != True:
            status = 503
        content[check["name"]] = check
    return content, status

@healthcheck_api.after_request
def add_header(response):
    response.cache_control.no_store = True
    return response
