package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Interface for security group-like objects.
type ISecurityGroup interface {
	IPeer
	awscdk.IResource
	// Add an egress rule for the current security group.
	//
	// `remoteRule` controls where the Rule object is created if the peer is also a
	// securityGroup and they are in different stack. If false (default) the
	// rule object is created under the current SecurityGroup object. If true and the
	// peer is also a SecurityGroup, the rule object is created under the remote
	// SecurityGroup object.
	AddEgressRule(peer IPeer, connection Port, description *string, remoteRule *bool)
	// Add an ingress rule for the current security group.
	//
	// `remoteRule` controls where the Rule object is created if the peer is also a
	// securityGroup and they are in different stack. If false (default) the
	// rule object is created under the current SecurityGroup object. If true and the
	// peer is also a SecurityGroup, the rule object is created under the remote
	// SecurityGroup object.
	AddIngressRule(peer IPeer, connection Port, description *string, remoteRule *bool)
	// Whether the SecurityGroup has been configured to allow all outbound traffic.
	AllowAllOutbound() *bool
	// ID for the current security group.
	SecurityGroupId() *string
}

// The jsii proxy for ISecurityGroup
type jsiiProxy_ISecurityGroup struct {
	jsiiProxy_IPeer
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_ISecurityGroup) AddEgressRule(peer IPeer, connection Port, description *string, remoteRule *bool) {
	if err := i.validateAddEgressRuleParameters(peer, connection); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addEgressRule",
		[]interface{}{peer, connection, description, remoteRule},
	)
}

func (i *jsiiProxy_ISecurityGroup) AddIngressRule(peer IPeer, connection Port, description *string, remoteRule *bool) {
	if err := i.validateAddIngressRuleParameters(peer, connection); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addIngressRule",
		[]interface{}{peer, connection, description, remoteRule},
	)
}

func (i *jsiiProxy_ISecurityGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_ISecurityGroup) ToEgressRuleConfig() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toEgressRuleConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ISecurityGroup) ToIngressRuleConfig() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toIngressRuleConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_ISecurityGroup) AllowAllOutbound() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"allowAllOutbound",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISecurityGroup) SecurityGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISecurityGroup) CanInlineRule() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"canInlineRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISecurityGroup) Connections() Connections {
	var returns Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISecurityGroup) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISecurityGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISecurityGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISecurityGroup) UniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uniqueId",
		&returns,
	)
	return returns
}

