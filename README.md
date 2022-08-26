# pokemon-cli
This is a little CLI to interact with the [pokemon api](https://pokeapi.co/).

## Build
```go
go build cmd/main.go
```

## Commands

### Get a pokemon
```
./main pokemon -n  # or ./main pokemon --name pikachu
```

### Get multiple pokemons
To use this command, a "from" and a "to" argument must be passed.
```
./main pokemons -f 0 -t 10 # or ./main pokemons --from 0 --to 10
```

### Get pokemon's moves
```
./main moves -n pikachu # or ./main moves --name pikachu
```

### Get pokemon's abilities
```
./main abilities -n pikachu # or ./main abilities --name pikachu
```

### Get version
```
./main version # 1.0.0
```

### Display help
```
./main help # or ./main -h
```
