package awsiot


// Describes an action to write data to an Amazon Kinesis stream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kinesisActionProperty := &kinesisActionProperty{
//   	roleArn: jsii.String("roleArn"),
//   	streamName: jsii.String("streamName"),
//
//   	// the properties below are optional
//   	partitionKey: jsii.String("partitionKey"),
//   }
//
type CfnTopicRule_KinesisActionProperty struct {
	// The ARN of the IAM role that grants access to the Amazon Kinesis stream.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The name of the Amazon Kinesis stream.
	StreamName *string `field:"required" json:"streamName" yaml:"streamName"`
	// The partition key.
	PartitionKey *string `field:"optional" json:"partitionKey" yaml:"partitionKey"`
}

