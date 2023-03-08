package awsiotsitewise


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   alarmsProperty := &alarmsProperty{
//   	alarmRoleArn: jsii.String("alarmRoleArn"),
//   	notificationLambdaArn: jsii.String("notificationLambdaArn"),
//   }
//
type CfnPortal_AlarmsProperty struct {
	// `CfnPortal.AlarmsProperty.AlarmRoleArn`.
	AlarmRoleArn *string `field:"optional" json:"alarmRoleArn" yaml:"alarmRoleArn"`
	// `CfnPortal.AlarmsProperty.NotificationLambdaArn`.
	NotificationLambdaArn *string `field:"optional" json:"notificationLambdaArn" yaml:"notificationLambdaArn"`
}

