package awsmediapackagev2alpha


// DRM settings keys for manifest filter configuration.
// Experimental.
type DrmSettingsKey string

const (
	// Exclude EXT-X-SESSION-KEY tags from HLS and LL-HLS multivariant playlists.
	//
	// This can improve compatibility with legacy HLS clients that have issues
	// processing session keys, or provide more granular access control when
	// using manifest filtering.
	// Experimental.
	DrmSettingsKey_EXCLUDE_SESSION_KEYS DrmSettingsKey = "EXCLUDE_SESSION_KEYS"
)

