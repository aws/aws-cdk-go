package awscodedeploy

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Configuration for CodeDeploy to deploy your application to one Availability Zone at a time within an AWS Region.
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
type ZonalConfig struct {
	// The period of time that CodeDeploy must wait after completing a deployment to the first Availability Zone.
	//
	// Accepted Values:
	//  * 0
	// * Greater than or equal to 1.
	// Default: - the same value as `monitorDuration`.
	//
	FirstZoneMonitorDuration awscdk.Duration `field:"optional" json:"firstZoneMonitorDuration" yaml:"firstZoneMonitorDuration"`
	// The number or percentage of instances that must remain available per Availability Zone during a deployment.
	//
	// This option works in conjunction with the `minimumHealthyHosts` option.
	// See: https://docs.aws.amazon.com/codedeploy/latest/userguide/instances-health.html#minimum-healthy-hosts-az
	//
	// Default: - 0 percent.
	//
	MinimumHealthyHostsPerZone MinimumHealthyHostsPerZone `field:"optional" json:"minimumHealthyHostsPerZone" yaml:"minimumHealthyHostsPerZone"`
	// The period of time that CodeDeploy must wait after completing a deployment to an Availability Zone.
	//
	// Accepted Values:
	//  * 0
	// * Greater than or equal to 1.
	// Default: - CodeDeploy starts deploying to the next Availability Zone immediately.
	//
	MonitorDuration awscdk.Duration `field:"optional" json:"monitorDuration" yaml:"monitorDuration"`
}

