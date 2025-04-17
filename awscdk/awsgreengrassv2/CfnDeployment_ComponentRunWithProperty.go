package awsgreengrassv2


// Contains information system user and group that the AWS IoT Greengrass Core software uses to run component processes on the core device.
//
// For more information, see [Configure the user and group that run components](https://docs.aws.amazon.com/greengrass/v2/developerguide/configure-greengrass-core-v2.html#configure-component-user) in the *AWS IoT Greengrass V2 Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   componentRunWithProperty := &ComponentRunWithProperty{
//   	PosixUser: jsii.String("posixUser"),
//   	SystemResourceLimits: &SystemResourceLimitsProperty{
//   		Cpus: jsii.Number(123),
//   		Memory: jsii.Number(123),
//   	},
//   	WindowsUser: jsii.String("windowsUser"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-deployment-componentrunwith.html
//
type CfnDeployment_ComponentRunWithProperty struct {
	// The POSIX system user and (optional) group to use to run this component.
	//
	// Specify the user and group separated by a colon ( `:` ) in the following format: `user:group` . The group is optional. If you don't specify a group, the AWS IoT Greengrass Core software uses the primary user for the group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-deployment-componentrunwith.html#cfn-greengrassv2-deployment-componentrunwith-posixuser
	//
	PosixUser *string `field:"optional" json:"posixUser" yaml:"posixUser"`
	// The system resource limits to apply to this component's process on the core device.
	//
	// AWS IoT Greengrass supports this feature only on Linux core devices.
	//
	// If you omit this parameter, the AWS IoT Greengrass Core software uses the default system resource limits that you configure on the AWS IoT Greengrass nucleus component. For more information, see [Configure system resource limits for components](https://docs.aws.amazon.com/greengrass/v2/developerguide/configure-greengrass-core-v2.html#configure-component-system-resource-limits) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-deployment-componentrunwith.html#cfn-greengrassv2-deployment-componentrunwith-systemresourcelimits
	//
	SystemResourceLimits interface{} `field:"optional" json:"systemResourceLimits" yaml:"systemResourceLimits"`
	// The Windows user to use to run this component on Windows core devices.
	//
	// The user must exist on each Windows core device, and its name and password must be in the LocalSystem account's Credentials Manager instance.
	//
	// If you omit this parameter, the AWS IoT Greengrass Core software uses the default Windows user that you configure on the AWS IoT Greengrass nucleus component. For more information, see [Configure the user and group that run components](https://docs.aws.amazon.com/greengrass/v2/developerguide/configure-greengrass-core-v2.html#configure-component-user) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-deployment-componentrunwith.html#cfn-greengrassv2-deployment-componentrunwith-windowsuser
	//
	WindowsUser *string `field:"optional" json:"windowsUser" yaml:"windowsUser"`
}

