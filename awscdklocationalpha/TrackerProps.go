package awscdklocationalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Properties for a tracker.
//
// Example:
//   var role role
//
//
//   tracker := location.NewTracker(this, jsii.String("Tracker"), &TrackerProps{
//   	TrackerName: jsii.String("MyTracker"),
//   })
//
//   tracker.GrantRead(role)
//
// Experimental.
type TrackerProps struct {
	// A description for the tracker.
	// Default: - no description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Send filtered device position updates to default EventBridge bus.
	// Default: false.
	//
	// Experimental.
	EventBridgeEnabled *bool `field:"optional" json:"eventBridgeEnabled" yaml:"eventBridgeEnabled"`
	// An optional list of geofence collections to associate with the tracker resource.
	// Default: - no geofence collections are associated.
	//
	// Experimental.
	GeofenceCollections *[]IGeofenceCollection `field:"optional" json:"geofenceCollections" yaml:"geofenceCollections"`
	// The customer managed key to encrypt data.
	//
	// If you set customer managed key, the Bounding Polygon Queries feature will be disabled by default.
	// You can choose to opt-in to the Bounding Polygon Queries feature by setting the kmsKeyEnableGeospatialQueries parameter to true.
	// Default: - Use an AWS managed key.
	//
	// Experimental.
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// Whether to opt-in to the Bounding Polygon Queries feature with customer managed key.
	// Default: false.
	//
	// Experimental.
	KmsKeyEnableGeospatialQueries *bool `field:"optional" json:"kmsKeyEnableGeospatialQueries" yaml:"kmsKeyEnableGeospatialQueries"`
	// The position filtering for the tracker resource.
	// Default: PositionFiltering.TIME_BASED
	//
	// Experimental.
	PositionFiltering PositionFiltering `field:"optional" json:"positionFiltering" yaml:"positionFiltering"`
	// A name for the tracker.
	//
	// Must be between 1 and 100 characters and contain only alphanumeric characters,
	// hyphens, periods and underscores.
	// Default: - A name is automatically generated.
	//
	// Experimental.
	TrackerName *string `field:"optional" json:"trackerName" yaml:"trackerName"`
}

