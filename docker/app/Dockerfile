FROM golang:1.14
COPY . /go/src/coffee-break
WORKDIR /go/src/coffee-break

RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u github.com/beego/bee
RUN go get -u github.com/astaxie/beego

# コンテナ実行時にコンパイルと実行を行う
CMD bee run
