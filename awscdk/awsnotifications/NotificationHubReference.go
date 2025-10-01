package awsnotifications


// A reference to a NotificationHub resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   notificationHubReference := &NotificationHubReference{
//   	Region: jsii.String("region"),
//   }
//
type NotificationHubReference struct {
	// The Region of the NotificationHub resource.
	Region *string `field:"required" json:"region" yaml:"region"`
}

