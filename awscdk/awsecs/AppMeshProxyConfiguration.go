package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
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
//   appMeshProxyConfiguration := awscdk.Aws_ecs.NewAppMeshProxyConfiguration(&AppMeshProxyConfigurationConfigProps{
//   	ContainerName: jsii.String("containerName"),
//   	Properties: &AppMeshProxyConfigurationProps{
//   		AppPorts: []*f64{
//   			jsii.Number(123),
//   		},
//   		ProxyEgressPort: jsii.Number(123),
//   		ProxyIngressPort: jsii.Number(123),
//
//   		// the properties below are optional
//   		EgressIgnoredIPs: []*string{
//   			jsii.String("egressIgnoredIPs"),
//   		},
//   		EgressIgnoredPorts: []*f64{
//   			jsii.Number(123),
//   		},
//   		IgnoredGID: jsii.Number(123),
//   		IgnoredUID: jsii.Number(123),
//   	},
//   })
//
type AppMeshProxyConfiguration interface {
	ProxyConfiguration
	// Called when the proxy configuration is configured on a task definition.
	Bind(_scope constructs.Construct, _taskDefinition TaskDefinition) *CfnTaskDefinition_ProxyConfigurationProperty
}

// The jsii proxy struct for AppMeshProxyConfiguration
type jsiiProxy_AppMeshProxyConfiguration struct {
	jsiiProxy_ProxyConfiguration
}

// Constructs a new instance of the AppMeshProxyConfiguration class.
func NewAppMeshProxyConfiguration(props *AppMeshProxyConfigurationConfigProps) AppMeshProxyConfiguration {
	_init_.Initialize()

	if err := validateNewAppMeshProxyConfigurationParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppMeshProxyConfiguration{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.AppMeshProxyConfiguration",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Constructs a new instance of the AppMeshProxyConfiguration class.
func NewAppMeshProxyConfiguration_Override(a AppMeshProxyConfiguration, props *AppMeshProxyConfigurationConfigProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.AppMeshProxyConfiguration",
		[]interface{}{props},
		a,
	)
}

func (a *jsiiProxy_AppMeshProxyConfiguration) Bind(_scope constructs.Construct, _taskDefinition TaskDefinition) *CfnTaskDefinition_ProxyConfigurationProperty {
	if err := a.validateBindParameters(_scope, _taskDefinition); err != nil {
		panic(err)
	}
	var returns *CfnTaskDefinition_ProxyConfigurationProperty

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{_scope, _taskDefinition},
		&returns,
	)

	return returns
}

