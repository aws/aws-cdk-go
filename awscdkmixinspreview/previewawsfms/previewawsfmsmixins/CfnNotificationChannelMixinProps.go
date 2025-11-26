package previewawsfmsmixins


// Properties for CfnNotificationChannelPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnNotificationChannelMixinProps := &CfnNotificationChannelMixinProps{
//   	SnsRoleName: jsii.String("snsRoleName"),
//   	SnsTopicArn: jsii.String("snsTopicArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fms-notificationchannel.html
//
type CfnNotificationChannelMixinProps struct {
	// The Amazon Resource Name (ARN) of the IAM role that allows Amazon  to record AWS Firewall Manager activity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fms-notificationchannel.html#cfn-fms-notificationchannel-snsrolename
	//
	SnsRoleName *string `field:"optional" json:"snsRoleName" yaml:"snsRoleName"`
	// The Amazon Resource Name (ARN) of the SNS topic that collects notifications from AWS Firewall Manager .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fms-notificationchannel.html#cfn-fms-notificationchannel-snstopicarn
	//
	SnsTopicArn *string `field:"optional" json:"snsTopicArn" yaml:"snsTopicArn"`
}

