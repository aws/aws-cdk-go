package awsec2


// The minimum and maximum number of vCPUs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vCpuCountRangeRequestProperty := &vCpuCountRangeRequestProperty{
//   	max: jsii.Number(123),
//   	min: jsii.Number(123),
//   }
//
type CfnEC2Fleet_VCpuCountRangeRequestProperty struct {
	// The maximum number of vCPUs.
	//
	// To specify no maximum limit, omit this parameter.
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// The minimum number of vCPUs.
	//
	// To specify no minimum limit, specify `0` .
	Min *float64 `field:"optional" json:"min" yaml:"min"`
}

