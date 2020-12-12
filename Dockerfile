FROM alpine

RUN apk add --no-cache go yarn
COPY . .
RUN cd api && go build -o ../server api/cmd && cd ..
RUN cd web && yarn install && yarn build && cd ..
CMD ./server
