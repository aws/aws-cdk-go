// An experiment to bundle the entire CDK into a single module
package awscdk


// Experimental.
type TagType string

const (
	// Experimental.
	TagType_STANDARD TagType = "STANDARD"
	// Experimental.
	TagType_AUTOSCALING_GROUP TagType = "AUTOSCALING_GROUP"
	// Experimental.
	TagType_MAP TagType = "MAP"
	// Experimental.
	TagType_KEY_VALUE TagType = "KEY_VALUE"
	// Experimental.
	TagType_NOT_TAGGABLE TagType = "NOT_TAGGABLE"
)

