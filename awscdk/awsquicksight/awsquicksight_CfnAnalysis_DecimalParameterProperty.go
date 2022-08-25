package awsquicksight


// A decimal parameter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   decimalParameterProperty := &decimalParameterProperty{
//   	name: jsii.String("name"),
//   	values: []interface{}{
//   		jsii.Number(123),
//   	},
//   }
//
type CfnAnalysis_DecimalParameterProperty struct {
	// A display name for the decimal parameter.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The values for the decimal parameter.
	Values interface{} `field:"required" json:"values" yaml:"values"`
}

