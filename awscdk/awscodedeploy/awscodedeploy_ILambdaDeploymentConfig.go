package awscodedeploy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The Deployment Configuration of a Lambda Deployment Group.
//
// The default, pre-defined Configurations are available as constants on the {@link LambdaDeploymentConfig} class
// (`LambdaDeploymentConfig.AllAtOnce`, `LambdaDeploymentConfig.Canary10Percent30Minutes`, etc.).
//
// Note: CloudFormation does not currently support creating custom lambda configs outside
// of using a custom resource. You can import custom deployment config created outside the
// CDK or via a custom resource with {@link LambdaDeploymentConfig#import}.
// Experimental.
type ILambdaDeploymentConfig interface {
	// Experimental.
	DeploymentConfigArn() *string
	// Experimental.
	DeploymentConfigName() *string
}

// The jsii proxy for ILambdaDeploymentConfig
type jsiiProxy_ILambdaDeploymentConfig struct {
	_ byte // padding
}

func (j *jsiiProxy_ILambdaDeploymentConfig) DeploymentConfigArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentConfigArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILambdaDeploymentConfig) DeploymentConfigName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentConfigName",
		&returns,
	)
	return returns
}

