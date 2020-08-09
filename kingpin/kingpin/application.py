from flask import Flask

application = Flask(__name__)

if __name__ == "__main__":
    application.run(host='0.0.0.0', port=4000)
