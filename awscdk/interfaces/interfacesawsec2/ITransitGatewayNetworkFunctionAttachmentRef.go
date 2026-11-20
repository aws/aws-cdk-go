package interfacesawsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TransitGatewayNetworkFunctionAttachment.
// Experimental.
type ITransitGatewayNetworkFunctionAttachmentRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a TransitGatewayNetworkFunctionAttachment resource.
	// Experimental.
	TransitGatewayNetworkFunctionAttachmentRef() *TransitGatewayNetworkFunctionAttachmentReference
}

// The jsii proxy for ITransitGatewayNetworkFunctionAttachmentRef
type jsiiProxy_ITransitGatewayNetworkFunctionAttachmentRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ITransitGatewayNetworkFunctionAttachmentRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_ITransitGatewayNetworkFunctionAttachmentRef) TransitGatewayNetworkFunctionAttachmentRef() *TransitGatewayNetworkFunctionAttachmentReference {
	var returns *TransitGatewayNetworkFunctionAttachmentReference
	_jsii_.Get(
		j,
		"transitGatewayNetworkFunctionAttachmentRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITransitGatewayNetworkFunctionAttachmentRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITransitGatewayNetworkFunctionAttachmentRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

