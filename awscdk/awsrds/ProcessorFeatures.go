package awsrds


// The processor features.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   processorFeatures := &ProcessorFeatures{
//   	CoreCount: jsii.Number(123),
//   	ThreadsPerCore: jsii.Number(123),
//   }
//
type ProcessorFeatures struct {
	// The number of CPU core.
	// Default: - the default number of CPU cores for the chosen instance class.
	//
	CoreCount *float64 `field:"optional" json:"coreCount" yaml:"coreCount"`
	// The number of threads per core.
	// Default: - the default number of threads per core for the chosen instance class.
	//
	ThreadsPerCore *float64 `field:"optional" json:"threadsPerCore" yaml:"threadsPerCore"`
}

