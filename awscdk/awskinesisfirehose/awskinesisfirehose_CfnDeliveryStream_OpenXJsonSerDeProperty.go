package awskinesisfirehose


// The OpenX SerDe.
//
// Used by Kinesis Data Firehose for deserializing data, which means converting it from the JSON format in preparation for serializing it to the Parquet or ORC format. This is one of two deserializers you can choose, depending on which one offers the functionality you need. The other option is the native Hive / HCatalog JsonSerDe.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   openXJsonSerDeProperty := &openXJsonSerDeProperty{
//   	caseInsensitive: jsii.Boolean(false),
//   	columnToJsonKeyMappings: map[string]*string{
//   		"columnToJsonKeyMappingsKey": jsii.String("columnToJsonKeyMappings"),
//   	},
//   	convertDotsInJsonKeysToUnderscores: jsii.Boolean(false),
//   }
//
type CfnDeliveryStream_OpenXJsonSerDeProperty struct {
	// When set to `true` , which is the default, Kinesis Data Firehose converts JSON keys to lowercase before deserializing them.
	CaseInsensitive interface{} `field:"optional" json:"caseInsensitive" yaml:"caseInsensitive"`
	// Maps column names to JSON keys that aren't identical to the column names.
	//
	// This is useful when the JSON contains keys that are Hive keywords. For example, `timestamp` is a Hive keyword. If you have a JSON key named `timestamp` , set this parameter to `{"ts": "timestamp"}` to map this key to a column named `ts` .
	ColumnToJsonKeyMappings interface{} `field:"optional" json:"columnToJsonKeyMappings" yaml:"columnToJsonKeyMappings"`
	// When set to `true` , specifies that the names of the keys include dots and that you want Kinesis Data Firehose to replace them with underscores.
	//
	// This is useful because Apache Hive does not allow dots in column names. For example, if the JSON contains a key whose name is "a.b", you can define the column name to be "a_b" when using this option.
	//
	// The default is `false` .
	ConvertDotsInJsonKeysToUnderscores interface{} `field:"optional" json:"convertDotsInJsonKeysToUnderscores" yaml:"convertDotsInJsonKeysToUnderscores"`
}

