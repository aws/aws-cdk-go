package awscodedeploy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodedeploy/internal"
)

// Interface for an ECS deployment group.
type IEcsDeploymentGroup interface {
	awscdk.IResource
	// The reference to the CodeDeploy ECS Application that this Deployment Group belongs to.
	Application() IEcsApplication
	// The Deployment Configuration this Group uses.
	DeploymentConfig() IEcsDeploymentConfig
	// The ARN of this Deployment Group.
	DeploymentGroupArn() *string
	// The physical name of the CodeDeploy Deployment Group.
	DeploymentGroupName() *string
}

// The jsii proxy for IEcsDeploymentGroup
type jsiiProxy_IEcsDeploymentGroup struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IEcsDeploymentGroup) Application() IEcsApplication {
	var returns IEcsApplication
	_jsii_.Get(
		j,
		"application",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEcsDeploymentGroup) DeploymentConfig() IEcsDeploymentConfig {
	var returns IEcsDeploymentConfig
	_jsii_.Get(
		j,
		"deploymentConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEcsDeploymentGroup) DeploymentGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEcsDeploymentGroup) DeploymentGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentGroupName",
		&returns,
	)
	return returns
}

