from flask import Flask, send_from_directory

#route of static
app = Flask(__name__, static_folder='src/static')

@app.route('/')
def serve_cv():
    return send_from_directory(app.static_folder, 'Alexandru_ALBU_CV.pdf')

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)
