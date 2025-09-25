package awsgreengrass


// A reference to a SubscriptionDefinitionVersion resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   subscriptionDefinitionVersionReference := &SubscriptionDefinitionVersionReference{
//   	SubscriptionDefinitionVersionId: jsii.String("subscriptionDefinitionVersionId"),
//   }
//
type SubscriptionDefinitionVersionReference struct {
	// The Id of the SubscriptionDefinitionVersion resource.
	SubscriptionDefinitionVersionId *string `field:"required" json:"subscriptionDefinitionVersionId" yaml:"subscriptionDefinitionVersionId"`
}

