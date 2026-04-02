package awsmediapackagev2alpha


// DRM systems available for TS encryption.
//
// SAMPLE_AES uses FairPlay, AES_128 uses Clear Key AES 128.
// Experimental.
type TsDrmSystem string

const (
	// FairPlay DRM - used with SAMPLE_AES encryption.
	// Experimental.
	TsDrmSystem_FAIRPLAY TsDrmSystem = "FAIRPLAY"
	// Clear Key AES 128 - used with AES_128 encryption.
	// Experimental.
	TsDrmSystem_CLEAR_KEY_AES_128 TsDrmSystem = "CLEAR_KEY_AES_128"
)

