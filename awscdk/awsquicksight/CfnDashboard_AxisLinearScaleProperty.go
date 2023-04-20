package awsquicksight


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
	// `CfnDashboard.AxisLinearScaleProperty.StepCount`.
	StepCount *float64 `field:"optional" json:"stepCount" yaml:"stepCount"`
	// `CfnDashboard.AxisLinearScaleProperty.StepSize`.
	StepSize *float64 `field:"optional" json:"stepSize" yaml:"stepSize"`
}

