FROM golang:1.15

WORKDIR /go/src/coffee-break
COPY . .

RUN go get -u github.com/beego/bee
RUN go get -u github.com/astaxie/beego

CMD bee run
