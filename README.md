### Go Quote Parser

Simple go app that parsers various file containing quotations and inserts them in a MongoDB database


## Running the parser
`./go run main.go`

## Parsing a new json file

- Set up the new parser configuration inside config.yaml
- Add new file with quotes inside the correct folder
- Register the new parser inside main.go
-  `./go run main.go`