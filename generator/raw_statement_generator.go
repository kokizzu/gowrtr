package gowrtr

type RawStatementGenerator struct {
	Statement   string
	WithNewline bool
}

func NewRawStatementGenerator(stmt string, withNewline ...bool) *RawStatementGenerator {
	nl := false
	if len(withNewline) > 0 {
		nl = withNewline[0]
	}

	return &RawStatementGenerator{
		Statement:   stmt,
		WithNewline: nl,
	}
}

func (r *RawStatementGenerator) Generate(indentLevel int) (string, error) {
	indent := buildIndent(indentLevel)

	newline := ""
	if r.WithNewline {
		newline = "\n"
	}

	return indent + r.Statement + newline, nil
}