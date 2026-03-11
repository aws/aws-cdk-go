package awsbedrockagentcore


// A numerical rating scale option.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   numericalScaleDefinitionProperty := &NumericalScaleDefinitionProperty{
//   	Definition: jsii.String("definition"),
//   	Label: jsii.String("label"),
//   	Value: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-evaluator-numericalscaledefinition.html
//
type CfnEvaluatorPropsMixin_NumericalScaleDefinitionProperty struct {
	// The description that explains what this numerical rating represents.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-evaluator-numericalscaledefinition.html#cfn-bedrockagentcore-evaluator-numericalscaledefinition-definition
	//
	Definition *string `field:"optional" json:"definition" yaml:"definition"`
	// The label that describes this numerical rating option.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-evaluator-numericalscaledefinition.html#cfn-bedrockagentcore-evaluator-numericalscaledefinition-label
	//
	Label *string `field:"optional" json:"label" yaml:"label"`
	// The numerical value for this rating scale option.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-evaluator-numericalscaledefinition.html#cfn-bedrockagentcore-evaluator-numericalscaledefinition-value
	//
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

