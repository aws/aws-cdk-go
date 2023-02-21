package awscodedeploy


// Information about blue/green deployment options for a deployment group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   blueGreenDeploymentConfigurationProperty := &blueGreenDeploymentConfigurationProperty{
//   	deploymentReadyOption: &deploymentReadyOptionProperty{
//   		actionOnTimeout: jsii.String("actionOnTimeout"),
//   		waitTimeInMinutes: jsii.Number(123),
//   	},
//   	greenFleetProvisioningOption: &greenFleetProvisioningOptionProperty{
//   		action: jsii.String("action"),
//   	},
//   	terminateBlueInstancesOnDeploymentSuccess: &blueInstanceTerminationOptionProperty{
//   		action: jsii.String("action"),
//   		terminationWaitTimeInMinutes: jsii.Number(123),
//   	},
//   }
//
type CfnDeploymentGroup_BlueGreenDeploymentConfigurationProperty struct {
	// Information about the action to take when newly provisioned instances are ready to receive traffic in a blue/green deployment.
	DeploymentReadyOption interface{} `field:"optional" json:"deploymentReadyOption" yaml:"deploymentReadyOption"`
	// Information about how instances are provisioned for a replacement environment in a blue/green deployment.
	GreenFleetProvisioningOption interface{} `field:"optional" json:"greenFleetProvisioningOption" yaml:"greenFleetProvisioningOption"`
	// Information about whether to terminate instances in the original fleet during a blue/green deployment.
	TerminateBlueInstancesOnDeploymentSuccess interface{} `field:"optional" json:"terminateBlueInstancesOnDeploymentSuccess" yaml:"terminateBlueInstancesOnDeploymentSuccess"`
}

