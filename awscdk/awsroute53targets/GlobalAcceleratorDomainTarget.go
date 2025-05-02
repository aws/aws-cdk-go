package awsroute53targets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53targets/internal"
)

// Use a Global Accelerator domain name as an alias record target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var aliasRecordTargetProps iAliasRecordTargetProps
//
//   globalAcceleratorDomainTarget := awscdk.Aws_route53_targets.NewGlobalAcceleratorDomainTarget(jsii.String("acceleratorDomainName"), aliasRecordTargetProps)
//
type GlobalAcceleratorDomainTarget interface {
	awsroute53.IAliasRecordTarget
	// Return hosted zone ID and DNS name, usable for Route53 alias targets.
	Bind(_record awsroute53.IRecordSet, _zone awsroute53.IHostedZone) *awsroute53.AliasRecordTargetConfig
}

// The jsii proxy struct for GlobalAcceleratorDomainTarget
type jsiiProxy_GlobalAcceleratorDomainTarget struct {
	internal.Type__awsroute53IAliasRecordTarget
}

// Create an Alias Target for a Global Accelerator domain name.
func NewGlobalAcceleratorDomainTarget(acceleratorDomainName *string, props IAliasRecordTargetProps) GlobalAcceleratorDomainTarget {
	_init_.Initialize()

	if err := validateNewGlobalAcceleratorDomainTargetParameters(acceleratorDomainName); err != nil {
		panic(err)
	}
	j := jsiiProxy_GlobalAcceleratorDomainTarget{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53_targets.GlobalAcceleratorDomainTarget",
		[]interface{}{acceleratorDomainName, props},
		&j,
	)

	return &j
}

// Create an Alias Target for a Global Accelerator domain name.
func NewGlobalAcceleratorDomainTarget_Override(g GlobalAcceleratorDomainTarget, acceleratorDomainName *string, props IAliasRecordTargetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53_targets.GlobalAcceleratorDomainTarget",
		[]interface{}{acceleratorDomainName, props},
		g,
	)
}

func GlobalAcceleratorDomainTarget_GLOBAL_ACCELERATOR_ZONE_ID() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_route53_targets.GlobalAcceleratorDomainTarget",
		"GLOBAL_ACCELERATOR_ZONE_ID",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GlobalAcceleratorDomainTarget) Bind(_record awsroute53.IRecordSet, _zone awsroute53.IHostedZone) *awsroute53.AliasRecordTargetConfig {
	if err := g.validateBindParameters(_record); err != nil {
		panic(err)
	}
	var returns *awsroute53.AliasRecordTargetConfig

	_jsii_.Invoke(
		g,
		"bind",
		[]interface{}{_record, _zone},
		&returns,
	)

	return returns
}

