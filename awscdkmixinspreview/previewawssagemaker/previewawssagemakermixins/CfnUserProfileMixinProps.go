package previewawssagemakermixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnUserProfilePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnUserProfileMixinProps := &CfnUserProfileMixinProps{
//   	DomainId: jsii.String("domainId"),
//   	SingleSignOnUserIdentifier: jsii.String("singleSignOnUserIdentifier"),
//   	SingleSignOnUserValue: jsii.String("singleSignOnUserValue"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	UserProfileName: jsii.String("userProfileName"),
//   	UserSettings: &UserSettingsProperty{
//   		AutoMountHomeEfs: jsii.String("autoMountHomeEfs"),
//   		CodeEditorAppSettings: &CodeEditorAppSettingsProperty{
//   			AppLifecycleManagement: &AppLifecycleManagementProperty{
//   				IdleSettings: &IdleSettingsProperty{
//   					IdleTimeoutInMinutes: jsii.Number(123),
//   					LifecycleManagement: jsii.String("lifecycleManagement"),
//   					MaxIdleTimeoutInMinutes: jsii.Number(123),
//   					MinIdleTimeoutInMinutes: jsii.Number(123),
//   				},
//   			},
//   			BuiltInLifecycleConfigArn: jsii.String("builtInLifecycleConfigArn"),
//   			CustomImages: []interface{}{
//   				&CustomImageProperty{
//   					AppImageConfigName: jsii.String("appImageConfigName"),
//   					ImageName: jsii.String("imageName"),
//   					ImageVersionNumber: jsii.Number(123),
//   				},
//   			},
//   			DefaultResourceSpec: &ResourceSpecProperty{
//   				InstanceType: jsii.String("instanceType"),
//   				LifecycleConfigArn: jsii.String("lifecycleConfigArn"),
//   				SageMakerImageArn: jsii.String("sageMakerImageArn"),
//   				SageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   			},
//   			LifecycleConfigArns: []*string{
//   				jsii.String("lifecycleConfigArns"),
//   			},
//   		},
//   		CustomFileSystemConfigs: []interface{}{
//   			&CustomFileSystemConfigProperty{
//   				EfsFileSystemConfig: &EFSFileSystemConfigProperty{
//   					FileSystemId: jsii.String("fileSystemId"),
//   					FileSystemPath: jsii.String("fileSystemPath"),
//   				},
//   				FSxLustreFileSystemConfig: &FSxLustreFileSystemConfigProperty{
//   					FileSystemId: jsii.String("fileSystemId"),
//   					FileSystemPath: jsii.String("fileSystemPath"),
//   				},
//   				S3FileSystemConfig: &S3FileSystemConfigProperty{
//   					MountPath: jsii.String("mountPath"),
//   					S3Uri: jsii.String("s3Uri"),
//   				},
//   			},
//   		},
//   		CustomPosixUserConfig: &CustomPosixUserConfigProperty{
//   			Gid: jsii.Number(123),
//   			Uid: jsii.Number(123),
//   		},
//   		DefaultLandingUri: jsii.String("defaultLandingUri"),
//   		ExecutionRole: jsii.String("executionRole"),
//   		JupyterLabAppSettings: &JupyterLabAppSettingsProperty{
//   			AppLifecycleManagement: &AppLifecycleManagementProperty{
//   				IdleSettings: &IdleSettingsProperty{
//   					IdleTimeoutInMinutes: jsii.Number(123),
//   					LifecycleManagement: jsii.String("lifecycleManagement"),
//   					MaxIdleTimeoutInMinutes: jsii.Number(123),
//   					MinIdleTimeoutInMinutes: jsii.Number(123),
//   				},
//   			},
//   			BuiltInLifecycleConfigArn: jsii.String("builtInLifecycleConfigArn"),
//   			CodeRepositories: []interface{}{
//   				&CodeRepositoryProperty{
//   					RepositoryUrl: jsii.String("repositoryUrl"),
//   				},
//   			},
//   			CustomImages: []interface{}{
//   				&CustomImageProperty{
//   					AppImageConfigName: jsii.String("appImageConfigName"),
//   					ImageName: jsii.String("imageName"),
//   					ImageVersionNumber: jsii.Number(123),
//   				},
//   			},
//   			DefaultResourceSpec: &ResourceSpecProperty{
//   				InstanceType: jsii.String("instanceType"),
//   				LifecycleConfigArn: jsii.String("lifecycleConfigArn"),
//   				SageMakerImageArn: jsii.String("sageMakerImageArn"),
//   				SageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   			},
//   			LifecycleConfigArns: []*string{
//   				jsii.String("lifecycleConfigArns"),
//   			},
//   		},
//   		JupyterServerAppSettings: &JupyterServerAppSettingsProperty{
//   			DefaultResourceSpec: &ResourceSpecProperty{
//   				InstanceType: jsii.String("instanceType"),
//   				LifecycleConfigArn: jsii.String("lifecycleConfigArn"),
//   				SageMakerImageArn: jsii.String("sageMakerImageArn"),
//   				SageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   			},
//   			LifecycleConfigArns: []*string{
//   				jsii.String("lifecycleConfigArns"),
//   			},
//   		},
//   		KernelGatewayAppSettings: &KernelGatewayAppSettingsProperty{
//   			CustomImages: []interface{}{
//   				&CustomImageProperty{
//   					AppImageConfigName: jsii.String("appImageConfigName"),
//   					ImageName: jsii.String("imageName"),
//   					ImageVersionNumber: jsii.Number(123),
//   				},
//   			},
//   			DefaultResourceSpec: &ResourceSpecProperty{
//   				InstanceType: jsii.String("instanceType"),
//   				LifecycleConfigArn: jsii.String("lifecycleConfigArn"),
//   				SageMakerImageArn: jsii.String("sageMakerImageArn"),
//   				SageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   			},
//   			LifecycleConfigArns: []*string{
//   				jsii.String("lifecycleConfigArns"),
//   			},
//   		},
//   		RStudioServerProAppSettings: &RStudioServerProAppSettingsProperty{
//   			AccessStatus: jsii.String("accessStatus"),
//   			UserGroup: jsii.String("userGroup"),
//   		},
//   		SecurityGroups: []*string{
//   			jsii.String("securityGroups"),
//   		},
//   		SharingSettings: &SharingSettingsProperty{
//   			NotebookOutputOption: jsii.String("notebookOutputOption"),
//   			S3KmsKeyId: jsii.String("s3KmsKeyId"),
//   			S3OutputPath: jsii.String("s3OutputPath"),
//   		},
//   		SpaceStorageSettings: &DefaultSpaceStorageSettingsProperty{
//   			DefaultEbsStorageSettings: &DefaultEbsStorageSettingsProperty{
//   				DefaultEbsVolumeSizeInGb: jsii.Number(123),
//   				MaximumEbsVolumeSizeInGb: jsii.Number(123),
//   			},
//   		},
//   		StudioWebPortal: jsii.String("studioWebPortal"),
//   		StudioWebPortalSettings: &StudioWebPortalSettingsProperty{
//   			HiddenAppTypes: []*string{
//   				jsii.String("hiddenAppTypes"),
//   			},
//   			HiddenInstanceTypes: []*string{
//   				jsii.String("hiddenInstanceTypes"),
//   			},
//   			HiddenMlTools: []*string{
//   				jsii.String("hiddenMlTools"),
//   			},
//   			HiddenSageMakerImageVersionAliases: []interface{}{
//   				&HiddenSageMakerImageProperty{
//   					SageMakerImageName: jsii.String("sageMakerImageName"),
//   					VersionAliases: []*string{
//   						jsii.String("versionAliases"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-userprofile.html
//
type CfnUserProfileMixinProps struct {
	// The domain ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-userprofile.html#cfn-sagemaker-userprofile-domainid
	//
	DomainId *string `field:"optional" json:"domainId" yaml:"domainId"`
	// A specifier for the type of value specified in SingleSignOnUserValue.
	//
	// Currently, the only supported value is "UserName". If the Domain's AuthMode is SSO , this field is required. If the Domain's AuthMode is not SSO , this field cannot be specified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-userprofile.html#cfn-sagemaker-userprofile-singlesignonuseridentifier
	//
	SingleSignOnUserIdentifier *string `field:"optional" json:"singleSignOnUserIdentifier" yaml:"singleSignOnUserIdentifier"`
	// The username of the associated AWS Single Sign-On User for this UserProfile.
	//
	// If the Domain's AuthMode is SSO , this field is required, and must match a valid username of a user in your directory. If the Domain's AuthMode is not SSO , this field cannot be specified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-userprofile.html#cfn-sagemaker-userprofile-singlesignonuservalue
	//
	SingleSignOnUserValue *string `field:"optional" json:"singleSignOnUserValue" yaml:"singleSignOnUserValue"`
	// An array of key-value pairs to apply to this resource.
	//
	// Tags that you specify for the User Profile are also added to all apps that the User Profile launches.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-userprofile.html#cfn-sagemaker-userprofile-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The user profile name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-userprofile.html#cfn-sagemaker-userprofile-userprofilename
	//
	UserProfileName *string `field:"optional" json:"userProfileName" yaml:"userProfileName"`
	// A collection of settings that apply to users of Amazon SageMaker Studio.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-userprofile.html#cfn-sagemaker-userprofile-usersettings
	//
	UserSettings interface{} `field:"optional" json:"userSettings" yaml:"userSettings"`
}

