package interfacesawsnetworkmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsnetworkmanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DirectConnectGatewayAttachment.
// Experimental.
type IDirectConnectGatewayAttachmentRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a DirectConnectGatewayAttachment resource.
	// Experimental.
	DirectConnectGatewayAttachmentRef() *DirectConnectGatewayAttachmentReference
}

// The jsii proxy for IDirectConnectGatewayAttachmentRef
type jsiiProxy_IDirectConnectGatewayAttachmentRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IDirectConnectGatewayAttachmentRef) DirectConnectGatewayAttachmentRef() *DirectConnectGatewayAttachmentReference {
	var returns *DirectConnectGatewayAttachmentReference
	_jsii_.Get(
		j,
		"directConnectGatewayAttachmentRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDirectConnectGatewayAttachmentRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDirectConnectGatewayAttachmentRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

