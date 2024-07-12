package awscodedeploy

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Minimum number of healthy hosts per availability zone for a server deployment.
//
// Example:
//   deploymentConfig := codedeploy.NewServerDeploymentConfig(this, jsii.String("DeploymentConfiguration"), &ServerDeploymentConfigProps{
//   	MinimumHealthyHosts: codedeploy.MinimumHealthyHosts_Count(jsii.Number(2)),
//   	ZonalConfig: &ZonalConfig{
//   		MonitorDuration: awscdk.Duration_Minutes(jsii.Number(30)),
//   		FirstZoneMonitorDuration: awscdk.Duration_*Minutes(jsii.Number(60)),
//   		MinimumHealthyHostsPerZone: codedeploy.MinimumHealthyHostsPerZone_Count(jsii.Number(1)),
//   	},
//   })
//
type MinimumHealthyHostsPerZone interface {
}

// The jsii proxy struct for MinimumHealthyHostsPerZone
type jsiiProxy_MinimumHealthyHostsPerZone struct {
	_ byte // padding
}

// The minimum healthy hosts threshold expressed as an absolute number.
func MinimumHealthyHostsPerZone_Count(value *float64) MinimumHealthyHostsPerZone {
	_init_.Initialize()

	if err := validateMinimumHealthyHostsPerZone_CountParameters(value); err != nil {
		panic(err)
	}
	var returns MinimumHealthyHostsPerZone

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codedeploy.MinimumHealthyHostsPerZone",
		"count",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// The minimum healthy hosts threshold expressed as a percentage of the fleet.
func MinimumHealthyHostsPerZone_Percentage(value *float64) MinimumHealthyHostsPerZone {
	_init_.Initialize()

	if err := validateMinimumHealthyHostsPerZone_PercentageParameters(value); err != nil {
		panic(err)
	}
	var returns MinimumHealthyHostsPerZone

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codedeploy.MinimumHealthyHostsPerZone",
		"percentage",
		[]interface{}{value},
		&returns,
	)

	return returns
}

