package awsec2


// Specifies the CPU options for the instance.
//
// When you specify CPU options, you must specify both the number of CPU cores and threads per core.
//
// For more information, see [Optimize CPU options](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-optimize-cpu.html) in the *Amazon Elastic Compute Cloud User Guide* .
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
type CfnInstance_CpuOptionsProperty struct {
	// The number of CPU cores for the instance.
	CoreCount *float64 `field:"optional" json:"coreCount" yaml:"coreCount"`
	// The number of threads per CPU core.
	ThreadsPerCore *float64 `field:"optional" json:"threadsPerCore" yaml:"threadsPerCore"`
}

