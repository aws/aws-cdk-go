package awsiot


// Describes an action to write data to an Amazon Kinesis stream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kinesisActionProperty := &KinesisActionProperty{
//   	RoleArn: jsii.String("roleArn"),
//   	StreamName: jsii.String("streamName"),
//
//   	// the properties below are optional
//   	PartitionKey: jsii.String("partitionKey"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-kinesisaction.html
//
type CfnTopicRule_KinesisActionProperty struct {
	// The ARN of the IAM role that grants access to the Amazon Kinesis stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-kinesisaction.html#cfn-iot-topicrule-kinesisaction-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The name of the Amazon Kinesis stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-kinesisaction.html#cfn-iot-topicrule-kinesisaction-streamname
	//
	StreamName *string `field:"required" json:"streamName" yaml:"streamName"`
	// The partition key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-kinesisaction.html#cfn-iot-topicrule-kinesisaction-partitionkey
	//
	PartitionKey *string `field:"optional" json:"partitionKey" yaml:"partitionKey"`
}

