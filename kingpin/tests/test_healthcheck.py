import pytest

from kingpin.application import application

@pytest.fixture
def client():
    application.config['TESTING'] = True
    with application.test_client() as client:
        yield client

def test_healthcheck_okay(client):
    rv = client.get('/health/check')
    assert rv.status_code == 200, "Status code should be 200"