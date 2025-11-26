package previewawsgreengrassmixins


// The owner setting for a downloaded machine learning resource.
//
// For more information, see [Access Machine Learning Resources from Lambda Functions](https://docs.aws.amazon.com/greengrass/v1/developerguide/access-ml-resources.html) in the *Developer Guide* .
//
// In an CloudFormation template, `ResourceDownloadOwnerSetting` is the property type of the `OwnerSetting` property for the [`S3MachineLearningModelResourceData`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinitionversion-s3machinelearningmodelresourcedata.html) and [`SageMakerMachineLearningModelResourceData`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinitionversion-sagemakermachinelearningmodelresourcedata.html) property types.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   resourceDownloadOwnerSettingProperty := &ResourceDownloadOwnerSettingProperty{
//   	GroupOwner: jsii.String("groupOwner"),
//   	GroupPermission: jsii.String("groupPermission"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinitionversion-resourcedownloadownersetting.html
//
type CfnResourceDefinitionVersionPropsMixin_ResourceDownloadOwnerSettingProperty struct {
	// The group owner of the machine learning resource.
	//
	// This is the group ID (GID) of an existing Linux OS group on the system. The group's permissions are added to the Lambda process.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinitionversion-resourcedownloadownersetting.html#cfn-greengrass-resourcedefinitionversion-resourcedownloadownersetting-groupowner
	//
	GroupOwner *string `field:"optional" json:"groupOwner" yaml:"groupOwner"`
	// The permissions that the group owner has to the machine learning resource.
	//
	// Valid values are `rw` (read-write) or `ro` (read-only).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinitionversion-resourcedownloadownersetting.html#cfn-greengrass-resourcedefinitionversion-resourcedownloadownersetting-grouppermission
	//
	GroupPermission *string `field:"optional" json:"groupPermission" yaml:"groupPermission"`
}

