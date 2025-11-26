package previewawsgreengrassmixins


// Properties for CfnSubscriptionDefinitionVersionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnSubscriptionDefinitionVersionMixinProps := &CfnSubscriptionDefinitionVersionMixinProps{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-subscriptiondefinitionversion.html
//
type CfnSubscriptionDefinitionVersionMixinProps struct {
	// The ID of the subscription definition associated with this version.
	//
	// This value is a GUID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-subscriptiondefinitionversion.html#cfn-greengrass-subscriptiondefinitionversion-subscriptiondefinitionid
	//
	SubscriptionDefinitionId *string `field:"optional" json:"subscriptionDefinitionId" yaml:"subscriptionDefinitionId"`
	// The subscriptions in this version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-subscriptiondefinitionversion.html#cfn-greengrass-subscriptiondefinitionversion-subscriptions
	//
	Subscriptions interface{} `field:"optional" json:"subscriptions" yaml:"subscriptions"`
}

