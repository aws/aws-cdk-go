package previewawsgreengrassmixins


// A container for resource data, which defines the resource type.
//
// The container takes only one of the following supported resource data types: `LocalDeviceResourceData` , `LocalVolumeResourceData` , `SageMakerMachineLearningModelResourceData` , `S3MachineLearningModelResourceData` , or `SecretsManagerSecretResourceData` .
//
// > Only one resource type can be defined for a `ResourceDataContainer` instance.
//
// In an CloudFormation template, `ResourceDataContainer` is a property of the [`ResourceInstance`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinitionversion-resourceinstance.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   resourceDataContainerProperty := &ResourceDataContainerProperty{
//   	LocalDeviceResourceData: &LocalDeviceResourceDataProperty{
//   		GroupOwnerSetting: &GroupOwnerSettingProperty{
//   			AutoAddGroupOwner: jsii.Boolean(false),
//   			GroupOwner: jsii.String("groupOwner"),
//   		},
//   		SourcePath: jsii.String("sourcePath"),
//   	},
//   	LocalVolumeResourceData: &LocalVolumeResourceDataProperty{
//   		DestinationPath: jsii.String("destinationPath"),
//   		GroupOwnerSetting: &GroupOwnerSettingProperty{
//   			AutoAddGroupOwner: jsii.Boolean(false),
//   			GroupOwner: jsii.String("groupOwner"),
//   		},
//   		SourcePath: jsii.String("sourcePath"),
//   	},
//   	S3MachineLearningModelResourceData: &S3MachineLearningModelResourceDataProperty{
//   		DestinationPath: jsii.String("destinationPath"),
//   		OwnerSetting: &ResourceDownloadOwnerSettingProperty{
//   			GroupOwner: jsii.String("groupOwner"),
//   			GroupPermission: jsii.String("groupPermission"),
//   		},
//   		S3Uri: jsii.String("s3Uri"),
//   	},
//   	SageMakerMachineLearningModelResourceData: &SageMakerMachineLearningModelResourceDataProperty{
//   		DestinationPath: jsii.String("destinationPath"),
//   		OwnerSetting: &ResourceDownloadOwnerSettingProperty{
//   			GroupOwner: jsii.String("groupOwner"),
//   			GroupPermission: jsii.String("groupPermission"),
//   		},
//   		SageMakerJobArn: jsii.String("sageMakerJobArn"),
//   	},
//   	SecretsManagerSecretResourceData: &SecretsManagerSecretResourceDataProperty{
//   		AdditionalStagingLabelsToDownload: []*string{
//   			jsii.String("additionalStagingLabelsToDownload"),
//   		},
//   		Arn: jsii.String("arn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinitionversion-resourcedatacontainer.html
//
type CfnResourceDefinitionVersionPropsMixin_ResourceDataContainerProperty struct {
	// Settings for a local device resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinitionversion-resourcedatacontainer.html#cfn-greengrass-resourcedefinitionversion-resourcedatacontainer-localdeviceresourcedata
	//
	LocalDeviceResourceData interface{} `field:"optional" json:"localDeviceResourceData" yaml:"localDeviceResourceData"`
	// Settings for a local volume resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinitionversion-resourcedatacontainer.html#cfn-greengrass-resourcedefinitionversion-resourcedatacontainer-localvolumeresourcedata
	//
	LocalVolumeResourceData interface{} `field:"optional" json:"localVolumeResourceData" yaml:"localVolumeResourceData"`
	// Settings for a machine learning resource stored in Amazon S3 .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinitionversion-resourcedatacontainer.html#cfn-greengrass-resourcedefinitionversion-resourcedatacontainer-s3machinelearningmodelresourcedata
	//
	S3MachineLearningModelResourceData interface{} `field:"optional" json:"s3MachineLearningModelResourceData" yaml:"s3MachineLearningModelResourceData"`
	// Settings for a machine learning resource saved as an SageMaker AI training job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinitionversion-resourcedatacontainer.html#cfn-greengrass-resourcedefinitionversion-resourcedatacontainer-sagemakermachinelearningmodelresourcedata
	//
	SageMakerMachineLearningModelResourceData interface{} `field:"optional" json:"sageMakerMachineLearningModelResourceData" yaml:"sageMakerMachineLearningModelResourceData"`
	// Settings for a secret resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinitionversion-resourcedatacontainer.html#cfn-greengrass-resourcedefinitionversion-resourcedatacontainer-secretsmanagersecretresourcedata
	//
	SecretsManagerSecretResourceData interface{} `field:"optional" json:"secretsManagerSecretResourceData" yaml:"secretsManagerSecretResourceData"`
}

