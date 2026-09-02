package awsec2


// The CPU options for the instance.
//
// Example:
//   ec2.NewLaunchTemplate(this, jsii.String("LaunchTemplate"), &LaunchTemplateProps{
//   	MachineImage: ec2.MachineImage_LatestAmazonLinux2023(),
//   	CpuOptions: &LaunchTemplateCpuOptions{
//   		CoreCount: jsii.Number(4),
//   		ThreadsPerCore: jsii.Number(1),
//   		NestedVirtualization: jsii.Boolean(true),
//   	},
//   })
//
// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-optimize-cpu.html
//
type LaunchTemplateCpuOptions struct {
	// Indicates whether to enable the instance for AMD SEV-SNP.
	//
	// AMD SEV-SNP is supported with M6a, R6a, and C6a instance types only.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/sev-snp.html
	//
	// Default: - AMD SEV-SNP is not specified in the launch template.
	//
	AmdSevSnp *bool `field:"optional" json:"amdSevSnp" yaml:"amdSevSnp"`
	// The number of CPU cores for the instance.
	// Default: - The default number of CPU cores for the selected instance type.
	//
	CoreCount *float64 `field:"optional" json:"coreCount" yaml:"coreCount"`
	// Indicates whether the instance is enabled for nested virtualization.
	// Default: - Nested virtualization is not specified in the launch template.
	//
	NestedVirtualization *bool `field:"optional" json:"nestedVirtualization" yaml:"nestedVirtualization"`
	// The number of threads per CPU core.
	//
	// To disable multithreading for the instance, specify a value of 1.
	// Otherwise, specify the default value of 2.
	// Default: - The default number of threads per core for the selected instance type.
	//
	ThreadsPerCore *float64 `field:"optional" json:"threadsPerCore" yaml:"threadsPerCore"`
}

