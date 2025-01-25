package awslogs


// This object defines one value to be copied with the [copyValue](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Processors.html#CloudWatch-Logs-Transformation-copyValue) processor.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   copyValueEntryProperty := &CopyValueEntryProperty{
//   	Source: jsii.String("source"),
//   	Target: jsii.String("target"),
//
//   	// the properties below are optional
//   	OverwriteIfExists: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-copyvalueentry.html
//
type CfnTransformer_CopyValueEntryProperty struct {
	// The key to copy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-copyvalueentry.html#cfn-logs-transformer-copyvalueentry-source
	//
	Source *string `field:"required" json:"source" yaml:"source"`
	// The key of the field to copy the value to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-copyvalueentry.html#cfn-logs-transformer-copyvalueentry-target
	//
	Target *string `field:"required" json:"target" yaml:"target"`
	// Specifies whether to overwrite the value if the destination key already exists.
	//
	// If you omit this, the default is `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-copyvalueentry.html#cfn-logs-transformer-copyvalueentry-overwriteifexists
	//
	OverwriteIfExists interface{} `field:"optional" json:"overwriteIfExists" yaml:"overwriteIfExists"`
}

