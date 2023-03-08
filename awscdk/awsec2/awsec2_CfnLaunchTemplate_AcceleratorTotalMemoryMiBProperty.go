package awsec2


// The minimum and maximum amount of total accelerator memory, in MiB.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   acceleratorTotalMemoryMiBProperty := &acceleratorTotalMemoryMiBProperty{
//   	max: jsii.Number(123),
//   	min: jsii.Number(123),
//   }
//
type CfnLaunchTemplate_AcceleratorTotalMemoryMiBProperty struct {
	// The maximum amount of accelerator memory, in MiB.
	//
	// To specify no maximum limit, omit this parameter.
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// The minimum amount of accelerator memory, in MiB.
	//
	// To specify no minimum limit, omit this parameter.
	Min *float64 `field:"optional" json:"min" yaml:"min"`
}

