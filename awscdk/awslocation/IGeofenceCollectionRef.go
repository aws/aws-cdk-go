package awslocation

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslocation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a GeofenceCollection.
// Experimental.
type IGeofenceCollectionRef interface {
	constructs.IConstruct
}

// The jsii proxy for IGeofenceCollectionRef
type jsiiProxy_IGeofenceCollectionRef struct {
	internal.Type__constructsIConstruct
}

