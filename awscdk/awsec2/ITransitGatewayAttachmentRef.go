package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TransitGatewayAttachment.
// Experimental.
type ITransitGatewayAttachmentRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a TransitGatewayAttachment resource.
	// Experimental.
	TransitGatewayAttachmentRef() *TransitGatewayAttachmentReference
}

// The jsii proxy for ITransitGatewayAttachmentRef
type jsiiProxy_ITransitGatewayAttachmentRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ITransitGatewayAttachmentRef) TransitGatewayAttachmentRef() *TransitGatewayAttachmentReference {
	var returns *TransitGatewayAttachmentReference
	_jsii_.Get(
		j,
		"transitGatewayAttachmentRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITransitGatewayAttachmentRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITransitGatewayAttachmentRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

