package previewawssammixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var pattern interface{}
//
//   eventBridgeRuleEventProperty := &EventBridgeRuleEventProperty{
//   	EventBusName: jsii.String("eventBusName"),
//   	Input: jsii.String("input"),
//   	InputPath: jsii.String("inputPath"),
//   	Pattern: pattern,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-eventbridgeruleevent.html
//
type CfnFunctionPropsMixin_EventBridgeRuleEventProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-eventbridgeruleevent.html#cfn-serverless-function-eventbridgeruleevent-eventbusname
	//
	EventBusName *string `field:"optional" json:"eventBusName" yaml:"eventBusName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-eventbridgeruleevent.html#cfn-serverless-function-eventbridgeruleevent-input
	//
	Input *string `field:"optional" json:"input" yaml:"input"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-eventbridgeruleevent.html#cfn-serverless-function-eventbridgeruleevent-inputpath
	//
	InputPath *string `field:"optional" json:"inputPath" yaml:"inputPath"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-eventbridgeruleevent.html#cfn-serverless-function-eventbridgeruleevent-pattern
	//
	Pattern interface{} `field:"optional" json:"pattern" yaml:"pattern"`
}

