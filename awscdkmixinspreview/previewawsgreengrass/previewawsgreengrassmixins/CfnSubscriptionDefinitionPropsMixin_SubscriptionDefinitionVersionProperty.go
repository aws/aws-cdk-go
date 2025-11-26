package previewawsgreengrassmixins


// A subscription definition version contains a list of [subscriptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-subscriptiondefinition-subscription.html) .
//
// > After you create a subscription definition version that contains the subscriptions you want to deploy, you must add it to your group version. For more information, see [`AWS::Greengrass::Group`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-group.html) .
//
// In an CloudFormation template, `SubscriptionDefinitionVersion` is the property type of the `InitialVersion` property in the [`AWS::Greengrass::SubscriptionDefinition`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-subscriptiondefinition.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   subscriptionDefinitionVersionProperty := &SubscriptionDefinitionVersionProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-subscriptiondefinition-subscriptiondefinitionversion.html
//
type CfnSubscriptionDefinitionPropsMixin_SubscriptionDefinitionVersionProperty struct {
	// The subscriptions in this version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-subscriptiondefinition-subscriptiondefinitionversion.html#cfn-greengrass-subscriptiondefinition-subscriptiondefinitionversion-subscriptions
	//
	Subscriptions interface{} `field:"optional" json:"subscriptions" yaml:"subscriptions"`
}

