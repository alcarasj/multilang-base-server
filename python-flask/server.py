from flask import Flask
server = Flask(__name__)

@server.route("/")
def index():
    return "Welcome! This is a base server written in Python using the Flask framework."

if __name__ == "__main__":
    server.run(host="0.0.0.0", port=8080)
