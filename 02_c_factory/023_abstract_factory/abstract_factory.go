package factory

type IRuleConfigParser interface {
	Parser(data []byte)
}

type jsonRuleConfigParser struct{}

func (j jsonRuleConfigParser) Parser(data []byte) {
	panic("implement me")
}

type ISystemConfigParser interface {
	ParseSystem(data []byte)
}

type jsonSystemConfigParser struct{}

func (j jsonSystemConfigParser) ParseSystem(data []byte) {
	panic("implement me")
}

type IConfigParserFactory interface {
	CreateRuleParser() IRuleConfigParser
	CreateSystemParser() ISystemConfigParser
}

type jsonConfigParserFactory struct{}

func (j jsonConfigParserFactory) CreateRuleParser() IRuleConfigParser {
	return jsonRuleConfigParser{}
}

func (j jsonConfigParserFactory) CreateSystemParser() ISystemConfigParser {
	return jsonSystemConfigParser{}
}
