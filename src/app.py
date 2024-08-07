from flask import Flask, render_template, request
from flask_socketio import SocketIO, emit
import base64

app = Flask(__name__)
app.config['SECRET_KEY'] = 'secret!'
socketio = SocketIO(app)

@app.route('/')
def index():
    return render_template('code.html')

@socketio.on('message')
def handle_message(message):
    print(f"Received message: {message}")  # לוג להדפסה
    if isinstance(message, dict) and 'image' in message:
        image_data = message['image'].split(",")[1]  # במידה והדאטה מכיל מידע על סוג התוכן
        emit('message', {'image': image_data}, broadcast=True)
    else:
        emit('message', {'text': message}, broadcast=True)

if __name__ == '__main__':
    socketio.run(app)
