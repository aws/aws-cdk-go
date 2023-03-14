package awsgreengrass


// Settings for a local volume resource, which represents a file or directory on the root file system.
//
// For more information, see [Access Local Resources with Lambda Functions](https://docs.aws.amazon.com/greengrass/latest/developerguide/access-local-resources.html) in the *Developer Guide* .
//
// In an AWS CloudFormation template, `LocalVolumeResourceData` can be used in the [`ResourceDataContainer`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinition-resourcedatacontainer.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   localVolumeResourceDataProperty := &localVolumeResourceDataProperty{
//   	destinationPath: jsii.String("destinationPath"),
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
type CfnResourceDefinition_LocalVolumeResourceDataProperty struct {
	// The absolute local path of the resource in the Lambda environment.
	DestinationPath *string `field:"required" json:"destinationPath" yaml:"destinationPath"`
	// The local absolute path of the volume resource on the host.
	//
	// The source path for a volume resource type cannot start with `/sys` .
	SourcePath *string `field:"required" json:"sourcePath" yaml:"sourcePath"`
	// Settings that define additional Linux OS group permissions to give to the Lambda function process.
	GroupOwnerSetting interface{} `field:"optional" json:"groupOwnerSetting" yaml:"groupOwnerSetting"`
}

