package interfacesawsbedrock


// A reference to a ApplicationInferenceProfile resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   applicationInferenceProfileReference := &ApplicationInferenceProfileReference{
//   	InferenceProfileIdentifier: jsii.String("inferenceProfileIdentifier"),
//   }
//
type ApplicationInferenceProfileReference struct {
	// The InferenceProfileIdentifier of the ApplicationInferenceProfile resource.
	InferenceProfileIdentifier *string `field:"required" json:"inferenceProfileIdentifier" yaml:"inferenceProfileIdentifier"`
}

