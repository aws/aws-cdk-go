package previewawsnotificationsmixins


// Properties for CfnManagedNotificationAdditionalChannelAssociationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnManagedNotificationAdditionalChannelAssociationMixinProps := &CfnManagedNotificationAdditionalChannelAssociationMixinProps{
//   	ChannelArn: jsii.String("channelArn"),
//   	ManagedNotificationConfigurationArn: jsii.String("managedNotificationConfigurationArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-notifications-managednotificationadditionalchannelassociation.html
//
type CfnManagedNotificationAdditionalChannelAssociationMixinProps struct {
	// The ARN of the `Channel` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-notifications-managednotificationadditionalchannelassociation.html#cfn-notifications-managednotificationadditionalchannelassociation-channelarn
	//
	ChannelArn *string `field:"optional" json:"channelArn" yaml:"channelArn"`
	// The ARN of the `ManagedNotificationAdditionalChannelAssociation` associated with the `Channel` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-notifications-managednotificationadditionalchannelassociation.html#cfn-notifications-managednotificationadditionalchannelassociation-managednotificationconfigurationarn
	//
	ManagedNotificationConfigurationArn *string `field:"optional" json:"managedNotificationConfigurationArn" yaml:"managedNotificationConfigurationArn"`
}

