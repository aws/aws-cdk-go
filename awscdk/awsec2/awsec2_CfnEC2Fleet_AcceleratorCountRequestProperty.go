package awsec2


// The minimum and maximum number of accelerators (GPUs, FPGAs, or AWS Inferentia chips) on an instance.
//
// To exclude accelerator-enabled instance types, set `Max` to `0` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   acceleratorCountRequestProperty := &acceleratorCountRequestProperty{
//   	max: jsii.Number(123),
//   	min: jsii.Number(123),
//   }
//
type CfnEC2Fleet_AcceleratorCountRequestProperty struct {
	// The maximum number of accelerators.
	//
	// To specify no maximum limit, omit this parameter. To exclude accelerator-enabled instance types, set `Max` to `0` .
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// The minimum number of accelerators.
	//
	// To specify no minimum limit, omit this parameter.
	Min *float64 `field:"optional" json:"min" yaml:"min"`
}

