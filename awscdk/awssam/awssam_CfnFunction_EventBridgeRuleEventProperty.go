package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var pattern interface{}
//
//   eventBridgeRuleEventProperty := &eventBridgeRuleEventProperty{
//   	pattern: pattern,
//
//   	// the properties below are optional
//   	eventBusName: jsii.String("eventBusName"),
//   	input: jsii.String("input"),
//   	inputPath: jsii.String("inputPath"),
//   }
//
type CfnFunction_EventBridgeRuleEventProperty struct {
	// `CfnFunction.EventBridgeRuleEventProperty.Pattern`.
	Pattern interface{} `field:"required" json:"pattern" yaml:"pattern"`
	// `CfnFunction.EventBridgeRuleEventProperty.EventBusName`.
	EventBusName *string `field:"optional" json:"eventBusName" yaml:"eventBusName"`
	// `CfnFunction.EventBridgeRuleEventProperty.Input`.
	Input *string `field:"optional" json:"input" yaml:"input"`
	// `CfnFunction.EventBridgeRuleEventProperty.InputPath`.
	InputPath *string `field:"optional" json:"inputPath" yaml:"inputPath"`
}

