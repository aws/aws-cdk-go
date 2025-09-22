package awsarczonalshift


// A reference to a AutoshiftObserverNotificationStatus resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   autoshiftObserverNotificationStatusReference := &AutoshiftObserverNotificationStatusReference{
//   	AccountId: jsii.String("accountId"),
//   	Region: jsii.String("region"),
//   }
//
type AutoshiftObserverNotificationStatusReference struct {
	// The AccountId of the AutoshiftObserverNotificationStatus resource.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The Region of the AutoshiftObserverNotificationStatus resource.
	Region *string `field:"required" json:"region" yaml:"region"`
}

