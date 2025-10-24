package awsbedrockagentcore


// The memory invocation configuration input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   invocationConfigurationInputProperty := &InvocationConfigurationInputProperty{
//   	PayloadDeliveryBucketName: jsii.String("payloadDeliveryBucketName"),
//   	TopicArn: jsii.String("topicArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-invocationconfigurationinput.html
//
type CfnMemory_InvocationConfigurationInputProperty struct {
	// The message invocation configuration information for the bucket name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-invocationconfigurationinput.html#cfn-bedrockagentcore-memory-invocationconfigurationinput-payloaddeliverybucketname
	//
	PayloadDeliveryBucketName *string `field:"optional" json:"payloadDeliveryBucketName" yaml:"payloadDeliveryBucketName"`
	// The memory trigger condition topic Amazon Resource Name (ARN).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-invocationconfigurationinput.html#cfn-bedrockagentcore-memory-invocationconfigurationinput-topicarn
	//
	TopicArn *string `field:"optional" json:"topicArn" yaml:"topicArn"`
}

