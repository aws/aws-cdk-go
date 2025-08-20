package awsnotifications


// Properties for defining a `CfnNotificationHub`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnNotificationHubProps := &CfnNotificationHubProps{
//   	Region: jsii.String("region"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-notifications-notificationhub.html
//
type CfnNotificationHubProps struct {
	// The `NotificationHub` Region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-notifications-notificationhub.html#cfn-notifications-notificationhub-region
	//
	Region *string `field:"required" json:"region" yaml:"region"`
}

