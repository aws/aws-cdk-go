package awsbatch


// An object representing a job timeout configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   timeoutProperty := &timeoutProperty{
//   	attemptDurationSeconds: jsii.Number(123),
//   }
//
type CfnJobDefinition_TimeoutProperty struct {
	// The time duration in seconds (measured from the job attempt's `startedAt` timestamp) after which AWS Batch terminates your jobs if they have not finished.
	//
	// The minimum value for the timeout is 60 seconds.
	AttemptDurationSeconds *float64 `field:"optional" json:"attemptDurationSeconds" yaml:"attemptDurationSeconds"`
}

