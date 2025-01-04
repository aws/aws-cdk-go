package awsroute53targets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53targets/internal"
)

// Use a user pool domain as an alias record target.
//
// Example:
//   import cognito "github.com/aws/aws-cdk-go/awscdk"
//
//   var zone hostedZone
//   var domain userPoolDomain
//
//   route53.NewARecord(this, jsii.String("AliasRecord"), &ARecordProps{
//   	Zone: Zone,
//   	Target: route53.RecordTarget_FromAlias(targets.NewUserPoolDomainTarget(domain)),
//   })
//
type UserPoolDomainTarget interface {
	awsroute53.IAliasRecordTarget
	// Return hosted zone ID and DNS name, usable for Route53 alias targets.
	Bind(record awsroute53.IRecordSet, _zone awsroute53.IHostedZone) *awsroute53.AliasRecordTargetConfig
}

// The jsii proxy struct for UserPoolDomainTarget
type jsiiProxy_UserPoolDomainTarget struct {
	internal.Type__awsroute53IAliasRecordTarget
}

func NewUserPoolDomainTarget(domain awscognito.UserPoolDomain) UserPoolDomainTarget {
	_init_.Initialize()

	if err := validateNewUserPoolDomainTargetParameters(domain); err != nil {
		panic(err)
	}
	j := jsiiProxy_UserPoolDomainTarget{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53_targets.UserPoolDomainTarget",
		[]interface{}{domain},
		&j,
	)

	return &j
}

func NewUserPoolDomainTarget_Override(u UserPoolDomainTarget, domain awscognito.UserPoolDomain) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53_targets.UserPoolDomainTarget",
		[]interface{}{domain},
		u,
	)
}

func (u *jsiiProxy_UserPoolDomainTarget) Bind(record awsroute53.IRecordSet, _zone awsroute53.IHostedZone) *awsroute53.AliasRecordTargetConfig {
	if err := u.validateBindParameters(record); err != nil {
		panic(err)
	}
	var returns *awsroute53.AliasRecordTargetConfig

	_jsii_.Invoke(
		u,
		"bind",
		[]interface{}{record, _zone},
		&returns,
	)

	return returns
}

