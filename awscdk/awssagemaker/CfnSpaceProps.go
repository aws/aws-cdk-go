package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnSpace`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSpaceProps := &CfnSpaceProps{
//   	DomainId: jsii.String("domainId"),
//   	SpaceName: jsii.String("spaceName"),
//
//   	// the properties below are optional
//   	OwnershipSettings: &OwnershipSettingsProperty{
//   		OwnerUserProfileName: jsii.String("ownerUserProfileName"),
//   	},
//   	SpaceDisplayName: jsii.String("spaceDisplayName"),
//   	SpaceSettings: &SpaceSettingsProperty{
//   		AppType: jsii.String("appType"),
//   		CodeEditorAppSettings: &SpaceCodeEditorAppSettingsProperty{
//   			DefaultResourceSpec: &ResourceSpecProperty{
//   				InstanceType: jsii.String("instanceType"),
//   				SageMakerImageArn: jsii.String("sageMakerImageArn"),
//   				SageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   			},
//   		},
//   		CustomFileSystems: []interface{}{
//   			&CustomFileSystemProperty{
//   				EfsFileSystem: &EFSFileSystemProperty{
//   					FileSystemId: jsii.String("fileSystemId"),
//   				},
//   			},
//   		},
//   		JupyterLabAppSettings: &SpaceJupyterLabAppSettingsProperty{
//   			CodeRepositories: []interface{}{
//   				&CodeRepositoryProperty{
//   					RepositoryUrl: jsii.String("repositoryUrl"),
//   				},
//   			},
//   			DefaultResourceSpec: &ResourceSpecProperty{
//   				InstanceType: jsii.String("instanceType"),
//   				SageMakerImageArn: jsii.String("sageMakerImageArn"),
//   				SageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   			},
//   		},
//   		JupyterServerAppSettings: &JupyterServerAppSettingsProperty{
//   			DefaultResourceSpec: &ResourceSpecProperty{
//   				InstanceType: jsii.String("instanceType"),
//   				SageMakerImageArn: jsii.String("sageMakerImageArn"),
//   				SageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   			},
//   		},
//   		KernelGatewayAppSettings: &KernelGatewayAppSettingsProperty{
//   			CustomImages: []interface{}{
//   				&CustomImageProperty{
//   					AppImageConfigName: jsii.String("appImageConfigName"),
//   					ImageName: jsii.String("imageName"),
//
//   					// the properties below are optional
//   					ImageVersionNumber: jsii.Number(123),
//   				},
//   			},
//   			DefaultResourceSpec: &ResourceSpecProperty{
//   				InstanceType: jsii.String("instanceType"),
//   				SageMakerImageArn: jsii.String("sageMakerImageArn"),
//   				SageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   			},
//   		},
//   		SpaceStorageSettings: &SpaceStorageSettingsProperty{
//   			EbsStorageSettings: &EbsStorageSettingsProperty{
//   				EbsVolumeSizeInGb: jsii.Number(123),
//   			},
//   		},
//   	},
//   	SpaceSharingSettings: &SpaceSharingSettingsProperty{
//   		SharingType: jsii.String("sharingType"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-space.html
//
type CfnSpaceProps struct {
	// The ID of the associated domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-space.html#cfn-sagemaker-space-domainid
	//
	DomainId *string `field:"required" json:"domainId" yaml:"domainId"`
	// The name of the space.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-space.html#cfn-sagemaker-space-spacename
	//
	SpaceName *string `field:"required" json:"spaceName" yaml:"spaceName"`
	// The collection of ownership settings for a space.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-space.html#cfn-sagemaker-space-ownershipsettings
	//
	OwnershipSettings interface{} `field:"optional" json:"ownershipSettings" yaml:"ownershipSettings"`
	// The name of the space that appears in the Studio UI.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-space.html#cfn-sagemaker-space-spacedisplayname
	//
	SpaceDisplayName *string `field:"optional" json:"spaceDisplayName" yaml:"spaceDisplayName"`
	// A collection of space settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-space.html#cfn-sagemaker-space-spacesettings
	//
	SpaceSettings interface{} `field:"optional" json:"spaceSettings" yaml:"spaceSettings"`
	// A collection of space sharing settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-space.html#cfn-sagemaker-space-spacesharingsettings
	//
	SpaceSharingSettings interface{} `field:"optional" json:"spaceSharingSettings" yaml:"spaceSharingSettings"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-space.html#cfn-sagemaker-space-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

