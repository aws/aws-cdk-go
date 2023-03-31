package awssagemaker


// Configures Amazon SNS notifications of available or expiring work items for work teams.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   notificationConfigurationProperty := &notificationConfigurationProperty{
//   	notificationTopicArn: jsii.String("notificationTopicArn"),
//   }
//
type CfnWorkteam_NotificationConfigurationProperty struct {
	// The ARN for the Amazon SNS topic to which notifications should be published.
	NotificationTopicArn *string `field:"required" json:"notificationTopicArn" yaml:"notificationTopicArn"`
}

