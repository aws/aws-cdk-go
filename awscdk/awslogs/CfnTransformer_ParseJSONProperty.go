package awslogs


// This processor parses log events that are in JSON format.
//
// It can extract JSON key-value pairs and place them under a destination that you specify.
//
// Additionally, because you must have at least one parse-type processor in a transformer, you can use `ParseJSON` as that processor for JSON-format logs, so that you can also apply other processors, such as mutate processors, to these logs.
//
// For more information about this processor including examples, see [parseJSON](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Processors.html#CloudWatch-Logs-Transformation-parseJSON) in the *CloudWatch Logs User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   parseJSONProperty := &ParseJSONProperty{
//   	Destination: jsii.String("destination"),
//   	Source: jsii.String("source"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-parsejson.html
//
type CfnTransformer_ParseJSONProperty struct {
	// The location to put the parsed key value pair into.
	//
	// If you omit this parameter, it is placed under the root node.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-parsejson.html#cfn-logs-transformer-parsejson-destination
	//
	Destination *string `field:"optional" json:"destination" yaml:"destination"`
	// Path to the field in the log event that will be parsed.
	//
	// Use dot notation to access child fields. For example, `store.book`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-parsejson.html#cfn-logs-transformer-parsejson-source
	//
	Source *string `field:"optional" json:"source" yaml:"source"`
}

