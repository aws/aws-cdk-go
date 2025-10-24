package awsroute53targets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsglobalaccelerator"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
)

// Use a Global Accelerator instance domain name as an alias record target.
//
// Example:
//   import globalaccelerator "github.com/aws/aws-cdk-go/awscdk"
//
//   var zone HostedZone
//   var accelerator Accelerator
//
//
//   route53.NewARecord(this, jsii.String("AliasRecord"), &ARecordProps{
//   	Zone: Zone,
//   	Target: route53.RecordTarget_FromAlias(
//   	targets.NewGlobalAcceleratorTarget(accelerator, map[string]*bool{
//   		"evaluateTargetHealth": jsii.Boolean(true),
//   	})),
//   })
//
type GlobalAcceleratorTarget interface {
	GlobalAcceleratorDomainTarget
	// Return hosted zone ID and DNS name, usable for Route53 alias targets.
	Bind(_record awsroute53.IRecordSet, _zone awsroute53.IHostedZone) *awsroute53.AliasRecordTargetConfig
}

// The jsii proxy struct for GlobalAcceleratorTarget
type jsiiProxy_GlobalAcceleratorTarget struct {
	jsiiProxy_GlobalAcceleratorDomainTarget
}

// Create an Alias Target for a Global Accelerator instance.
func NewGlobalAcceleratorTarget(accelerator awsglobalaccelerator.IAccelerator, props IAliasRecordTargetProps) GlobalAcceleratorTarget {
	_init_.Initialize()

	if err := validateNewGlobalAcceleratorTargetParameters(accelerator); err != nil {
		panic(err)
	}
	j := jsiiProxy_GlobalAcceleratorTarget{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53_targets.GlobalAcceleratorTarget",
		[]interface{}{accelerator, props},
		&j,
	)

	return &j
}

// Create an Alias Target for a Global Accelerator instance.
func NewGlobalAcceleratorTarget_Override(g GlobalAcceleratorTarget, accelerator awsglobalaccelerator.IAccelerator, props IAliasRecordTargetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53_targets.GlobalAcceleratorTarget",
		[]interface{}{accelerator, props},
		g,
	)
}

func GlobalAcceleratorTarget_GLOBAL_ACCELERATOR_ZONE_ID() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_route53_targets.GlobalAcceleratorTarget",
		"GLOBAL_ACCELERATOR_ZONE_ID",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GlobalAcceleratorTarget) Bind(_record awsroute53.IRecordSet, _zone awsroute53.IHostedZone) *awsroute53.AliasRecordTargetConfig {
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

