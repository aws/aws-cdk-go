package awscodedeploy


// Construction properties of `ServerDeploymentConfig`.
//
// Example:
//   deploymentConfig := codedeploy.NewServerDeploymentConfig(this, jsii.String("DeploymentConfiguration"), &ServerDeploymentConfigProps{
//   	DeploymentConfigName: jsii.String("MyDeploymentConfiguration"),
//   	 // optional property
//   	// one of these is required, but both cannot be specified at the same time
//   	MinimumHealthyHosts: codedeploy.MinimumHealthyHosts_Count(jsii.Number(2)),
//   })
//
type ServerDeploymentConfigProps struct {
	// The physical, human-readable name of the Deployment Configuration.
	// Default: - automatically generated name.
	//
	DeploymentConfigName *string `field:"optional" json:"deploymentConfigName" yaml:"deploymentConfigName"`
	// Minimum number of healthy hosts.
	MinimumHealthyHosts MinimumHealthyHosts `field:"required" json:"minimumHealthyHosts" yaml:"minimumHealthyHosts"`
	// Configure CodeDeploy to deploy your application to one Availability Zone at a time within an AWS Region.
	// Default: - deploy your application to a random selection of hosts across a Region.
	//
	ZonalConfig *ZonalConfig `field:"optional" json:"zonalConfig" yaml:"zonalConfig"`
}

