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
//   		DefaultResourceSpec: &ResourceSpecProperty{
//   			InstanceType: jsii.String("instanceType"),
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
//   			SageMakerImageArn: jsii.String("sageMakerImageArn"),
//   			SageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
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
//   			SageMakerImageArn: jsii.String("sageMakerImageArn"),
//   			SageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-usersettings.html
//
type CfnUserProfile_UserSettingsProperty struct {
	// The CodeEditor app settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-usersettings.html#cfn-sagemaker-userprofile-usersettings-codeeditorappsettings
	//
	CodeEditorAppSettings interface{} `field:"optional" json:"codeEditorAppSettings" yaml:"codeEditorAppSettings"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-usersettings.html#cfn-sagemaker-userprofile-usersettings-customfilesystemconfigs
	//
	CustomFileSystemConfigs interface{} `field:"optional" json:"customFileSystemConfigs" yaml:"customFileSystemConfigs"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-usersettings.html#cfn-sagemaker-userprofile-usersettings-customposixuserconfig
	//
	CustomPosixUserConfig interface{} `field:"optional" json:"customPosixUserConfig" yaml:"customPosixUserConfig"`
	// Defines which Amazon SageMaker application users are directed to by default.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-usersettings.html#cfn-sagemaker-userprofile-usersettings-defaultlandinguri
	//
	DefaultLandingUri *string `field:"optional" json:"defaultLandingUri" yaml:"defaultLandingUri"`
	// The execution role for the user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-usersettings.html#cfn-sagemaker-userprofile-usersettings-executionrole
	//
	ExecutionRole *string `field:"optional" json:"executionRole" yaml:"executionRole"`
	// The JupyterLab app settings.
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-usersettings.html#cfn-sagemaker-userprofile-usersettings-securitygroups
	//
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Specifies options for sharing Amazon SageMaker Studio notebooks.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-usersettings.html#cfn-sagemaker-userprofile-usersettings-sharingsettings
	//
	SharingSettings interface{} `field:"optional" json:"sharingSettings" yaml:"sharingSettings"`
	// Default storage settings for a space.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-usersettings.html#cfn-sagemaker-userprofile-usersettings-spacestoragesettings
	//
	SpaceStorageSettings interface{} `field:"optional" json:"spaceStorageSettings" yaml:"spaceStorageSettings"`
	// Indicates whether the Studio experience is available to users.
	//
	// If not, users cannot access Studio.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-usersettings.html#cfn-sagemaker-userprofile-usersettings-studiowebportal
	//
	StudioWebPortal *string `field:"optional" json:"studioWebPortal" yaml:"studioWebPortal"`
}

