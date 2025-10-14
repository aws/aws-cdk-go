package awsneptunegraph

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsneptunegraph/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PrivateGraphEndpoint.
// Experimental.
type IPrivateGraphEndpointRef interface {
	constructs.IConstruct
	// A reference to a PrivateGraphEndpoint resource.
	// Experimental.
	PrivateGraphEndpointRef() *PrivateGraphEndpointReference
}

// The jsii proxy for IPrivateGraphEndpointRef
type jsiiProxy_IPrivateGraphEndpointRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IPrivateGraphEndpointRef) PrivateGraphEndpointRef() *PrivateGraphEndpointReference {
	var returns *PrivateGraphEndpointReference
	_jsii_.Get(
		j,
		"privateGraphEndpointRef",
		&returns,
	)
	return returns
}

