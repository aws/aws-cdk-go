package awsgreengrassv2


// Contains information about a component to deploy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   componentDeploymentSpecificationProperty := &componentDeploymentSpecificationProperty{
//   	componentVersion: jsii.String("componentVersion"),
//   	configurationUpdate: &componentConfigurationUpdateProperty{
//   		merge: jsii.String("merge"),
//   		reset: []*string{
//   			jsii.String("reset"),
//   		},
//   	},
//   	runWith: &componentRunWithProperty{
//   		posixUser: jsii.String("posixUser"),
//   		systemResourceLimits: &systemResourceLimitsProperty{
//   			cpus: jsii.Number(123),
//   			memory: jsii.Number(123),
//   		},
//   		windowsUser: jsii.String("windowsUser"),
//   	},
//   }
//
type CfnDeployment_ComponentDeploymentSpecificationProperty struct {
	// The version of the component.
	ComponentVersion *string `field:"optional" json:"componentVersion" yaml:"componentVersion"`
	// The configuration updates to deploy for the component.
	//
	// You can define reset updates and merge updates. A reset updates the keys that you specify to the default configuration for the component. A merge updates the core device's component configuration with the keys and values that you specify. The AWS IoT Greengrass Core software applies reset updates before it applies merge updates. For more information, see [Update component configuration](https://docs.aws.amazon.com/greengrass/v2/developerguide/update-component-configurations.html) .
	ConfigurationUpdate interface{} `field:"optional" json:"configurationUpdate" yaml:"configurationUpdate"`
	// The system user and group that the  software uses to run component processes on the core device.
	//
	// If you omit this parameter, the  software uses the system user and group that you configure for the core device. For more information, see [Configure the user and group that run components](https://docs.aws.amazon.com/greengrass/v2/developerguide/configure-greengrass-core-v2.html#configure-component-user) in the *AWS IoT Greengrass V2 Developer Guide* .
	RunWith interface{} `field:"optional" json:"runWith" yaml:"runWith"`
}

