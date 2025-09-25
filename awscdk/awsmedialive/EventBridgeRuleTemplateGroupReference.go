package awsmedialive


// A reference to a EventBridgeRuleTemplateGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventBridgeRuleTemplateGroupReference := &EventBridgeRuleTemplateGroupReference{
//   	EventBridgeRuleTemplateGroupArn: jsii.String("eventBridgeRuleTemplateGroupArn"),
//   	Identifier: jsii.String("identifier"),
//   }
//
type EventBridgeRuleTemplateGroupReference struct {
	// The ARN of the EventBridgeRuleTemplateGroup resource.
	EventBridgeRuleTemplateGroupArn *string `field:"required" json:"eventBridgeRuleTemplateGroupArn" yaml:"eventBridgeRuleTemplateGroupArn"`
	// The Identifier of the EventBridgeRuleTemplateGroup resource.
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
}

