package awslogs


// This processor adds new key-value pairs to the log event.
//
// For more information about this processor including examples, see addKeys in the CloudWatch Logs User Guide.
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
type AddKeysProperty struct {
	// An array of objects, where each object contains information about one key to add to the log event.
	Entries *[]*AddKeyEntryProperty `field:"required" json:"entries" yaml:"entries"`
}

