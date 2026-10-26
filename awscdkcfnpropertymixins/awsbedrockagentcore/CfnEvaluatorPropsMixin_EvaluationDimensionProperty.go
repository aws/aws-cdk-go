package awsbedrockagentcore


// One dimension of an LLM-as-a-Judge evaluator.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   evaluationDimensionProperty := &EvaluationDimensionProperty{
//   	DimensionInstruction: jsii.String("dimensionInstruction"),
//   	DimensionName: jsii.String("dimensionName"),
//   	DimensionRatingScale: &DimensionRatingScaleProperty{
//   		Categorical: []interface{}{
//   			&CategoricalScaleDefinitionProperty{
//   				Definition: jsii.String("definition"),
//   				Label: jsii.String("label"),
//   			},
//   		},
//   		Numerical: []interface{}{
//   			&NumericalScaleDefinitionProperty{
//   				Definition: jsii.String("definition"),
//   				Label: jsii.String("label"),
//   				Value: jsii.Number(123),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-evaluator-evaluationdimension.html
//
type CfnEvaluatorPropsMixin_EvaluationDimensionProperty struct {
	// The instruction to the evaluator model describing what to judge for this dimension.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-evaluator-evaluationdimension.html#cfn-bedrockagentcore-evaluator-evaluationdimension-dimensioninstruction
	//
	DimensionInstruction *string `field:"optional" json:"dimensionInstruction" yaml:"dimensionInstruction"`
	// The name of the dimension.
	//
	// Must be unique within the evaluator.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-evaluator-evaluationdimension.html#cfn-bedrockagentcore-evaluator-evaluationdimension-dimensionname
	//
	DimensionName *string `field:"optional" json:"dimensionName" yaml:"dimensionName"`
	// The rating scale for one dimension of an LLM-as-a-Judge evaluator.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-evaluator-evaluationdimension.html#cfn-bedrockagentcore-evaluator-evaluationdimension-dimensionratingscale
	//
	DimensionRatingScale interface{} `field:"optional" json:"dimensionRatingScale" yaml:"dimensionRatingScale"`
}

