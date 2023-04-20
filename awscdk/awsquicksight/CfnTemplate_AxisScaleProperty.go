package awsquicksight


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
type CfnTemplate_AxisScaleProperty struct {
	// `CfnTemplate.AxisScaleProperty.Linear`.
	Linear interface{} `field:"optional" json:"linear" yaml:"linear"`
	// `CfnTemplate.AxisScaleProperty.Logarithmic`.
	Logarithmic interface{} `field:"optional" json:"logarithmic" yaml:"logarithmic"`
}

