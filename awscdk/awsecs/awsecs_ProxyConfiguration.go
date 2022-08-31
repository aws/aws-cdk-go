package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
)

// The base class for proxy configurations.
// Experimental.
type ProxyConfiguration interface {
	// Called when the proxy configuration is configured on a task definition.
	// Experimental.
	Bind(_scope awscdk.Construct, _taskDefinition TaskDefinition) *CfnTaskDefinition_ProxyConfigurationProperty
}

// The jsii proxy struct for ProxyConfiguration
type jsiiProxy_ProxyConfiguration struct {
	_ byte // padding
}

// Experimental.
func NewProxyConfiguration_Override(p ProxyConfiguration) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ecs.ProxyConfiguration",
		nil, // no parameters
		p,
	)
}

func (p *jsiiProxy_ProxyConfiguration) Bind(_scope awscdk.Construct, _taskDefinition TaskDefinition) *CfnTaskDefinition_ProxyConfigurationProperty {
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

