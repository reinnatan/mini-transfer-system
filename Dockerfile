FROM golang:1.14.3-alpine AS build
#RUN apk --no-cache add ca-certificates git
RUN apk add curl nano
ENV GO111MODULE=on
RUN CGO_ENABLED=0
RUN mkdir /app
WORKDIR /app
EXPOSE 8000
COPY . .
ADD . /app
RUN go mod download
RUN go build -o api main.go
#COPY --from=builder api /app
ENTRYPOINT ["/app/api"]
#CMD ["/app/api"]

# add these two lines
# ADD go.mod go.sum /app/

#WORKDIR /
#WORKDIR /app
#ADD go.mod go.sum /app/

#COPY  . .
#RUN CGO_ENABLED=0 GOOS=linux go mod download
#RUN CGO_ENABLED=0 GOOS=linux go build -o main main.go

#RUN go mod download
#ADD . .
#RUN go build -o main main.go

#CMD ["go", "mod" "download"]
#EXPOSE 8000 8000    
#CMD ["go", "run" "main.go"]
