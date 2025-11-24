package mixinsawslogs


// This processor parses a specified field in the original log event into key-value pairs.
//
// For more information about this processor including examples, see [parseKeyValue](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation.html#CloudWatch-Logs-Transformation-parseKeyValue) in the *CloudWatch Logs User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   parseKeyValueProperty := &ParseKeyValueProperty{
//   	Destination: jsii.String("destination"),
//   	FieldDelimiter: jsii.String("fieldDelimiter"),
//   	KeyPrefix: jsii.String("keyPrefix"),
//   	KeyValueDelimiter: jsii.String("keyValueDelimiter"),
//   	NonMatchValue: jsii.String("nonMatchValue"),
//   	OverwriteIfExists: jsii.Boolean(false),
//   	Source: jsii.String("source"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-parsekeyvalue.html
//
type CfnTransformerPropsMixin_ParseKeyValueProperty struct {
	// The destination field to put the extracted key-value pairs into.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-parsekeyvalue.html#cfn-logs-transformer-parsekeyvalue-destination
	//
	Destination *string `field:"optional" json:"destination" yaml:"destination"`
	// The field delimiter string that is used between key-value pairs in the original log events.
	//
	// If you omit this, the ampersand `&` character is used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-parsekeyvalue.html#cfn-logs-transformer-parsekeyvalue-fielddelimiter
	//
	FieldDelimiter *string `field:"optional" json:"fieldDelimiter" yaml:"fieldDelimiter"`
	// If you want to add a prefix to all transformed keys, specify it here.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-parsekeyvalue.html#cfn-logs-transformer-parsekeyvalue-keyprefix
	//
	KeyPrefix *string `field:"optional" json:"keyPrefix" yaml:"keyPrefix"`
	// The delimiter string to use between the key and value in each pair in the transformed log event.
	//
	// If you omit this, the equal `=` character is used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-parsekeyvalue.html#cfn-logs-transformer-parsekeyvalue-keyvaluedelimiter
	//
	KeyValueDelimiter *string `field:"optional" json:"keyValueDelimiter" yaml:"keyValueDelimiter"`
	// A value to insert into the value field in the result, when a key-value pair is not successfully split.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-parsekeyvalue.html#cfn-logs-transformer-parsekeyvalue-nonmatchvalue
	//
	NonMatchValue *string `field:"optional" json:"nonMatchValue" yaml:"nonMatchValue"`
	// Specifies whether to overwrite the value if the destination key already exists.
	//
	// If you omit this, the default is `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-parsekeyvalue.html#cfn-logs-transformer-parsekeyvalue-overwriteifexists
	//
	OverwriteIfExists interface{} `field:"optional" json:"overwriteIfExists" yaml:"overwriteIfExists"`
	// Path to the field in the log event that will be parsed.
	//
	// Use dot notation to access child fields. For example, `store.book`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-parsekeyvalue.html#cfn-logs-transformer-parsekeyvalue-source
	//
	Source *string `field:"optional" json:"source" yaml:"source"`
}

