package awspersonalize


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   categoricalHyperParameterRangeProperty := &categoricalHyperParameterRangeProperty{
//   	name: jsii.String("name"),
//   	values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
type CfnSolution_CategoricalHyperParameterRangeProperty struct {
	// `CfnSolution.CategoricalHyperParameterRangeProperty.Name`.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// `CfnSolution.CategoricalHyperParameterRangeProperty.Values`.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

