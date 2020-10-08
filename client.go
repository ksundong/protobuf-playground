package main

import (
        protos "./protos"
        proto "github.com/golang/protobuf/proto"
        http "net/http"
        ioutil "io/ioutil"
        bytes "bytes"
        fmt "fmt"
)

func main() {
        // HelloRequest 프로토 값 생성
        request := protos.HelloRequest{
                Name: "World",
        }
        requestInBinary, _ := proto.Marshal(&request)

        // localhost:8080 으로 프로토 전송
        resp, _ := http.Post("http://localhost:8080", "application/octet-stream", bytes.NewBuffer(requestInBinary))

        // HelloResponse 프로토 값 읽음
        var helloResponse protos.HelloResponse
        body, _ := ioutil.ReadAll(resp.Body)
        proto.Unmarshal(body, &helloResponse)

        fmt.Printf("Received %v from server\n", helloResponse.GetText())
}

