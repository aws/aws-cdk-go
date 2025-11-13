package interfacesawsb2bi

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsb2bi/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Partnership.
// Experimental.
type IPartnershipRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Partnership resource.
	// Experimental.
	PartnershipRef() *PartnershipReference
}

// The jsii proxy for IPartnershipRef
type jsiiProxy_IPartnershipRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IPartnershipRef) PartnershipRef() *PartnershipReference {
	var returns *PartnershipReference
	_jsii_.Get(
		j,
		"partnershipRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPartnershipRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPartnershipRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

