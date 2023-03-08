package awsinspectorv2


// An object that describes the details of a string filter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stringFilterProperty := &stringFilterProperty{
//   	comparison: jsii.String("comparison"),
//   	value: jsii.String("value"),
//   }
//
type CfnFilter_StringFilterProperty struct {
	// The operator to use when comparing values in the filter.
	Comparison *string `field:"required" json:"comparison" yaml:"comparison"`
	// The value to filter on.
	Value *string `field:"required" json:"value" yaml:"value"`
}

