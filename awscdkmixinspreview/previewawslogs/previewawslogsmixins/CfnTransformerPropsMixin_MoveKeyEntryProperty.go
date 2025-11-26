package previewawslogsmixins


// This object defines one key that will be moved with the [moveKey](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation.html#CloudWatch-Logs-Transformation-moveKey) processor.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   moveKeyEntryProperty := &MoveKeyEntryProperty{
//   	OverwriteIfExists: jsii.Boolean(false),
//   	Source: jsii.String("source"),
//   	Target: jsii.String("target"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-movekeyentry.html
//
type CfnTransformerPropsMixin_MoveKeyEntryProperty struct {
	// Specifies whether to overwrite the value if the destination key already exists.
	//
	// If you omit this, the default is `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-movekeyentry.html#cfn-logs-transformer-movekeyentry-overwriteifexists
	//
	OverwriteIfExists interface{} `field:"optional" json:"overwriteIfExists" yaml:"overwriteIfExists"`
	// The key to move.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-movekeyentry.html#cfn-logs-transformer-movekeyentry-source
	//
	Source *string `field:"optional" json:"source" yaml:"source"`
	// The key to move to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-movekeyentry.html#cfn-logs-transformer-movekeyentry-target
	//
	Target *string `field:"optional" json:"target" yaml:"target"`
}

