package awsstepfunctionstasks


// An object representing an AWS Batch job dependency.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   jobDependency := &jobDependency{
//   	jobId: jsii.String("jobId"),
//   	type: jsii.String("type"),
//   }
//
type JobDependency struct {
	// The job ID of the AWS Batch job associated with this dependency.
	JobId *string `field:"optional" json:"jobId" yaml:"jobId"`
	// The type of the job dependency.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

