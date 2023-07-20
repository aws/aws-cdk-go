package awsgreengrass


// A resource definition version contains a list of resources. (In AWS CloudFormation , resources are named *resource instances* .).
//
// > After you create a resource definition version that contains the resources you want to deploy, you must add it to your group version. For more information, see [`AWS::Greengrass::Group`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-group.html) .
//
// In an AWS CloudFormation template, `ResourceDefinitionVersion` is the property type of the `InitialVersion` property in the [`AWS::Greengrass::ResourceDefinition`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-resourcedefinition.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceDefinitionVersionProperty := &ResourceDefinitionVersionProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinition-resourcedefinitionversion.html
//
type CfnResourceDefinition_ResourceDefinitionVersionProperty struct {
	// The resources in this version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinition-resourcedefinitionversion.html#cfn-greengrass-resourcedefinition-resourcedefinitionversion-resources
	//
	Resources interface{} `field:"required" json:"resources" yaml:"resources"`
}

