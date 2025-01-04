package awslogs


// This processor uses pattern matching to parse and structure unstructured data.
//
// This processor can also extract fields from log messages.
//
// For more information about this processor including examples, see [grok](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation.html#CloudWatch-Logs-Transformation-Grok) in the *CloudWatch Logs User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   grokProperty := &GrokProperty{
//   	Match: jsii.String("match"),
//
//   	// the properties below are optional
//   	Source: jsii.String("source"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-grok.html
//
type CfnTransformer_GrokProperty struct {
	// The grok pattern to match against the log event.
	//
	// For a list of supported grok patterns, see [Supported grok patterns](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation.html#Grok-Patterns) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-grok.html#cfn-logs-transformer-grok-match
	//
	Match *string `field:"required" json:"match" yaml:"match"`
	// The path to the field in the log event that you want to parse.
	//
	// If you omit this value, the whole log message is parsed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-grok.html#cfn-logs-transformer-grok-source
	//
	Source *string `field:"optional" json:"source" yaml:"source"`
}

