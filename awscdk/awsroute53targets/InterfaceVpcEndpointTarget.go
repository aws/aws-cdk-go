package awsroute53targets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53targets/internal"
)

// Set an InterfaceVpcEndpoint as a target for an ARecord.
//
// Example:
//   import ec2 "github.com/aws/aws-cdk-go/awscdk"
//
//   var zone hostedZone
//   var interfaceVpcEndpoint interfaceVpcEndpoint
//
//
//   route53.NewARecord(this, jsii.String("AliasRecord"), &ARecordProps{
//   	Zone: Zone,
//   	Target: route53.RecordTarget_FromAlias(targets.NewInterfaceVpcEndpointTarget(interfaceVpcEndpoint)),
//   })
//
type InterfaceVpcEndpointTarget interface {
	awsroute53.IAliasRecordTarget
	// Return hosted zone ID and DNS name, usable for Route53 alias targets.
	Bind(_record awsroute53.IRecordSet, _zone awsroute53.IHostedZone) *awsroute53.AliasRecordTargetConfig
}

// The jsii proxy struct for InterfaceVpcEndpointTarget
type jsiiProxy_InterfaceVpcEndpointTarget struct {
	internal.Type__awsroute53IAliasRecordTarget
}

func NewInterfaceVpcEndpointTarget(vpcEndpoint awsec2.InterfaceVpcEndpoint) InterfaceVpcEndpointTarget {
	_init_.Initialize()

	if err := validateNewInterfaceVpcEndpointTargetParameters(vpcEndpoint); err != nil {
		panic(err)
	}
	j := jsiiProxy_InterfaceVpcEndpointTarget{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53_targets.InterfaceVpcEndpointTarget",
		[]interface{}{vpcEndpoint},
		&j,
	)

	return &j
}

func NewInterfaceVpcEndpointTarget_Override(i InterfaceVpcEndpointTarget, vpcEndpoint awsec2.InterfaceVpcEndpoint) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53_targets.InterfaceVpcEndpointTarget",
		[]interface{}{vpcEndpoint},
		i,
	)
}

func (i *jsiiProxy_InterfaceVpcEndpointTarget) Bind(_record awsroute53.IRecordSet, _zone awsroute53.IHostedZone) *awsroute53.AliasRecordTargetConfig {
	if err := i.validateBindParameters(_record); err != nil {
		panic(err)
	}
	var returns *awsroute53.AliasRecordTargetConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{_record, _zone},
		&returns,
	)

	return returns
}

