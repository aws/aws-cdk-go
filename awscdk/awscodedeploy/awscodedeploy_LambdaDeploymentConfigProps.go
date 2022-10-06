package awscodedeploy


// Construction properties of {@link LambdaDeploymentConfig}.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var application lambdaApplication
//   var alias alias
//   config := codedeploy.NewLambdaDeploymentConfig(this, jsii.String("CustomConfig"), &lambdaDeploymentConfigProps{
//   	trafficRoutingConfig: codedeploy.NewTimeBasedCanaryTrafficRoutingConfig(map[string]interface{}{
//   		"interval": cdk.Duration_minutes(jsii.Number(15)),
//   		"percentage": jsii.Number(5),
//   	}),
//   })
//   deploymentGroup := codedeploy.NewLambdaDeploymentGroup(this, jsii.String("BlueGreenDeployment"), &lambdaDeploymentGroupProps{
//   	application: application,
//   	alias: alias,
//   	deploymentConfig: config,
//   })
//
type LambdaDeploymentConfigProps struct {
	// The physical, human-readable name of the Deployment Configuration.
	DeploymentConfigName *string `field:"optional" json:"deploymentConfigName" yaml:"deploymentConfigName"`
	// The configuration that specifies how traffic is shifted from the 'blue' target group to the 'green' target group during a deployment.
	TrafficRouting TrafficRouting `field:"optional" json:"trafficRouting" yaml:"trafficRouting"`
}

