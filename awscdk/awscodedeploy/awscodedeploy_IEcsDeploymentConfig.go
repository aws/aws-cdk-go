package awscodedeploy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The Deployment Configuration of an ECS Deployment Group.
//
// The default, pre-defined Configurations are available as constants on the {@link EcsDeploymentConfig} class
// (for example, `EcsDeploymentConfig.AllAtOnce`).
//
// Note: CloudFormation does not currently support creating custom ECS configs outside
// of using a custom resource. You can import custom deployment config created outside the
// CDK or via a custom resource with {@link EcsDeploymentConfig#fromEcsDeploymentConfigName}.
type IEcsDeploymentConfig interface {
	DeploymentConfigArn() *string
	DeploymentConfigName() *string
}

// The jsii proxy for IEcsDeploymentConfig
type jsiiProxy_IEcsDeploymentConfig struct {
	_ byte // padding
}

func (j *jsiiProxy_IEcsDeploymentConfig) DeploymentConfigArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentConfigArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEcsDeploymentConfig) DeploymentConfigName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentConfigName",
		&returns,
	)
	return returns
}

