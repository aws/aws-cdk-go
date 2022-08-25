package awsroute53targets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsroute53"
	"github.com/aws/aws-cdk-go/awscdk/awsroute53targets/internal"
)

// Use a Global Accelerator domain name as an alias record target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   globalAcceleratorDomainTarget := awscdk.Aws_route53_targets.NewGlobalAcceleratorDomainTarget(jsii.String("acceleratorDomainName"))
//
// Experimental.
type GlobalAcceleratorDomainTarget interface {
	awsroute53.IAliasRecordTarget
	// Return hosted zone ID and DNS name, usable for Route53 alias targets.
	// Experimental.
	Bind(_record awsroute53.IRecordSet, _zone awsroute53.IHostedZone) *awsroute53.AliasRecordTargetConfig
}

// The jsii proxy struct for GlobalAcceleratorDomainTarget
type jsiiProxy_GlobalAcceleratorDomainTarget struct {
	internal.Type__awsroute53IAliasRecordTarget
}

// Create an Alias Target for a Global Accelerator domain name.
// Experimental.
func NewGlobalAcceleratorDomainTarget(acceleratorDomainName *string) GlobalAcceleratorDomainTarget {
	_init_.Initialize()

	j := jsiiProxy_GlobalAcceleratorDomainTarget{}

	_jsii_.Create(
		"monocdk.aws_route53_targets.GlobalAcceleratorDomainTarget",
		[]interface{}{acceleratorDomainName},
		&j,
	)

	return &j
}

// Create an Alias Target for a Global Accelerator domain name.
// Experimental.
func NewGlobalAcceleratorDomainTarget_Override(g GlobalAcceleratorDomainTarget, acceleratorDomainName *string) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_route53_targets.GlobalAcceleratorDomainTarget",
		[]interface{}{acceleratorDomainName},
		g,
	)
}

func GlobalAcceleratorDomainTarget_GLOBAL_ACCELERATOR_ZONE_ID() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_route53_targets.GlobalAcceleratorDomainTarget",
		"GLOBAL_ACCELERATOR_ZONE_ID",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GlobalAcceleratorDomainTarget) Bind(_record awsroute53.IRecordSet, _zone awsroute53.IHostedZone) *awsroute53.AliasRecordTargetConfig {
	var returns *awsroute53.AliasRecordTargetConfig

	_jsii_.Invoke(
		g,
		"bind",
		[]interface{}{_record, _zone},
		&returns,
	)

	return returns
}

