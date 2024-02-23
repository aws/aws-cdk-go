package awsappconfig

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Defines the deployment strategy ID's of AWS AppConfig deployment strategies.
//
// Example:
//   appconfig.DeploymentStrategy_FromDeploymentStrategyId(this, jsii.String("MyImportedDeploymentStrategy"), appconfig.DeploymentStrategyId_FromString(jsii.String("abc123")))
//
// See: https://docs.aws.amazon.com/appconfig/latest/userguide/appconfig-creating-deployment-strategy.html
//
type DeploymentStrategyId interface {
	// The deployment strategy ID.
	Id() *string
}

// The jsii proxy struct for DeploymentStrategyId
type jsiiProxy_DeploymentStrategyId struct {
	_ byte // padding
}

func (j *jsiiProxy_DeploymentStrategyId) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}


func NewDeploymentStrategyId_Override(d DeploymentStrategyId) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appconfig.DeploymentStrategyId",
		nil, // no parameters
		d,
	)
}

// Builds a deployment strategy ID from a string.
func DeploymentStrategyId_FromString(deploymentStrategyId *string) DeploymentStrategyId {
	_init_.Initialize()

	if err := validateDeploymentStrategyId_FromStringParameters(deploymentStrategyId); err != nil {
		panic(err)
	}
	var returns DeploymentStrategyId

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appconfig.DeploymentStrategyId",
		"fromString",
		[]interface{}{deploymentStrategyId},
		&returns,
	)

	return returns
}

func DeploymentStrategyId_ALL_AT_ONCE() DeploymentStrategyId {
	_init_.Initialize()
	var returns DeploymentStrategyId
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_appconfig.DeploymentStrategyId",
		"ALL_AT_ONCE",
		&returns,
	)
	return returns
}

func DeploymentStrategyId_CANARY_10_PERCENT_20_MINUTES() DeploymentStrategyId {
	_init_.Initialize()
	var returns DeploymentStrategyId
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_appconfig.DeploymentStrategyId",
		"CANARY_10_PERCENT_20_MINUTES",
		&returns,
	)
	return returns
}

func DeploymentStrategyId_LINEAR_20_PERCENT_EVERY_6_MINUTES() DeploymentStrategyId {
	_init_.Initialize()
	var returns DeploymentStrategyId
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_appconfig.DeploymentStrategyId",
		"LINEAR_20_PERCENT_EVERY_6_MINUTES",
		&returns,
	)
	return returns
}

func DeploymentStrategyId_LINEAR_50_PERCENT_EVERY_30_SECONDS() DeploymentStrategyId {
	_init_.Initialize()
	var returns DeploymentStrategyId
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_appconfig.DeploymentStrategyId",
		"LINEAR_50_PERCENT_EVERY_30_SECONDS",
		&returns,
	)
	return returns
}

