package previewawslogsmixins


// This object defines one log field that will be split with the [splitString](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Processors.html#CloudWatch-Logs-Transformation-splitString) processor.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   splitStringEntryProperty := &SplitStringEntryProperty{
//   	Delimiter: jsii.String("delimiter"),
//   	Source: jsii.String("source"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-splitstringentry.html
//
type CfnTransformerPropsMixin_SplitStringEntryProperty struct {
	// The separator characters to split the string entry on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-splitstringentry.html#cfn-logs-transformer-splitstringentry-delimiter
	//
	Delimiter *string `field:"optional" json:"delimiter" yaml:"delimiter"`
	// The key of the field to split.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-splitstringentry.html#cfn-logs-transformer-splitstringentry-source
	//
	Source *string `field:"optional" json:"source" yaml:"source"`
}

