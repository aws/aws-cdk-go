package awsvpclattice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsvpclattice/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ServiceNetworkVpcAssociation.
// Experimental.
type IServiceNetworkVpcAssociationRef interface {
	constructs.IConstruct
	// A reference to a ServiceNetworkVpcAssociation resource.
	// Experimental.
	ServiceNetworkVpcAssociationRef() *ServiceNetworkVpcAssociationReference
}

// The jsii proxy for IServiceNetworkVpcAssociationRef
type jsiiProxy_IServiceNetworkVpcAssociationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IServiceNetworkVpcAssociationRef) ServiceNetworkVpcAssociationRef() *ServiceNetworkVpcAssociationReference {
	var returns *ServiceNetworkVpcAssociationReference
	_jsii_.Get(
		j,
		"serviceNetworkVpcAssociationRef",
		&returns,
	)
	return returns
}

