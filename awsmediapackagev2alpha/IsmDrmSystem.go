package awsmediapackagev2alpha


// DRM systems available for ISM encryption.
//
// ISM only supports PlayReady with CENC.
// Experimental.
type IsmDrmSystem string

const (
	// PlayReady DRM.
	// Experimental.
	IsmDrmSystem_PLAYREADY IsmDrmSystem = "PLAYREADY"
)

