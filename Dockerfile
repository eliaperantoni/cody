FROM alpine

RUN apk add --no-cache go
COPY . .
RUN cd api && go build -o ../server api/cmd && cd ..

CMD ./server
