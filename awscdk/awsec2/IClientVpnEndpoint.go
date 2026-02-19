package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsec2"
	"github.com/aws/constructs-go/constructs/v10"
)

// A client VPN endpoint.
type IClientVpnEndpoint interface {
	interfacesawsec2.IClientVpnEndpointRef
	IConnectable
	awscdk.IResource
	// The endpoint ID.
	EndpointId() *string
	// Dependable that can be depended upon to force target networks associations.
	TargetNetworksAssociated() constructs.IDependable
}

// The jsii proxy for IClientVpnEndpoint
type jsiiProxy_IClientVpnEndpoint struct {
	internal.Type__interfacesawsec2IClientVpnEndpointRef
	jsiiProxy_IConnectable
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IClientVpnEndpoint) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IClientVpnEndpoint) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IClientVpnEndpoint) EndpointId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IClientVpnEndpoint) TargetNetworksAssociated() constructs.IDependable {
	var returns constructs.IDependable
	_jsii_.Get(
		j,
		"targetNetworksAssociated",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IClientVpnEndpoint) ClientVpnEndpointRef() *interfacesawsec2.ClientVpnEndpointReference {
	var returns *interfacesawsec2.ClientVpnEndpointReference
	_jsii_.Get(
		j,
		"clientVpnEndpointRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IClientVpnEndpoint) Connections() Connections {
	var returns Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IClientVpnEndpoint) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IClientVpnEndpoint) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IClientVpnEndpoint) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

