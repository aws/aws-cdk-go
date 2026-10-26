package awsbedrockagentcore


// The rating scale for one dimension of an LLM-as-a-Judge evaluator.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   dimensionRatingScaleProperty := &DimensionRatingScaleProperty{
//   	Categorical: []interface{}{
//   		&CategoricalScaleDefinitionProperty{
//   			Definition: jsii.String("definition"),
//   			Label: jsii.String("label"),
//   		},
//   	},
//   	Numerical: []interface{}{
//   		&NumericalScaleDefinitionProperty{
//   			Definition: jsii.String("definition"),
//   			Label: jsii.String("label"),
//   			Value: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-evaluator-dimensionratingscale.html
//
type CfnEvaluatorPropsMixin_DimensionRatingScaleProperty struct {
	// The categorical rating scale for this dimension.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-evaluator-dimensionratingscale.html#cfn-bedrockagentcore-evaluator-dimensionratingscale-categorical
	//
	Categorical interface{} `field:"optional" json:"categorical" yaml:"categorical"`
	// The numerical rating scale for this dimension.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-evaluator-dimensionratingscale.html#cfn-bedrockagentcore-evaluator-dimensionratingscale-numerical
	//
	Numerical interface{} `field:"optional" json:"numerical" yaml:"numerical"`
}

