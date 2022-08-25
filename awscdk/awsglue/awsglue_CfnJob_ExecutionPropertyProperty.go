package awsglue


// An execution property of a job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   executionPropertyProperty := &executionPropertyProperty{
//   	maxConcurrentRuns: jsii.Number(123),
//   }
//
type CfnJob_ExecutionPropertyProperty struct {
	// The maximum number of concurrent runs allowed for the job.
	//
	// The default is 1. An error is returned when this threshold is reached. The maximum value you can specify is controlled by a service limit.
	MaxConcurrentRuns *float64 `field:"optional" json:"maxConcurrentRuns" yaml:"maxConcurrentRuns"`
}

