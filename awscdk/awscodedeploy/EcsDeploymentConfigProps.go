package awscodedeploy


// Construction properties of `EcsDeploymentConfig`.
//
// Example:
//   codedeploy.NewEcsDeploymentConfig(this, jsii.String("CustomConfig"), &EcsDeploymentConfigProps{
//   	TrafficRouting: codedeploy.NewTimeBasedCanaryTrafficRouting(&TimeBasedCanaryTrafficRoutingProps{
//   		Interval: awscdk.Duration_Minutes(jsii.Number(15)),
//   		Percentage: jsii.Number(5),
//   	}),
//   })
//
type EcsDeploymentConfigProps struct {
	// The physical, human-readable name of the Deployment Configuration.
	// Default: - automatically generated name.
	//
	DeploymentConfigName *string `field:"optional" json:"deploymentConfigName" yaml:"deploymentConfigName"`
	// The configuration that specifies how traffic is shifted from the 'blue' target group to the 'green' target group during a deployment.
	// Default: AllAtOnce.
	//
	TrafficRouting TrafficRouting `field:"optional" json:"trafficRouting" yaml:"trafficRouting"`
}

