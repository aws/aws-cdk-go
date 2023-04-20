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
type CfnTemplate_AxisLinearScaleProperty struct {
	// `CfnTemplate.AxisLinearScaleProperty.StepCount`.
	StepCount *float64 `field:"optional" json:"stepCount" yaml:"stepCount"`
	// `CfnTemplate.AxisLinearScaleProperty.StepSize`.
	StepSize *float64 `field:"optional" json:"stepSize" yaml:"stepSize"`
}

