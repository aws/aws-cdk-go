package awsstepfunctionstasks


// An object representing an AWS Batch job dependency.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   batchJobDependency := &BatchJobDependency{
//   	JobId: jsii.String("jobId"),
//   	Type: jsii.String("type"),
//   }
//
// Experimental.
type BatchJobDependency struct {
	// The job ID of the AWS Batch job associated with this dependency.
	// Experimental.
	JobId *string `field:"optional" json:"jobId" yaml:"jobId"`
	// The type of the job dependency.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

