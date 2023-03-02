package awsgreengrass


// A container for resource data, which defines the resource type.
//
// The container takes only one of the following supported resource data types: `LocalDeviceResourceData` , `LocalVolumeResourceData` , `SageMakerMachineLearningModelResourceData` , `S3MachineLearningModelResourceData` , or `SecretsManagerSecretResourceData` .
//
// > Only one resource type can be defined for a `ResourceDataContainer` instance.
//
// In an AWS CloudFormation template, `ResourceDataContainer` is a property of the [`ResourceInstance`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinition-resourceinstance.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceDataContainerProperty := &resourceDataContainerProperty{
//   	localDeviceResourceData: &localDeviceResourceDataProperty{
//   		sourcePath: jsii.String("sourcePath"),
//
//   		// the properties below are optional
//   		groupOwnerSetting: &groupOwnerSettingProperty{
//   			autoAddGroupOwner: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			groupOwner: jsii.String("groupOwner"),
//   		},
//   	},
//   	localVolumeResourceData: &localVolumeResourceDataProperty{
//   		destinationPath: jsii.String("destinationPath"),
//   		sourcePath: jsii.String("sourcePath"),
//
//   		// the properties below are optional
//   		groupOwnerSetting: &groupOwnerSettingProperty{
//   			autoAddGroupOwner: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			groupOwner: jsii.String("groupOwner"),
//   		},
//   	},
//   	s3MachineLearningModelResourceData: &s3MachineLearningModelResourceDataProperty{
//   		destinationPath: jsii.String("destinationPath"),
//   		s3Uri: jsii.String("s3Uri"),
//
//   		// the properties below are optional
//   		ownerSetting: &resourceDownloadOwnerSettingProperty{
//   			groupOwner: jsii.String("groupOwner"),
//   			groupPermission: jsii.String("groupPermission"),
//   		},
//   	},
//   	sageMakerMachineLearningModelResourceData: &sageMakerMachineLearningModelResourceDataProperty{
//   		destinationPath: jsii.String("destinationPath"),
//   		sageMakerJobArn: jsii.String("sageMakerJobArn"),
//
//   		// the properties below are optional
//   		ownerSetting: &resourceDownloadOwnerSettingProperty{
//   			groupOwner: jsii.String("groupOwner"),
//   			groupPermission: jsii.String("groupPermission"),
//   		},
//   	},
//   	secretsManagerSecretResourceData: &secretsManagerSecretResourceDataProperty{
//   		arn: jsii.String("arn"),
//
//   		// the properties below are optional
//   		additionalStagingLabelsToDownload: []*string{
//   			jsii.String("additionalStagingLabelsToDownload"),
//   		},
//   	},
//   }
//
type CfnResourceDefinition_ResourceDataContainerProperty struct {
	// Settings for a local device resource.
	LocalDeviceResourceData interface{} `field:"optional" json:"localDeviceResourceData" yaml:"localDeviceResourceData"`
	// Settings for a local volume resource.
	LocalVolumeResourceData interface{} `field:"optional" json:"localVolumeResourceData" yaml:"localVolumeResourceData"`
	// Settings for a machine learning resource stored in Amazon S3 .
	S3MachineLearningModelResourceData interface{} `field:"optional" json:"s3MachineLearningModelResourceData" yaml:"s3MachineLearningModelResourceData"`
	// Settings for a machine learning resource saved as an SageMaker training job.
	SageMakerMachineLearningModelResourceData interface{} `field:"optional" json:"sageMakerMachineLearningModelResourceData" yaml:"sageMakerMachineLearningModelResourceData"`
	// Settings for a secret resource.
	SecretsManagerSecretResourceData interface{} `field:"optional" json:"secretsManagerSecretResourceData" yaml:"secretsManagerSecretResourceData"`
}

