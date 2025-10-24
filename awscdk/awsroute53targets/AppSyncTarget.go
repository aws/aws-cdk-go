package awsroute53targets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53targets/internal"
)

// Defines an AppSync Graphql API as the alias target.
//
// Requires that the domain
// name will be defined through `GraphqlApiProps.domainName`.
//
// Example:
//   import appsync "github.com/aws/aws-cdk-go/awscdk"
//
//   var zone HostedZone
//   var graphqlApi GraphqlApi
//
//
//   route53.NewARecord(this, jsii.String("AliasRecord"), &ARecordProps{
//   	Zone: Zone,
//   	Target: route53.RecordTarget_FromAlias(targets.NewAppSyncTarget(graphqlApi)),
//   })
//
type AppSyncTarget interface {
	awsroute53.IAliasRecordTarget
	// Return hosted zone ID and DNS name, usable for Route53 alias targets.
	Bind(_record awsroute53.IRecordSet, _zone awsroute53.IHostedZone) *awsroute53.AliasRecordTargetConfig
}

// The jsii proxy struct for AppSyncTarget
type jsiiProxy_AppSyncTarget struct {
	internal.Type__awsroute53IAliasRecordTarget
}

func NewAppSyncTarget(graphqlApi awsappsync.GraphqlApi) AppSyncTarget {
	_init_.Initialize()

	if err := validateNewAppSyncTargetParameters(graphqlApi); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppSyncTarget{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53_targets.AppSyncTarget",
		[]interface{}{graphqlApi},
		&j,
	)

	return &j
}

func NewAppSyncTarget_Override(a AppSyncTarget, graphqlApi awsappsync.GraphqlApi) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53_targets.AppSyncTarget",
		[]interface{}{graphqlApi},
		a,
	)
}

func (a *jsiiProxy_AppSyncTarget) Bind(_record awsroute53.IRecordSet, _zone awsroute53.IHostedZone) *awsroute53.AliasRecordTargetConfig {
	if err := a.validateBindParameters(_record); err != nil {
		panic(err)
	}
	var returns *awsroute53.AliasRecordTargetConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{_record, _zone},
		&returns,
	)

	return returns
}

