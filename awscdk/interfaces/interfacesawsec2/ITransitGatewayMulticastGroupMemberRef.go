package interfacesawsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TransitGatewayMulticastGroupMember.
// Experimental.
type ITransitGatewayMulticastGroupMemberRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a TransitGatewayMulticastGroupMember resource.
	// Experimental.
	TransitGatewayMulticastGroupMemberRef() *TransitGatewayMulticastGroupMemberReference
}

// The jsii proxy for ITransitGatewayMulticastGroupMemberRef
type jsiiProxy_ITransitGatewayMulticastGroupMemberRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_ITransitGatewayMulticastGroupMemberRef) TransitGatewayMulticastGroupMemberRef() *TransitGatewayMulticastGroupMemberReference {
	var returns *TransitGatewayMulticastGroupMemberReference
	_jsii_.Get(
		j,
		"transitGatewayMulticastGroupMemberRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITransitGatewayMulticastGroupMemberRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITransitGatewayMulticastGroupMemberRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

