package awspipes


// An object that represents an AWS Batch job dependency.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   batchJobDependencyProperty := &BatchJobDependencyProperty{
//   	JobId: jsii.String("jobId"),
//   	Type: jsii.String("type"),
//   }
//
type CfnPipe_BatchJobDependencyProperty struct {
	// The job ID of the AWS Batch job that's associated with this dependency.
	JobId *string `field:"optional" json:"jobId" yaml:"jobId"`
	// The type of the job dependency.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

