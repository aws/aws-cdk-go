package awsmedialive


// A reference to a EventBridgeRuleTemplate resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventBridgeRuleTemplateReference := &EventBridgeRuleTemplateReference{
//   	EventBridgeRuleTemplateArn: jsii.String("eventBridgeRuleTemplateArn"),
//   	Identifier: jsii.String("identifier"),
//   }
//
type EventBridgeRuleTemplateReference struct {
	// The ARN of the EventBridgeRuleTemplate resource.
	EventBridgeRuleTemplateArn *string `field:"required" json:"eventBridgeRuleTemplateArn" yaml:"eventBridgeRuleTemplateArn"`
	// The Identifier of the EventBridgeRuleTemplate resource.
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
}

