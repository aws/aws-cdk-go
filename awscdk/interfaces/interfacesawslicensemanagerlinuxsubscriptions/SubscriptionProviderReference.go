package interfacesawslicensemanagerlinuxsubscriptions


// A reference to a SubscriptionProvider resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   subscriptionProviderReference := &SubscriptionProviderReference{
//   	SubscriptionProviderArn: jsii.String("subscriptionProviderArn"),
//   }
//
type SubscriptionProviderReference struct {
	// The SubscriptionProviderArn of the SubscriptionProvider resource.
	SubscriptionProviderArn *string `field:"required" json:"subscriptionProviderArn" yaml:"subscriptionProviderArn"`
}

