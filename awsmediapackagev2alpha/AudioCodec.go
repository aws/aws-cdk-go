package awsmediapackagev2alpha


// Accepted audio codec values for manifest filtering.
// See: https://docs.aws.amazon.com/mediapackage/latest/userguide/manifest-filter-query-parameters.html
//
// Experimental.
type AudioCodec string

const (
	// AAC-LC audio codec.
	// Experimental.
	AudioCodec_AACL AudioCodec = "AACL"
	// HE-AAC audio codec.
	// Experimental.
	AudioCodec_AACH AudioCodec = "AACH"
	// Dolby Digital audio codec (include the hyphen).
	// Experimental.
	AudioCodec_AC_3 AudioCodec = "AC_3"
	// Dolby Digital Plus audio codec (include the hyphen).
	// Experimental.
	AudioCodec_EC_3 AudioCodec = "EC_3"
)

