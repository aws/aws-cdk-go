package previewawsgreengrassmixins


// A local resource, machine learning resource, or secret resource.
//
// For more information, see [Access Local Resources with Lambda Functions](https://docs.aws.amazon.com/greengrass/v1/developerguide/access-local-resources.html) , [Perform Machine Learning Inference](https://docs.aws.amazon.com/greengrass/v1/developerguide/ml-inference.html) , and [Deploy Secrets to the AWS IoT Greengrass Core](https://docs.aws.amazon.com/greengrass/v1/developerguide/secrets.html) in the *Developer Guide* .
//
// In an CloudFormation template, the `Resources` property of the [`AWS::Greengrass::ResourceDefinition`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-resourcedefinition.html) resource contains a list of `ResourceInstance` property types.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   resourceInstanceProperty := &ResourceInstanceProperty{
//   	Id: jsii.String("id"),
//   	Name: jsii.String("name"),
//   	ResourceDataContainer: &ResourceDataContainerProperty{
//   		LocalDeviceResourceData: &LocalDeviceResourceDataProperty{
//   			GroupOwnerSetting: &GroupOwnerSettingProperty{
//   				AutoAddGroupOwner: jsii.Boolean(false),
//   				GroupOwner: jsii.String("groupOwner"),
//   			},
//   			SourcePath: jsii.String("sourcePath"),
//   		},
//   		LocalVolumeResourceData: &LocalVolumeResourceDataProperty{
//   			DestinationPath: jsii.String("destinationPath"),
//   			GroupOwnerSetting: &GroupOwnerSettingProperty{
//   				AutoAddGroupOwner: jsii.Boolean(false),
//   				GroupOwner: jsii.String("groupOwner"),
//   			},
//   			SourcePath: jsii.String("sourcePath"),
//   		},
//   		S3MachineLearningModelResourceData: &S3MachineLearningModelResourceDataProperty{
//   			DestinationPath: jsii.String("destinationPath"),
//   			OwnerSetting: &ResourceDownloadOwnerSettingProperty{
//   				GroupOwner: jsii.String("groupOwner"),
//   				GroupPermission: jsii.String("groupPermission"),
//   			},
//   			S3Uri: jsii.String("s3Uri"),
//   		},
//   		SageMakerMachineLearningModelResourceData: &SageMakerMachineLearningModelResourceDataProperty{
//   			DestinationPath: jsii.String("destinationPath"),
//   			OwnerSetting: &ResourceDownloadOwnerSettingProperty{
//   				GroupOwner: jsii.String("groupOwner"),
//   				GroupPermission: jsii.String("groupPermission"),
//   			},
//   			SageMakerJobArn: jsii.String("sageMakerJobArn"),
//   		},
//   		SecretsManagerSecretResourceData: &SecretsManagerSecretResourceDataProperty{
//   			AdditionalStagingLabelsToDownload: []*string{
//   				jsii.String("additionalStagingLabelsToDownload"),
//   			},
//   			Arn: jsii.String("arn"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinition-resourceinstance.html
//
type CfnResourceDefinitionPropsMixin_ResourceInstanceProperty struct {
	// A descriptive or arbitrary ID for the resource.
	//
	// This value must be unique within the resource definition version. Maximum length is 128 characters with pattern `[a-zA-Z0-9:_-]+` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinition-resourceinstance.html#cfn-greengrass-resourcedefinition-resourceinstance-id
	//
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The descriptive resource name, which is displayed on the AWS IoT Greengrass console.
	//
	// Maximum length 128 characters with pattern [a-zA-Z0-9:_-]+. This must be unique within a Greengrass group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinition-resourceinstance.html#cfn-greengrass-resourcedefinition-resourceinstance-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A container for resource data.
	//
	// The container takes only one of the following supported resource data types: `LocalDeviceResourceData` , `LocalVolumeResourceData` , `SageMakerMachineLearningModelResourceData` , `S3MachineLearningModelResourceData` , or `SecretsManagerSecretResourceData` .
	//
	// > Only one resource type can be defined for a `ResourceDataContainer` instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinition-resourceinstance.html#cfn-greengrass-resourcedefinition-resourceinstance-resourcedatacontainer
	//
	ResourceDataContainer interface{} `field:"optional" json:"resourceDataContainer" yaml:"resourceDataContainer"`
}

