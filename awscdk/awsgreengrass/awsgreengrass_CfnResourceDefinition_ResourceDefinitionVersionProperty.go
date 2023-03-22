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
//   resourceDefinitionVersionProperty := &resourceDefinitionVersionProperty{
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
type CfnResourceDefinition_ResourceDefinitionVersionProperty struct {
	// The resources in this version.
	Resources interface{} `field:"required" json:"resources" yaml:"resources"`
}

