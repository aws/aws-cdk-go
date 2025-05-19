package awscloudwatch


// Fill shading options that will be used with a vertical annotation.
type VerticalShading string

const (
	// Don't add shading.
	VerticalShading_NONE VerticalShading = "NONE"
	// Add shading before the annotation.
	VerticalShading_BEFORE VerticalShading = "BEFORE"
	// Add shading after the annotation.
	VerticalShading_AFTER VerticalShading = "AFTER"
)

