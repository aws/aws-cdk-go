package awsroute53targets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53targets/internal"
)

// Use another Route 53 record as an alias record target.
//
// Example:
//   var zone hostedZone
//   var record aRecord
//
//   route53.NewARecord(this, jsii.String("AliasRecord"), &ARecordProps{
//   	Zone: Zone,
//   	Target: route53.RecordTarget_FromAlias(targets.NewRoute53RecordTarget(record)),
//   })
//
type Route53RecordTarget interface {
	awsroute53.IAliasRecordTarget
	// Return hosted zone ID and DNS name, usable for Route53 alias targets.
	Bind(record awsroute53.IRecordSet, zone awsroute53.IHostedZone) *awsroute53.AliasRecordTargetConfig
}

// The jsii proxy struct for Route53RecordTarget
type jsiiProxy_Route53RecordTarget struct {
	internal.Type__awsroute53IAliasRecordTarget
}

func NewRoute53RecordTarget(record awsroute53.IRecordSet) Route53RecordTarget {
	_init_.Initialize()

	if err := validateNewRoute53RecordTargetParameters(record); err != nil {
		panic(err)
	}
	j := jsiiProxy_Route53RecordTarget{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53_targets.Route53RecordTarget",
		[]interface{}{record},
		&j,
	)

	return &j
}

func NewRoute53RecordTarget_Override(r Route53RecordTarget, record awsroute53.IRecordSet) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53_targets.Route53RecordTarget",
		[]interface{}{record},
		r,
	)
}

func (r *jsiiProxy_Route53RecordTarget) Bind(record awsroute53.IRecordSet, zone awsroute53.IHostedZone) *awsroute53.AliasRecordTargetConfig {
	if err := r.validateBindParameters(record); err != nil {
		panic(err)
	}
	var returns *awsroute53.AliasRecordTargetConfig

	_jsii_.Invoke(
		r,
		"bind",
		[]interface{}{record, zone},
		&returns,
	)

	return returns
}

