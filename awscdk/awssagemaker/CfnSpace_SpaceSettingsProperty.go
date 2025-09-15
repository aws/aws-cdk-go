package awssagemaker


// A collection of space settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   spaceSettingsProperty := &SpaceSettingsProperty{
//   	AppType: jsii.String("appType"),
//   	CodeEditorAppSettings: &SpaceCodeEditorAppSettingsProperty{
//   		AppLifecycleManagement: &SpaceAppLifecycleManagementProperty{
//   			IdleSettings: &SpaceIdleSettingsProperty{
//   				IdleTimeoutInMinutes: jsii.Number(123),
//   			},
//   		},
//   		DefaultResourceSpec: &ResourceSpecProperty{
//   			InstanceType: jsii.String("instanceType"),
//   			LifecycleConfigArn: jsii.String("lifecycleConfigArn"),
//   			SageMakerImageArn: jsii.String("sageMakerImageArn"),
//   			SageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   		},
//   	},
//   	CustomFileSystems: []interface{}{
//   		&CustomFileSystemProperty{
//   			EfsFileSystem: &EFSFileSystemProperty{
//   				FileSystemId: jsii.String("fileSystemId"),
//   			},
//   			FSxLustreFileSystem: &FSxLustreFileSystemProperty{
//   				FileSystemId: jsii.String("fileSystemId"),
//   			},
//   			S3FileSystem: &S3FileSystemProperty{
//   				S3Uri: jsii.String("s3Uri"),
//   			},
//   		},
//   	},
//   	JupyterLabAppSettings: &SpaceJupyterLabAppSettingsProperty{
//   		AppLifecycleManagement: &SpaceAppLifecycleManagementProperty{
//   			IdleSettings: &SpaceIdleSettingsProperty{
//   				IdleTimeoutInMinutes: jsii.Number(123),
//   			},
//   		},
//   		CodeRepositories: []interface{}{
//   			&CodeRepositoryProperty{
//   				RepositoryUrl: jsii.String("repositoryUrl"),
//   			},
//   		},
//   		DefaultResourceSpec: &ResourceSpecProperty{
//   			InstanceType: jsii.String("instanceType"),
//   			LifecycleConfigArn: jsii.String("lifecycleConfigArn"),
//   			SageMakerImageArn: jsii.String("sageMakerImageArn"),
//   			SageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
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
//   	RemoteAccess: jsii.String("remoteAccess"),
//   	SpaceManagedResources: jsii.String("spaceManagedResources"),
//   	SpaceStorageSettings: &SpaceStorageSettingsProperty{
//   		EbsStorageSettings: &EbsStorageSettingsProperty{
//   			EbsVolumeSizeInGb: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-space-spacesettings.html
//
type CfnSpace_SpaceSettingsProperty struct {
	// The type of app created within the space.
	//
	// If using the [UpdateSpace](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_UpdateSpace.html) API, you can't change the app type of your space by specifying a different value for this field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-space-spacesettings.html#cfn-sagemaker-space-spacesettings-apptype
	//
	AppType *string `field:"optional" json:"appType" yaml:"appType"`
	// The Code Editor application settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-space-spacesettings.html#cfn-sagemaker-space-spacesettings-codeeditorappsettings
	//
	CodeEditorAppSettings interface{} `field:"optional" json:"codeEditorAppSettings" yaml:"codeEditorAppSettings"`
	// A file system, created by you, that you assign to a space for an Amazon SageMaker AI Domain.
	//
	// Permitted users can access this file system in Amazon SageMaker AI Studio.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-space-spacesettings.html#cfn-sagemaker-space-spacesettings-customfilesystems
	//
	CustomFileSystems interface{} `field:"optional" json:"customFileSystems" yaml:"customFileSystems"`
	// The settings for the JupyterLab application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-space-spacesettings.html#cfn-sagemaker-space-spacesettings-jupyterlabappsettings
	//
	JupyterLabAppSettings interface{} `field:"optional" json:"jupyterLabAppSettings" yaml:"jupyterLabAppSettings"`
	// The JupyterServer app settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-space-spacesettings.html#cfn-sagemaker-space-spacesettings-jupyterserverappsettings
	//
	JupyterServerAppSettings interface{} `field:"optional" json:"jupyterServerAppSettings" yaml:"jupyterServerAppSettings"`
	// The KernelGateway app settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-space-spacesettings.html#cfn-sagemaker-space-spacesettings-kernelgatewayappsettings
	//
	KernelGatewayAppSettings interface{} `field:"optional" json:"kernelGatewayAppSettings" yaml:"kernelGatewayAppSettings"`
	// A setting that enables or disables remote access for a SageMaker space.
	//
	// When enabled, this allows you to connect to the remote space from your local IDE.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-space-spacesettings.html#cfn-sagemaker-space-spacesettings-remoteaccess
	//
	RemoteAccess *string `field:"optional" json:"remoteAccess" yaml:"remoteAccess"`
	// If you enable this option, SageMaker AI creates the following resources on your behalf when you create the space:.
	//
	// - The user profile that possesses the space.
	// - The app that the space contains.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-space-spacesettings.html#cfn-sagemaker-space-spacesettings-spacemanagedresources
	//
	SpaceManagedResources *string `field:"optional" json:"spaceManagedResources" yaml:"spaceManagedResources"`
	// The storage settings for a space.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-space-spacesettings.html#cfn-sagemaker-space-spacesettings-spacestoragesettings
	//
	SpaceStorageSettings interface{} `field:"optional" json:"spaceStorageSettings" yaml:"spaceStorageSettings"`
}

