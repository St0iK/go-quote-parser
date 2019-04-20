package parser

type Parser interface {
	Process() (string, error)
}
