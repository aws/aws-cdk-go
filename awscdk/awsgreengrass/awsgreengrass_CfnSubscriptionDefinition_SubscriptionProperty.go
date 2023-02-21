package awsgreengrass


// Subscriptions define how MQTT messages can be exchanged between devices, functions, and connectors in the group, and with AWS IoT or the local shadow service.
//
// A subscription defines a message source, message target, and a topic (or subject) that's used to route messages from the source to the target. A subscription defines the message flow in one direction, from the source to the target. For two-way communication, you must set up two subscriptions, one for each direction.
//
// In an AWS CloudFormation template, the `Subscriptions` property of the [`SubscriptionDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-subscriptiondefinition-subscriptiondefinitionversion.html) property type contains a list of `Subscription` property types.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   subscriptionProperty := &subscriptionProperty{
//   	id: jsii.String("id"),
//   	source: jsii.String("source"),
//   	subject: jsii.String("subject"),
//   	target: jsii.String("target"),
//   }
//
type CfnSubscriptionDefinition_SubscriptionProperty struct {
	// A descriptive or arbitrary ID for the subscription.
	//
	// This value must be unique within the subscription definition version. Maximum length is 128 characters with pattern `[a-zA-Z0-9:_-]+` .
	Id *string `field:"required" json:"id" yaml:"id"`
	// The originator of the message.
	//
	// The value can be a thing ARN, the ARN of a Lambda function alias (recommended) or version, a connector ARN, `cloud` (which represents the AWS IoT cloud), or `GGShadowService` .
	Source *string `field:"required" json:"source" yaml:"source"`
	// The MQTT topic used to route the message.
	Subject *string `field:"required" json:"subject" yaml:"subject"`
	// The destination of the message.
	//
	// The value can be a thing ARN, the ARN of a Lambda function alias (recommended) or version, a connector ARN, `cloud` (which represents the AWS IoT cloud), or `GGShadowService` .
	Target *string `field:"required" json:"target" yaml:"target"`
}

