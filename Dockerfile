FROM golang:1.21
ENV TZ=America/Sao_Paulo

RUN mkdir /home/babychain
WORKDIR /home/babychain

CMD tail -f /dev/null