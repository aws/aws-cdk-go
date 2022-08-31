package awsroute53targets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsroute53"
	"github.com/aws/aws-cdk-go/awscdk/awsroute53targets/internal"
)

// Use another Route 53 record as an alias record target.
//
// Example:
//   var zone hostedZone
//   var record aRecord
//
//   route53.NewARecord(this, jsii.String("AliasRecord"), &aRecordProps{
//   	zone: zone,
//   	target: route53.recordTarget.fromAlias(targets.NewRoute53RecordTarget(record)),
//   })
//
// Experimental.
type Route53RecordTarget interface {
	awsroute53.IAliasRecordTarget
	// Return hosted zone ID and DNS name, usable for Route53 alias targets.
	// Experimental.
	Bind(_record awsroute53.IRecordSet, zone awsroute53.IHostedZone) *awsroute53.AliasRecordTargetConfig
}

// The jsii proxy struct for Route53RecordTarget
type jsiiProxy_Route53RecordTarget struct {
	internal.Type__awsroute53IAliasRecordTarget
}

// Experimental.
func NewRoute53RecordTarget(record awsroute53.IRecordSet) Route53RecordTarget {
	_init_.Initialize()

	if err := validateNewRoute53RecordTargetParameters(record); err != nil {
		panic(err)
	}
	j := jsiiProxy_Route53RecordTarget{}

	_jsii_.Create(
		"monocdk.aws_route53_targets.Route53RecordTarget",
		[]interface{}{record},
		&j,
	)

	return &j
}

// Experimental.
func NewRoute53RecordTarget_Override(r Route53RecordTarget, record awsroute53.IRecordSet) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_route53_targets.Route53RecordTarget",
		[]interface{}{record},
		r,
	)
}

func (r *jsiiProxy_Route53RecordTarget) Bind(_record awsroute53.IRecordSet, zone awsroute53.IHostedZone) *awsroute53.AliasRecordTargetConfig {
	if err := r.validateBindParameters(_record); err != nil {
		panic(err)
	}
	var returns *awsroute53.AliasRecordTargetConfig

	_jsii_.Invoke(
		r,
		"bind",
		[]interface{}{_record, zone},
		&returns,
	)

	return returns
}

