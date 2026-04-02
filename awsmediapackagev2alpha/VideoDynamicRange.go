package awsmediapackagev2alpha


// Accepted video dynamic range values for manifest filtering.
// See: https://docs.aws.amazon.com/mediapackage/latest/userguide/manifest-filter-query-parameters.html
//
// Experimental.
type VideoDynamicRange string

const (
	// Dolby Vision.
	// Experimental.
	VideoDynamicRange_DV VideoDynamicRange = "DV"
	// HDR10.
	// Experimental.
	VideoDynamicRange_HDR10 VideoDynamicRange = "HDR10"
	// Hybrid Log-Gamma.
	// Experimental.
	VideoDynamicRange_HLG VideoDynamicRange = "HLG"
	// Standard Dynamic Range.
	// Experimental.
	VideoDynamicRange_SDR VideoDynamicRange = "SDR"
)

