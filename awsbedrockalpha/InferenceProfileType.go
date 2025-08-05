package awsbedrockalpha


// These are the values used by the API when using aws bedrock get-inference-profile --inference-profile-identifier XXXXXXX.
// Experimental.
type InferenceProfileType string

const (
	// An inference profile that is created by AWS.
	//
	// These are profiles such as cross-region
	// which help you distribute traffic across a geographic region.
	// Experimental.
	InferenceProfileType_SYSTEM_DEFINED InferenceProfileType = "SYSTEM_DEFINED"
	// An inference profile that is user-created.
	//
	// These are profiles that help
	// you track costs or metrics.
	// Experimental.
	InferenceProfileType_APPLICATION InferenceProfileType = "APPLICATION"
)

