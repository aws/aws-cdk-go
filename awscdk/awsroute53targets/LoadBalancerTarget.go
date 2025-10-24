package awsroute53targets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53targets/internal"
)

// Use an ELBv2 as an alias record target.
//
// Example:
//   import elbv2 "github.com/aws/aws-cdk-go/awscdk"
//
//   var zone HostedZone
//   var lb ApplicationLoadBalancer
//
//
//   route53.NewARecord(this, jsii.String("AliasRecord"), &ARecordProps{
//   	Zone: Zone,
//   	Target: route53.RecordTarget_FromAlias(
//   	targets.NewLoadBalancerTarget(lb, map[string]*bool{
//   		"evaluateTargetHealth": jsii.Boolean(true),
//   	})),
//   })
//
type LoadBalancerTarget interface {
	awsroute53.IAliasRecordTarget
	// Return hosted zone ID and DNS name, usable for Route53 alias targets.
	Bind(_record awsroute53.IRecordSet, _zone awsroute53.IHostedZone) *awsroute53.AliasRecordTargetConfig
}

// The jsii proxy struct for LoadBalancerTarget
type jsiiProxy_LoadBalancerTarget struct {
	internal.Type__awsroute53IAliasRecordTarget
}

func NewLoadBalancerTarget(loadBalancer awselasticloadbalancingv2.ILoadBalancerV2, props IAliasRecordTargetProps) LoadBalancerTarget {
	_init_.Initialize()

	if err := validateNewLoadBalancerTargetParameters(loadBalancer); err != nil {
		panic(err)
	}
	j := jsiiProxy_LoadBalancerTarget{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53_targets.LoadBalancerTarget",
		[]interface{}{loadBalancer, props},
		&j,
	)

	return &j
}

func NewLoadBalancerTarget_Override(l LoadBalancerTarget, loadBalancer awselasticloadbalancingv2.ILoadBalancerV2, props IAliasRecordTargetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53_targets.LoadBalancerTarget",
		[]interface{}{loadBalancer, props},
		l,
	)
}

func (l *jsiiProxy_LoadBalancerTarget) Bind(_record awsroute53.IRecordSet, _zone awsroute53.IHostedZone) *awsroute53.AliasRecordTargetConfig {
	if err := l.validateBindParameters(_record); err != nil {
		panic(err)
	}
	var returns *awsroute53.AliasRecordTargetConfig

	_jsii_.Invoke(
		l,
		"bind",
		[]interface{}{_record, _zone},
		&returns,
	)

	return returns
}

