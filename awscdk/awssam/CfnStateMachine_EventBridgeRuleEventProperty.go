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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-statemachine-eventbridgeruleevent.html
//
type CfnStateMachine_EventBridgeRuleEventProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-statemachine-eventbridgeruleevent.html#cfn-serverless-statemachine-eventbridgeruleevent-pattern
	//
	Pattern interface{} `field:"required" json:"pattern" yaml:"pattern"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-statemachine-eventbridgeruleevent.html#cfn-serverless-statemachine-eventbridgeruleevent-eventbusname
	//
	EventBusName *string `field:"optional" json:"eventBusName" yaml:"eventBusName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-statemachine-eventbridgeruleevent.html#cfn-serverless-statemachine-eventbridgeruleevent-input
	//
	Input *string `field:"optional" json:"input" yaml:"input"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-statemachine-eventbridgeruleevent.html#cfn-serverless-statemachine-eventbridgeruleevent-inputpath
	//
	InputPath *string `field:"optional" json:"inputPath" yaml:"inputPath"`
}

