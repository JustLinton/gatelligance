FROM golang:latest

WORKDIR /app/light-shadow
COPY . .

#宿主编译
#RUN go build main.go
#交叉编译
RUN env GOOS=linux GOARCH=386 go build main.go

EXPOSE 8082
ENTRYPOINT ["./main"]
