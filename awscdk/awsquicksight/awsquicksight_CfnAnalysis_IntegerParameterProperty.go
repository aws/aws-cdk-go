package awsquicksight


// An integer parameter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   integerParameterProperty := &integerParameterProperty{
//   	name: jsii.String("name"),
//   	values: []interface{}{
//   		jsii.Number(123),
//   	},
//   }
//
type CfnAnalysis_IntegerParameterProperty struct {
	// The name of the integer parameter.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The values for the integer parameter.
	Values interface{} `field:"required" json:"values" yaml:"values"`
}

