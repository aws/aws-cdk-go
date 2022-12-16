package awsec2


// Specifies the CPU options for an instance.
//
// For more information, see [Optimize CPU options](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-optimize-cpu.html) in the *Amazon Elastic Compute Cloud User Guide* .
//
// `CpuOptions` is a property of [AWS::EC2::LaunchTemplate LaunchTemplateData](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cpuOptionsProperty := &cpuOptionsProperty{
//   	coreCount: jsii.Number(123),
//   	threadsPerCore: jsii.Number(123),
//   }
//
type CfnLaunchTemplate_CpuOptionsProperty struct {
	// The number of CPU cores for the instance.
	CoreCount *float64 `field:"optional" json:"coreCount" yaml:"coreCount"`
	// The number of threads per CPU core.
	//
	// To disable multithreading for the instance, specify a value of 1. Otherwise, specify the default value of 2.
	ThreadsPerCore *float64 `field:"optional" json:"threadsPerCore" yaml:"threadsPerCore"`
}

