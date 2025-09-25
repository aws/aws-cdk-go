package awsdatazone


// A reference to a SubscriptionTarget resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   subscriptionTargetReference := &SubscriptionTargetReference{
//   	DomainId: jsii.String("domainId"),
//   	EnvironmentId: jsii.String("environmentId"),
//   	SubscriptionTargetId: jsii.String("subscriptionTargetId"),
//   }
//
type SubscriptionTargetReference struct {
	// The DomainId of the SubscriptionTarget resource.
	DomainId *string `field:"required" json:"domainId" yaml:"domainId"`
	// The EnvironmentId of the SubscriptionTarget resource.
	EnvironmentId *string `field:"required" json:"environmentId" yaml:"environmentId"`
	// The Id of the SubscriptionTarget resource.
	SubscriptionTargetId *string `field:"required" json:"subscriptionTargetId" yaml:"subscriptionTargetId"`
}

