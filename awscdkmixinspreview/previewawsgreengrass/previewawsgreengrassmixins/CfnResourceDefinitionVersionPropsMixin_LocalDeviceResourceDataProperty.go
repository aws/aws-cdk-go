package previewawsgreengrassmixins


// Settings for a local device resource, which represents a file under `/dev` .
//
// For more information, see [Access Local Resources with Lambda Functions](https://docs.aws.amazon.com/greengrass/v1/developerguide/access-local-resources.html) in the *Developer Guide* .
//
// In an CloudFormation template, `LocalDeviceResourceData` can be used in the [`ResourceDataContainer`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinitionversion-resourcedatacontainer.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   localDeviceResourceDataProperty := &LocalDeviceResourceDataProperty{
//   	GroupOwnerSetting: &GroupOwnerSettingProperty{
//   		AutoAddGroupOwner: jsii.Boolean(false),
//   		GroupOwner: jsii.String("groupOwner"),
//   	},
//   	SourcePath: jsii.String("sourcePath"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinitionversion-localdeviceresourcedata.html
//
type CfnResourceDefinitionVersionPropsMixin_LocalDeviceResourceDataProperty struct {
	// Settings that define additional Linux OS group permissions to give to the Lambda function process.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinitionversion-localdeviceresourcedata.html#cfn-greengrass-resourcedefinitionversion-localdeviceresourcedata-groupownersetting
	//
	GroupOwnerSetting interface{} `field:"optional" json:"groupOwnerSetting" yaml:"groupOwnerSetting"`
	// The local absolute path of the device resource.
	//
	// The source path for a device resource can refer only to a character device or block device under `/dev` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinitionversion-localdeviceresourcedata.html#cfn-greengrass-resourcedefinitionversion-localdeviceresourcedata-sourcepath
	//
	SourcePath *string `field:"optional" json:"sourcePath" yaml:"sourcePath"`
}

