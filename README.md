
# Daniel's Chess game written in GO
## Milestones:
1. MVP: Be able to play a game of chess against a super dumb computer player. ✅
2. Create AI agent and replace the dumbness with the AI agent. The computer becomes very smart.  🌕
3. Create a web socket capability that allows the react app render/communicate with the go game. 🌕
4. Game Simulator, Castling, Player can choose to be black/white, can choose difficulty level 🌕

##### Go installation
1. In the terminal after homebrew is installed: ```brew install golang```
2. Check the version in the terminal by running ```go version```

##### Run all unit tests
```go test ./...```

##### Start game
```go run cmd/main.go```

##### Compile program
```
go build cmd/main.go
```

##### Run executable
```
./main
```
##### About this executable
The executable in this project was built on apple silicon. If you're not on apple silicon you'll have to rebuild.
