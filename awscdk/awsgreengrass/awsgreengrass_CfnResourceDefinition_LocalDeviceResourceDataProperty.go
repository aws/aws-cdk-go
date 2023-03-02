package awsgreengrass


// Settings for a local device resource, which represents a file under `/dev` .
//
// For more information, see [Access Local Resources with Lambda Functions](https://docs.aws.amazon.com/greengrass/latest/developerguide/access-local-resources.html) in the *Developer Guide* .
//
// In an AWS CloudFormation template, `LocalDeviceResourceData` can be used in the [`ResourceDataContainer`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinition-resourcedatacontainer.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   localDeviceResourceDataProperty := &localDeviceResourceDataProperty{
//   	sourcePath: jsii.String("sourcePath"),
//
//   	// the properties below are optional
//   	groupOwnerSetting: &groupOwnerSettingProperty{
//   		autoAddGroupOwner: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		groupOwner: jsii.String("groupOwner"),
//   	},
//   }
//
type CfnResourceDefinition_LocalDeviceResourceDataProperty struct {
	// The local absolute path of the device resource.
	//
	// The source path for a device resource can refer only to a character device or block device under `/dev` .
	SourcePath *string `field:"required" json:"sourcePath" yaml:"sourcePath"`
	// Settings that define additional Linux OS group permissions to give to the Lambda function process.
	GroupOwnerSetting interface{} `field:"optional" json:"groupOwnerSetting" yaml:"groupOwnerSetting"`
}

