package awsroute53targets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awselasticloadbalancing"
	"github.com/aws/aws-cdk-go/awscdk/awsroute53"
	"github.com/aws/aws-cdk-go/awscdk/awsroute53targets/internal"
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
//   route53.NewARecord(this, jsii.String("AliasRecord"), &ARecordProps{
//   	Zone: Zone,
//   	Target: route53.RecordTarget_FromAlias(targets.NewClassicLoadBalancerTarget(lb)),
//   })
//
// Experimental.
type ClassicLoadBalancerTarget interface {
	awsroute53.IAliasRecordTarget
	// Return hosted zone ID and DNS name, usable for Route53 alias targets.
	// Experimental.
	Bind(_record awsroute53.IRecordSet, _zone awsroute53.IHostedZone) *awsroute53.AliasRecordTargetConfig
}

// The jsii proxy struct for ClassicLoadBalancerTarget
type jsiiProxy_ClassicLoadBalancerTarget struct {
	internal.Type__awsroute53IAliasRecordTarget
}

// Experimental.
func NewClassicLoadBalancerTarget(loadBalancer awselasticloadbalancing.LoadBalancer) ClassicLoadBalancerTarget {
	_init_.Initialize()

	if err := validateNewClassicLoadBalancerTargetParameters(loadBalancer); err != nil {
		panic(err)
	}
	j := jsiiProxy_ClassicLoadBalancerTarget{}

	_jsii_.Create(
		"monocdk.aws_route53_targets.ClassicLoadBalancerTarget",
		[]interface{}{loadBalancer},
		&j,
	)

	return &j
}

// Experimental.
func NewClassicLoadBalancerTarget_Override(c ClassicLoadBalancerTarget, loadBalancer awselasticloadbalancing.LoadBalancer) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_route53_targets.ClassicLoadBalancerTarget",
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

