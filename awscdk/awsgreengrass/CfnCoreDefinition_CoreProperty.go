package awsgreengrass


// A core is an AWS IoT device that runs the AWS IoT Greengrass core software and manages local processes for a Greengrass group.
//
// For more information, see [What Is AWS IoT Greengrass ?](https://docs.aws.amazon.com/greengrass/v1/developerguide/what-is-gg.html) in the *Developer Guide* .
//
// In an AWS CloudFormation template, the `Cores` property of the [`CoreDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-coredefinition-coredefinitionversion.html) property type contains a list of `Core` property types. Currently, the list can contain only one core.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   coreProperty := &CoreProperty{
//   	CertificateArn: jsii.String("certificateArn"),
//   	Id: jsii.String("id"),
//   	ThingArn: jsii.String("thingArn"),
//
//   	// the properties below are optional
//   	SyncShadow: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-coredefinition-core.html
//
type CfnCoreDefinition_CoreProperty struct {
	// The Amazon Resource Name (ARN) of the device certificate for the core.
	//
	// This X.509 certificate is used to authenticate the core with AWS IoT and AWS IoT Greengrass services.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-coredefinition-core.html#cfn-greengrass-coredefinition-core-certificatearn
	//
	CertificateArn *string `field:"required" json:"certificateArn" yaml:"certificateArn"`
	// A descriptive or arbitrary ID for the core.
	//
	// This value must be unique within the core definition version. Maximum length is 128 characters with pattern `[a-zA-Z0-9:_-]+` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-coredefinition-core.html#cfn-greengrass-coredefinition-core-id
	//
	Id *string `field:"required" json:"id" yaml:"id"`
	// The ARN of the core, which is an AWS IoT device (thing).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-coredefinition-core.html#cfn-greengrass-coredefinition-core-thingarn
	//
	ThingArn *string `field:"required" json:"thingArn" yaml:"thingArn"`
	// Indicates whether the core's local shadow is synced with the cloud automatically.
	//
	// The default is false.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-coredefinition-core.html#cfn-greengrass-coredefinition-core-syncshadow
	//
	SyncShadow interface{} `field:"optional" json:"syncShadow" yaml:"syncShadow"`
}

