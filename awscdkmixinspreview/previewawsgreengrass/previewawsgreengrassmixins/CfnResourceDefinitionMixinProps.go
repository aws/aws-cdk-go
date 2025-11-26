package previewawsgreengrassmixins


// Properties for CfnResourceDefinitionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var tags interface{}
//
//   cfnResourceDefinitionMixinProps := &CfnResourceDefinitionMixinProps{
//   	InitialVersion: &ResourceDefinitionVersionProperty{
//   		Resources: []interface{}{
//   			&ResourceInstanceProperty{
//   				Id: jsii.String("id"),
//   				Name: jsii.String("name"),
//   				ResourceDataContainer: &ResourceDataContainerProperty{
//   					LocalDeviceResourceData: &LocalDeviceResourceDataProperty{
//   						GroupOwnerSetting: &GroupOwnerSettingProperty{
//   							AutoAddGroupOwner: jsii.Boolean(false),
//   							GroupOwner: jsii.String("groupOwner"),
//   						},
//   						SourcePath: jsii.String("sourcePath"),
//   					},
//   					LocalVolumeResourceData: &LocalVolumeResourceDataProperty{
//   						DestinationPath: jsii.String("destinationPath"),
//   						GroupOwnerSetting: &GroupOwnerSettingProperty{
//   							AutoAddGroupOwner: jsii.Boolean(false),
//   							GroupOwner: jsii.String("groupOwner"),
//   						},
//   						SourcePath: jsii.String("sourcePath"),
//   					},
//   					S3MachineLearningModelResourceData: &S3MachineLearningModelResourceDataProperty{
//   						DestinationPath: jsii.String("destinationPath"),
//   						OwnerSetting: &ResourceDownloadOwnerSettingProperty{
//   							GroupOwner: jsii.String("groupOwner"),
//   							GroupPermission: jsii.String("groupPermission"),
//   						},
//   						S3Uri: jsii.String("s3Uri"),
//   					},
//   					SageMakerMachineLearningModelResourceData: &SageMakerMachineLearningModelResourceDataProperty{
//   						DestinationPath: jsii.String("destinationPath"),
//   						OwnerSetting: &ResourceDownloadOwnerSettingProperty{
//   							GroupOwner: jsii.String("groupOwner"),
//   							GroupPermission: jsii.String("groupPermission"),
//   						},
//   						SageMakerJobArn: jsii.String("sageMakerJobArn"),
//   					},
//   					SecretsManagerSecretResourceData: &SecretsManagerSecretResourceDataProperty{
//   						AdditionalStagingLabelsToDownload: []*string{
//   							jsii.String("additionalStagingLabelsToDownload"),
//   						},
//   						Arn: jsii.String("arn"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Tags: tags,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-resourcedefinition.html
//
type CfnResourceDefinitionMixinProps struct {
	// The resource definition version to include when the resource definition is created.
	//
	// A resource definition version contains a list of [`resource instance`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinition-resourceinstance.html) property types.
	//
	// > To associate a resource definition version after the resource definition is created, create an [`AWS::Greengrass::ResourceDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-resourcedefinitionversion.html) resource and specify the ID of this resource definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-resourcedefinition.html#cfn-greengrass-resourcedefinition-initialversion
	//
	InitialVersion interface{} `field:"optional" json:"initialVersion" yaml:"initialVersion"`
	// The name of the resource definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-resourcedefinition.html#cfn-greengrass-resourcedefinition-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
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

