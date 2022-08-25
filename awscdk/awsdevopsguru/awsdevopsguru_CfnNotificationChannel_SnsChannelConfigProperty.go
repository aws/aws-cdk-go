package awsdevopsguru


// Contains the Amazon Resource Name (ARN) of an Amazon Simple Notification Service topic.
//
// If you use an Amazon SNS topic in another account, you must attach a policy to it that grants DevOps Guru permission to it notifications. DevOps Guru adds the required policy on your behalf to send notifications using Amazon SNS in your account. DevOps Guru only supports standard SNS topics. For more information, see [Permissions for cross account Amazon SNS topics](https://docs.aws.amazon.com/devops-guru/latest/userguide/sns-required-permissions.html) .
//
// If you use an Amazon SNS topic in another account, you must attach a policy to it that grants DevOps Guru permission to it notifications. DevOps Guru adds the required policy on your behalf to send notifications using Amazon SNS in your account. For more information, see Permissions for cross account Amazon SNS topics.
//
// If you use an Amazon SNS topic that is encrypted by an AWS Key Management Service customer-managed key (CMK), then you must add permissions to the CMK. For more information, see [Permissions for AWS KMS–encrypted Amazon SNS topics](https://docs.aws.amazon.com/devops-guru/latest/userguide/sns-kms-permissions.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   snsChannelConfigProperty := &snsChannelConfigProperty{
//   	topicArn: jsii.String("topicArn"),
//   }
//
type CfnNotificationChannel_SnsChannelConfigProperty struct {
	// The Amazon Resource Name (ARN) of an Amazon Simple Notification Service topic.
	TopicArn *string `field:"optional" json:"topicArn" yaml:"topicArn"`
}

