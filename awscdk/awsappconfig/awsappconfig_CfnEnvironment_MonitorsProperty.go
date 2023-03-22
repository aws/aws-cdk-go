package awsappconfig


// Amazon CloudWatch alarms to monitor during the deployment process.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   monitorsProperty := &monitorsProperty{
//   	alarmArn: jsii.String("alarmArn"),
//   	alarmRoleArn: jsii.String("alarmRoleArn"),
//   }
//
type CfnEnvironment_MonitorsProperty struct {
	// Amazon Resource Name (ARN) of the Amazon CloudWatch alarm.
	AlarmArn *string `field:"optional" json:"alarmArn" yaml:"alarmArn"`
	// ARN of an AWS Identity and Access Management (IAM) role for AWS AppConfig to monitor `AlarmArn` .
	AlarmRoleArn *string `field:"optional" json:"alarmRoleArn" yaml:"alarmRoleArn"`
}

