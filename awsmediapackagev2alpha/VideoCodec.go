package awsmediapackagev2alpha


// Accepted video codec values for manifest filtering.
// See: https://docs.aws.amazon.com/mediapackage/latest/userguide/manifest-filter-query-parameters.html
//
// Experimental.
type VideoCodec string

const (
	// H.264/AVC.
	// Experimental.
	VideoCodec_H264 VideoCodec = "H264"
	// H.265/HEVC.
	// Experimental.
	VideoCodec_H265 VideoCodec = "H265"
	// AV1.
	// Experimental.
	VideoCodec_AV1 VideoCodec = "AV1"
)

