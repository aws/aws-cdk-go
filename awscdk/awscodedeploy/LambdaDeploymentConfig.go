package awscodedeploy

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v3"
)

// A custom Deployment Configuration for a Lambda Deployment Group.
//
// Note: This class currently stands as namespaced container of the default configurations
// until CloudFormation supports custom Lambda Deployment Configs. Until then it is closed
// (private constructor) and does not extend {@link cdk.Construct}
//
// Example:
//   var myApplication lambdaApplication
//   var func function
//
//   version := func.currentVersion
//   version1Alias := lambda.NewAlias(this, jsii.String("alias"), &AliasProps{
//   	AliasName: jsii.String("prod"),
//   	Version: Version,
//   })
//
//   deploymentGroup := codedeploy.NewLambdaDeploymentGroup(this, jsii.String("BlueGreenDeployment"), &LambdaDeploymentGroupProps{
//   	Application: myApplication,
//   	 // optional property: one will be created for you if not provided
//   	Alias: version1Alias,
//   	DeploymentConfig: codedeploy.LambdaDeploymentConfig_LINEAR_10PERCENT_EVERY_1MINUTE(),
//   })
//
// Experimental.
type LambdaDeploymentConfig interface {
}

// The jsii proxy struct for LambdaDeploymentConfig
type jsiiProxy_LambdaDeploymentConfig struct {
	_ byte // padding
}

// Import a custom Deployment Configuration for a Lambda Deployment Group defined outside the CDK.
//
// Returns: a Construct representing a reference to an existing custom Deployment Configuration.
// Experimental.
func LambdaDeploymentConfig_Import(_scope constructs.Construct, _id *string, props *LambdaDeploymentConfigImportProps) ILambdaDeploymentConfig {
	_init_.Initialize()

	if err := validateLambdaDeploymentConfig_ImportParameters(_scope, _id, props); err != nil {
		panic(err)
	}
	var returns ILambdaDeploymentConfig

	_jsii_.StaticInvoke(
		"monocdk.aws_codedeploy.LambdaDeploymentConfig",
		"import",
		[]interface{}{_scope, _id, props},
		&returns,
	)

	return returns
}

func LambdaDeploymentConfig_ALL_AT_ONCE() ILambdaDeploymentConfig {
	_init_.Initialize()
	var returns ILambdaDeploymentConfig
	_jsii_.StaticGet(
		"monocdk.aws_codedeploy.LambdaDeploymentConfig",
		"ALL_AT_ONCE",
		&returns,
	)
	return returns
}

func LambdaDeploymentConfig_CANARY_10PERCENT_10MINUTES() ILambdaDeploymentConfig {
	_init_.Initialize()
	var returns ILambdaDeploymentConfig
	_jsii_.StaticGet(
		"monocdk.aws_codedeploy.LambdaDeploymentConfig",
		"CANARY_10PERCENT_10MINUTES",
		&returns,
	)
	return returns
}

func LambdaDeploymentConfig_CANARY_10PERCENT_15MINUTES() ILambdaDeploymentConfig {
	_init_.Initialize()
	var returns ILambdaDeploymentConfig
	_jsii_.StaticGet(
		"monocdk.aws_codedeploy.LambdaDeploymentConfig",
		"CANARY_10PERCENT_15MINUTES",
		&returns,
	)
	return returns
}

func LambdaDeploymentConfig_CANARY_10PERCENT_30MINUTES() ILambdaDeploymentConfig {
	_init_.Initialize()
	var returns ILambdaDeploymentConfig
	_jsii_.StaticGet(
		"monocdk.aws_codedeploy.LambdaDeploymentConfig",
		"CANARY_10PERCENT_30MINUTES",
		&returns,
	)
	return returns
}

func LambdaDeploymentConfig_CANARY_10PERCENT_5MINUTES() ILambdaDeploymentConfig {
	_init_.Initialize()
	var returns ILambdaDeploymentConfig
	_jsii_.StaticGet(
		"monocdk.aws_codedeploy.LambdaDeploymentConfig",
		"CANARY_10PERCENT_5MINUTES",
		&returns,
	)
	return returns
}

func LambdaDeploymentConfig_LINEAR_10PERCENT_EVERY_10MINUTES() ILambdaDeploymentConfig {
	_init_.Initialize()
	var returns ILambdaDeploymentConfig
	_jsii_.StaticGet(
		"monocdk.aws_codedeploy.LambdaDeploymentConfig",
		"LINEAR_10PERCENT_EVERY_10MINUTES",
		&returns,
	)
	return returns
}

func LambdaDeploymentConfig_LINEAR_10PERCENT_EVERY_1MINUTE() ILambdaDeploymentConfig {
	_init_.Initialize()
	var returns ILambdaDeploymentConfig
	_jsii_.StaticGet(
		"monocdk.aws_codedeploy.LambdaDeploymentConfig",
		"LINEAR_10PERCENT_EVERY_1MINUTE",
		&returns,
	)
	return returns
}

func LambdaDeploymentConfig_LINEAR_10PERCENT_EVERY_2MINUTES() ILambdaDeploymentConfig {
	_init_.Initialize()
	var returns ILambdaDeploymentConfig
	_jsii_.StaticGet(
		"monocdk.aws_codedeploy.LambdaDeploymentConfig",
		"LINEAR_10PERCENT_EVERY_2MINUTES",
		&returns,
	)
	return returns
}

func LambdaDeploymentConfig_LINEAR_10PERCENT_EVERY_3MINUTES() ILambdaDeploymentConfig {
	_init_.Initialize()
	var returns ILambdaDeploymentConfig
	_jsii_.StaticGet(
		"monocdk.aws_codedeploy.LambdaDeploymentConfig",
		"LINEAR_10PERCENT_EVERY_3MINUTES",
		&returns,
	)
	return returns
}

