package awsroute53targets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awselasticloadbalancingv2"
	"github.com/aws/aws-cdk-go/awscdk/awsroute53"
	"github.com/aws/aws-cdk-go/awscdk/awsroute53targets/internal"
)

// Use an ELBv2 as an alias record target.
//
// Example:
//   import elbv2 "github.com/aws/aws-cdk-go/awscdk"
//
//   var zone hostedZone
//   var lb applicationLoadBalancer
//
//
//   route53.NewARecord(this, jsii.String("AliasRecord"), &ARecordProps{
//   	Zone: Zone,
//   	Target: route53.RecordTarget_FromAlias(targets.NewLoadBalancerTarget(lb)),
//   })
//
// Experimental.
type LoadBalancerTarget interface {
	awsroute53.IAliasRecordTarget
	// Return hosted zone ID and DNS name, usable for Route53 alias targets.
	// Experimental.
	Bind(_record awsroute53.IRecordSet, _zone awsroute53.IHostedZone) *awsroute53.AliasRecordTargetConfig
}

// The jsii proxy struct for LoadBalancerTarget
type jsiiProxy_LoadBalancerTarget struct {
	internal.Type__awsroute53IAliasRecordTarget
}

// Experimental.
func NewLoadBalancerTarget(loadBalancer awselasticloadbalancingv2.ILoadBalancerV2) LoadBalancerTarget {
	_init_.Initialize()

	if err := validateNewLoadBalancerTargetParameters(loadBalancer); err != nil {
		panic(err)
	}
	j := jsiiProxy_LoadBalancerTarget{}

	_jsii_.Create(
		"monocdk.aws_route53_targets.LoadBalancerTarget",
		[]interface{}{loadBalancer},
		&j,
	)

	return &j
}

// Experimental.
func NewLoadBalancerTarget_Override(l LoadBalancerTarget, loadBalancer awselasticloadbalancingv2.ILoadBalancerV2) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_route53_targets.LoadBalancerTarget",
		[]interface{}{loadBalancer},
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

