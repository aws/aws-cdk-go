package awscodedeploy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The Deployment Configuration of an EC2/on-premise Deployment Group.
//
// The default, pre-defined Configurations are available as constants on the {@link ServerDeploymentConfig} class
// (`ServerDeploymentConfig.HALF_AT_A_TIME`, `ServerDeploymentConfig.ALL_AT_ONCE`, etc.).
// To create a custom Deployment Configuration,
// instantiate the {@link ServerDeploymentConfig} Construct.
type IServerDeploymentConfig interface {
	DeploymentConfigArn() *string
	DeploymentConfigName() *string
}

// The jsii proxy for IServerDeploymentConfig
type jsiiProxy_IServerDeploymentConfig struct {
	_ byte // padding
}

func (j *jsiiProxy_IServerDeploymentConfig) DeploymentConfigArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentConfigArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServerDeploymentConfig) DeploymentConfigName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentConfigName",
		&returns,
	)
	return returns
}

