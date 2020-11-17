FROM golang:1.12.13

ENV REPO_URL=github.com/themiken/GoApi

ENV GOPATH=/app

ENV APP_PATH=$GOPATH/src/$REPO_URL

ENV WORKPATH=$APP_PATH/src
COPY src $WORKPATH
WORKDIR $WORKPATH

RUN go build -o mikenapi .

EXPOSE 8081

CMD ["./mikenapi"]



