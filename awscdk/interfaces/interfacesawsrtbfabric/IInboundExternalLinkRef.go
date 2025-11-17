package interfacesawsrtbfabric

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsrtbfabric/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a InboundExternalLink.
// Experimental.
type IInboundExternalLinkRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a InboundExternalLink resource.
	// Experimental.
	InboundExternalLinkRef() *InboundExternalLinkReference
}

// The jsii proxy for IInboundExternalLinkRef
type jsiiProxy_IInboundExternalLinkRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IInboundExternalLinkRef) InboundExternalLinkRef() *InboundExternalLinkReference {
	var returns *InboundExternalLinkReference
	_jsii_.Get(
		j,
		"inboundExternalLinkRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInboundExternalLinkRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInboundExternalLinkRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

