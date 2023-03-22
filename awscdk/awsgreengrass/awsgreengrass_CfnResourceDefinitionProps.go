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
//   cfnResourceDefinitionProps := &cfnResourceDefinitionProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	initialVersion: &resourceDefinitionVersionProperty{
//   		resources: []interface{}{
//   			&resourceInstanceProperty{
//   				id: jsii.String("id"),
//   				name: jsii.String("name"),
//   				resourceDataContainer: &resourceDataContainerProperty{
//   					localDeviceResourceData: &localDeviceResourceDataProperty{
//   						sourcePath: jsii.String("sourcePath"),
//
//   						// the properties below are optional
//   						groupOwnerSetting: &groupOwnerSettingProperty{
//   							autoAddGroupOwner: jsii.Boolean(false),
//
//   							// the properties below are optional
//   							groupOwner: jsii.String("groupOwner"),
//   						},
//   					},
//   					localVolumeResourceData: &localVolumeResourceDataProperty{
//   						destinationPath: jsii.String("destinationPath"),
//   						sourcePath: jsii.String("sourcePath"),
//
//   						// the properties below are optional
//   						groupOwnerSetting: &groupOwnerSettingProperty{
//   							autoAddGroupOwner: jsii.Boolean(false),
//
//   							// the properties below are optional
//   							groupOwner: jsii.String("groupOwner"),
//   						},
//   					},
//   					s3MachineLearningModelResourceData: &s3MachineLearningModelResourceDataProperty{
//   						destinationPath: jsii.String("destinationPath"),
//   						s3Uri: jsii.String("s3Uri"),
//
//   						// the properties below are optional
//   						ownerSetting: &resourceDownloadOwnerSettingProperty{
//   							groupOwner: jsii.String("groupOwner"),
//   							groupPermission: jsii.String("groupPermission"),
//   						},
//   					},
//   					sageMakerMachineLearningModelResourceData: &sageMakerMachineLearningModelResourceDataProperty{
//   						destinationPath: jsii.String("destinationPath"),
//   						sageMakerJobArn: jsii.String("sageMakerJobArn"),
//
//   						// the properties below are optional
//   						ownerSetting: &resourceDownloadOwnerSettingProperty{
//   							groupOwner: jsii.String("groupOwner"),
//   							groupPermission: jsii.String("groupPermission"),
//   						},
//   					},
//   					secretsManagerSecretResourceData: &secretsManagerSecretResourceDataProperty{
//   						arn: jsii.String("arn"),
//
//   						// the properties below are optional
//   						additionalStagingLabelsToDownload: []*string{
//   							jsii.String("additionalStagingLabelsToDownload"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	tags: tags,
//   }
//
type CfnResourceDefinitionProps struct {
	// The name of the resource definition.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The resource definition version to include when the resource definition is created.
	//
	// A resource definition version contains a list of [`resource instance`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinition-resourceinstance.html) property types.
	//
	// > To associate a resource definition version after the resource definition is created, create an [`AWS::Greengrass::ResourceDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-resourcedefinitionversion.html) resource and specify the ID of this resource definition.
	InitialVersion interface{} `field:"optional" json:"initialVersion" yaml:"initialVersion"`
	// Application-specific metadata to attach to the resource definition.
	//
	// You can use tags in IAM policies to control access to AWS IoT Greengrass resources. You can also use tags to categorize your resources. For more information, see [Tagging Your AWS IoT Greengrass Resources](https://docs.aws.amazon.com/greengrass/latest/developerguide/tagging.html) in the *Developer Guide* .
	//
	// This `Json` property type is processed as a map of key-value pairs. It uses the following format, which is different from most `Tags` implementations in AWS CloudFormation templates.
	//
	// ```json
	// "Tags": { "KeyName0": "value", "KeyName1": "value", "KeyName2": "value"
	// }
	// ```.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

