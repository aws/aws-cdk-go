package mixinsawsgreengrass


// Properties for CfnResourceDefinitionVersionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnResourceDefinitionVersionMixinProps := &CfnResourceDefinitionVersionMixinProps{
//   	ResourceDefinitionId: jsii.String("resourceDefinitionId"),
//   	Resources: []interface{}{
//   		&ResourceInstanceProperty{
//   			Id: jsii.String("id"),
//   			Name: jsii.String("name"),
//   			ResourceDataContainer: &ResourceDataContainerProperty{
//   				LocalDeviceResourceData: &LocalDeviceResourceDataProperty{
//   					GroupOwnerSetting: &GroupOwnerSettingProperty{
//   						AutoAddGroupOwner: jsii.Boolean(false),
//   						GroupOwner: jsii.String("groupOwner"),
//   					},
//   					SourcePath: jsii.String("sourcePath"),
//   				},
//   				LocalVolumeResourceData: &LocalVolumeResourceDataProperty{
//   					DestinationPath: jsii.String("destinationPath"),
//   					GroupOwnerSetting: &GroupOwnerSettingProperty{
//   						AutoAddGroupOwner: jsii.Boolean(false),
//   						GroupOwner: jsii.String("groupOwner"),
//   					},
//   					SourcePath: jsii.String("sourcePath"),
//   				},
//   				S3MachineLearningModelResourceData: &S3MachineLearningModelResourceDataProperty{
//   					DestinationPath: jsii.String("destinationPath"),
//   					OwnerSetting: &ResourceDownloadOwnerSettingProperty{
//   						GroupOwner: jsii.String("groupOwner"),
//   						GroupPermission: jsii.String("groupPermission"),
//   					},
//   					S3Uri: jsii.String("s3Uri"),
//   				},
//   				SageMakerMachineLearningModelResourceData: &SageMakerMachineLearningModelResourceDataProperty{
//   					DestinationPath: jsii.String("destinationPath"),
//   					OwnerSetting: &ResourceDownloadOwnerSettingProperty{
//   						GroupOwner: jsii.String("groupOwner"),
//   						GroupPermission: jsii.String("groupPermission"),
//   					},
//   					SageMakerJobArn: jsii.String("sageMakerJobArn"),
//   				},
//   				SecretsManagerSecretResourceData: &SecretsManagerSecretResourceDataProperty{
//   					AdditionalStagingLabelsToDownload: []*string{
//   						jsii.String("additionalStagingLabelsToDownload"),
//   					},
//   					Arn: jsii.String("arn"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-resourcedefinitionversion.html
//
type CfnResourceDefinitionVersionMixinProps struct {
	// The ID of the resource definition associated with this version.
	//
	// This value is a GUID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-resourcedefinitionversion.html#cfn-greengrass-resourcedefinitionversion-resourcedefinitionid
	//
	ResourceDefinitionId *string `field:"optional" json:"resourceDefinitionId" yaml:"resourceDefinitionId"`
	// The resources in this version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-resourcedefinitionversion.html#cfn-greengrass-resourcedefinitionversion-resources
	//
	Resources interface{} `field:"optional" json:"resources" yaml:"resources"`
}

