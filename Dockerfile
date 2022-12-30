FROM golang:1.20rc1-alpine3.17

WORKDIR /usr/scrap-servise

COPY go.mod ./

COPY go.sum ./

RUN go mod download

COPY ./ ./

RUN go build -o scrap-servise ./cmd/scraping-app/

CMD [ "./scrap-servise"]