package awslogs


// This processor parses a specified field in the original log event into key-value pairs.
//
// For more information about this processor including examples, see parseKeyValue in the CloudWatch Logs User Guide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   parseKeyValueProperty := &ParseKeyValueProperty{
//   	Destination: jsii.String("destination"),
//   	FieldDelimiter: awscdk.Aws_logs.KeyValuePairDelimiter_AMPERSAND,
//   	KeyPrefix: jsii.String("keyPrefix"),
//   	KeyValueDelimiter: awscdk.*Aws_logs.KeyValueDelimiter_EQUAL,
//   	NonMatchValue: jsii.String("nonMatchValue"),
//   	OverwriteIfExists: jsii.Boolean(false),
//   	Source: jsii.String("source"),
//   }
//
type ParseKeyValueProperty struct {
	// The destination field to put the extracted key-value pairs into.
	// Default: - Places at the root of the JSON input.
	//
	Destination *string `field:"optional" json:"destination" yaml:"destination"`
	// The field delimiter string that is used between key-value pairs in the original log events.
	// Default: KeyValuePairDelimiter.AMPERSAND
	//
	FieldDelimiter KeyValuePairDelimiter `field:"optional" json:"fieldDelimiter" yaml:"fieldDelimiter"`
	// If you want to add a prefix to all transformed keys, specify it here.
	// Default: - No prefix is added to the keys.
	//
	KeyPrefix *string `field:"optional" json:"keyPrefix" yaml:"keyPrefix"`
	// The delimiter string to use between the key and value in each pair in the transformed log event.
	// Default: KeyValueDelimiter.EQUAL
	//
	KeyValueDelimiter KeyValueDelimiter `field:"optional" json:"keyValueDelimiter" yaml:"keyValueDelimiter"`
	// A value to insert into the value field in the result, when a key-value pair is not successfully split.
	// Default: - No values is inserted when split is not successful.
	//
	NonMatchValue *string `field:"optional" json:"nonMatchValue" yaml:"nonMatchValue"`
	// Specifies whether to overwrite the value if the destination key already exists.
	// Default: false.
	//
	OverwriteIfExists *bool `field:"optional" json:"overwriteIfExists" yaml:"overwriteIfExists"`
	// Path to the field in the log event that will be parsed.
	//
	// Use dot notation to access child fields.
	// Default: '@message'.
	//
	Source *string `field:"optional" json:"source" yaml:"source"`
}

