package awscdkredshiftalpha


// The maintenance track for the cluster.
// See: https://docs.aws.amazon.com/redshift/latest/mgmt/managing-cluster-considerations.html#rs-mgmt-maintenance-tracks
//
// Experimental.
type MaintenanceTrackName string

const (
	// Updated to the most recently certified maintenance release.
	// Experimental.
	MaintenanceTrackName_CURRENT MaintenanceTrackName = "CURRENT"
	// Update to the previously certified maintenance release.
	// Experimental.
	MaintenanceTrackName_TRAILING MaintenanceTrackName = "TRAILING"
)

