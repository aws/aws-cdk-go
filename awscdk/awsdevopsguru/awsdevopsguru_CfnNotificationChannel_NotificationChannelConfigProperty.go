package awsdevopsguru


// Information about notification channels you have configured with DevOps Guru.
//
// The one supported notification channel is Amazon Simple Notification Service (Amazon SNS).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   notificationChannelConfigProperty := &notificationChannelConfigProperty{
//   	filters: &notificationFilterConfigProperty{
//   		messageTypes: []*string{
//   			jsii.String("messageTypes"),
//   		},
//   		severities: []*string{
//   			jsii.String("severities"),
//   		},
//   	},
//   	sns: &snsChannelConfigProperty{
//   		topicArn: jsii.String("topicArn"),
//   	},
//   }
//
type CfnNotificationChannel_NotificationChannelConfigProperty struct {
	// `CfnNotificationChannel.NotificationChannelConfigProperty.Filters`.
	Filters interface{} `field:"optional" json:"filters" yaml:"filters"`
	// Information about a notification channel configured in DevOps Guru to send notifications when insights are created.
	//
	// If you use an Amazon SNS topic in another account, you must attach a policy to it that grants DevOps Guru permission to it notifications. DevOps Guru adds the required policy on your behalf to send notifications using Amazon SNS in your account. DevOps Guru only supports standard SNS topics. For more information, see [Permissions for cross account Amazon SNS topics](https://docs.aws.amazon.com/devops-guru/latest/userguide/sns-required-permissions.html) .
	//
	// If you use an Amazon SNS topic in another account, you must attach a policy to it that grants DevOps Guru permission to it notifications. DevOps Guru adds the required policy on your behalf to send notifications using Amazon SNS in your account. For more information, see Permissions for cross account Amazon SNS topics.
	//
	// If you use an Amazon SNS topic that is encrypted by an AWS Key Management Service customer-managed key (CMK), then you must add permissions to the CMK. For more information, see [Permissions for AWS KMS–encrypted Amazon SNS topics](https://docs.aws.amazon.com/devops-guru/latest/userguide/sns-kms-permissions.html) .
	Sns interface{} `field:"optional" json:"sns" yaml:"sns"`
}

