package awsquicksight


// The range setup of a numeric axis display range.
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var dataDriven interface{}
//
//   axisDisplayRangeProperty := &AxisDisplayRangeProperty{
//   	DataDriven: dataDriven,
//   	MinMax: &AxisDisplayMinMaxRangeProperty{
//   		Maximum: jsii.Number(123),
//   		Minimum: jsii.Number(123),
//   	},
//   }
//
type CfnAnalysis_AxisDisplayRangeProperty struct {
	// The data-driven setup of an axis display range.
	DataDriven interface{} `field:"optional" json:"dataDriven" yaml:"dataDriven"`
	// The minimum and maximum setup of an axis display range.
	MinMax interface{} `field:"optional" json:"minMax" yaml:"minMax"`
}

