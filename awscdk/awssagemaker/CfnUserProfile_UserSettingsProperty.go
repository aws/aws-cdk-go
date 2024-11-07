package awssagemaker


// A collection of settings that apply to users of Amazon SageMaker Studio.
//
// These settings are specified when the [CreateUserProfile](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateUserProfile.html) API is called, and as `DefaultUserSettings` when the [CreateDomain](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateDomain.html) API is called.
//
// `SecurityGroups` is aggregated when specified in both calls. For all other settings in `UserSettings` , the values specified in `CreateUserProfile` take precedence over those specified in `CreateDomain` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   userSettingsProperty := &UserSettingsProperty{
//   	CodeEditorAppSettings: &CodeEditorAppSettingsProperty{
//   		AppLifecycleManagement: &AppLifecycleManagementProperty{
//   			IdleSettings: &IdleSettingsProperty{
//   				IdleTimeoutInMinutes: jsii.Number(123),
//   				LifecycleManagement: jsii.String("lifecycleManagement"),
//   				MaxIdleTimeoutInMinutes: jsii.Number(123),
//   				MinIdleTimeoutInMinutes: jsii.Number(123),
//   			},
//   		},
//   		CustomImages: []interface{}{
//   			&CustomImageProperty{
//   				AppImageConfigName: jsii.String("appImageConfigName"),
//   				ImageName: jsii.String("imageName"),
//
//   				// the properties below are optional
//   				ImageVersionNumber: jsii.Number(123),
//   			},
//   		},
//   		DefaultResourceSpec: &ResourceSpecProperty{
//   			InstanceType: jsii.String("instanceType"),
//   			LifecycleConfigArn: jsii.String("lifecycleConfigArn"),
//   			SageMakerImageArn: jsii.String("sageMakerImageArn"),
//   			SageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   		},
//   		LifecycleConfigArns: []*string{
//   			jsii.String("lifecycleConfigArns"),
//   		},
//   	},
//   	CustomFileSystemConfigs: []interface{}{
//   		&CustomFileSystemConfigProperty{
//   			EfsFileSystemConfig: &EFSFileSystemConfigProperty{
//   				FileSystemId: jsii.String("fileSystemId"),
//
//   				// the properties below are optional
//   				FileSystemPath: jsii.String("fileSystemPath"),
//   			},
//   		},
//   	},
//   	CustomPosixUserConfig: &CustomPosixUserConfigProperty{
//   		Gid: jsii.Number(123),
//   		Uid: jsii.Number(123),
//   	},
//   	DefaultLandingUri: jsii.String("defaultLandingUri"),
//   	ExecutionRole: jsii.String("executionRole"),
//   	JupyterLabAppSettings: &JupyterLabAppSettingsProperty{
//   		AppLifecycleManagement: &AppLifecycleManagementProperty{
//   			IdleSettings: &IdleSettingsProperty{
//   				IdleTimeoutInMinutes: jsii.Number(123),
//   				LifecycleManagement: jsii.String("lifecycleManagement"),
//   				MaxIdleTimeoutInMinutes: jsii.Number(123),
//   				MinIdleTimeoutInMinutes: jsii.Number(123),
//   			},
//   		},
//   		CodeRepositories: []interface{}{
//   			&CodeRepositoryProperty{
//   				RepositoryUrl: jsii.String("repositoryUrl"),
//   			},
//   		},
//   		CustomImages: []interface{}{
//   			&CustomImageProperty{
//   				AppImageConfigName: jsii.String("appImageConfigName"),
//   				ImageName: jsii.String("imageName"),
//
//   				// the properties below are optional
//   				ImageVersionNumber: jsii.Number(123),
//   			},
//   		},
//   		DefaultResourceSpec: &ResourceSpecProperty{
//   			InstanceType: jsii.String("instanceType"),
//   			LifecycleConfigArn: jsii.String("lifecycleConfigArn"),
//   			SageMakerImageArn: jsii.String("sageMakerImageArn"),
//   			SageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   		},
//   		LifecycleConfigArns: []*string{
//   			jsii.String("lifecycleConfigArns"),
//   		},
//   	},
//   	JupyterServerAppSettings: &JupyterServerAppSettingsProperty{
//   		DefaultResourceSpec: &ResourceSpecProperty{
//   			InstanceType: jsii.String("instanceType"),
//   			LifecycleConfigArn: jsii.String("lifecycleConfigArn"),
//   			SageMakerImageArn: jsii.String("sageMakerImageArn"),
//   			SageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   		},
//   		LifecycleConfigArns: []*string{
//   			jsii.String("lifecycleConfigArns"),
//   		},
//   	},
//   	KernelGatewayAppSettings: &KernelGatewayAppSettingsProperty{
//   		CustomImages: []interface{}{
//   			&CustomImageProperty{
//   				AppImageConfigName: jsii.String("appImageConfigName"),
//   				ImageName: jsii.String("imageName"),
//
//   				// the properties below are optional
//   				ImageVersionNumber: jsii.Number(123),
//   			},
//   		},
//   		DefaultResourceSpec: &ResourceSpecProperty{
//   			InstanceType: jsii.String("instanceType"),
//   			LifecycleConfigArn: jsii.String("lifecycleConfigArn"),
//   			SageMakerImageArn: jsii.String("sageMakerImageArn"),
//   			SageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   		},
//   		LifecycleConfigArns: []*string{
//   			jsii.String("lifecycleConfigArns"),
//   		},
//   	},
//   	RStudioServerProAppSettings: &RStudioServerProAppSettingsProperty{
//   		AccessStatus: jsii.String("accessStatus"),
//   		UserGroup: jsii.String("userGroup"),
//   	},
//   	SecurityGroups: []*string{
//   		jsii.String("securityGroups"),
//   	},
//   	SharingSettings: &SharingSettingsProperty{
//   		NotebookOutputOption: jsii.String("notebookOutputOption"),
//   		S3KmsKeyId: jsii.String("s3KmsKeyId"),
//   		S3OutputPath: jsii.String("s3OutputPath"),
//   	},
//   	SpaceStorageSettings: &DefaultSpaceStorageSettingsProperty{
//   		DefaultEbsStorageSettings: &DefaultEbsStorageSettingsProperty{
//   			DefaultEbsVolumeSizeInGb: jsii.Number(123),
//   			MaximumEbsVolumeSizeInGb: jsii.Number(123),
//   		},
//   	},
//   	StudioWebPortal: jsii.String("studioWebPortal"),
//   	StudioWebPortalSettings: &StudioWebPortalSettingsProperty{
//   		HiddenAppTypes: []*string{
//   			jsii.String("hiddenAppTypes"),
//   		},
//   		HiddenMlTools: []*string{
//   			jsii.String("hiddenMlTools"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-usersettings.html
//
type CfnUserProfile_UserSettingsProperty struct {
	// The Code Editor application settings.
	//
	// SageMaker applies these settings only to private spaces that the user creates in the domain. SageMaker doesn't apply these settings to shared spaces.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-usersettings.html#cfn-sagemaker-userprofile-usersettings-codeeditorappsettings
	//
	CodeEditorAppSettings interface{} `field:"optional" json:"codeEditorAppSettings" yaml:"codeEditorAppSettings"`
	// The settings for assigning a custom file system to a user profile.
	//
	// Permitted users can access this file system in Amazon SageMaker Studio.
	//
	// SageMaker applies these settings only to private spaces that the user creates in the domain. SageMaker doesn't apply these settings to shared spaces.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-usersettings.html#cfn-sagemaker-userprofile-usersettings-customfilesystemconfigs
	//
	CustomFileSystemConfigs interface{} `field:"optional" json:"customFileSystemConfigs" yaml:"customFileSystemConfigs"`
	// Details about the POSIX identity that is used for file system operations.
	//
	// SageMaker applies these settings only to private spaces that the user creates in the domain. SageMaker doesn't apply these settings to shared spaces.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-usersettings.html#cfn-sagemaker-userprofile-usersettings-customposixuserconfig
	//
	CustomPosixUserConfig interface{} `field:"optional" json:"customPosixUserConfig" yaml:"customPosixUserConfig"`
	// The default experience that the user is directed to when accessing the domain. The supported values are:.
	//
	// - `studio::` : Indicates that Studio is the default experience. This value can only be passed if `StudioWebPortal` is set to `ENABLED` .
	// - `app:JupyterServer:` : Indicates that Studio Classic is the default experience.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-usersettings.html#cfn-sagemaker-userprofile-usersettings-defaultlandinguri
	//
	DefaultLandingUri *string `field:"optional" json:"defaultLandingUri" yaml:"defaultLandingUri"`
	// The execution role for the user.
	//
	// SageMaker applies this setting only to private spaces that the user creates in the domain. SageMaker doesn't apply this setting to shared spaces.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-usersettings.html#cfn-sagemaker-userprofile-usersettings-executionrole
	//
	ExecutionRole *string `field:"optional" json:"executionRole" yaml:"executionRole"`
	// The settings for the JupyterLab application.
	//
	// SageMaker applies these settings only to private spaces that the user creates in the domain. SageMaker doesn't apply these settings to shared spaces.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-usersettings.html#cfn-sagemaker-userprofile-usersettings-jupyterlabappsettings
	//
	JupyterLabAppSettings interface{} `field:"optional" json:"jupyterLabAppSettings" yaml:"jupyterLabAppSettings"`
	// The Jupyter server's app settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-usersettings.html#cfn-sagemaker-userprofile-usersettings-jupyterserverappsettings
	//
	JupyterServerAppSettings interface{} `field:"optional" json:"jupyterServerAppSettings" yaml:"jupyterServerAppSettings"`
	// The kernel gateway app settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-usersettings.html#cfn-sagemaker-userprofile-usersettings-kernelgatewayappsettings
	//
	KernelGatewayAppSettings interface{} `field:"optional" json:"kernelGatewayAppSettings" yaml:"kernelGatewayAppSettings"`
	// A collection of settings that configure user interaction with the `RStudioServerPro` app.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-usersettings.html#cfn-sagemaker-userprofile-usersettings-rstudioserverproappsettings
	//
	RStudioServerProAppSettings interface{} `field:"optional" json:"rStudioServerProAppSettings" yaml:"rStudioServerProAppSettings"`
	// The security groups for the Amazon Virtual Private Cloud (VPC) that the domain uses for communication.
	//
	// Optional when the `CreateDomain.AppNetworkAccessType` parameter is set to `PublicInternetOnly` .
	//
	// Required when the `CreateDomain.AppNetworkAccessType` parameter is set to `VpcOnly` , unless specified as part of the `DefaultUserSettings` for the domain.
	//
	// Amazon SageMaker adds a security group to allow NFS traffic from Amazon SageMaker Studio. Therefore, the number of security groups that you can specify is one less than the maximum number shown.
	//
	// SageMaker applies these settings only to private spaces that the user creates in the domain. SageMaker doesn't apply these settings to shared spaces.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-usersettings.html#cfn-sagemaker-userprofile-usersettings-securitygroups
	//
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Specifies options for sharing Amazon SageMaker Studio notebooks.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-usersettings.html#cfn-sagemaker-userprofile-usersettings-sharingsettings
	//
	SharingSettings interface{} `field:"optional" json:"sharingSettings" yaml:"sharingSettings"`
	// The storage settings for a space.
	//
	// SageMaker applies these settings only to private spaces that the user creates in the domain. SageMaker doesn't apply these settings to shared spaces.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-usersettings.html#cfn-sagemaker-userprofile-usersettings-spacestoragesettings
	//
	SpaceStorageSettings interface{} `field:"optional" json:"spaceStorageSettings" yaml:"spaceStorageSettings"`
	// Whether the user can access Studio.
	//
	// If this value is set to `DISABLED` , the user cannot access Studio, even if that is the default experience for the domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-usersettings.html#cfn-sagemaker-userprofile-usersettings-studiowebportal
	//
	StudioWebPortal *string `field:"optional" json:"studioWebPortal" yaml:"studioWebPortal"`
	// Studio settings.
	//
	// If these settings are applied on a user level, they take priority over the settings applied on a domain level.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-usersettings.html#cfn-sagemaker-userprofile-usersettings-studiowebportalsettings
	//
	StudioWebPortalSettings interface{} `field:"optional" json:"studioWebPortalSettings" yaml:"studioWebPortalSettings"`
}

