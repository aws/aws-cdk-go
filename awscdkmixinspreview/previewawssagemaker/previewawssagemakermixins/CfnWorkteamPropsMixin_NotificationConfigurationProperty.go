package previewawssagemakermixins


// Configures Amazon SNS notifications of available or expiring work items for work teams.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   notificationConfigurationProperty := &NotificationConfigurationProperty{
//   	NotificationTopicArn: jsii.String("notificationTopicArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-workteam-notificationconfiguration.html
//
type CfnWorkteamPropsMixin_NotificationConfigurationProperty struct {
	// The ARN for the Amazon SNS topic to which notifications should be published.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-workteam-notificationconfiguration.html#cfn-sagemaker-workteam-notificationconfiguration-notificationtopicarn
	//
	NotificationTopicArn *string `field:"optional" json:"notificationTopicArn" yaml:"notificationTopicArn"`
}

