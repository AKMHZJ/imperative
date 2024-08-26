FROM golang:1.22.5-alpine

WORKDIR /home/app

COPY . .

LABEL projectName="ascii-art-web(dockerize)"
LABEL Authors="Mehdi.moulabbi-hamza.boutalbe-mohamed.ghdaigui"

RUN go build -o main .

CMD [ "./main" ]