package awsiotsitewise


// Contains the configuration information of an alarm created in an AWS IoT SiteWise Monitor portal.
//
// You can use the alarm to monitor an asset property and get notified when the asset property value is outside a specified range. For more information, see [Monitoring with alarms](https://docs.aws.amazon.com/iot-sitewise/latest/appguide/monitor-alarms.html) in the *AWS IoT SiteWise Application Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   alarmsProperty := &AlarmsProperty{
//   	AlarmRoleArn: jsii.String("alarmRoleArn"),
//   	NotificationLambdaArn: jsii.String("notificationLambdaArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-portal-alarms.html
//
type CfnPortal_AlarmsProperty struct {
	// The [ARN](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the IAM role that allows the alarm to perform actions and access AWS resources and services, such as AWS IoT Events .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-portal-alarms.html#cfn-iotsitewise-portal-alarms-alarmrolearn
	//
	AlarmRoleArn *string `field:"optional" json:"alarmRoleArn" yaml:"alarmRoleArn"`
	// The [ARN](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the Lambda function that manages alarm notifications. For more information, see [Managing alarm notifications](https://docs.aws.amazon.com/iotevents/latest/developerguide/lambda-support.html) in the *AWS IoT Events Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-portal-alarms.html#cfn-iotsitewise-portal-alarms-notificationlambdaarn
	//
	NotificationLambdaArn *string `field:"optional" json:"notificationLambdaArn" yaml:"notificationLambdaArn"`
}

