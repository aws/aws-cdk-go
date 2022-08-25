package awscodedeploy

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// A custom Deployment Configuration for an ECS Deployment Group.
//
// Note: This class currently stands as namespaced container of the default configurations
// until CloudFormation supports custom ECS Deployment Configs. Until then it is closed
// (private constructor) and does not extend {@link Construct}.
type EcsDeploymentConfig interface {
}

// The jsii proxy struct for EcsDeploymentConfig
type jsiiProxy_EcsDeploymentConfig struct {
	_ byte // padding
}

// Import a custom Deployment Configuration for an ECS Deployment Group defined outside the CDK.
//
// Returns: a Construct representing a reference to an existing custom Deployment Configuration.
func EcsDeploymentConfig_FromEcsDeploymentConfigName(_scope constructs.Construct, _id *string, ecsDeploymentConfigName *string) IEcsDeploymentConfig {
	_init_.Initialize()

	var returns IEcsDeploymentConfig

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codedeploy.EcsDeploymentConfig",
		"fromEcsDeploymentConfigName",
		[]interface{}{_scope, _id, ecsDeploymentConfigName},
		&returns,
	)

	return returns
}

func EcsDeploymentConfig_ALL_AT_ONCE() IEcsDeploymentConfig {
	_init_.Initialize()
	var returns IEcsDeploymentConfig
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codedeploy.EcsDeploymentConfig",
		"ALL_AT_ONCE",
		&returns,
	)
	return returns
}

