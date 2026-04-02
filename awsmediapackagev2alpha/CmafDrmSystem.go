package awsmediapackagev2alpha


// DRM systems available for CMAF encryption.
//
// CENC supports PlayReady, Widevine, and Irdeto.
// CBCS supports PlayReady, Widevine, and FairPlay.
// Experimental.
type CmafDrmSystem string

const (
	// FairPlay DRM - used with CBCS encryption.
	// Experimental.
	CmafDrmSystem_FAIRPLAY CmafDrmSystem = "FAIRPLAY"
	// PlayReady DRM - used with CENC and CBCS encryption.
	// Experimental.
	CmafDrmSystem_PLAYREADY CmafDrmSystem = "PLAYREADY"
	// Widevine DRM - used with CENC and CBCS encryption.
	// Experimental.
	CmafDrmSystem_WIDEVINE CmafDrmSystem = "WIDEVINE"
	// Irdeto DRM - used with CENC encryption.
	// Experimental.
	CmafDrmSystem_IRDETO CmafDrmSystem = "IRDETO"
)

