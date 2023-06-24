FROM golang as builder

LABEL autor="Arty"
LABEL description="for skillfactory"

RUN apt-get update
RUN apt-get install -y git
WORKDIR /go/src/

RUN git clone https://github.com/ta01rus/SkillCh26A 
RUN go env -w GO111MODULE=auto
WORKDIR /go/src/SkillCh26A

RUN go mod download
RUN go build -o cmd/service

FROM alpine:latest
WORKDIR /app
COPY --from=builder /go/src/SkillCh26A/cmd /app/ 
ENTRYPOINT ["./service"]
