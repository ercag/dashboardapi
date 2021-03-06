FROM golang:1.17

ENV APP_NAME myproject
ENV PORT 5000

ADD . /go/src/${APP_NAME}

COPY . /go/src/${APP_NAME}
WORKDIR /go/src/${APP_NAME}

RUN go get ./
RUN go build -o ${APP_NAME}

CMD ./${APP_NAME}

EXPOSE ${PORT}