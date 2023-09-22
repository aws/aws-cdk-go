package awscloudwatch


// Fill shading options that will be used with a horizontal annotation.
type Shading string

const (
	// Don't add shading.
	Shading_NONE Shading = "NONE"
	// Add shading above the annotation.
	Shading_ABOVE Shading = "ABOVE"
	// Add shading below the annotation.
	Shading_BELOW Shading = "BELOW"
)

