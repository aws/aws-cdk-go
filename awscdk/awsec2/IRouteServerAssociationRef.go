package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RouteServerAssociation.
// Experimental.
type IRouteServerAssociationRef interface {
	constructs.IConstruct
	// A reference to a RouteServerAssociation resource.
	// Experimental.
	RouteServerAssociationRef() *RouteServerAssociationReference
}

// The jsii proxy for IRouteServerAssociationRef
type jsiiProxy_IRouteServerAssociationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IRouteServerAssociationRef) RouteServerAssociationRef() *RouteServerAssociationReference {
	var returns *RouteServerAssociationReference
	_jsii_.Get(
		j,
		"routeServerAssociationRef",
		&returns,
	)
	return returns
}

