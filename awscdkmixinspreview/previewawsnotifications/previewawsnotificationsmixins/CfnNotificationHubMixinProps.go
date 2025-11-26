package previewawsnotificationsmixins


// Properties for CfnNotificationHubPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnNotificationHubMixinProps := &CfnNotificationHubMixinProps{
//   	Region: jsii.String("region"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-notifications-notificationhub.html
//
type CfnNotificationHubMixinProps struct {
	// The `NotificationHub` Region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-notifications-notificationhub.html#cfn-notifications-notificationhub-region
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
}

