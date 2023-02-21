package awspersonalize


// Provides the name and range of a categorical hyperparameter.
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
type CfnSolution_CategoricalHyperParameterRangeProperty struct {
	// The name of the hyperparameter.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A list of the categories for the hyperparameter.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

