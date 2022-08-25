package awsrds


// The processor features.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   processorFeatures := &processorFeatures{
//   	coreCount: jsii.Number(123),
//   	threadsPerCore: jsii.Number(123),
//   }
//
type ProcessorFeatures struct {
	// The number of CPU core.
	CoreCount *float64 `field:"optional" json:"coreCount" yaml:"coreCount"`
	// The number of threads per core.
	ThreadsPerCore *float64 `field:"optional" json:"threadsPerCore" yaml:"threadsPerCore"`
}

