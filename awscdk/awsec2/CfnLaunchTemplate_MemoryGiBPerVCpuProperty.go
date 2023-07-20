package awsec2


// The minimum and maximum amount of memory per vCPU, in GiB.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   memoryGiBPerVCpuProperty := &MemoryGiBPerVCpuProperty{
//   	Max: jsii.Number(123),
//   	Min: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-memorygibpervcpu.html
//
type CfnLaunchTemplate_MemoryGiBPerVCpuProperty struct {
	// The maximum amount of memory per vCPU, in GiB.
	//
	// To specify no maximum limit, omit this parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-memorygibpervcpu.html#cfn-ec2-launchtemplate-memorygibpervcpu-max
	//
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// The minimum amount of memory per vCPU, in GiB.
	//
	// To specify no minimum limit, omit this parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-memorygibpervcpu.html#cfn-ec2-launchtemplate-memorygibpervcpu-min
	//
	Min *float64 `field:"optional" json:"min" yaml:"min"`
}

