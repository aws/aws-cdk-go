package awselasticloadbalancingv2targets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2targets/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

// Example:
//   import lambda "github.com/aws/aws-cdk-go/awscdk"
//   import targets "github.com/aws/aws-cdk-go/awscdk"
//
//   var lambdaFunction function
//   var lb applicationLoadBalancer
//
//
//   listener := lb.AddListener(jsii.String("Listener"), &BaseApplicationListenerProps{
//   	Port: jsii.Number(80),
//   })
//   listener.AddTargets(jsii.String("Targets"), &AddApplicationTargetsProps{
//   	Targets: []iApplicationLoadBalancerTarget{
//   		targets.NewLambdaTarget(lambdaFunction),
//   	},
//
//   	// For Lambda Targets, you need to explicitly enable health checks if you
//   	// want them.
//   	HealthCheck: &HealthCheck{
//   		Enabled: jsii.Boolean(true),
//   	},
//   })
//
type LambdaTarget interface {
	awselasticloadbalancingv2.IApplicationLoadBalancerTarget
	// Register this instance target with a load balancer.
	//
	// Don't call this, it is called automatically when you add the target to a
	// load balancer.
	AttachToApplicationTargetGroup(targetGroup awselasticloadbalancingv2.IApplicationTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps
	// Register this instance target with a load balancer.
	//
	// Don't call this, it is called automatically when you add the target to a
	// load balancer.
	AttachToNetworkTargetGroup(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps
}

// The jsii proxy struct for LambdaTarget
type jsiiProxy_LambdaTarget struct {
	internal.Type__awselasticloadbalancingv2IApplicationLoadBalancerTarget
}

// Create a new Lambda target.
func NewLambdaTarget(fn awslambda.IFunction) LambdaTarget {
	_init_.Initialize()

	if err := validateNewLambdaTargetParameters(fn); err != nil {
		panic(err)
	}
	j := jsiiProxy_LambdaTarget{}

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2_targets.LambdaTarget",
		[]interface{}{fn},
		&j,
	)

	return &j
}

// Create a new Lambda target.
func NewLambdaTarget_Override(l LambdaTarget, fn awslambda.IFunction) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2_targets.LambdaTarget",
		[]interface{}{fn},
		l,
	)
}

func (l *jsiiProxy_LambdaTarget) AttachToApplicationTargetGroup(targetGroup awselasticloadbalancingv2.IApplicationTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps {
	if err := l.validateAttachToApplicationTargetGroupParameters(targetGroup); err != nil {
		panic(err)
	}
	var returns *awselasticloadbalancingv2.LoadBalancerTargetProps

	_jsii_.Invoke(
		l,
		"attachToApplicationTargetGroup",
		[]interface{}{targetGroup},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaTarget) AttachToNetworkTargetGroup(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps {
	if err := l.validateAttachToNetworkTargetGroupParameters(targetGroup); err != nil {
		panic(err)
	}
	var returns *awselasticloadbalancingv2.LoadBalancerTargetProps

	_jsii_.Invoke(
		l,
		"attachToNetworkTargetGroup",
		[]interface{}{targetGroup},
		&returns,
	)

	return returns
}

