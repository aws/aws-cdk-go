package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A client VPN endpoint.
type IClientVpnEndpoint interface {
	IConnectable
	awscdk.IResource
	// The endpoint ID.
	EndpointId() *string
	// Dependable that can be depended upon to force target networks associations.
	TargetNetworksAssociated() constructs.IDependable
}

// The jsii proxy for IClientVpnEndpoint
type jsiiProxy_IClientVpnEndpoint struct {
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

func (j *jsiiProxy_IClientVpnEndpoint) Connections() Connections {
	var returns Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IClientVpnEndpoint) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
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

