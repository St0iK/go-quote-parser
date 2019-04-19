package parser

type Parser interface {
	Process(conf map[string]string) (string, error)
}
