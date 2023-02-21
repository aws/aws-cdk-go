package awsgreengrass


// Properties for defining a `CfnResourceDefinitionVersion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResourceDefinitionVersionProps := &CfnResourceDefinitionVersionProps{
//   	ResourceDefinitionId: jsii.String("resourceDefinitionId"),
//   	Resources: []interface{}{
//   		&ResourceInstanceProperty{
//   			Id: jsii.String("id"),
//   			Name: jsii.String("name"),
//   			ResourceDataContainer: &ResourceDataContainerProperty{
//   				LocalDeviceResourceData: &LocalDeviceResourceDataProperty{
//   					SourcePath: jsii.String("sourcePath"),
//
//   					// the properties below are optional
//   					GroupOwnerSetting: &GroupOwnerSettingProperty{
//   						AutoAddGroupOwner: jsii.Boolean(false),
//
//   						// the properties below are optional
//   						GroupOwner: jsii.String("groupOwner"),
//   					},
//   				},
//   				LocalVolumeResourceData: &LocalVolumeResourceDataProperty{
//   					DestinationPath: jsii.String("destinationPath"),
//   					SourcePath: jsii.String("sourcePath"),
//
//   					// the properties below are optional
//   					GroupOwnerSetting: &GroupOwnerSettingProperty{
//   						AutoAddGroupOwner: jsii.Boolean(false),
//
//   						// the properties below are optional
//   						GroupOwner: jsii.String("groupOwner"),
//   					},
//   				},
//   				S3MachineLearningModelResourceData: &S3MachineLearningModelResourceDataProperty{
//   					DestinationPath: jsii.String("destinationPath"),
//   					S3Uri: jsii.String("s3Uri"),
//
//   					// the properties below are optional
//   					OwnerSetting: &ResourceDownloadOwnerSettingProperty{
//   						GroupOwner: jsii.String("groupOwner"),
//   						GroupPermission: jsii.String("groupPermission"),
//   					},
//   				},
//   				SageMakerMachineLearningModelResourceData: &SageMakerMachineLearningModelResourceDataProperty{
//   					DestinationPath: jsii.String("destinationPath"),
//   					SageMakerJobArn: jsii.String("sageMakerJobArn"),
//
//   					// the properties below are optional
//   					OwnerSetting: &ResourceDownloadOwnerSettingProperty{
//   						GroupOwner: jsii.String("groupOwner"),
//   						GroupPermission: jsii.String("groupPermission"),
//   					},
//   				},
//   				SecretsManagerSecretResourceData: &SecretsManagerSecretResourceDataProperty{
//   					Arn: jsii.String("arn"),
//
//   					// the properties below are optional
//   					AdditionalStagingLabelsToDownload: []*string{
//   						jsii.String("additionalStagingLabelsToDownload"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnResourceDefinitionVersionProps struct {
	// The ID of the resource definition associated with this version.
	//
	// This value is a GUID.
	ResourceDefinitionId *string `field:"required" json:"resourceDefinitionId" yaml:"resourceDefinitionId"`
	// The resources in this version.
	Resources interface{} `field:"required" json:"resources" yaml:"resources"`
}

