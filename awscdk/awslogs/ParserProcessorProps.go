package awslogs


// Properties for creating configurable parser processors.
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
//   		Entries: []addKeyEntryProperty{
//   			&addKeyEntryProperty{
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
//   	TransformerConfig: []iProcessor{
//   		jsonParser,
//   		addKeysProcessor,
//   	},
//   })
//
type ParserProcessorProps struct {
	// The type of parser processor.
	Type ParserProcessorType `field:"required" json:"type" yaml:"type"`
	// Options for CSV parser.
	//
	// Required when type is CSV.
	// Default: - No CSV parser is created if props not set.
	//
	CsvOptions *CsvProperty `field:"optional" json:"csvOptions" yaml:"csvOptions"`
	// Options for Grok parser.
	//
	// Required when type is GROK.
	// Default: - No Grok parser is created if props not set.
	//
	GrokOptions *GrokProperty `field:"optional" json:"grokOptions" yaml:"grokOptions"`
	// Options for JSON parser.
	//
	// Required when type is JSON.
	// Default: - No JSON parser is created if props not set.
	//
	JsonOptions *ParseJSONProperty `field:"optional" json:"jsonOptions" yaml:"jsonOptions"`
	// Options for key-value parser.
	//
	// Required when type is KEY_VALUE.
	// Default: - No key-value parser is created if props not set.
	//
	KeyValueOptions *ParseKeyValueProperty `field:"optional" json:"keyValueOptions" yaml:"keyValueOptions"`
	// Options for ParseToOCSF parser.
	//
	// Required when type is set to OCSF.
	// Default: - no OCSF parser is created.
	//
	ParseToOCSFOptions *ParseToOCSFProperty `field:"optional" json:"parseToOCSFOptions" yaml:"parseToOCSFOptions"`
}

