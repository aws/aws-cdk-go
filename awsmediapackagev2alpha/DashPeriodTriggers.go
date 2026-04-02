package awsmediapackagev2alpha


// Options for triggers which cause AWS Elemental MediaPackage to create media presentation description (MPD) periods in the output manifest.
// Experimental.
type DashPeriodTriggers string

const (
	// Option for AVAILS.
	// Experimental.
	DashPeriodTriggers_AVAILS DashPeriodTriggers = "AVAILS"
	// Option for DRM_KEY_ROTATION.
	// Experimental.
	DashPeriodTriggers_DRM_KEY_ROTATION DashPeriodTriggers = "DRM_KEY_ROTATION"
	// Option for SOURCE_CHANGES.
	// Experimental.
	DashPeriodTriggers_SOURCE_CHANGES DashPeriodTriggers = "SOURCE_CHANGES"
	// Option for SOURCE_DISRUPTIONS.
	// Experimental.
	DashPeriodTriggers_SOURCE_DISRUPTIONS DashPeriodTriggers = "SOURCE_DISRUPTIONS"
	// Option for NONE.
	// Experimental.
	DashPeriodTriggers_NONE DashPeriodTriggers = "NONE"
)

