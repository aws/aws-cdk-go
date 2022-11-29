package awscodedeploy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsautoscaling"
	"github.com/aws/aws-cdk-go/awscdk/awscodedeploy/internal"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// Experimental.
type IServerDeploymentGroup interface {
	awscdk.IResource
	// Experimental.
	Application() IServerApplication
	// Experimental.
	AutoScalingGroups() *[]awsautoscaling.IAutoScalingGroup
	// Experimental.
	DeploymentConfig() IServerDeploymentConfig
	// Experimental.
	DeploymentGroupArn() *string
	// Experimental.
	DeploymentGroupName() *string
	// Experimental.
	Role() awsiam.IRole
}

// The jsii proxy for IServerDeploymentGroup
type jsiiProxy_IServerDeploymentGroup struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IServerDeploymentGroup) Application() IServerApplication {
	var returns IServerApplication
	_jsii_.Get(
		j,
		"application",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServerDeploymentGroup) AutoScalingGroups() *[]awsautoscaling.IAutoScalingGroup {
	var returns *[]awsautoscaling.IAutoScalingGroup
	_jsii_.Get(
		j,
		"autoScalingGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServerDeploymentGroup) DeploymentConfig() IServerDeploymentConfig {
	var returns IServerDeploymentConfig
	_jsii_.Get(
		j,
		"deploymentConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServerDeploymentGroup) DeploymentGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServerDeploymentGroup) DeploymentGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServerDeploymentGroup) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

