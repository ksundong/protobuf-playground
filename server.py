from protos import service_pb2
from flask import Flask, request

app = Flask(__name__)

@app.route('/', methods=['POST'])
def hello():
    data = request.get_data()

    # HelloRequest 프로토 타입으로 변환
    req = service_pb2.HelloRequest()
    req.ParseFromString(data)

    # HelloResponse 프로토 생성
    resp = service_pb2.HelloResponse()
    resp.text = f"Hello {req.name}"

    return resp.SerializeToString()

if __name__ == "__main__":
    app.run(debug=True, port=8080)

