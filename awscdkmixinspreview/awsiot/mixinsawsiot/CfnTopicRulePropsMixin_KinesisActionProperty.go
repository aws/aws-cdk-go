package mixinsawsiot


// Describes an action to write data to an Amazon Kinesis stream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   kinesisActionProperty := &KinesisActionProperty{
//   	PartitionKey: jsii.String("partitionKey"),
//   	RoleArn: jsii.String("roleArn"),
//   	StreamName: jsii.String("streamName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-kinesisaction.html
//
type CfnTopicRulePropsMixin_KinesisActionProperty struct {
	// The partition key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-kinesisaction.html#cfn-iot-topicrule-kinesisaction-partitionkey
	//
	PartitionKey *string `field:"optional" json:"partitionKey" yaml:"partitionKey"`
	// The ARN of the IAM role that grants access to the Amazon Kinesis stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-kinesisaction.html#cfn-iot-topicrule-kinesisaction-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The name of the Amazon Kinesis stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-kinesisaction.html#cfn-iot-topicrule-kinesisaction-streamname
	//
	StreamName *string `field:"optional" json:"streamName" yaml:"streamName"`
}

