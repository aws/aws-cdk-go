package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ClientVpnAuthorizationRule.
// Experimental.
type IClientVpnAuthorizationRuleRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ClientVpnAuthorizationRule resource.
	// Experimental.
	ClientVpnAuthorizationRuleRef() *ClientVpnAuthorizationRuleReference
}

// The jsii proxy for IClientVpnAuthorizationRuleRef
type jsiiProxy_IClientVpnAuthorizationRuleRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IClientVpnAuthorizationRuleRef) ClientVpnAuthorizationRuleRef() *ClientVpnAuthorizationRuleReference {
	var returns *ClientVpnAuthorizationRuleReference
	_jsii_.Get(
		j,
		"clientVpnAuthorizationRuleRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IClientVpnAuthorizationRuleRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IClientVpnAuthorizationRuleRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

