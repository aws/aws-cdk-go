package awsgreengrassv2


// Contains information about a deployment's update to a component's configuration on AWS IoT Greengrass core devices.
//
// For more information, see [Update component configurations](https://docs.aws.amazon.com/greengrass/v2/developerguide/update-component-configurations.html) in the *AWS IoT Greengrass V2 Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   componentConfigurationUpdateProperty := &ComponentConfigurationUpdateProperty{
//   	Merge: jsii.String("merge"),
//   	Reset: []*string{
//   		jsii.String("reset"),
//   	},
//   }
//
type CfnDeployment_ComponentConfigurationUpdateProperty struct {
	// A serialized JSON string that contains the configuration object to merge to target devices.
	//
	// The core device merges this configuration with the component's existing configuration. If this is the first time a component deploys on a device, the core device merges this configuration with the component's default configuration. This means that the core device keeps it's existing configuration for keys and values that you don't specify in this object. For more information, see [Merge configuration updates](https://docs.aws.amazon.com/greengrass/v2/developerguide/update-component-configurations.html#merge-configuration-update) in the *AWS IoT Greengrass V2 Developer Guide* .
	Merge *string `field:"optional" json:"merge" yaml:"merge"`
	// The list of configuration nodes to reset to default values on target devices.
	//
	// Use JSON pointers to specify each node to reset. JSON pointers start with a forward slash ( `/` ) and use forward slashes to separate the key for each level in the object. For more information, see the [JSON pointer specification](https://docs.aws.amazon.com/https://tools.ietf.org/html/rfc6901) and [Reset configuration updates](https://docs.aws.amazon.com/greengrass/v2/developerguide/update-component-configurations.html#reset-configuration-update) in the *AWS IoT Greengrass V2 Developer Guide* .
	Reset *[]*string `field:"optional" json:"reset" yaml:"reset"`
}

