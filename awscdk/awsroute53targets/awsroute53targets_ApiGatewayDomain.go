package awsroute53targets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53targets/internal"
)

// Defines an API Gateway domain name as the alias target.
//
// Use the `ApiGateway` class if you wish to map the alias to an REST API with a
// domain name defined through the `RestApiProps.domainName` prop.
//
// Example:
//   var hostedZoneForExampleCom interface{}
//   var domainName domainName
//
//   import route53 "github.com/aws/aws-cdk-go/awscdk"
//   import targets "github.com/aws/aws-cdk-go/awscdk"
//
//
//   route53.NewARecord(this, jsii.String("CustomDomainAliasRecord"), &aRecordProps{
//   	zone: hostedZoneForExampleCom,
//   	target: route53.recordTarget.fromAlias(targets.NewApiGatewayDomain(domainName)),
//   })
//
type ApiGatewayDomain interface {
	awsroute53.IAliasRecordTarget
	// Return hosted zone ID and DNS name, usable for Route53 alias targets.
	Bind(_record awsroute53.IRecordSet, _zone awsroute53.IHostedZone) *awsroute53.AliasRecordTargetConfig
}

// The jsii proxy struct for ApiGatewayDomain
type jsiiProxy_ApiGatewayDomain struct {
	internal.Type__awsroute53IAliasRecordTarget
}

func NewApiGatewayDomain(domainName awsapigateway.IDomainName) ApiGatewayDomain {
	_init_.Initialize()

	j := jsiiProxy_ApiGatewayDomain{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53_targets.ApiGatewayDomain",
		[]interface{}{domainName},
		&j,
	)

	return &j
}

func NewApiGatewayDomain_Override(a ApiGatewayDomain, domainName awsapigateway.IDomainName) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53_targets.ApiGatewayDomain",
		[]interface{}{domainName},
		a,
	)
}

func (a *jsiiProxy_ApiGatewayDomain) Bind(_record awsroute53.IRecordSet, _zone awsroute53.IHostedZone) *awsroute53.AliasRecordTargetConfig {
	var returns *awsroute53.AliasRecordTargetConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{_record, _zone},
		&returns,
	)

	return returns
}

