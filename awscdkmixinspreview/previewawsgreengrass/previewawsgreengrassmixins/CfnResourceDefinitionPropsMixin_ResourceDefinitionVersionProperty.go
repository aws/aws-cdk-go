package previewawsgreengrassmixins


// A resource definition version contains a list of resources. (In CloudFormation , resources are named *resource instances* .).
//
// > After you create a resource definition version that contains the resources you want to deploy, you must add it to your group version. For more information, see [`AWS::Greengrass::Group`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-group.html) .
//
// In an CloudFormation template, `ResourceDefinitionVersion` is the property type of the `InitialVersion` property in the [`AWS::Greengrass::ResourceDefinition`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-resourcedefinition.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   resourceDefinitionVersionProperty := &ResourceDefinitionVersionProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinition-resourcedefinitionversion.html
//
type CfnResourceDefinitionPropsMixin_ResourceDefinitionVersionProperty struct {
	// The resources in this version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinition-resourcedefinitionversion.html#cfn-greengrass-resourcedefinition-resourcedefinitionversion-resources
	//
	Resources interface{} `field:"optional" json:"resources" yaml:"resources"`
}

