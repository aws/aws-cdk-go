package awslogs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Parser processor for common data formats.
//
// Example:
//   // Create a log group
//   logGroup := logs.NewLogGroup(this, jsii.String("MyLogGroup"))
//
//   // Create a JSON parser processor
//   jsonParser := logs.NewParserProcessor(&ParserProcessorProps{
//   	Type: logs.ParserProcessorType_JSON,
//   })
//
//   // Create a processor to add keys
//   addKeysProcessor := logs.NewJsonMutatorProcessor(&JsonMutatorProps{
//   	Type: logs.JsonMutatorType_ADD_KEYS,
//   	AddKeysOptions: &AddKeysProperty{
//   		Entries: []AddKeyEntryProperty{
//   			&AddKeyEntryProperty{
//   				Key: jsii.String("metadata.transformed_in"),
//   				Value: jsii.String("CloudWatchLogs"),
//   			},
//   		},
//   	},
//   })
//
//   // Create a transformer with these processors
//   // Create a transformer with these processors
//   logs.NewTransformer(this, jsii.String("Transformer"), &TransformerProps{
//   	TransformerName: jsii.String("MyTransformer"),
//   	LogGroup: logGroup,
//   	TransformerConfig: []IProcessor{
//   		jsonParser,
//   		addKeysProcessor,
//   	},
//   })
//
type ParserProcessor interface {
	IProcessor
	// The type of parser.
	Type() ParserProcessorType
	SetType(val ParserProcessorType)
}

// The jsii proxy struct for ParserProcessor
type jsiiProxy_ParserProcessor struct {
	jsiiProxy_IProcessor
}

func (j *jsiiProxy_ParserProcessor) Type() ParserProcessorType {
	var returns ParserProcessorType
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Creates a new parser processor.
func NewParserProcessor(props *ParserProcessorProps) ParserProcessor {
	_init_.Initialize()

	if err := validateNewParserProcessorParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ParserProcessor{}

	_jsii_.Create(
		"aws-cdk-lib.aws_logs.ParserProcessor",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Creates a new parser processor.
func NewParserProcessor_Override(p ParserProcessor, props *ParserProcessorProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_logs.ParserProcessor",
		[]interface{}{props},
		p,
	)
}

func (j *jsiiProxy_ParserProcessor)SetType(val ParserProcessorType) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

