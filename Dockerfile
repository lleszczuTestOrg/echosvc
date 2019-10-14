FROM golang:alpine
WORKDIR /app
ADD . /app
RUN cd /app && go build -o echosvc
ENTRYPOINT ./echosvc

