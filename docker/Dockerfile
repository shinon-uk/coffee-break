# @see::https://hub.docker.com/_/amazonlinux
FROM amazonlinux:2

# システムアップデート
RUN yum update -y

# gitのインストール
RUN yum install -y git

# @see::https://fedoraproject.org/wiki/EPEL
RUN amazon-linux-extras install -y epel

# amazon-linux-extrasでインストールできる最新のgolang
RUN amazon-linux-extras list | grep golang
RUN amazon-linux-extras install -y golang1.11
RUN go version

# beegoのインストール
ENV GOPATH /go
ENV PATH $PATH:$GOPATH/bin
RUN go get -u github.com/beego/bee
RUN go get -u github.com/astaxie/beego

# カレントディレクトリをコンテナに追加する
COPY . /go/src/coffee-break

# 作業ディレクトリを指定する
WORKDIR /go/src/coffee-break

# コンテナ実行時にコンパイルと実行を行う
CMD bee run
