package awssns


// The `LoggingConfig` property type specifies the `Delivery` status logging configuration for an [`AWS::SNS::Topic`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-topic.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loggingConfigProperty := &LoggingConfigProperty{
//   	Protocol: jsii.String("protocol"),
//
//   	// the properties below are optional
//   	FailureFeedbackRoleArn: jsii.String("failureFeedbackRoleArn"),
//   	SuccessFeedbackRoleArn: jsii.String("successFeedbackRoleArn"),
//   	SuccessFeedbackSampleRate: jsii.String("successFeedbackSampleRate"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-topic-loggingconfig.html
//
type CfnTopic_LoggingConfigProperty struct {
	// Indicates one of the supported protocols for the Amazon SNS topic.
	//
	// > At least one of the other three `LoggingConfig` properties is recommend along with `Protocol` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-topic-loggingconfig.html#cfn-sns-topic-loggingconfig-protocol
	//
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// The IAM role ARN to be used when logging failed message deliveries in Amazon CloudWatch.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-topic-loggingconfig.html#cfn-sns-topic-loggingconfig-failurefeedbackrolearn
	//
	FailureFeedbackRoleArn *string `field:"optional" json:"failureFeedbackRoleArn" yaml:"failureFeedbackRoleArn"`
	// The IAM role ARN to be used when logging successful message deliveries in Amazon CloudWatch.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-topic-loggingconfig.html#cfn-sns-topic-loggingconfig-successfeedbackrolearn
	//
	SuccessFeedbackRoleArn *string `field:"optional" json:"successFeedbackRoleArn" yaml:"successFeedbackRoleArn"`
	// The percentage of successful message deliveries to be logged in Amazon CloudWatch.
	//
	// Valid percentage values range from 0 to 100.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-topic-loggingconfig.html#cfn-sns-topic-loggingconfig-successfeedbacksamplerate
	//
	SuccessFeedbackSampleRate *string `field:"optional" json:"successFeedbackSampleRate" yaml:"successFeedbackSampleRate"`
}

