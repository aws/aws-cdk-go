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
//   resourceDataContainerProperty := &ResourceDataContainerProperty{
//   	LocalDeviceResourceData: &LocalDeviceResourceDataProperty{
//   		SourcePath: jsii.String("sourcePath"),
//
//   		// the properties below are optional
//   		GroupOwnerSetting: &GroupOwnerSettingProperty{
//   			AutoAddGroupOwner: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			GroupOwner: jsii.String("groupOwner"),
//   		},
//   	},
//   	LocalVolumeResourceData: &LocalVolumeResourceDataProperty{
//   		DestinationPath: jsii.String("destinationPath"),
//   		SourcePath: jsii.String("sourcePath"),
//
//   		// the properties below are optional
//   		GroupOwnerSetting: &GroupOwnerSettingProperty{
//   			AutoAddGroupOwner: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			GroupOwner: jsii.String("groupOwner"),
//   		},
//   	},
//   	S3MachineLearningModelResourceData: &S3MachineLearningModelResourceDataProperty{
//   		DestinationPath: jsii.String("destinationPath"),
//   		S3Uri: jsii.String("s3Uri"),
//
//   		// the properties below are optional
//   		OwnerSetting: &ResourceDownloadOwnerSettingProperty{
//   			GroupOwner: jsii.String("groupOwner"),
//   			GroupPermission: jsii.String("groupPermission"),
//   		},
//   	},
//   	SageMakerMachineLearningModelResourceData: &SageMakerMachineLearningModelResourceDataProperty{
//   		DestinationPath: jsii.String("destinationPath"),
//   		SageMakerJobArn: jsii.String("sageMakerJobArn"),
//
//   		// the properties below are optional
//   		OwnerSetting: &ResourceDownloadOwnerSettingProperty{
//   			GroupOwner: jsii.String("groupOwner"),
//   			GroupPermission: jsii.String("groupPermission"),
//   		},
//   	},
//   	SecretsManagerSecretResourceData: &SecretsManagerSecretResourceDataProperty{
//   		Arn: jsii.String("arn"),
//
//   		// the properties below are optional
//   		AdditionalStagingLabelsToDownload: []*string{
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

