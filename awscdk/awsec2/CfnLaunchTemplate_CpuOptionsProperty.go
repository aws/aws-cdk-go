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
//   cpuOptionsProperty := &CpuOptionsProperty{
//   	AmdSevSnp: jsii.String("amdSevSnp"),
//   	CoreCount: jsii.Number(123),
//   	ThreadsPerCore: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-cpuoptions.html
//
type CfnLaunchTemplate_CpuOptionsProperty struct {
	// Indicates whether to enable the instance for AMD SEV-SNP.
	//
	// AMD SEV-SNP is supported with M6a, R6a, and C6a instance types only. For more information, see [AMD SEV-SNP](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/sev-snp.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-cpuoptions.html#cfn-ec2-launchtemplate-cpuoptions-amdsevsnp
	//
	AmdSevSnp *string `field:"optional" json:"amdSevSnp" yaml:"amdSevSnp"`
	// The number of CPU cores for the instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-cpuoptions.html#cfn-ec2-launchtemplate-cpuoptions-corecount
	//
	CoreCount *float64 `field:"optional" json:"coreCount" yaml:"coreCount"`
	// The number of threads per CPU core.
	//
	// To disable multithreading for the instance, specify a value of `1` . Otherwise, specify the default value of `2` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-cpuoptions.html#cfn-ec2-launchtemplate-cpuoptions-threadspercore
	//
	ThreadsPerCore *float64 `field:"optional" json:"threadsPerCore" yaml:"threadsPerCore"`
}

