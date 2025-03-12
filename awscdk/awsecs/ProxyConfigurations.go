package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The base class for proxy configurations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   proxyConfigurations := awscdk.Aws_ecs.NewProxyConfigurations()
//
type ProxyConfigurations interface {
}

// The jsii proxy struct for ProxyConfigurations
type jsiiProxy_ProxyConfigurations struct {
	_ byte // padding
}

func NewProxyConfigurations() ProxyConfigurations {
	_init_.Initialize()

	j := jsiiProxy_ProxyConfigurations{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.ProxyConfigurations",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewProxyConfigurations_Override(p ProxyConfigurations) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.ProxyConfigurations",
		nil, // no parameters
		p,
	)
}

// Constructs a new instance of the ProxyConfiguration class.
func ProxyConfigurations_AppMeshProxyConfiguration(props *AppMeshProxyConfigurationConfigProps) ProxyConfiguration {
	_init_.Initialize()

	if err := validateProxyConfigurations_AppMeshProxyConfigurationParameters(props); err != nil {
		panic(err)
	}
	var returns ProxyConfiguration

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.ProxyConfigurations",
		"appMeshProxyConfiguration",
		[]interface{}{props},
		&returns,
	)

	return returns
}

