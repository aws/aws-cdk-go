package interfacesawsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VerifiedAccessEndpoint.
// Experimental.
type IVerifiedAccessEndpointRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a VerifiedAccessEndpoint resource.
	// Experimental.
	VerifiedAccessEndpointRef() *VerifiedAccessEndpointReference
}

// The jsii proxy for IVerifiedAccessEndpointRef
type jsiiProxy_IVerifiedAccessEndpointRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IVerifiedAccessEndpointRef) VerifiedAccessEndpointRef() *VerifiedAccessEndpointReference {
	var returns *VerifiedAccessEndpointReference
	_jsii_.Get(
		j,
		"verifiedAccessEndpointRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVerifiedAccessEndpointRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVerifiedAccessEndpointRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

