package awsgreengrass


// Properties for defining a `CfnSubscriptionDefinitionVersion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSubscriptionDefinitionVersionProps := &cfnSubscriptionDefinitionVersionProps{
//   	subscriptionDefinitionId: jsii.String("subscriptionDefinitionId"),
//   	subscriptions: []interface{}{
//   		&subscriptionProperty{
//   			id: jsii.String("id"),
//   			source: jsii.String("source"),
//   			subject: jsii.String("subject"),
//   			target: jsii.String("target"),
//   		},
//   	},
//   }
//
type CfnSubscriptionDefinitionVersionProps struct {
	// The ID of the subscription definition associated with this version.
	//
	// This value is a GUID.
	SubscriptionDefinitionId *string `field:"required" json:"subscriptionDefinitionId" yaml:"subscriptionDefinitionId"`
	// The subscriptions in this version.
	Subscriptions interface{} `field:"required" json:"subscriptions" yaml:"subscriptions"`
}

