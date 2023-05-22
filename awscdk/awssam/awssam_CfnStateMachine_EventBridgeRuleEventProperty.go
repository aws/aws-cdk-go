package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var pattern interface{}
//
//   eventBridgeRuleEventProperty := &EventBridgeRuleEventProperty{
//   	Pattern: pattern,
//
//   	// the properties below are optional
//   	EventBusName: jsii.String("eventBusName"),
//   	Input: jsii.String("input"),
//   	InputPath: jsii.String("inputPath"),
//   }
//
type CfnStateMachine_EventBridgeRuleEventProperty struct {
	// `CfnStateMachine.EventBridgeRuleEventProperty.Pattern`.
	Pattern interface{} `field:"required" json:"pattern" yaml:"pattern"`
	// `CfnStateMachine.EventBridgeRuleEventProperty.EventBusName`.
	EventBusName *string `field:"optional" json:"eventBusName" yaml:"eventBusName"`
	// `CfnStateMachine.EventBridgeRuleEventProperty.Input`.
	Input *string `field:"optional" json:"input" yaml:"input"`
	// `CfnStateMachine.EventBridgeRuleEventProperty.InputPath`.
	InputPath *string `field:"optional" json:"inputPath" yaml:"inputPath"`
}

