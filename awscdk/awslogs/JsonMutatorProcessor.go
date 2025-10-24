package awslogs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Processor for JSON mutation operations.
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
type JsonMutatorProcessor interface {
	IProcessor
	// The type of JSON mutation operation.
	Type() JsonMutatorType
	SetType(val JsonMutatorType)
}

// The jsii proxy struct for JsonMutatorProcessor
type jsiiProxy_JsonMutatorProcessor struct {
	jsiiProxy_IProcessor
}

func (j *jsiiProxy_JsonMutatorProcessor) Type() JsonMutatorType {
	var returns JsonMutatorType
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Creates a new JSON mutator processor.
func NewJsonMutatorProcessor(props *JsonMutatorProps) JsonMutatorProcessor {
	_init_.Initialize()

	if err := validateNewJsonMutatorProcessorParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_JsonMutatorProcessor{}

	_jsii_.Create(
		"aws-cdk-lib.aws_logs.JsonMutatorProcessor",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Creates a new JSON mutator processor.
func NewJsonMutatorProcessor_Override(j JsonMutatorProcessor, props *JsonMutatorProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_logs.JsonMutatorProcessor",
		[]interface{}{props},
		j,
	)
}

func (j *jsiiProxy_JsonMutatorProcessor)SetType(val JsonMutatorType) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

