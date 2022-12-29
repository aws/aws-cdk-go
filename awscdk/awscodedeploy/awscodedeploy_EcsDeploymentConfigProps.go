package awscodedeploy


// Construction properties of {@link EcsDeploymentConfig}.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   codedeploy.NewEcsDeploymentConfig(this, jsii.String("CustomConfig"), &ecsDeploymentConfigProps{
//   	trafficRoutingConfig: codedeploy.NewTimeBasedCanaryTrafficRoutingConfig(map[string]interface{}{
//   		"interval": cdk.Duration_minutes(jsii.Number(15)),
//   		"percentage": jsii.Number(5),
//   	}),
//   })
//
type EcsDeploymentConfigProps struct {
	// The physical, human-readable name of the Deployment Configuration.
	DeploymentConfigName *string `field:"optional" json:"deploymentConfigName" yaml:"deploymentConfigName"`
	// The configuration that specifies how traffic is shifted from the 'blue' target group to the 'green' target group during a deployment.
	TrafficRouting TrafficRouting `field:"optional" json:"trafficRouting" yaml:"trafficRouting"`
}

