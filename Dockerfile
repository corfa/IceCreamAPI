FROM golang:latest
EXPOSE 8080
RUN git clone https://github.com/corfa/IceCreamApi.git
COPY ./ ./
RUN go build -o app ./main.go
CMD ["./app"]