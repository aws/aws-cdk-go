package awsroute53targets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsroute53"
	"github.com/aws/aws-cdk-go/awscdk/awsroute53targets/internal"
)

// Use an Elastic Beanstalk environment URL as an alias record target. E.g. mysampleenvironment.xyz.us-east-1.elasticbeanstalk.com or mycustomcnameprefix.us-east-1.elasticbeanstalk.com.
//
// Only supports Elastic Beanstalk environments created after 2016 that have a regional endpoint.
//
// Example:
//   var zone hostedZone
//   var ebsEnvironmentUrl string
//
//
//   route53.NewARecord(this, jsii.String("AliasRecord"), &ARecordProps{
//   	Zone: Zone,
//   	Target: route53.RecordTarget_FromAlias(targets.NewElasticBeanstalkEnvironmentEndpointTarget(ebsEnvironmentUrl)),
//   })
//
// Experimental.
type ElasticBeanstalkEnvironmentEndpointTarget interface {
	awsroute53.IAliasRecordTarget
	// Return hosted zone ID and DNS name, usable for Route53 alias targets.
	// Experimental.
	Bind(_record awsroute53.IRecordSet, _zone awsroute53.IHostedZone) *awsroute53.AliasRecordTargetConfig
}

// The jsii proxy struct for ElasticBeanstalkEnvironmentEndpointTarget
type jsiiProxy_ElasticBeanstalkEnvironmentEndpointTarget struct {
	internal.Type__awsroute53IAliasRecordTarget
}

// Experimental.
func NewElasticBeanstalkEnvironmentEndpointTarget(environmentEndpoint *string) ElasticBeanstalkEnvironmentEndpointTarget {
	_init_.Initialize()

	if err := validateNewElasticBeanstalkEnvironmentEndpointTargetParameters(environmentEndpoint); err != nil {
		panic(err)
	}
	j := jsiiProxy_ElasticBeanstalkEnvironmentEndpointTarget{}

	_jsii_.Create(
		"monocdk.aws_route53_targets.ElasticBeanstalkEnvironmentEndpointTarget",
		[]interface{}{environmentEndpoint},
		&j,
	)

	return &j
}

// Experimental.
func NewElasticBeanstalkEnvironmentEndpointTarget_Override(e ElasticBeanstalkEnvironmentEndpointTarget, environmentEndpoint *string) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_route53_targets.ElasticBeanstalkEnvironmentEndpointTarget",
		[]interface{}{environmentEndpoint},
		e,
	)
}

func (e *jsiiProxy_ElasticBeanstalkEnvironmentEndpointTarget) Bind(_record awsroute53.IRecordSet, _zone awsroute53.IHostedZone) *awsroute53.AliasRecordTargetConfig {
	if err := e.validateBindParameters(_record); err != nil {
		panic(err)
	}
	var returns *awsroute53.AliasRecordTargetConfig

	_jsii_.Invoke(
		e,
		"bind",
		[]interface{}{_record, _zone},
		&returns,
	)

	return returns
}

