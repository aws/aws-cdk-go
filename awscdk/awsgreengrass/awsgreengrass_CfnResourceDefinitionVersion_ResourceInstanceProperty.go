package awsgreengrass


// A local resource, machine learning resource, or secret resource.
//
// For more information, see [Access Local Resources with Lambda Functions](https://docs.aws.amazon.com/greengrass/latest/developerguide/access-local-resources.html) , [Perform Machine Learning Inference](https://docs.aws.amazon.com/greengrass/latest/developerguide/ml-inference.html) , and [Deploy Secrets to the AWS IoT Greengrass Core](https://docs.aws.amazon.com/greengrass/latest/developerguide/secrets.html) in the *Developer Guide* .
//
// In an AWS CloudFormation template, the `Resources` property of the [`AWS::Greengrass::ResourceDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-resourcedefinitionversion.html) resource contains a list of `ResourceInstance` property types.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceInstanceProperty := &resourceInstanceProperty{
//   	id: jsii.String("id"),
//   	name: jsii.String("name"),
//   	resourceDataContainer: &resourceDataContainerProperty{
//   		localDeviceResourceData: &localDeviceResourceDataProperty{
//   			sourcePath: jsii.String("sourcePath"),
//
//   			// the properties below are optional
//   			groupOwnerSetting: &groupOwnerSettingProperty{
//   				autoAddGroupOwner: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				groupOwner: jsii.String("groupOwner"),
//   			},
//   		},
//   		localVolumeResourceData: &localVolumeResourceDataProperty{
//   			destinationPath: jsii.String("destinationPath"),
//   			sourcePath: jsii.String("sourcePath"),
//
//   			// the properties below are optional
//   			groupOwnerSetting: &groupOwnerSettingProperty{
//   				autoAddGroupOwner: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				groupOwner: jsii.String("groupOwner"),
//   			},
//   		},
//   		s3MachineLearningModelResourceData: &s3MachineLearningModelResourceDataProperty{
//   			destinationPath: jsii.String("destinationPath"),
//   			s3Uri: jsii.String("s3Uri"),
//
//   			// the properties below are optional
//   			ownerSetting: &resourceDownloadOwnerSettingProperty{
//   				groupOwner: jsii.String("groupOwner"),
//   				groupPermission: jsii.String("groupPermission"),
//   			},
//   		},
//   		sageMakerMachineLearningModelResourceData: &sageMakerMachineLearningModelResourceDataProperty{
//   			destinationPath: jsii.String("destinationPath"),
//   			sageMakerJobArn: jsii.String("sageMakerJobArn"),
//
//   			// the properties below are optional
//   			ownerSetting: &resourceDownloadOwnerSettingProperty{
//   				groupOwner: jsii.String("groupOwner"),
//   				groupPermission: jsii.String("groupPermission"),
//   			},
//   		},
//   		secretsManagerSecretResourceData: &secretsManagerSecretResourceDataProperty{
//   			arn: jsii.String("arn"),
//
//   			// the properties below are optional
//   			additionalStagingLabelsToDownload: []*string{
//   				jsii.String("additionalStagingLabelsToDownload"),
//   			},
//   		},
//   	},
//   }
//
type CfnResourceDefinitionVersion_ResourceInstanceProperty struct {
	// A descriptive or arbitrary ID for the resource.
	//
	// This value must be unique within the resource definition version. Maximum length is 128 characters with pattern `[a-zA-Z0-9:_-]+` .
	Id *string `field:"required" json:"id" yaml:"id"`
	// The descriptive resource name, which is displayed on the AWS IoT Greengrass console.
	//
	// Maximum length 128 characters with pattern [a-zA-Z0-9:_-]+. This must be unique within a Greengrass group.
	Name *string `field:"required" json:"name" yaml:"name"`
	// A container for resource data.
	//
	// The container takes only one of the following supported resource data types: `LocalDeviceResourceData` , `LocalVolumeResourceData` , `SageMakerMachineLearningModelResourceData` , `S3MachineLearningModelResourceData` , or `SecretsManagerSecretResourceData` .
	//
	// > Only one resource type can be defined for a `ResourceDataContainer` instance.
	ResourceDataContainer interface{} `field:"required" json:"resourceDataContainer" yaml:"resourceDataContainer"`
}

