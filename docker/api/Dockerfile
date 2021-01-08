FROM golang:latest

COPY src/api /go/src/api

WORKDIR /go/src/api/

RUN go get -u github.com/gin-gonic/gin && \
    go get -u github.com/jinzhu/gorm && \
    go get -u github.com/go-sql-driver/mysql && \
    go get -u github.com/sclevine/agouti 

CMD ["sh", "./start_app.sh"] 