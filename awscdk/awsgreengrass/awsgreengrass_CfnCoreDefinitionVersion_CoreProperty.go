package awsgreengrass


// A core is an AWS IoT device that runs the AWS IoT Greengrass core software and manages local processes for a Greengrass group.
//
// For more information, see [What Is AWS IoT Greengrass ?](https://docs.aws.amazon.com/greengrass/latest/developerguide/what-is-gg.html) in the *Developer Guide* .
//
// In an AWS CloudFormation template, the `Cores` property of the [`AWS::Greengrass::CoreDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-coredefinitionversion.html) resource contains a list of `Core` property types. Currently, the list can contain only one core.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   coreProperty := &coreProperty{
//   	certificateArn: jsii.String("certificateArn"),
//   	id: jsii.String("id"),
//   	thingArn: jsii.String("thingArn"),
//
//   	// the properties below are optional
//   	syncShadow: jsii.Boolean(false),
//   }
//
type CfnCoreDefinitionVersion_CoreProperty struct {
	// The ARN of the device certificate for the core.
	//
	// This X.509 certificate is used to authenticate the core with AWS IoT and AWS IoT Greengrass services.
	CertificateArn *string `field:"required" json:"certificateArn" yaml:"certificateArn"`
	// A descriptive or arbitrary ID for the core.
	//
	// This value must be unique within the core definition version. Maximum length is 128 characters with pattern `[a-zA-Z0-9:_-]+` .
	Id *string `field:"required" json:"id" yaml:"id"`
	// The Amazon Resource Name (ARN) of the core, which is an AWS IoT device (thing).
	ThingArn *string `field:"required" json:"thingArn" yaml:"thingArn"`
	// Indicates whether the core's local shadow is synced with the cloud automatically.
	//
	// The default is false.
	SyncShadow interface{} `field:"optional" json:"syncShadow" yaml:"syncShadow"`
}

