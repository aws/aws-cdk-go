package awsappconfig


// Amazon CloudWatch alarms to monitor during the deployment process.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   monitorProperty := &MonitorProperty{
//   	AlarmArn: jsii.String("alarmArn"),
//
//   	// the properties below are optional
//   	AlarmRoleArn: jsii.String("alarmRoleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-environment-monitor.html
//
type CfnEnvironment_MonitorProperty struct {
	// Amazon Resource Name (ARN) of the Amazon CloudWatch alarm.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-environment-monitor.html#cfn-appconfig-environment-monitor-alarmarn
	//
	AlarmArn *string `field:"required" json:"alarmArn" yaml:"alarmArn"`
	// ARN of an AWS Identity and Access Management (IAM) role for AWS AppConfig to monitor `AlarmArn` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-environment-monitor.html#cfn-appconfig-environment-monitor-alarmrolearn
	//
	AlarmRoleArn *string `field:"optional" json:"alarmRoleArn" yaml:"alarmRoleArn"`
}

