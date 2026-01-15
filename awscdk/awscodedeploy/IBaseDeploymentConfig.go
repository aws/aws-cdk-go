package awscodedeploy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodedeploy/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscodedeploy"
)

// The base class for ServerDeploymentConfig, EcsDeploymentConfig, and LambdaDeploymentConfig deployment configurations.
type IBaseDeploymentConfig interface {
	interfacesawscodedeploy.IDeploymentConfigRef
	// The ARN of the Deployment Configuration.
	DeploymentConfigArn() *string
	// The physical, human-readable name of the Deployment Configuration.
	DeploymentConfigName() *string
}

// The jsii proxy for IBaseDeploymentConfig
type jsiiProxy_IBaseDeploymentConfig struct {
	internal.Type__interfacesawscodedeployIDeploymentConfigRef
}

func (j *jsiiProxy_IBaseDeploymentConfig) DeploymentConfigArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentConfigArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBaseDeploymentConfig) DeploymentConfigName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentConfigName",
		&returns,
	)
	return returns
}

