# Introdução GO
Atividade final da aula ministrada por Elton Minetto na Pós-Graduação da UniALFA Umuarama de Desenvolvimento de Aplicações para Internet e Dispositivos Móveis.

## Executando
```
go run main.go
```

## Executando testes
```
go test ./...
```

## Compilando

```
go build -o bin/curso main.go
GOOS=darwin GOARCH=amd64 go build -o bin/curso main.go
GOOS=darwin GOARCH=arm64 go build -o bin/curso main.go
GOOS=windows GOARCH=amd64 go build -o bin/curso main.go
GOOS=linux GOARCH=amd64 go build -o bin/curso main.go
```