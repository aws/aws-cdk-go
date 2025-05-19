package awselasticloadbalancingv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Properties to reference an existing listener.
type IApplicationListener interface {
	awsec2.IConnectable
	IListener
	// Perform the given action on incoming requests.
	//
	// This allows full control of the default action of the load balancer,
	// including Action chaining, fixed responses and redirect responses. See
	// the `ListenerAction` class for all options.
	//
	// It's possible to add routing conditions to the Action added in this way.
	//
	// It is not possible to add a default action to an imported IApplicationListener.
	// In order to add actions to an imported IApplicationListener a `priority`
	// must be provided.
	AddAction(id *string, props *AddApplicationActionProps)
	// Add one or more certificates to this listener.
	AddCertificates(id *string, certificates *[]IListenerCertificate)
	// Load balance incoming requests to the given target groups.
	//
	// It's possible to add conditions to the TargetGroups added in this way.
	// At least one TargetGroup must be added without conditions.
	AddTargetGroups(id *string, props *AddApplicationTargetGroupsProps)
	// Load balance incoming requests to the given load balancing targets.
	//
	// This method implicitly creates an ApplicationTargetGroup for the targets
	// involved.
	//
	// It's possible to add conditions to the targets added in this way. At least
	// one set of targets must be added without conditions.
	//
	// Returns: The newly created target group.
	AddTargets(id *string, props *AddApplicationTargetsProps) ApplicationTargetGroup
	// Register that a connectable that has been added to this load balancer.
	//
	// Don't call this directly. It is called by ApplicationTargetGroup.
	RegisterConnectable(connectable awsec2.IConnectable, portRange awsec2.Port)
}

// The jsii proxy for IApplicationListener
type jsiiProxy_IApplicationListener struct {
	internal.Type__awsec2IConnectable
	jsiiProxy_IListener
}

func (i *jsiiProxy_IApplicationListener) AddAction(id *string, props *AddApplicationActionProps) {
	if err := i.validateAddActionParameters(id, props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addAction",
		[]interface{}{id, props},
	)
}

func (i *jsiiProxy_IApplicationListener) AddCertificates(id *string, certificates *[]IListenerCertificate) {
	if err := i.validateAddCertificatesParameters(id, certificates); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addCertificates",
		[]interface{}{id, certificates},
	)
}

func (i *jsiiProxy_IApplicationListener) AddTargetGroups(id *string, props *AddApplicationTargetGroupsProps) {
	if err := i.validateAddTargetGroupsParameters(id, props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addTargetGroups",
		[]interface{}{id, props},
	)
}

func (i *jsiiProxy_IApplicationListener) AddTargets(id *string, props *AddApplicationTargetsProps) ApplicationTargetGroup {
	if err := i.validateAddTargetsParameters(id, props); err != nil {
		panic(err)
	}
	var returns ApplicationTargetGroup

	_jsii_.Invoke(
		i,
		"addTargets",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IApplicationListener) RegisterConnectable(connectable awsec2.IConnectable, portRange awsec2.Port) {
	if err := i.validateRegisterConnectableParameters(connectable, portRange); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"registerConnectable",
		[]interface{}{connectable, portRange},
	)
}

func (i *jsiiProxy_IApplicationListener) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_IApplicationListener) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplicationListener) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplicationListener) ListenerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"listenerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplicationListener) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplicationListener) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

