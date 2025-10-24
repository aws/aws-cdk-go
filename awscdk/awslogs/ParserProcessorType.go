package awslogs


// Types of configurable parser processors.
//
// Defines the various parser types that can be used to process log events.
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
type ParserProcessorType string

const (
	// Parse log entries as JSON.
	ParserProcessorType_JSON ParserProcessorType = "JSON"
	// Parse log entries as key-value pairs.
	ParserProcessorType_KEY_VALUE ParserProcessorType = "KEY_VALUE"
	// Parse log entries in CSV format.
	ParserProcessorType_CSV ParserProcessorType = "CSV"
	// Parse log entries using Grok patterns.
	ParserProcessorType_GROK ParserProcessorType = "GROK"
	// Parse logs to OCSF format.
	ParserProcessorType_OCSF ParserProcessorType = "OCSF"
)

