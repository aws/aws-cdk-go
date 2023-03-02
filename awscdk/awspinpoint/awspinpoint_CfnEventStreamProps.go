package awspinpoint


// Properties for defining a `CfnEventStream`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEventStreamProps := &cfnEventStreamProps{
//   	applicationId: jsii.String("applicationId"),
//   	destinationStreamArn: jsii.String("destinationStreamArn"),
//   	roleArn: jsii.String("roleArn"),
//   }
//
type CfnEventStreamProps struct {
	// The unique identifier for the Amazon Pinpoint application that you want to export data from.
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// The Amazon Resource Name (ARN) of the Amazon Kinesis data stream or Amazon Kinesis Data Firehose delivery stream that you want to publish event data to.
	//
	// For a Kinesis data stream, the ARN format is: `arn:aws:kinesis: region : account-id :stream/ stream_name`
	//
	// For a Kinesis Data Firehose delivery stream, the ARN format is: `arn:aws:firehose: region : account-id :deliverystream/ stream_name`.
	DestinationStreamArn *string `field:"required" json:"destinationStreamArn" yaml:"destinationStreamArn"`
	// The AWS Identity and Access Management (IAM) role that authorizes Amazon Pinpoint to publish event data to the stream in your AWS account.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
}

