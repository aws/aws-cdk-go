package awslogs


// Properties for creating JSON mutator processors.
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
type JsonMutatorProps struct {
	// The type of JSON mutation operation.
	Type JsonMutatorType `field:"required" json:"type" yaml:"type"`
	// Options for adding keys.
	//
	// Required when type is ADD_KEYS.
	// Default: - No adding keys processor is created if props not set.
	//
	AddKeysOptions *AddKeysProperty `field:"optional" json:"addKeysOptions" yaml:"addKeysOptions"`
	// Options for copying values.
	//
	// Required when type is COPY_VALUE.
	// Default: - No copy value processor is created if props not set.
	//
	CopyValueOptions *CopyValueProperty `field:"optional" json:"copyValueOptions" yaml:"copyValueOptions"`
	// Keys to delete.
	//
	// Required when type is DELETE_KEYS.
	// Default: - No delete key processor is created if props not set.
	//
	DeleteKeysOptions *ProcessorDeleteKeysProperty `field:"optional" json:"deleteKeysOptions" yaml:"deleteKeysOptions"`
	// Options for converting lists to maps.
	//
	// Required when type is LIST_TO_MAP.
	// Default: - No list-to-map processor is created if props not set.
	//
	ListToMapOptions *ListToMapProperty `field:"optional" json:"listToMapOptions" yaml:"listToMapOptions"`
	// Options for moving keys.
	//
	// Required when type is MOVE_KEYS.
	// Default: - No move key processor is created if props not set.
	//
	MoveKeysOptions *MoveKeysProperty `field:"optional" json:"moveKeysOptions" yaml:"moveKeysOptions"`
	// Options for renaming keys.
	//
	// Required when type is RENAME_KEYS.
	// Default: - No rename key processor is created if props not set.
	//
	RenameKeysOptions *RenameKeysProperty `field:"optional" json:"renameKeysOptions" yaml:"renameKeysOptions"`
}

