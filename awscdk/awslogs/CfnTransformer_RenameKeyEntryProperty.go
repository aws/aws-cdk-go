package awslogs


// This object defines one key that will be renamed with the [renameKey](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Processors.html#CloudWatch-Logs-Transformation-renameKey) processor.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   renameKeyEntryProperty := &RenameKeyEntryProperty{
//   	Key: jsii.String("key"),
//   	RenameTo: jsii.String("renameTo"),
//
//   	// the properties below are optional
//   	OverwriteIfExists: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-renamekeyentry.html
//
type CfnTransformer_RenameKeyEntryProperty struct {
	// The key to rename.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-renamekeyentry.html#cfn-logs-transformer-renamekeyentry-key
	//
	Key *string `field:"required" json:"key" yaml:"key"`
	// The string to use for the new key name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-renamekeyentry.html#cfn-logs-transformer-renamekeyentry-renameto
	//
	RenameTo *string `field:"required" json:"renameTo" yaml:"renameTo"`
	// Specifies whether to overwrite the existing value if the destination key already exists.
	//
	// The default is `false`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-renamekeyentry.html#cfn-logs-transformer-renamekeyentry-overwriteifexists
	//
	OverwriteIfExists interface{} `field:"optional" json:"overwriteIfExists" yaml:"overwriteIfExists"`
}

