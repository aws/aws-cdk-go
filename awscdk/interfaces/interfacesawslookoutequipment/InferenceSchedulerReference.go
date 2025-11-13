package interfacesawslookoutequipment


// A reference to a InferenceScheduler resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inferenceSchedulerReference := &InferenceSchedulerReference{
//   	InferenceSchedulerArn: jsii.String("inferenceSchedulerArn"),
//   	InferenceSchedulerName: jsii.String("inferenceSchedulerName"),
//   }
//
type InferenceSchedulerReference struct {
	// The ARN of the InferenceScheduler resource.
	InferenceSchedulerArn *string `field:"required" json:"inferenceSchedulerArn" yaml:"inferenceSchedulerArn"`
	// The InferenceSchedulerName of the InferenceScheduler resource.
	InferenceSchedulerName *string `field:"required" json:"inferenceSchedulerName" yaml:"inferenceSchedulerName"`
}

