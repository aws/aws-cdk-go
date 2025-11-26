package previewawsbedrockmixins


// Contains configurations for the controller node of a DoWhile loop in the flow.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   loopControllerFlowNodeConfigurationProperty := &LoopControllerFlowNodeConfigurationProperty{
//   	ContinueCondition: &FlowConditionProperty{
//   		Expression: jsii.String("expression"),
//   		Name: jsii.String("name"),
//   	},
//   	MaxIterations: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-loopcontrollerflownodeconfiguration.html
//
type CfnFlowVersionPropsMixin_LoopControllerFlowNodeConfigurationProperty struct {
	// Specifies the condition that determines when the flow exits the DoWhile loop.
	//
	// The loop executes until this condition evaluates to true.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-loopcontrollerflownodeconfiguration.html#cfn-bedrock-flowversion-loopcontrollerflownodeconfiguration-continuecondition
	//
	ContinueCondition interface{} `field:"optional" json:"continueCondition" yaml:"continueCondition"`
	// Specifies the maximum number of times the DoWhile loop can iterate before the flow exits the loop.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-loopcontrollerflownodeconfiguration.html#cfn-bedrock-flowversion-loopcontrollerflownodeconfiguration-maxiterations
	//
	// Default: - 10.
	//
	MaxIterations *float64 `field:"optional" json:"maxIterations" yaml:"maxIterations"`
}

