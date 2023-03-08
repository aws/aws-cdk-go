package awsgreengrass


// Properties for defining a `CfnSubscriptionDefinition`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tags interface{}
//
//   cfnSubscriptionDefinitionProps := &cfnSubscriptionDefinitionProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	initialVersion: &subscriptionDefinitionVersionProperty{
//   		subscriptions: []interface{}{
//   			&subscriptionProperty{
//   				id: jsii.String("id"),
//   				source: jsii.String("source"),
//   				subject: jsii.String("subject"),
//   				target: jsii.String("target"),
//   			},
//   		},
//   	},
//   	tags: tags,
//   }
//
type CfnSubscriptionDefinitionProps struct {
	// The name of the subscription definition.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The subscription definition version to include when the subscription definition is created.
	//
	// A subscription definition version contains a list of [`subscription`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-subscriptiondefinition-subscription.html) property types.
	//
	// > To associate a subscription definition version after the subscription definition is created, create an [`AWS::Greengrass::SubscriptionDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-subscriptiondefinitionversion.html) resource and specify the ID of this subscription definition.
	InitialVersion interface{} `field:"optional" json:"initialVersion" yaml:"initialVersion"`
	// Application-specific metadata to attach to the subscription definition.
	//
	// You can use tags in IAM policies to control access to AWS IoT Greengrass resources. You can also use tags to categorize your resources. For more information, see [Tagging Your AWS IoT Greengrass Resources](https://docs.aws.amazon.com/greengrass/latest/developerguide/tagging.html) in the *Developer Guide* .
	//
	// This `Json` property type is processed as a map of key-value pairs. It uses the following format, which is different from most `Tags` implementations in AWS CloudFormation templates.
	//
	// ```json
	// "Tags": { "KeyName0": "value", "KeyName1": "value", "KeyName2": "value"
	// }
	// ```.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

