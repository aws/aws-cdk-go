package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// The base class for proxy configurations.
type ProxyConfiguration interface {
	// Called when the proxy configuration is configured on a task definition.
	Bind(_scope constructs.Construct, _taskDefinition TaskDefinition) *CfnTaskDefinition_ProxyConfigurationProperty
}

// The jsii proxy struct for ProxyConfiguration
type jsiiProxy_ProxyConfiguration struct {
	_ byte // padding
}

func NewProxyConfiguration_Override(p ProxyConfiguration) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.ProxyConfiguration",
		nil, // no parameters
		p,
	)
}

func (p *jsiiProxy_ProxyConfiguration) Bind(_scope constructs.Construct, _taskDefinition TaskDefinition) *CfnTaskDefinition_ProxyConfigurationProperty {
	if err := p.validateBindParameters(_scope, _taskDefinition); err != nil {
		panic(err)
	}
	var returns *CfnTaskDefinition_ProxyConfigurationProperty

	_jsii_.Invoke(
		p,
		"bind",
		[]interface{}{_scope, _taskDefinition},
		&returns,
	)

	return returns
}

