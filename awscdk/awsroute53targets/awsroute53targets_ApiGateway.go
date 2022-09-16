package awsroute53targets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
)

// Defines an API Gateway REST API as the alias target. Requires that the domain name will be defined through `RestApiProps.domainName`.
//
// You can direct the alias to any `apigateway.DomainName` resource through the
// `ApiGatewayDomain` class.
//
// Example:
//   import route53 "github.com/aws/aws-cdk-go/awscdk"
//   import targets "github.com/aws/aws-cdk-go/awscdk"
//
//   var api restApi
//   var hostedZoneForExampleCom interface{}
//
//
//   route53.NewARecord(this, jsii.String("CustomDomainAliasRecord"), &aRecordProps{
//   	zone: hostedZoneForExampleCom,
//   	target: route53.recordTarget.fromAlias(targets.NewApiGateway(api)),
//   })
//
type ApiGateway interface {
	ApiGatewayDomain
	// Return hosted zone ID and DNS name, usable for Route53 alias targets.
	Bind(_record awsroute53.IRecordSet, _zone awsroute53.IHostedZone) *awsroute53.AliasRecordTargetConfig
}

// The jsii proxy struct for ApiGateway
type jsiiProxy_ApiGateway struct {
	jsiiProxy_ApiGatewayDomain
}

func NewApiGateway(api awsapigateway.RestApiBase) ApiGateway {
	_init_.Initialize()

	if err := validateNewApiGatewayParameters(api); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApiGateway{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53_targets.ApiGateway",
		[]interface{}{api},
		&j,
	)

	return &j
}

func NewApiGateway_Override(a ApiGateway, api awsapigateway.RestApiBase) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53_targets.ApiGateway",
		[]interface{}{api},
		a,
	)
}

func (a *jsiiProxy_ApiGateway) Bind(_record awsroute53.IRecordSet, _zone awsroute53.IHostedZone) *awsroute53.AliasRecordTargetConfig {
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

