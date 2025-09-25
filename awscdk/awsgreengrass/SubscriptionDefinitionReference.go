package awsgreengrass


// A reference to a SubscriptionDefinition resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   subscriptionDefinitionReference := &SubscriptionDefinitionReference{
//   	SubscriptionDefinitionArn: jsii.String("subscriptionDefinitionArn"),
//   	SubscriptionDefinitionId: jsii.String("subscriptionDefinitionId"),
//   }
//
type SubscriptionDefinitionReference struct {
	// The ARN of the SubscriptionDefinition resource.
	SubscriptionDefinitionArn *string `field:"required" json:"subscriptionDefinitionArn" yaml:"subscriptionDefinitionArn"`
	// The Id of the SubscriptionDefinition resource.
	SubscriptionDefinitionId *string `field:"required" json:"subscriptionDefinitionId" yaml:"subscriptionDefinitionId"`
}

