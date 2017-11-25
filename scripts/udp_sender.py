
import socket
import json

UDP_IP = "0.0.0.0"
UDP_PORT = 5003

MESSAGE = json.dumps({
    "to": "Magnus",
    "from": "Bjorn",
    "msg": "hello",
    "test-id": "MAG"})

print(MESSAGE)

sock = socket.socket(socket.AF_INET,     # Internet
                     socket.SOCK_DGRAM)  # UDP
sock.sendto(MESSAGE, (UDP_IP, UDP_PORT))
