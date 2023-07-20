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
//   acceleratorCountRequestProperty := &AcceleratorCountRequestProperty{
//   	Max: jsii.Number(123),
//   	Min: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-acceleratorcountrequest.html
//
type CfnEC2Fleet_AcceleratorCountRequestProperty struct {
	// The maximum number of accelerators.
	//
	// To specify no maximum limit, omit this parameter. To exclude accelerator-enabled instance types, set `Max` to `0` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-acceleratorcountrequest.html#cfn-ec2-ec2fleet-acceleratorcountrequest-max
	//
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// The minimum number of accelerators.
	//
	// To specify no minimum limit, omit this parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-acceleratorcountrequest.html#cfn-ec2-ec2fleet-acceleratorcountrequest-min
	//
	Min *float64 `field:"optional" json:"min" yaml:"min"`
}

