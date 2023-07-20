package awsappconfig


// Amazon CloudWatch alarms to monitor during the deployment process.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   monitorsProperty := &MonitorsProperty{
//   	AlarmArn: jsii.String("alarmArn"),
//   	AlarmRoleArn: jsii.String("alarmRoleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-environment-monitors.html
//
type CfnEnvironment_MonitorsProperty struct {
	// Amazon Resource Name (ARN) of the Amazon CloudWatch alarm.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-environment-monitors.html#cfn-appconfig-environment-monitors-alarmarn
	//
	AlarmArn *string `field:"optional" json:"alarmArn" yaml:"alarmArn"`
	// ARN of an AWS Identity and Access Management (IAM) role for AWS AppConfig to monitor `AlarmArn` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-environment-monitors.html#cfn-appconfig-environment-monitors-alarmrolearn
	//
	AlarmRoleArn *string `field:"optional" json:"alarmRoleArn" yaml:"alarmRoleArn"`
}

