package awscdklocationalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdklocationalpha/v2/internal"
)

// A Geofence Collection.
// Experimental.
type IGeofenceCollection interface {
	awscdk.IResource
	// The Amazon Resource Name (ARN) of the geofence collection resource.
	// Experimental.
	GeofenceCollectionArn() *string
	// The name of the geofence collection.
	// Experimental.
	GeofenceCollectionName() *string
}

// The jsii proxy for IGeofenceCollection
type jsiiProxy_IGeofenceCollection struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IGeofenceCollection) GeofenceCollectionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"geofenceCollectionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGeofenceCollection) GeofenceCollectionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"geofenceCollectionName",
		&returns,
	)
	return returns
}

