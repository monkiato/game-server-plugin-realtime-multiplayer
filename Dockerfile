FROM golang:1.13.5 as serverBuilder
WORKDIR /app

COPY . .

RUN go mod download
RUN go build --trimpath --buildmode=plugin -o /plugin.so plugin.go

FROM monkiato/game-server:0.0.1 as server

COPY --from=serverBuilder /plugin.so /app/plugins/realtime_multiplayer_plugin.so