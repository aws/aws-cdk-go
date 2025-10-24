package awskinesisfirehose


// Props for OpenX JSON input format for data record format conversion.
//
// Example:
//   inputFormat := firehose.NewOpenXJsonInputFormat(&OpenXJsonInputFormatProps{
//   	LowercaseColumnNames: jsii.Boolean(false),
//   	ColumnToJsonKeyMappings: map[string]*string{
//   		"ts": jsii.String("timestamp"),
//   	},
//   	ConvertDotsInJsonKeysToUnderscores: jsii.Boolean(true),
//   })
//
type OpenXJsonInputFormatProps struct {
	// Maps column names to JSON keys that aren't identical to the column names.
	//
	// This is useful when the JSON contains keys that are Hive keywords.
	// For example, `timestamp` is a Hive keyword. If you have a JSON key named `timestamp`, set this parameter to `{"ts": "timestamp"}` to map this key to a column named `ts`
	// Default: JSON keys are not renamed.
	//
	ColumnToJsonKeyMappings *map[string]*string `field:"optional" json:"columnToJsonKeyMappings" yaml:"columnToJsonKeyMappings"`
	// When set to `true`, specifies that the names of the keys include dots and that you want Firehose to replace them with underscores.
	//
	// This is useful because Apache Hive does not allow dots in column names.
	// For example, if the JSON contains a key whose name is "a.b", you can define the column name to be "a_b" when using this option.
	// Default: `false`.
	//
	ConvertDotsInJsonKeysToUnderscores *bool `field:"optional" json:"convertDotsInJsonKeysToUnderscores" yaml:"convertDotsInJsonKeysToUnderscores"`
	// Whether the JSON keys should be lowercased when written as column names.
	// Default: `true`.
	//
	LowercaseColumnNames *bool `field:"optional" json:"lowercaseColumnNames" yaml:"lowercaseColumnNames"`
}

