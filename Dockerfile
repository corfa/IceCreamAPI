FROM golang:latest
EXPOSE 8080
RUN git clone https://github.com/corfa/IceCreamApi.git
RUN go mod download