package awsiot


// Describes an action that updates a CloudWatch log.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudwatchLogsActionProperty := &CloudwatchLogsActionProperty{
//   	LogGroupName: jsii.String("logGroupName"),
//   	RoleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	BatchMode: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-cloudwatchlogsaction.html
//
type CfnTopicRule_CloudwatchLogsActionProperty struct {
	// The CloudWatch log name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-cloudwatchlogsaction.html#cfn-iot-topicrule-cloudwatchlogsaction-loggroupname
	//
	LogGroupName *string `field:"required" json:"logGroupName" yaml:"logGroupName"`
	// The IAM role that allows access to the CloudWatch log.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-cloudwatchlogsaction.html#cfn-iot-topicrule-cloudwatchlogsaction-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Indicates whether batches of log records will be extracted and uploaded into CloudWatch.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-cloudwatchlogsaction.html#cfn-iot-topicrule-cloudwatchlogsaction-batchmode
	//
	BatchMode interface{} `field:"optional" json:"batchMode" yaml:"batchMode"`
}

