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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-axisdisplayrange.html
//
type CfnDashboard_AxisDisplayRangeProperty struct {
	// The data-driven setup of an axis display range.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-axisdisplayrange.html#cfn-quicksight-dashboard-axisdisplayrange-datadriven
	//
	DataDriven interface{} `field:"optional" json:"dataDriven" yaml:"dataDriven"`
	// The minimum and maximum setup of an axis display range.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-axisdisplayrange.html#cfn-quicksight-dashboard-axisdisplayrange-minmax
	//
	MinMax interface{} `field:"optional" json:"minMax" yaml:"minMax"`
}

