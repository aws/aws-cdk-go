package awsquicksight


// The scale setup options for a numeric axis display.
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   axisScaleProperty := &AxisScaleProperty{
//   	Linear: &AxisLinearScaleProperty{
//   		StepCount: jsii.Number(123),
//   		StepSize: jsii.Number(123),
//   	},
//   	Logarithmic: &AxisLogarithmicScaleProperty{
//   		Base: jsii.Number(123),
//   	},
//   }
//
type CfnDashboard_AxisScaleProperty struct {
	// The linear axis scale setup.
	Linear interface{} `field:"optional" json:"linear" yaml:"linear"`
	// The logarithmic axis scale setup.
	Logarithmic interface{} `field:"optional" json:"logarithmic" yaml:"logarithmic"`
}

