package awsinspectorv2


// An object that describes details of a map filter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mapFilterProperty := &mapFilterProperty{
//   	comparison: jsii.String("comparison"),
//
//   	// the properties below are optional
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//   }
//
type CfnFilter_MapFilterProperty struct {
	// The operator to use when comparing values in the filter.
	Comparison *string `field:"required" json:"comparison" yaml:"comparison"`
	// The tag key used in the filter.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The tag value used in the filter.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

