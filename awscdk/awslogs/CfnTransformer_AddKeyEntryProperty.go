package awslogs


// This object defines one key that will be added with the [addKeys](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Processors.html#CloudWatch-Logs-Transformation-addKey) processor.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   addKeyEntryProperty := &AddKeyEntryProperty{
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//
//   	// the properties below are optional
//   	OverwriteIfExists: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-addkeyentry.html
//
type CfnTransformer_AddKeyEntryProperty struct {
	// The key of the new entry to be added to the log event.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-addkeyentry.html#cfn-logs-transformer-addkeyentry-key
	//
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value of the new entry to be added to the log event.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-addkeyentry.html#cfn-logs-transformer-addkeyentry-value
	//
	Value *string `field:"required" json:"value" yaml:"value"`
	// Specifies whether to overwrite the value if the key already exists in the log event.
	//
	// If you omit this, the default is `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-addkeyentry.html#cfn-logs-transformer-addkeyentry-overwriteifexists
	//
	OverwriteIfExists interface{} `field:"optional" json:"overwriteIfExists" yaml:"overwriteIfExists"`
}

