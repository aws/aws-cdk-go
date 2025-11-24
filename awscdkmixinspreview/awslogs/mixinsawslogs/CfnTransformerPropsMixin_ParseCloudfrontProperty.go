package mixinsawslogs


// This processor parses CloudFront vended logs, extract fields, and convert them into JSON format.
//
// Encoded field values are decoded. Values that are integers and doubles are treated as such. For more information about this processor including examples, see [parseCloudfront](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Processors.html#CloudWatch-Logs-Transformation-parseCloudfront)
//
// For more information about CloudFront log format, see [Configure and use standard logs (access logs)](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/AccessLogs.html) .
//
// If you use this processor, it must be the first processor in your transformer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   parseCloudfrontProperty := &ParseCloudfrontProperty{
//   	Source: jsii.String("source"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-parsecloudfront.html
//
type CfnTransformerPropsMixin_ParseCloudfrontProperty struct {
	// Omit this parameter and the whole log message will be processed by this processor.
	//
	// No other value than `@message` is allowed for `source` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-parsecloudfront.html#cfn-logs-transformer-parsecloudfront-source
	//
	Source *string `field:"optional" json:"source" yaml:"source"`
}

