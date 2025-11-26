package previewawsiotmixins


// Describes an action that updates a CloudWatch log.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cloudwatchLogsActionProperty := &CloudwatchLogsActionProperty{
//   	BatchMode: jsii.Boolean(false),
//   	LogGroupName: jsii.String("logGroupName"),
//   	RoleArn: jsii.String("roleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-cloudwatchlogsaction.html
//
type CfnTopicRulePropsMixin_CloudwatchLogsActionProperty struct {
	// Indicates whether batches of log records will be extracted and uploaded into CloudWatch.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-cloudwatchlogsaction.html#cfn-iot-topicrule-cloudwatchlogsaction-batchmode
	//
	BatchMode interface{} `field:"optional" json:"batchMode" yaml:"batchMode"`
	// The CloudWatch log name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-cloudwatchlogsaction.html#cfn-iot-topicrule-cloudwatchlogsaction-loggroupname
	//
	LogGroupName *string `field:"optional" json:"logGroupName" yaml:"logGroupName"`
	// The IAM role that allows access to the CloudWatch log.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-cloudwatchlogsaction.html#cfn-iot-topicrule-cloudwatchlogsaction-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

