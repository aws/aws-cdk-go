package awscodedeploy


// Complete base deployment config properties that are required to be supplied by the implementation of the BaseDeploymentConfig class.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var minimumHealthyHosts minimumHealthyHosts
//   var trafficRouting trafficRouting
//
//   baseDeploymentConfigProps := &baseDeploymentConfigProps{
//   	computePlatform: awscdk.Aws_codedeploy.computePlatform_SERVER,
//   	deploymentConfigName: jsii.String("deploymentConfigName"),
//   	minimumHealthyHosts: minimumHealthyHosts,
//   	trafficRouting: trafficRouting,
//   }
//
type BaseDeploymentConfigProps struct {
	// The physical, human-readable name of the Deployment Configuration.
	DeploymentConfigName *string `field:"optional" json:"deploymentConfigName" yaml:"deploymentConfigName"`
	// The destination compute platform for the deployment.
	ComputePlatform ComputePlatform `field:"optional" json:"computePlatform" yaml:"computePlatform"`
	// Minimum number of healthy hosts.
	MinimumHealthyHosts MinimumHealthyHosts `field:"optional" json:"minimumHealthyHosts" yaml:"minimumHealthyHosts"`
	// The configuration that specifies how traffic is shifted during a deployment.
	//
	// Only applicable to ECS and Lambda deployments, and must not be specified for Server deployments.
	TrafficRouting TrafficRouting `field:"optional" json:"trafficRouting" yaml:"trafficRouting"`
}

