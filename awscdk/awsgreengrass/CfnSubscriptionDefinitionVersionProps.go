package awsgreengrass


// Properties for defining a `CfnSubscriptionDefinitionVersion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSubscriptionDefinitionVersionProps := &CfnSubscriptionDefinitionVersionProps{
//   	SubscriptionDefinitionId: jsii.String("subscriptionDefinitionId"),
//   	Subscriptions: []interface{}{
//   		&SubscriptionProperty{
//   			Id: jsii.String("id"),
//   			Source: jsii.String("source"),
//   			Subject: jsii.String("subject"),
//   			Target: jsii.String("target"),
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

