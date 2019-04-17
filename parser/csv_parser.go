package parser

type Parser interface {
	Process(filename string) (string, error)
}
