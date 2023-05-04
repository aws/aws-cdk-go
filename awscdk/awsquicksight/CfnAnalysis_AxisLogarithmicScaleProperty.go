package awsquicksight


// The logarithmic axis scale setup.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   axisLogarithmicScaleProperty := &AxisLogarithmicScaleProperty{
//   	Base: jsii.Number(123),
//   }
//
type CfnAnalysis_AxisLogarithmicScaleProperty struct {
	// The base setup of a logarithmic axis scale.
	Base *float64 `field:"optional" json:"base" yaml:"base"`
}

