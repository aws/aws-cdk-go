package awsquicksight


// The liner axis scale setup.
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   axisLinearScaleProperty := &AxisLinearScaleProperty{
//   	StepCount: jsii.Number(123),
//   	StepSize: jsii.Number(123),
//   }
//
type CfnDashboard_AxisLinearScaleProperty struct {
	// The step count setup of a linear axis.
	StepCount *float64 `field:"optional" json:"stepCount" yaml:"stepCount"`
	// The step size setup of a linear axis.
	StepSize *float64 `field:"optional" json:"stepSize" yaml:"stepSize"`
}

