package awscodedeploy


// Construction properties of {@link ServerDeploymentConfig}.
//
// Example:
//   deploymentConfig := codedeploy.NewServerDeploymentConfig(this, jsii.String("DeploymentConfiguration"), &serverDeploymentConfigProps{
//   	deploymentConfigName: jsii.String("MyDeploymentConfiguration"),
//   	 // optional property
//   	// one of these is required, but both cannot be specified at the same time
//   	minimumHealthyHosts: codedeploy.minimumHealthyHosts.count(jsii.Number(2)),
//   })
//
type ServerDeploymentConfigProps struct {
	// The physical, human-readable name of the Deployment Configuration.
	DeploymentConfigName *string `field:"optional" json:"deploymentConfigName" yaml:"deploymentConfigName"`
	// Minimum number of healthy hosts.
	MinimumHealthyHosts MinimumHealthyHosts `field:"required" json:"minimumHealthyHosts" yaml:"minimumHealthyHosts"`
}

