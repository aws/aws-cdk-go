package awslocation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslocation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a GeofenceCollection.
// Experimental.
type IGeofenceCollectionRef interface {
	constructs.IConstruct
	// A reference to a GeofenceCollection resource.
	// Experimental.
	GeofenceCollectionRef() *GeofenceCollectionReference
}

// The jsii proxy for IGeofenceCollectionRef
type jsiiProxy_IGeofenceCollectionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IGeofenceCollectionRef) GeofenceCollectionRef() *GeofenceCollectionReference {
	var returns *GeofenceCollectionReference
	_jsii_.Get(
		j,
		"geofenceCollectionRef",
		&returns,
	)
	return returns
}

