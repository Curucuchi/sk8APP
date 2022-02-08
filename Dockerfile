FROM golang

RUN mkdir -p /myapp

COPY . /myapp

CMD [ "go run", "/myapp/main.go" ]