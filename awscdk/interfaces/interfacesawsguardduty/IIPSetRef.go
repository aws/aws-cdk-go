package interfacesawsguardduty

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsguardduty/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a IPSet.
// Experimental.
type IIPSetRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a IPSet resource.
	// Experimental.
	IpSetRef() *IPSetReference
}

// The jsii proxy for IIPSetRef
type jsiiProxy_IIPSetRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IIPSetRef) IpSetRef() *IPSetReference {
	var returns *IPSetReference
	_jsii_.Get(
		j,
		"ipSetRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIPSetRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIPSetRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

