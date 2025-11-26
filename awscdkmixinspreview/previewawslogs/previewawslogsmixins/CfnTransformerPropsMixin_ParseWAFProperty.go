package previewawslogsmixins


// Use this processor to parse AWS WAF vended logs, extract fields, and and convert them into a JSON format.
//
// This processor always processes the entire log event message. For more information about this processor including examples, see [parseWAF](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Processors.html#CloudWatch-Logs-Transformation-parsePostGres) .
//
// For more information about AWS WAF log format, see [Log examples for web ACL traffic](https://docs.aws.amazon.com/waf/latest/developerguide/logging-examples.html) .
//
// > If you use this processor, it must be the first processor in your transformer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   parseWAFProperty := &ParseWAFProperty{
//   	Source: jsii.String("source"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-parsewaf.html
//
type CfnTransformerPropsMixin_ParseWAFProperty struct {
	// Omit this parameter and the whole log message will be processed by this processor.
	//
	// No other value than `@message` is allowed for `source` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-parsewaf.html#cfn-logs-transformer-parsewaf-source
	//
	Source *string `field:"optional" json:"source" yaml:"source"`
}

