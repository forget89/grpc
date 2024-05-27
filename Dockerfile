FROM golang:latest

RUN apk add build-base
COPY go.mod go.sum* ./
RUN go mod download
EXPOSE 80
EXPOSE 8080
EXPOSE 3000
EXPOSE 50051
COPY . .
RUN go build phase_server.go 
CMD ["./server"]

# ADD . /go/src/phase/go.mod

# RUN go install go.mod

# ENTRYPOINT ["/phase/phase_server"]

# EXPOSE 50051

# RUN apk update

# RUN mkdir /phase
# RUN mkdir -p /phase/proto
# WORKDIR /phase
# COPY .phase/phase.pb.go /phase
# COPY ./phase_client/phase_client.go /phase

# COPY /phase/go.mod .
# COPY /phase/go.mod .

# RUN go mod download 

# RUN go build -o phase .
# CMD ./phase