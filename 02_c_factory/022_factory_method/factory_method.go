package factory

type IRuleConfigParser interface {
	Parse(data []byte)
}

type jsonRuleConfigParser struct{}

func (j jsonRuleConfigParser) Parse(data []byte) {
	panic("implement me")
}

type yamlRuleConfigParser struct{}

func (y yamlRuleConfigParser) Parse(data []byte) {
	panic("implement me")
}

type IRuleConfigParserFactory interface {
	CreateParser() IRuleConfigParser
}

type yamlRuleConfigParserFactory struct{}

func (y yamlRuleConfigParserFactory) CreateParser() IRuleConfigParser {
	return yamlRuleConfigParser{}
}

type jsonRuleConfigParserFactory struct{}

func (j jsonRuleConfigParserFactory) CreateParser() IRuleConfigParser {
	return jsonRuleConfigParser{}
}

func NewIRuleConfigParserFactory(t string) IRuleConfigParserFactory {
	switch t {
	case "yaml":
		return yamlRuleConfigParserFactory{}
	case "json":
		return jsonRuleConfigParserFactory{}
	}
	return nil
}
