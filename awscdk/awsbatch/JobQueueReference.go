package awsbatch


// A reference to a JobQueue resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   jobQueueReference := &JobQueueReference{
//   	JobQueueArn: jsii.String("jobQueueArn"),
//   }
//
type JobQueueReference struct {
	// The JobQueueArn of the JobQueue resource.
	JobQueueArn *string `field:"required" json:"jobQueueArn" yaml:"jobQueueArn"`
}

