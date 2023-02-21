package awscodedeploy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodedeploy/internal"
)

// Interface for a Lambda deployment groups.
type ILambdaDeploymentGroup interface {
	awscdk.IResource
	// The reference to the CodeDeploy Lambda Application that this Deployment Group belongs to.
	Application() ILambdaApplication
	// The Deployment Configuration this Group uses.
	DeploymentConfig() ILambdaDeploymentConfig
	// The ARN of this Deployment Group.
	DeploymentGroupArn() *string
	// The physical name of the CodeDeploy Deployment Group.
	DeploymentGroupName() *string
}

// The jsii proxy for ILambdaDeploymentGroup
type jsiiProxy_ILambdaDeploymentGroup struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_ILambdaDeploymentGroup) Application() ILambdaApplication {
	var returns ILambdaApplication
	_jsii_.Get(
		j,
		"application",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILambdaDeploymentGroup) DeploymentConfig() ILambdaDeploymentConfig {
	var returns ILambdaDeploymentConfig
	_jsii_.Get(
		j,
		"deploymentConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILambdaDeploymentGroup) DeploymentGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILambdaDeploymentGroup) DeploymentGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentGroupName",
		&returns,
	)
	return returns
}

