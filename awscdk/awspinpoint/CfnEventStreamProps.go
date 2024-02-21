package awspinpoint


// Properties for defining a `CfnEventStream`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEventStreamProps := &CfnEventStreamProps{
//   	ApplicationId: jsii.String("applicationId"),
//   	DestinationStreamArn: jsii.String("destinationStreamArn"),
//   	RoleArn: jsii.String("roleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-eventstream.html
//
type CfnEventStreamProps struct {
	// The unique identifier for the Amazon Pinpoint application that you want to export data from.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-eventstream.html#cfn-pinpoint-eventstream-applicationid
	//
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// The Amazon Resource Name (ARN) of the Amazon Kinesis Data Stream or Amazon Data Firehose delivery stream that you want to publish event data to.
	//
	// For a Kinesis Data Stream, the ARN format is: `arn:aws:kinesis: region : account-id :stream/ stream_name`
	//
	// For a Firehose delivery stream, the ARN format is: `arn:aws:firehose: region : account-id :deliverystream/ stream_name`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-eventstream.html#cfn-pinpoint-eventstream-destinationstreamarn
	//
	DestinationStreamArn *string `field:"required" json:"destinationStreamArn" yaml:"destinationStreamArn"`
	// The AWS Identity and Access Management (IAM) role that authorizes Amazon Pinpoint to publish event data to the stream in your AWS account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-eventstream.html#cfn-pinpoint-eventstream-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
}

