package awsec2


// The minimum and maximum amount of total accelerator memory, in MiB.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   acceleratorTotalMemoryMiBProperty := &AcceleratorTotalMemoryMiBProperty{
//   	Max: jsii.Number(123),
//   	Min: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-acceleratortotalmemorymib.html
//
type CfnLaunchTemplate_AcceleratorTotalMemoryMiBProperty struct {
	// The maximum amount of accelerator memory, in MiB.
	//
	// To specify no maximum limit, omit this parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-acceleratortotalmemorymib.html#cfn-ec2-launchtemplate-acceleratortotalmemorymib-max
	//
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// The minimum amount of accelerator memory, in MiB.
	//
	// To specify no minimum limit, omit this parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-acceleratortotalmemorymib.html#cfn-ec2-launchtemplate-acceleratortotalmemorymib-min
	//
	Min *float64 `field:"optional" json:"min" yaml:"min"`
}

