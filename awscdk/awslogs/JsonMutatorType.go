package awslogs


// Types of JSON mutation operations.
//
// Defines operations that can be performed to modify the JSON structure of log events.
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
type JsonMutatorType string

const (
	// Add new keys to the log event.
	JsonMutatorType_ADD_KEYS JsonMutatorType = "ADD_KEYS"
	// Delete keys from the log event.
	JsonMutatorType_DELETE_KEYS JsonMutatorType = "DELETE_KEYS"
	// Move keys to different locations.
	JsonMutatorType_MOVE_KEYS JsonMutatorType = "MOVE_KEYS"
	// Rename keys in the log event.
	JsonMutatorType_RENAME_KEYS JsonMutatorType = "RENAME_KEYS"
	// Copy values between keys.
	JsonMutatorType_COPY_VALUE JsonMutatorType = "COPY_VALUE"
	// Convert a list to a map.
	JsonMutatorType_LIST_TO_MAP JsonMutatorType = "LIST_TO_MAP"
)

