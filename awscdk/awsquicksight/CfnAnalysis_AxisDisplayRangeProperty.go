package awsquicksight


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
	// `CfnAnalysis.AxisDisplayRangeProperty.DataDriven`.
	DataDriven interface{} `field:"optional" json:"dataDriven" yaml:"dataDriven"`
	// `CfnAnalysis.AxisDisplayRangeProperty.MinMax`.
	MinMax interface{} `field:"optional" json:"minMax" yaml:"minMax"`
}

