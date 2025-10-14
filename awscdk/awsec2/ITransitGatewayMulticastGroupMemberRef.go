package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TransitGatewayMulticastGroupMember.
// Experimental.
type ITransitGatewayMulticastGroupMemberRef interface {
	constructs.IConstruct
	// A reference to a TransitGatewayMulticastGroupMember resource.
	// Experimental.
	TransitGatewayMulticastGroupMemberRef() *TransitGatewayMulticastGroupMemberReference
}

// The jsii proxy for ITransitGatewayMulticastGroupMemberRef
type jsiiProxy_ITransitGatewayMulticastGroupMemberRef struct {
	internal.Type__constructsIConstruct
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

