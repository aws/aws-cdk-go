package awsgreengrass


// The owner setting for a downloaded machine learning resource.
//
// For more information, see [Access Machine Learning Resources from Lambda Functions](https://docs.aws.amazon.com/greengrass/latest/developerguide/access-ml-resources.html) in the *Developer Guide* .
//
// In an AWS CloudFormation template, `ResourceDownloadOwnerSetting` is the property type of the `OwnerSetting` property for the [`S3MachineLearningModelResourceData`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinition-s3machinelearningmodelresourcedata.html) and [`SageMakerMachineLearningModelResourceData`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinition-sagemakermachinelearningmodelresourcedata.html) property types.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceDownloadOwnerSettingProperty := &resourceDownloadOwnerSettingProperty{
//   	groupOwner: jsii.String("groupOwner"),
//   	groupPermission: jsii.String("groupPermission"),
//   }
//
type CfnResourceDefinition_ResourceDownloadOwnerSettingProperty struct {
	// The group owner of the machine learning resource.
	//
	// This is the group ID (GID) of an existing Linux OS group on the system. The group's permissions are added to the Lambda process.
	GroupOwner *string `field:"required" json:"groupOwner" yaml:"groupOwner"`
	// The permissions that the group owner has to the machine learning resource.
	//
	// Valid values are `rw` (read-write) or `ro` (read-only).
	GroupPermission *string `field:"required" json:"groupPermission" yaml:"groupPermission"`
}

