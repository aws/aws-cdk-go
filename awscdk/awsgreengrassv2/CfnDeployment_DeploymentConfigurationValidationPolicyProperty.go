package awsgreengrassv2


// Contains information about how long a component on a core device can validate its configuration updates before it times out.
//
// Components can use the [SubscribeToValidateConfigurationUpdates](https://docs.aws.amazon.com/greengrass/v2/developerguide/interprocess-communication.html#ipc-operation-subscribetovalidateconfigurationupdates) IPC operation to receive notifications when a deployment specifies a configuration update. Then, components can respond with the [SendConfigurationValidityReport](https://docs.aws.amazon.com/greengrass/v2/developerguide/interprocess-communication.html#ipc-operation-sendconfigurationvalidityreport) IPC operation. For more information, see the [Create deployments](https://docs.aws.amazon.com/greengrass/v2/latest/developerguide/create-deployments.html) in the *AWS IoT Greengrass V2 Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deploymentConfigurationValidationPolicyProperty := &DeploymentConfigurationValidationPolicyProperty{
//   	TimeoutInSeconds: jsii.Number(123),
//   }
//
type CfnDeployment_DeploymentConfigurationValidationPolicyProperty struct {
	// The amount of time in seconds that a component can validate its configuration updates.
	//
	// If the validation time exceeds this timeout, then the deployment proceeds for the device.
	//
	// Default: `30`.
	TimeoutInSeconds *float64 `field:"optional" json:"timeoutInSeconds" yaml:"timeoutInSeconds"`
}

