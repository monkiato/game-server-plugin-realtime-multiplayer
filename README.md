# Game Server Plugin - Realtime Multiplayer

**[in-progress]**

Plugin to include realtime multiplayer support for the game server.

### Available features

 - Create match
 - Join match
 - leave match
 - user data stream

### Generate Go files from schema

Run `flatc -g --grpc schema/multiplayer.fbs`

This must 
### Build plugin

Run `go build --buildmode=plugin -o out/realtime_multiplayer_plugin.so plugin.go`