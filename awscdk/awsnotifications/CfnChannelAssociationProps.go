package awsnotifications


// Properties for defining a `CfnChannelAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnChannelAssociationProps := &CfnChannelAssociationProps{
//   	Arn: jsii.String("arn"),
//   	NotificationConfigurationArn: jsii.String("notificationConfigurationArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-notifications-channelassociation.html
//
type CfnChannelAssociationProps struct {
	// ARN identifier of the channel.
	//
	// Example: arn:aws:chatbot::123456789012:chat-configuration/slack-channel/security-ops.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-notifications-channelassociation.html#cfn-notifications-channelassociation-arn
	//
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// ARN identifier of the NotificationConfiguration.
	//
	// Example: arn:aws:notifications::123456789012:configuration/a01jes88qxwkbj05xv9c967pgm1.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-notifications-channelassociation.html#cfn-notifications-channelassociation-notificationconfigurationarn
	//
	NotificationConfigurationArn *string `field:"required" json:"notificationConfigurationArn" yaml:"notificationConfigurationArn"`
}

