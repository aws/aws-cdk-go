package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ClientVpnTargetNetworkAssociation.
// Experimental.
type IClientVpnTargetNetworkAssociationRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ClientVpnTargetNetworkAssociation resource.
	// Experimental.
	ClientVpnTargetNetworkAssociationRef() *ClientVpnTargetNetworkAssociationReference
}

// The jsii proxy for IClientVpnTargetNetworkAssociationRef
type jsiiProxy_IClientVpnTargetNetworkAssociationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IClientVpnTargetNetworkAssociationRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IClientVpnTargetNetworkAssociationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

