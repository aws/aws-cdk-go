package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
)

// The class for App Mesh proxy configurations.
//
// For tasks using the EC2 launch type, the container instances require at least version 1.26.0 of the container agent and at least version
// 1.26.0-1 of the ecs-init package to enable a proxy configuration. If your container instances are launched from the Amazon ECS-optimized
// AMI version 20190301 or later, then they contain the required versions of the container agent and ecs-init.
// For more information, see [Amazon ECS-optimized AMIs](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-optimized_AMI.html).
//
// For tasks using the Fargate launch type, the task or service requires platform version 1.3.0 or later.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   appMeshProxyConfiguration := awscdk.Aws_ecs.NewAppMeshProxyConfiguration(&appMeshProxyConfigurationConfigProps{
//   	containerName: jsii.String("containerName"),
//   	properties: &appMeshProxyConfigurationProps{
//   		appPorts: []*f64{
//   			jsii.Number(123),
//   		},
//   		proxyEgressPort: jsii.Number(123),
//   		proxyIngressPort: jsii.Number(123),
//
//   		// the properties below are optional
//   		egressIgnoredIPs: []*string{
//   			jsii.String("egressIgnoredIPs"),
//   		},
//   		egressIgnoredPorts: []*f64{
//   			jsii.Number(123),
//   		},
//   		ignoredGID: jsii.Number(123),
//   		ignoredUID: jsii.Number(123),
//   	},
//   })
//
// Experimental.
type AppMeshProxyConfiguration interface {
	ProxyConfiguration
	// Called when the proxy configuration is configured on a task definition.
	// Experimental.
	Bind(_scope awscdk.Construct, _taskDefinition TaskDefinition) *CfnTaskDefinition_ProxyConfigurationProperty
}

// The jsii proxy struct for AppMeshProxyConfiguration
type jsiiProxy_AppMeshProxyConfiguration struct {
	jsiiProxy_ProxyConfiguration
}

// Constructs a new instance of the AppMeshProxyConfiguration class.
// Experimental.
func NewAppMeshProxyConfiguration(props *AppMeshProxyConfigurationConfigProps) AppMeshProxyConfiguration {
	_init_.Initialize()

	j := jsiiProxy_AppMeshProxyConfiguration{}

	_jsii_.Create(
		"monocdk.aws_ecs.AppMeshProxyConfiguration",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Constructs a new instance of the AppMeshProxyConfiguration class.
// Experimental.
func NewAppMeshProxyConfiguration_Override(a AppMeshProxyConfiguration, props *AppMeshProxyConfigurationConfigProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ecs.AppMeshProxyConfiguration",
		[]interface{}{props},
		a,
	)
}

func (a *jsiiProxy_AppMeshProxyConfiguration) Bind(_scope awscdk.Construct, _taskDefinition TaskDefinition) *CfnTaskDefinition_ProxyConfigurationProperty {
	var returns *CfnTaskDefinition_ProxyConfigurationProperty

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{_scope, _taskDefinition},
		&returns,
	)

	return returns
}

