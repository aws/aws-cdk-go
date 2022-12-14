package awspipes


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   batchJobDependencyProperty := &batchJobDependencyProperty{
//   	jobId: jsii.String("jobId"),
//   	type: jsii.String("type"),
//   }
//
type CfnPipe_BatchJobDependencyProperty struct {
	// `CfnPipe.BatchJobDependencyProperty.JobId`.
	JobId *string `field:"optional" json:"jobId" yaml:"jobId"`
	// `CfnPipe.BatchJobDependencyProperty.Type`.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

