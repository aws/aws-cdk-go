package awsgreengrass


// Settings that define additional Linux OS group permissions to give to the Lambda function process.
//
// You can give the permissions of the Linux group that owns the resource or choose another Linux group. These permissions are in addition to the function's `RunAs` permissions.
//
// In an AWS CloudFormation template, `GroupOwnerSetting` is a property of the [`LocalDeviceResourceData`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinition-localdeviceresourcedata.html) and [`LocalVolumeResourceData`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinition-localvolumeresourcedata.html) property types.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   groupOwnerSettingProperty := &groupOwnerSettingProperty{
//   	autoAddGroupOwner: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	groupOwner: jsii.String("groupOwner"),
//   }
//
type CfnResourceDefinition_GroupOwnerSettingProperty struct {
	// Indicates whether to give the privileges of the Linux group that owns the resource to the Lambda process.
	//
	// This gives the Lambda process the file access permissions of the Linux group.
	AutoAddGroupOwner interface{} `field:"required" json:"autoAddGroupOwner" yaml:"autoAddGroupOwner"`
	// The name of the Linux group whose privileges you want to add to the Lambda process.
	//
	// This value is ignored if `AutoAddGroupOwner` is true.
	GroupOwner *string `field:"optional" json:"groupOwner" yaml:"groupOwner"`
}

