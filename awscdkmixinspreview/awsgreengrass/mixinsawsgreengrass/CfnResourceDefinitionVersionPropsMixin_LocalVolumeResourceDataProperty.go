package mixinsawsgreengrass


// Settings for a local volume resource, which represents a file or directory on the root file system.
//
// For more information, see [Access Local Resources with Lambda Functions](https://docs.aws.amazon.com/greengrass/v1/developerguide/access-local-resources.html) in the *Developer Guide* .
//
// In an CloudFormation template, `LocalVolumeResourceData` can be used in the [`ResourceDataContainer`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinitionversion-resourcedatacontainer.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   localVolumeResourceDataProperty := &LocalVolumeResourceDataProperty{
//   	DestinationPath: jsii.String("destinationPath"),
//   	GroupOwnerSetting: &GroupOwnerSettingProperty{
//   		AutoAddGroupOwner: jsii.Boolean(false),
//   		GroupOwner: jsii.String("groupOwner"),
//   	},
//   	SourcePath: jsii.String("sourcePath"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinitionversion-localvolumeresourcedata.html
//
type CfnResourceDefinitionVersionPropsMixin_LocalVolumeResourceDataProperty struct {
	// The absolute local path of the resource in the Lambda environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinitionversion-localvolumeresourcedata.html#cfn-greengrass-resourcedefinitionversion-localvolumeresourcedata-destinationpath
	//
	DestinationPath *string `field:"optional" json:"destinationPath" yaml:"destinationPath"`
	// Settings that define additional Linux OS group permissions to give to the Lambda function process.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinitionversion-localvolumeresourcedata.html#cfn-greengrass-resourcedefinitionversion-localvolumeresourcedata-groupownersetting
	//
	GroupOwnerSetting interface{} `field:"optional" json:"groupOwnerSetting" yaml:"groupOwnerSetting"`
	// The local absolute path of the volume resource on the host.
	//
	// The source path for a volume resource type cannot start with `/sys` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinitionversion-localvolumeresourcedata.html#cfn-greengrass-resourcedefinitionversion-localvolumeresourcedata-sourcepath
	//
	SourcePath *string `field:"optional" json:"sourcePath" yaml:"sourcePath"`
}

