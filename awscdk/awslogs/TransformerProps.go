package awslogs


// The Resource properties for AWS::Logs::Transformer resource.
//
// This
// interface defines all configuration options for the CfnTransformer construct.
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
type TransformerProps struct {
	// Existing log group that you want to associate with this transformer.
	LogGroup ILogGroup `field:"required" json:"logGroup" yaml:"logGroup"`
	// List of processors in a transformer.
	TransformerConfig *[]IProcessor `field:"required" json:"transformerConfig" yaml:"transformerConfig"`
	// Name of the transformer.
	TransformerName *string `field:"required" json:"transformerName" yaml:"transformerName"`
}

