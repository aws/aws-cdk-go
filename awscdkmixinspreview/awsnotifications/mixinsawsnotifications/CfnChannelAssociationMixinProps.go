package mixinsawsnotifications


// Properties for CfnChannelAssociationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnChannelAssociationMixinProps := &CfnChannelAssociationMixinProps{
//   	Arn: jsii.String("arn"),
//   	NotificationConfigurationArn: jsii.String("notificationConfigurationArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-notifications-channelassociation.html
//
type CfnChannelAssociationMixinProps struct {
	// The Amazon Resource Name (ARN) of the `Channel` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-notifications-channelassociation.html#cfn-notifications-channelassociation-arn
	//
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// The ARN of the `NotificationConfiguration` associated with the `Channel` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-notifications-channelassociation.html#cfn-notifications-channelassociation-notificationconfigurationarn
	//
	NotificationConfigurationArn *string `field:"optional" json:"notificationConfigurationArn" yaml:"notificationConfigurationArn"`
}

