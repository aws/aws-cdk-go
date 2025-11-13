package interfacesawssagemaker


// A reference to a InferenceExperiment resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inferenceExperimentReference := &InferenceExperimentReference{
//   	InferenceExperimentArn: jsii.String("inferenceExperimentArn"),
//   	InferenceExperimentName: jsii.String("inferenceExperimentName"),
//   }
//
type InferenceExperimentReference struct {
	// The ARN of the InferenceExperiment resource.
	InferenceExperimentArn *string `field:"required" json:"inferenceExperimentArn" yaml:"inferenceExperimentArn"`
	// The Name of the InferenceExperiment resource.
	InferenceExperimentName *string `field:"required" json:"inferenceExperimentName" yaml:"inferenceExperimentName"`
}

