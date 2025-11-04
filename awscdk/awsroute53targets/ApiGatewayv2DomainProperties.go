package awsroute53targets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53targets/internal"
)

// Defines an API Gateway V2 domain name as the alias target.
//
// Example:
//   import apigwv2 "github.com/aws/aws-cdk-go/awscdk"
//
//   var zone HostedZone
//   var domainName DomainName
//
//
//   route53.NewARecord(this, jsii.String("AliasRecord"), &ARecordProps{
//   	Zone: Zone,
//   	Target: route53.RecordTarget_FromAlias(targets.NewApiGatewayv2DomainProperties(domainName.RegionalDomainName, domainName.RegionalHostedZoneId)),
//   })
//
type ApiGatewayv2DomainProperties interface {
	awsroute53.IAliasRecordTarget
	// Return hosted zone ID and DNS name, usable for Route53 alias targets.
	Bind(record awsroute53.IRecordSet, zone awsroute53.IHostedZone) *awsroute53.AliasRecordTargetConfig
}

// The jsii proxy struct for ApiGatewayv2DomainProperties
type jsiiProxy_ApiGatewayv2DomainProperties struct {
	internal.Type__awsroute53IAliasRecordTarget
}

func NewApiGatewayv2DomainProperties(regionalDomainName *string, regionalHostedZoneId *string) ApiGatewayv2DomainProperties {
	_init_.Initialize()

	if err := validateNewApiGatewayv2DomainPropertiesParameters(regionalDomainName, regionalHostedZoneId); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApiGatewayv2DomainProperties{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53_targets.ApiGatewayv2DomainProperties",
		[]interface{}{regionalDomainName, regionalHostedZoneId},
		&j,
	)

	return &j
}

func NewApiGatewayv2DomainProperties_Override(a ApiGatewayv2DomainProperties, regionalDomainName *string, regionalHostedZoneId *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53_targets.ApiGatewayv2DomainProperties",
		[]interface{}{regionalDomainName, regionalHostedZoneId},
		a,
	)
}

func (a *jsiiProxy_ApiGatewayv2DomainProperties) Bind(record awsroute53.IRecordSet, zone awsroute53.IHostedZone) *awsroute53.AliasRecordTargetConfig {
	if err := a.validateBindParameters(record); err != nil {
		panic(err)
	}
	var returns *awsroute53.AliasRecordTargetConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{record, zone},
		&returns,
	)

	return returns
}

