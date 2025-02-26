package awselasticloadbalancingv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/constructs-go/constructs/v10"
)

// A Target Group for Application Load Balancers.
type IApplicationTargetGroup interface {
	ITargetGroup
	// Add a load balancing target to this target group.
	AddTarget(targets ...IApplicationLoadBalancerTarget)
	// Register a connectable as a member of this target group.
	//
	// Don't call this directly. It will be called by load balancing targets.
	RegisterConnectable(connectable awsec2.IConnectable, portRange awsec2.Port)
	// Register a listener that is load balancing to this target group.
	//
	// Don't call this directly. It will be called by listeners.
	RegisterListener(listener IApplicationListener, associatingConstruct constructs.IConstruct)
	// All metrics available for this target group.
	Metrics() IApplicationTargetGroupMetrics
}

// The jsii proxy for IApplicationTargetGroup
type jsiiProxy_IApplicationTargetGroup struct {
	jsiiProxy_ITargetGroup
}

func (i *jsiiProxy_IApplicationTargetGroup) AddTarget(targets ...IApplicationLoadBalancerTarget) {
	args := []interface{}{}
	for _, a := range targets {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		i,
		"addTarget",
		args,
	)
}

func (i *jsiiProxy_IApplicationTargetGroup) RegisterConnectable(connectable awsec2.IConnectable, portRange awsec2.Port) {
	if err := i.validateRegisterConnectableParameters(connectable); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"registerConnectable",
		[]interface{}{connectable, portRange},
	)
}

func (i *jsiiProxy_IApplicationTargetGroup) RegisterListener(listener IApplicationListener, associatingConstruct constructs.IConstruct) {
	if err := i.validateRegisterListenerParameters(listener); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"registerListener",
		[]interface{}{listener, associatingConstruct},
	)
}

func (j *jsiiProxy_IApplicationTargetGroup) Metrics() IApplicationTargetGroupMetrics {
	var returns IApplicationTargetGroupMetrics
	_jsii_.Get(
		j,
		"metrics",
		&returns,
	)
	return returns
}

