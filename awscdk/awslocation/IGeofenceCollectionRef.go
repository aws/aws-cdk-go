package awslocation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslocation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a GeofenceCollection.
// Experimental.
type IGeofenceCollectionRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a GeofenceCollection resource.
	// Experimental.
	GeofenceCollectionRef() *GeofenceCollectionReference
}

// The jsii proxy for IGeofenceCollectionRef
type jsiiProxy_IGeofenceCollectionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IGeofenceCollectionRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGeofenceCollectionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

