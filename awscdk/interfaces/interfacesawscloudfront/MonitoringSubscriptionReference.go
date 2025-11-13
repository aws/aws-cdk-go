package interfacesawscloudfront


// A reference to a MonitoringSubscription resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   monitoringSubscriptionReference := &MonitoringSubscriptionReference{
//   	DistributionId: jsii.String("distributionId"),
//   }
//
type MonitoringSubscriptionReference struct {
	// The DistributionId of the MonitoringSubscription resource.
	DistributionId *string `field:"required" json:"distributionId" yaml:"distributionId"`
}

