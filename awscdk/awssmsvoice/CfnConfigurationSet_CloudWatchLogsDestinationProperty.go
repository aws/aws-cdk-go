package awssmsvoice


// Contains the destination configuration to use when publishing message sending events.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudWatchLogsDestinationProperty := &CloudWatchLogsDestinationProperty{
//   	IamRoleArn: jsii.String("iamRoleArn"),
//   	LogGroupArn: jsii.String("logGroupArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-smsvoice-configurationset-cloudwatchlogsdestination.html
//
type CfnConfigurationSet_CloudWatchLogsDestinationProperty struct {
	// The Amazon Resource Name (ARN) of an AWS Identity and Access Management role that is able to write event data to an Amazon CloudWatch destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-smsvoice-configurationset-cloudwatchlogsdestination.html#cfn-smsvoice-configurationset-cloudwatchlogsdestination-iamrolearn
	//
	IamRoleArn *string `field:"required" json:"iamRoleArn" yaml:"iamRoleArn"`
	// The name of the Amazon CloudWatch log group that you want to record events in.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-smsvoice-configurationset-cloudwatchlogsdestination.html#cfn-smsvoice-configurationset-cloudwatchlogsdestination-loggrouparn
	//
	LogGroupArn *string `field:"required" json:"logGroupArn" yaml:"logGroupArn"`
}

