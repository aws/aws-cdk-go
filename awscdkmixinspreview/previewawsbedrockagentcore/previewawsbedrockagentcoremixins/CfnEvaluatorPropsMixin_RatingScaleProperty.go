package previewawsbedrockagentcoremixins


// The rating scale that defines how evaluators should score agent performance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ratingScaleProperty := &RatingScaleProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-evaluator-ratingscale.html
//
type CfnEvaluatorPropsMixin_RatingScaleProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-evaluator-ratingscale.html#cfn-bedrockagentcore-evaluator-ratingscale-categorical
	//
	Categorical interface{} `field:"optional" json:"categorical" yaml:"categorical"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-evaluator-ratingscale.html#cfn-bedrockagentcore-evaluator-ratingscale-numerical
	//
	Numerical interface{} `field:"optional" json:"numerical" yaml:"numerical"`
}

