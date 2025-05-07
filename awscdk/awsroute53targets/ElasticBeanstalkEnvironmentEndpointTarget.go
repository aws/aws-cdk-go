package awsroute53targets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53targets/internal"
)

// Use an Elastic Beanstalk environment URL as an alias record target. E.g. mysampleenvironment.xyz.us-east-1.elasticbeanstalk.com or mycustomcnameprefix.us-east-1.elasticbeanstalk.com.
//
// Only supports Elastic Beanstalk environments created after 2016 that have a regional endpoint.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var zone hostedZone
//   var ebsEnvironmentUrl string
//
//
//   route53.NewARecord(this, jsii.String("AliasRecord"), &ARecordProps{
//   	Zone: Zone,
//   	Target: route53.RecordTarget_FromAlias(
//   	targets.NewElasticBeanstalkEnvironmentEndpointTarget(ebsEnvironmentUrl, map[string]*string{
//   		"hostedZoneId": awscdk.RegionInfo_get(jsii.String("us-east-1")).ebsEnvEndpointHostedZoneId,
//   	})),
//   })
//
type ElasticBeanstalkEnvironmentEndpointTarget interface {
	awsroute53.IAliasRecordTarget
	// Return hosted zone ID and DNS name, usable for Route53 alias targets.
	Bind(record awsroute53.IRecordSet, _zone awsroute53.IHostedZone) *awsroute53.AliasRecordTargetConfig
}

// The jsii proxy struct for ElasticBeanstalkEnvironmentEndpointTarget
type jsiiProxy_ElasticBeanstalkEnvironmentEndpointTarget struct {
	internal.Type__awsroute53IAliasRecordTarget
}

func NewElasticBeanstalkEnvironmentEndpointTarget(environmentEndpoint *string, props IAliasRecordTargetProps) ElasticBeanstalkEnvironmentEndpointTarget {
	_init_.Initialize()

	if err := validateNewElasticBeanstalkEnvironmentEndpointTargetParameters(environmentEndpoint); err != nil {
		panic(err)
	}
	j := jsiiProxy_ElasticBeanstalkEnvironmentEndpointTarget{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53_targets.ElasticBeanstalkEnvironmentEndpointTarget",
		[]interface{}{environmentEndpoint, props},
		&j,
	)

	return &j
}

func NewElasticBeanstalkEnvironmentEndpointTarget_Override(e ElasticBeanstalkEnvironmentEndpointTarget, environmentEndpoint *string, props IAliasRecordTargetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53_targets.ElasticBeanstalkEnvironmentEndpointTarget",
		[]interface{}{environmentEndpoint, props},
		e,
	)
}

func (e *jsiiProxy_ElasticBeanstalkEnvironmentEndpointTarget) Bind(record awsroute53.IRecordSet, _zone awsroute53.IHostedZone) *awsroute53.AliasRecordTargetConfig {
	if err := e.validateBindParameters(record); err != nil {
		panic(err)
	}
	var returns *awsroute53.AliasRecordTargetConfig

	_jsii_.Invoke(
		e,
		"bind",
		[]interface{}{record, _zone},
		&returns,
	)

	return returns
}

