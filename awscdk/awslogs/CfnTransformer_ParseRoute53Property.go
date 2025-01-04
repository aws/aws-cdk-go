package awslogs


// Use this processor to parse Route 53 vended logs, extract fields, and and convert them into a JSON format.
//
// This processor always processes the entire log event message. For more information about this processor including examples, see [parseRoute53](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation.html#CloudWatch-Logs-Transformation-parseRoute53) .
//
// > If you use this processor, it must be the first processor in your transformer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   parseRoute53Property := &ParseRoute53Property{
//   	Source: jsii.String("source"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-parseroute53.html
//
type CfnTransformer_ParseRoute53Property struct {
	// Omit this parameter and the whole log message will be processed by this processor.
	//
	// No other value than `@message` is allowed for `source` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-parseroute53.html#cfn-logs-transformer-parseroute53-source
	//
	Source *string `field:"optional" json:"source" yaml:"source"`
}

