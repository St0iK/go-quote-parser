package parser

//The second implementation.
type JsonV2Parser struct {
	filename string
}

func (pds *JsonV2Parser) Process(file string) (string, error) {
	return "Processing......JsonV2Parser", nil
}

func NewJsonV2Factory(conf map[string]string) (Parser, error) {

	return &JsonV2Parser{
		"something 2",
	}, nil
}
