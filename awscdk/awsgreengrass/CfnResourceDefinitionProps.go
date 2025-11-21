package awsgreengrass


// Properties for defining a `CfnResourceDefinition`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tags interface{}
//
//   cfnResourceDefinitionProps := &CfnResourceDefinitionProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	InitialVersion: &ResourceDefinitionVersionProperty{
//   		Resources: []interface{}{
//   			&ResourceInstanceProperty{
//   				Id: jsii.String("id"),
//   				Name: jsii.String("name"),
//   				ResourceDataContainer: &ResourceDataContainerProperty{
//   					LocalDeviceResourceData: &LocalDeviceResourceDataProperty{
//   						SourcePath: jsii.String("sourcePath"),
//
//   						// the properties below are optional
//   						GroupOwnerSetting: &GroupOwnerSettingProperty{
//   							AutoAddGroupOwner: jsii.Boolean(false),
//
//   							// the properties below are optional
//   							GroupOwner: jsii.String("groupOwner"),
//   						},
//   					},
//   					LocalVolumeResourceData: &LocalVolumeResourceDataProperty{
//   						DestinationPath: jsii.String("destinationPath"),
//   						SourcePath: jsii.String("sourcePath"),
//
//   						// the properties below are optional
//   						GroupOwnerSetting: &GroupOwnerSettingProperty{
//   							AutoAddGroupOwner: jsii.Boolean(false),
//
//   							// the properties below are optional
//   							GroupOwner: jsii.String("groupOwner"),
//   						},
//   					},
//   					S3MachineLearningModelResourceData: &S3MachineLearningModelResourceDataProperty{
//   						DestinationPath: jsii.String("destinationPath"),
//   						S3Uri: jsii.String("s3Uri"),
//
//   						// the properties below are optional
//   						OwnerSetting: &ResourceDownloadOwnerSettingProperty{
//   							GroupOwner: jsii.String("groupOwner"),
//   							GroupPermission: jsii.String("groupPermission"),
//   						},
//   					},
//   					SageMakerMachineLearningModelResourceData: &SageMakerMachineLearningModelResourceDataProperty{
//   						DestinationPath: jsii.String("destinationPath"),
//   						SageMakerJobArn: jsii.String("sageMakerJobArn"),
//
//   						// the properties below are optional
//   						OwnerSetting: &ResourceDownloadOwnerSettingProperty{
//   							GroupOwner: jsii.String("groupOwner"),
//   							GroupPermission: jsii.String("groupPermission"),
//   						},
//   					},
//   					SecretsManagerSecretResourceData: &SecretsManagerSecretResourceDataProperty{
//   						Arn: jsii.String("arn"),
//
//   						// the properties below are optional
//   						AdditionalStagingLabelsToDownload: []*string{
//   							jsii.String("additionalStagingLabelsToDownload"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	Tags: tags,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-resourcedefinition.html
//
type CfnResourceDefinitionProps struct {
	// The name of the resource definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-resourcedefinition.html#cfn-greengrass-resourcedefinition-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The resource definition version to include when the resource definition is created.
	//
	// A resource definition version contains a list of [`resource instance`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinition-resourceinstance.html) property types.
	//
	// > To associate a resource definition version after the resource definition is created, create an [`AWS::Greengrass::ResourceDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-resourcedefinitionversion.html) resource and specify the ID of this resource definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-resourcedefinition.html#cfn-greengrass-resourcedefinition-initialversion
	//
	InitialVersion interface{} `field:"optional" json:"initialVersion" yaml:"initialVersion"`
	// Application-specific metadata to attach to the resource definition.
	//
	// You can use tags in IAM policies to control access to AWS IoT Greengrass resources. You can also use tags to categorize your resources. For more information, see [Tagging Your AWS IoT Greengrass Resources](https://docs.aws.amazon.com/greengrass/v1/developerguide/tagging.html) in the *Developer Guide* .
	//
	// This `Json` property type is processed as a map of key-value pairs. It uses the following format, which is different from most `Tags` implementations in CloudFormation templates.
	//
	// ```json
	// "Tags": { "KeyName0": "value", "KeyName1": "value", "KeyName2": "value"
	// }
	// ```.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-resourcedefinition.html#cfn-greengrass-resourcedefinition-tags
	//
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

