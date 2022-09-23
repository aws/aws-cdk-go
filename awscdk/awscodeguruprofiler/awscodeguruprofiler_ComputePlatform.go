package awscodeguruprofiler


// The compute platform of the profiling group.
//
// Example:
//   profilingGroup := codeguruprofiler.NewProfilingGroup(this, jsii.String("MyProfilingGroup"), &profilingGroupProps{
//   	computePlatform: codeguruprofiler.computePlatform_AWS_LAMBDA,
//   })
//
type ComputePlatform string

const (
	// Use AWS_LAMBDA if your application runs on AWS Lambda.
	ComputePlatform_AWS_LAMBDA ComputePlatform = "AWS_LAMBDA"
	// Use Default if your application runs on a compute platform that is not AWS Lambda, such an Amazon EC2 instance, an on-premises server, or a different platform.
	ComputePlatform_DEFAULT ComputePlatform = "DEFAULT"
)

