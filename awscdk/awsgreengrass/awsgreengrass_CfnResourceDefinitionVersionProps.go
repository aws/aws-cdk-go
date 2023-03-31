package awsgreengrass


// Properties for defining a `CfnResourceDefinitionVersion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResourceDefinitionVersionProps := &cfnResourceDefinitionVersionProps{
//   	resourceDefinitionId: jsii.String("resourceDefinitionId"),
//   	resources: []interface{}{
//   		&resourceInstanceProperty{
//   			id: jsii.String("id"),
//   			name: jsii.String("name"),
//   			resourceDataContainer: &resourceDataContainerProperty{
//   				localDeviceResourceData: &localDeviceResourceDataProperty{
//   					sourcePath: jsii.String("sourcePath"),
//
//   					// the properties below are optional
//   					groupOwnerSetting: &groupOwnerSettingProperty{
//   						autoAddGroupOwner: jsii.Boolean(false),
//
//   						// the properties below are optional
//   						groupOwner: jsii.String("groupOwner"),
//   					},
//   				},
//   				localVolumeResourceData: &localVolumeResourceDataProperty{
//   					destinationPath: jsii.String("destinationPath"),
//   					sourcePath: jsii.String("sourcePath"),
//
//   					// the properties below are optional
//   					groupOwnerSetting: &groupOwnerSettingProperty{
//   						autoAddGroupOwner: jsii.Boolean(false),
//
//   						// the properties below are optional
//   						groupOwner: jsii.String("groupOwner"),
//   					},
//   				},
//   				s3MachineLearningModelResourceData: &s3MachineLearningModelResourceDataProperty{
//   					destinationPath: jsii.String("destinationPath"),
//   					s3Uri: jsii.String("s3Uri"),
//
//   					// the properties below are optional
//   					ownerSetting: &resourceDownloadOwnerSettingProperty{
//   						groupOwner: jsii.String("groupOwner"),
//   						groupPermission: jsii.String("groupPermission"),
//   					},
//   				},
//   				sageMakerMachineLearningModelResourceData: &sageMakerMachineLearningModelResourceDataProperty{
//   					destinationPath: jsii.String("destinationPath"),
//   					sageMakerJobArn: jsii.String("sageMakerJobArn"),
//
//   					// the properties below are optional
//   					ownerSetting: &resourceDownloadOwnerSettingProperty{
//   						groupOwner: jsii.String("groupOwner"),
//   						groupPermission: jsii.String("groupPermission"),
//   					},
//   				},
//   				secretsManagerSecretResourceData: &secretsManagerSecretResourceDataProperty{
//   					arn: jsii.String("arn"),
//
//   					// the properties below are optional
//   					additionalStagingLabelsToDownload: []*string{
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

