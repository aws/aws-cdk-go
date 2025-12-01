package interfacesawsrtbfabric

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsrtbfabric/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a OutboundExternalLink.
// Experimental.
type IOutboundExternalLinkRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a OutboundExternalLink resource.
	// Experimental.
	OutboundExternalLinkRef() *OutboundExternalLinkReference
}

// The jsii proxy for IOutboundExternalLinkRef
type jsiiProxy_IOutboundExternalLinkRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IOutboundExternalLinkRef) OutboundExternalLinkRef() *OutboundExternalLinkReference {
	var returns *OutboundExternalLinkReference
	_jsii_.Get(
		j,
		"outboundExternalLinkRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOutboundExternalLinkRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOutboundExternalLinkRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

