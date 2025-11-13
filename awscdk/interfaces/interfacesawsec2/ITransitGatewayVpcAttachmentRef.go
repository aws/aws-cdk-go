package interfacesawsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TransitGatewayVpcAttachment.
// Experimental.
type ITransitGatewayVpcAttachmentRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a TransitGatewayVpcAttachment resource.
	// Experimental.
	TransitGatewayVpcAttachmentRef() *TransitGatewayVpcAttachmentReference
}

// The jsii proxy for ITransitGatewayVpcAttachmentRef
type jsiiProxy_ITransitGatewayVpcAttachmentRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_ITransitGatewayVpcAttachmentRef) TransitGatewayVpcAttachmentRef() *TransitGatewayVpcAttachmentReference {
	var returns *TransitGatewayVpcAttachmentReference
	_jsii_.Get(
		j,
		"transitGatewayVpcAttachmentRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITransitGatewayVpcAttachmentRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITransitGatewayVpcAttachmentRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

