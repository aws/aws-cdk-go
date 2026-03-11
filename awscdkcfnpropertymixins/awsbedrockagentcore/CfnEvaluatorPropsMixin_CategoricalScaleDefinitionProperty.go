package awsbedrockagentcore


// A categorical rating scale option.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   categoricalScaleDefinitionProperty := &CategoricalScaleDefinitionProperty{
//   	Definition: jsii.String("definition"),
//   	Label: jsii.String("label"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-evaluator-categoricalscaledefinition.html
//
type CfnEvaluatorPropsMixin_CategoricalScaleDefinitionProperty struct {
	// The description that explains what this categorical rating represents.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-evaluator-categoricalscaledefinition.html#cfn-bedrockagentcore-evaluator-categoricalscaledefinition-definition
	//
	Definition *string `field:"optional" json:"definition" yaml:"definition"`
	// The label of this categorical rating option.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-evaluator-categoricalscaledefinition.html#cfn-bedrockagentcore-evaluator-categoricalscaledefinition-label
	//
	Label *string `field:"optional" json:"label" yaml:"label"`
}

