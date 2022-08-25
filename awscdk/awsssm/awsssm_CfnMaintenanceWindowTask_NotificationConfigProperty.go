package awsssm


// The `NotificationConfig` property type specifies configurations for sending notifications for a maintenance window task in AWS Systems Manager .
//
// `NotificationConfig` is a property of the [MaintenanceWindowRunCommandParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-maintenancewindowruncommandparameters.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   notificationConfigProperty := &notificationConfigProperty{
//   	notificationArn: jsii.String("notificationArn"),
//
//   	// the properties below are optional
//   	notificationEvents: []*string{
//   		jsii.String("notificationEvents"),
//   	},
//   	notificationType: jsii.String("notificationType"),
//   }
//
type CfnMaintenanceWindowTask_NotificationConfigProperty struct {
	// An Amazon Resource Name (ARN) for an Amazon Simple Notification Service (Amazon SNS) topic.
	//
	// Run Command pushes notifications about command status changes to this topic.
	NotificationArn *string `field:"required" json:"notificationArn" yaml:"notificationArn"`
	// The different events that you can receive notifications for.
	//
	// These events include the following: `All` (events), `InProgress` , `Success` , `TimedOut` , `Cancelled` , `Failed` . To learn more about these events, see [Configuring Amazon SNS Notifications for AWS Systems Manager](https://docs.aws.amazon.com/systems-manager/latest/userguide/monitoring-sns-notifications.html) in the *AWS Systems Manager User Guide* .
	NotificationEvents *[]*string `field:"optional" json:"notificationEvents" yaml:"notificationEvents"`
	// The notification type.
	//
	// - `Command` : Receive notification when the status of a command changes.
	// - `Invocation` : For commands sent to multiple instances, receive notification on a per-instance basis when the status of a command changes.
	NotificationType *string `field:"optional" json:"notificationType" yaml:"notificationType"`
}

