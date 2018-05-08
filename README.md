# ServiceGenerator
Simple Golang cmd to generate dotnetCore Service / Controller in MVC pattern


## usage

`go build Generator && go install Generator`

inside dotnetcore Project:
`Generator service --name {ItemName} --context {DatabaseContext} --postgres true`

Runs in postgres or Mssql flavor for the pagination


Using inside a dotnet core project to generate Controller / Service

Some attributes used are specific for my needs
