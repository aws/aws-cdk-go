package awscodedeploy

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Minimum number of healthy hosts for a server deployment.
//
// Example:
//   deploymentConfig := codedeploy.NewServerDeploymentConfig(this, jsii.String("DeploymentConfiguration"), &serverDeploymentConfigProps{
//   	deploymentConfigName: jsii.String("MyDeploymentConfiguration"),
//   	 // optional property
//   	// one of these is required, but both cannot be specified at the same time
//   	minimumHealthyHosts: codedeploy.minimumHealthyHosts.count(jsii.Number(2)),
//   })
//
type MinimumHealthyHosts interface {
}

// The jsii proxy struct for MinimumHealthyHosts
type jsiiProxy_MinimumHealthyHosts struct {
	_ byte // padding
}

// The minimum healhty hosts threshold expressed as an absolute number.
func MinimumHealthyHosts_Count(value *float64) MinimumHealthyHosts {
	_init_.Initialize()

	if err := validateMinimumHealthyHosts_CountParameters(value); err != nil {
		panic(err)
	}
	var returns MinimumHealthyHosts

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codedeploy.MinimumHealthyHosts",
		"count",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// The minmum healhty hosts threshold expressed as a percentage of the fleet.
func MinimumHealthyHosts_Percentage(value *float64) MinimumHealthyHosts {
	_init_.Initialize()

	if err := validateMinimumHealthyHosts_PercentageParameters(value); err != nil {
		panic(err)
	}
	var returns MinimumHealthyHosts

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codedeploy.MinimumHealthyHosts",
		"percentage",
		[]interface{}{value},
		&returns,
	)

	return returns
}

