package awsroute53targets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancing"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53targets/internal"
)

// Use a classic ELB as an alias record target.
//
// Example:
//   import elb "github.com/aws/aws-cdk-go/awscdk"
//
//   var zone hostedZone
//   var lb loadBalancer
//
//
//   route53.NewARecord(this, jsii.String("AliasRecord"), &aRecordProps{
//   	zone: zone,
//   	target: route53.recordTarget.fromAlias(targets.NewClassicLoadBalancerTarget(lb)),
//   })
//
type ClassicLoadBalancerTarget interface {
	awsroute53.IAliasRecordTarget
	// Return hosted zone ID and DNS name, usable for Route53 alias targets.
	Bind(_record awsroute53.IRecordSet, _zone awsroute53.IHostedZone) *awsroute53.AliasRecordTargetConfig
}

// The jsii proxy struct for ClassicLoadBalancerTarget
type jsiiProxy_ClassicLoadBalancerTarget struct {
	internal.Type__awsroute53IAliasRecordTarget
}

func NewClassicLoadBalancerTarget(loadBalancer awselasticloadbalancing.LoadBalancer) ClassicLoadBalancerTarget {
	_init_.Initialize()

	if err := validateNewClassicLoadBalancerTargetParameters(loadBalancer); err != nil {
		panic(err)
	}
	j := jsiiProxy_ClassicLoadBalancerTarget{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53_targets.ClassicLoadBalancerTarget",
		[]interface{}{loadBalancer},
		&j,
	)

	return &j
}

func NewClassicLoadBalancerTarget_Override(c ClassicLoadBalancerTarget, loadBalancer awselasticloadbalancing.LoadBalancer) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53_targets.ClassicLoadBalancerTarget",
		[]interface{}{loadBalancer},
		c,
	)
}

func (c *jsiiProxy_ClassicLoadBalancerTarget) Bind(_record awsroute53.IRecordSet, _zone awsroute53.IHostedZone) *awsroute53.AliasRecordTargetConfig {
	if err := c.validateBindParameters(_record); err != nil {
		panic(err)
	}
	var returns *awsroute53.AliasRecordTargetConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{_record, _zone},
		&returns,
	)

	return returns
}

