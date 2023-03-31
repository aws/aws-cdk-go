package awskinesisfirehose


// The `KinesisStreamSourceConfiguration` property type specifies the stream and role Amazon Resource Names (ARNs) for a Kinesis stream used as the source for a delivery stream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kinesisStreamSourceConfigurationProperty := &kinesisStreamSourceConfigurationProperty{
//   	kinesisStreamArn: jsii.String("kinesisStreamArn"),
//   	roleArn: jsii.String("roleArn"),
//   }
//
type CfnDeliveryStream_KinesisStreamSourceConfigurationProperty struct {
	// The ARN of the source Kinesis data stream.
	KinesisStreamArn *string `field:"required" json:"kinesisStreamArn" yaml:"kinesisStreamArn"`
	// The ARN of the role that provides access to the source Kinesis data stream.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
}

