package awsbedrockalpha


// Attributes for specifying an imported Application Inference Profile.
//
// Example:
//   // Import an inference profile through attributes
//   importedProfile := bedrock.ApplicationInferenceProfile_FromApplicationInferenceProfileAttributes(this, jsii.String("ImportedProfile"), &ApplicationInferenceProfileAttributes{
//   	InferenceProfileArn: jsii.String("arn:aws:bedrock:us-east-1:123456789012:application-inference-profile/my-profile-id"),
//   	InferenceProfileIdentifier: jsii.String("my-profile-id"),
//   })
//
// Experimental.
type ApplicationInferenceProfileAttributes struct {
	// The ARN of the application inference profile.
	// Experimental.
	InferenceProfileArn *string `field:"required" json:"inferenceProfileArn" yaml:"inferenceProfileArn"`
	// The ID or Amazon Resource Name (ARN) of the inference profile.
	//
	// This can be either the profile ID or the full ARN.
	// Experimental.
	InferenceProfileIdentifier *string `field:"required" json:"inferenceProfileIdentifier" yaml:"inferenceProfileIdentifier"`
}

