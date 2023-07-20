package awspersonalize


// Provides the name and values of a Categorical hyperparameter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   categoricalHyperParameterRangeProperty := &CategoricalHyperParameterRangeProperty{
//   	Name: jsii.String("name"),
//   	Values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-solution-categoricalhyperparameterrange.html
//
type CfnSolution_CategoricalHyperParameterRangeProperty struct {
	// The name of the hyperparameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-solution-categoricalhyperparameterrange.html#cfn-personalize-solution-categoricalhyperparameterrange-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A list of the categories for the hyperparameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-solution-categoricalhyperparameterrange.html#cfn-personalize-solution-categoricalhyperparameterrange-values
	//
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

