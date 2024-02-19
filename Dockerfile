FROM golang:latest
WORKDIR /src
COPY . /src
RUN go build -o main cmd/app/main.go 
EXPOSE 8080
CMD [ "./main" ]