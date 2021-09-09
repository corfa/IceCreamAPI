FROM golang:latest
EXPOSE 8080
RUN git clone https://github.com/corfa/IceCreamApi.git
COPY ./ ./
RUN go build -o app ./cmd/main.go
CMD ["./app"]