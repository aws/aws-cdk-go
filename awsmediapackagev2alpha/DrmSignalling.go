package awsmediapackagev2alpha


// DRM signaling determines the way DASH manifest signals the DRM content.
// Experimental.
type DrmSignalling string

const (
	// Option for INDIVIDUAL.
	// Experimental.
	DrmSignalling_INDIVIDUAL DrmSignalling = "INDIVIDUAL"
	// Option for REFERENCED.
	// Experimental.
	DrmSignalling_REFERENCED DrmSignalling = "REFERENCED"
)

