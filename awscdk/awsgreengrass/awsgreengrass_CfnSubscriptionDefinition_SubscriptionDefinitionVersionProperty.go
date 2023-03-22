package awsgreengrass


// A subscription definition version contains a list of [subscriptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-subscriptiondefinition-subscription.html) .
//
// > After you create a subscription definition version that contains the subscriptions you want to deploy, you must add it to your group version. For more information, see [`AWS::Greengrass::Group`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-group.html) .
//
// In an AWS CloudFormation template, `SubscriptionDefinitionVersion` is the property type of the `InitialVersion` property in the [`AWS::Greengrass::SubscriptionDefinition`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-subscriptiondefinition.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   subscriptionDefinitionVersionProperty := &subscriptionDefinitionVersionProperty{
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
type CfnSubscriptionDefinition_SubscriptionDefinitionVersionProperty struct {
	// The subscriptions in this version.
	Subscriptions interface{} `field:"required" json:"subscriptions" yaml:"subscriptions"`
}

