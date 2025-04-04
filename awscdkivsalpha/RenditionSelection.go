package awscdkivsalpha


// Rendition selection mode.
// Experimental.
type RenditionSelection string

const (
	// Record all available renditions.
	// Experimental.
	RenditionSelection_ALL RenditionSelection = "ALL"
	// Does not record any video.
	//
	// This option is useful if you just want to record thumbnails.
	// Experimental.
	RenditionSelection_NONE RenditionSelection = "NONE"
	// Select a subset of video renditions to record.
	// Experimental.
	RenditionSelection_CUSTOM RenditionSelection = "CUSTOM"
)

