package awssagemaker


// A reference to a ProcessingJob resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   processingJobReference := &ProcessingJobReference{
//   	ProcessingJobArn: jsii.String("processingJobArn"),
//   }
//
type ProcessingJobReference struct {
	// The ProcessingJobArn of the ProcessingJob resource.
	ProcessingJobArn *string `field:"required" json:"processingJobArn" yaml:"processingJobArn"`
}

