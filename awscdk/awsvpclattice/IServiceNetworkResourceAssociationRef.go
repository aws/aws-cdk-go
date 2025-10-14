package awsvpclattice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsvpclattice/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ServiceNetworkResourceAssociation.
// Experimental.
type IServiceNetworkResourceAssociationRef interface {
	constructs.IConstruct
	// A reference to a ServiceNetworkResourceAssociation resource.
	// Experimental.
	ServiceNetworkResourceAssociationRef() *ServiceNetworkResourceAssociationReference
}

// The jsii proxy for IServiceNetworkResourceAssociationRef
type jsiiProxy_IServiceNetworkResourceAssociationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IServiceNetworkResourceAssociationRef) ServiceNetworkResourceAssociationRef() *ServiceNetworkResourceAssociationReference {
	var returns *ServiceNetworkResourceAssociationReference
	_jsii_.Get(
		j,
		"serviceNetworkResourceAssociationRef",
		&returns,
	)
	return returns
}

