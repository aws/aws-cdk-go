package awscdklocationalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Properties for a geofence collection.
//
// Example:
//   var key key
//
//
//   location.NewGeofenceCollection(this, jsii.String("GeofenceCollection"), &GeofenceCollectionProps{
//   	GeofenceCollectionName: jsii.String("MyGeofenceCollection"),
//   	 // optional, defaults to a generated name
//   	KmsKey: key,
//   })
//
// Experimental.
type GeofenceCollectionProps struct {
	// A description for the geofence collection.
	// Default: - no description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A name for the geofence collection.
	//
	// Must be between 1 and 100 characters and contain only alphanumeric characters,
	// hyphens, periods and underscores.
	// Default: - A name is automatically generated.
	//
	// Experimental.
	GeofenceCollectionName *string `field:"optional" json:"geofenceCollectionName" yaml:"geofenceCollectionName"`
	// The customer managed to encrypt your data.
	// See: https://docs.aws.amazon.com/location/latest/developerguide/encryption-at-rest.html
	//
	// Default: - Use an AWS managed key.
	//
	// Experimental.
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
}

