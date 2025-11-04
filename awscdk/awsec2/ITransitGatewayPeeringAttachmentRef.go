package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TransitGatewayPeeringAttachment.
// Experimental.
type ITransitGatewayPeeringAttachmentRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a TransitGatewayPeeringAttachment resource.
	// Experimental.
	TransitGatewayPeeringAttachmentRef() *TransitGatewayPeeringAttachmentReference
}

// The jsii proxy for ITransitGatewayPeeringAttachmentRef
type jsiiProxy_ITransitGatewayPeeringAttachmentRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ITransitGatewayPeeringAttachmentRef) TransitGatewayPeeringAttachmentRef() *TransitGatewayPeeringAttachmentReference {
	var returns *TransitGatewayPeeringAttachmentReference
	_jsii_.Get(
		j,
		"transitGatewayPeeringAttachmentRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITransitGatewayPeeringAttachmentRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITransitGatewayPeeringAttachmentRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

