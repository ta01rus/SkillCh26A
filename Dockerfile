FROM golang as builder

LABEL autor="Arty"
LABEL description="for skillfactory"

RUN apt-get update
RUN apt-get install -y git

RUN mkdir -p build/

RUN git clone git@github.com:ta01rus/SkillCh26A.git 
WORKDIR /build/SkillCh26A/

RUN go build -o cmd/app

FROM alpine:latest

USER gofer

WORKDIR /app
COPY --from=builder build/cmd/  /app/app 
ENTRYPOINT ["./app"]
