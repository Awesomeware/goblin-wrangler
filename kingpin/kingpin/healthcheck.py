""" Healthcheck API, for load balancing """

from flask import Blueprint

healthcheck_api = Blueprint('healthcheck_api', __name__)

circuitbreakers = []


@healthcheck_api.route("/check")
def health():
    """ The health endpoint, showing if we should accept traffic """
    status = 200
    content = {}
    for func in circuitbreakers:
        check = func()
        if not check["healthy"]:
            status = 503
        content[check["name"]] = check
    return content, status


@healthcheck_api.after_request
def add_header(response):
    """ Adds no-cache headers to healthcheck endpoints """
    response.cache_control.no_store = True
    return response
