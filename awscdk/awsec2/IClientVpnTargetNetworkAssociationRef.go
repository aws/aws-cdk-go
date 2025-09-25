package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ClientVpnTargetNetworkAssociation.
// Experimental.
type IClientVpnTargetNetworkAssociationRef interface {
	constructs.IConstruct
	// A reference to a ClientVpnTargetNetworkAssociation resource.
	// Experimental.
	ClientVpnTargetNetworkAssociationRef() *ClientVpnTargetNetworkAssociationReference
}

// The jsii proxy for IClientVpnTargetNetworkAssociationRef
type jsiiProxy_IClientVpnTargetNetworkAssociationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IClientVpnTargetNetworkAssociationRef) ClientVpnTargetNetworkAssociationRef() *ClientVpnTargetNetworkAssociationReference {
	var returns *ClientVpnTargetNetworkAssociationReference
	_jsii_.Get(
		j,
		"clientVpnTargetNetworkAssociationRef",
		&returns,
	)
	return returns
}

