package interfacesawssagemaker


// A reference to a InferenceComponent resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inferenceComponentReference := &InferenceComponentReference{
//   	InferenceComponentArn: jsii.String("inferenceComponentArn"),
//   }
//
type InferenceComponentReference struct {
	// The InferenceComponentArn of the InferenceComponent resource.
	InferenceComponentArn *string `field:"required" json:"inferenceComponentArn" yaml:"inferenceComponentArn"`
}

