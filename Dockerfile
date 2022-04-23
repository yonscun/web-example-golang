FROM golang:latest
WORKDIR /app
COPY . ./
RUN go mod download
RUN go build -o /goapp
EXPOSE 8080
CMD [ "/goapp" ]
