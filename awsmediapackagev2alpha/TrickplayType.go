package awsmediapackagev2alpha


// Accepted trickplay type values for manifest filtering.
// See: https://docs.aws.amazon.com/mediapackage/latest/userguide/manifest-filter-query-parameters.html
//
// Experimental.
type TrickplayType string

const (
	// I-frame only trick-play.
	// Experimental.
	TrickplayType_IFRAME TrickplayType = "IFRAME"
	// Image-based trick-play.
	// Experimental.
	TrickplayType_IMAGE TrickplayType = "IMAGE"
	// No trick-play.
	// Experimental.
	TrickplayType_NONE TrickplayType = "NONE"
)

