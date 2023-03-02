package awscloudwatch


// Fill shading options that will be used with an annotation.
// Experimental.
type Shading string

const (
	// Don't add shading.
	// Experimental.
	Shading_NONE Shading = "NONE"
	// Add shading above the annotation.
	// Experimental.
	Shading_ABOVE Shading = "ABOVE"
	// Add shading below the annotation.
	// Experimental.
	Shading_BELOW Shading = "BELOW"
)

