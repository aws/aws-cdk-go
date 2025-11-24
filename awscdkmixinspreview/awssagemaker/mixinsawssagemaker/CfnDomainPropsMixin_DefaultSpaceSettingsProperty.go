package mixinsawssagemaker


// The default settings for shared spaces that users create in the domain.
//
// SageMaker applies these settings only to shared spaces. It doesn't apply them to private spaces.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   defaultSpaceSettingsProperty := &DefaultSpaceSettingsProperty{
//   	CustomFileSystemConfigs: []interface{}{
//   		&CustomFileSystemConfigProperty{
//   			EfsFileSystemConfig: &EFSFileSystemConfigProperty{
//   				FileSystemId: jsii.String("fileSystemId"),
//   				FileSystemPath: jsii.String("fileSystemPath"),
//   			},
//   			FSxLustreFileSystemConfig: &FSxLustreFileSystemConfigProperty{
//   				FileSystemId: jsii.String("fileSystemId"),
//   				FileSystemPath: jsii.String("fileSystemPath"),
//   			},
//   			S3FileSystemConfig: &S3FileSystemConfigProperty{
//   				MountPath: jsii.String("mountPath"),
//   				S3Uri: jsii.String("s3Uri"),
//   			},
//   		},
//   	},
//   	CustomPosixUserConfig: &CustomPosixUserConfigProperty{
//   		Gid: jsii.Number(123),
//   		Uid: jsii.Number(123),
//   	},
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
//   		BuiltInLifecycleConfigArn: jsii.String("builtInLifecycleConfigArn"),
//   		CodeRepositories: []interface{}{
//   			&CodeRepositoryProperty{
//   				RepositoryUrl: jsii.String("repositoryUrl"),
//   			},
//   		},
//   		CustomImages: []interface{}{
//   			&CustomImageProperty{
//   				AppImageConfigName: jsii.String("appImageConfigName"),
//   				ImageName: jsii.String("imageName"),
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
//   	SecurityGroups: []*string{
//   		jsii.String("securityGroups"),
//   	},
//   	SpaceStorageSettings: &DefaultSpaceStorageSettingsProperty{
//   		DefaultEbsStorageSettings: &DefaultEbsStorageSettingsProperty{
//   			DefaultEbsVolumeSizeInGb: jsii.Number(123),
//   			MaximumEbsVolumeSizeInGb: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-defaultspacesettings.html
//
type CfnDomainPropsMixin_DefaultSpaceSettingsProperty struct {
	// The settings for assigning a custom file system to a domain.
	//
	// Permitted users can access this file system in Amazon SageMaker AI Studio.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-defaultspacesettings.html#cfn-sagemaker-domain-defaultspacesettings-customfilesystemconfigs
	//
	CustomFileSystemConfigs interface{} `field:"optional" json:"customFileSystemConfigs" yaml:"customFileSystemConfigs"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-defaultspacesettings.html#cfn-sagemaker-domain-defaultspacesettings-customposixuserconfig
	//
	CustomPosixUserConfig interface{} `field:"optional" json:"customPosixUserConfig" yaml:"customPosixUserConfig"`
	// The ARN of the execution role for the space.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-defaultspacesettings.html#cfn-sagemaker-domain-defaultspacesettings-executionrole
	//
	ExecutionRole *string `field:"optional" json:"executionRole" yaml:"executionRole"`
	// The JupyterLab app settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-defaultspacesettings.html#cfn-sagemaker-domain-defaultspacesettings-jupyterlabappsettings
	//
	JupyterLabAppSettings interface{} `field:"optional" json:"jupyterLabAppSettings" yaml:"jupyterLabAppSettings"`
	// The JupyterServer app settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-defaultspacesettings.html#cfn-sagemaker-domain-defaultspacesettings-jupyterserverappsettings
	//
	JupyterServerAppSettings interface{} `field:"optional" json:"jupyterServerAppSettings" yaml:"jupyterServerAppSettings"`
	// The KernelGateway app settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-defaultspacesettings.html#cfn-sagemaker-domain-defaultspacesettings-kernelgatewayappsettings
	//
	KernelGatewayAppSettings interface{} `field:"optional" json:"kernelGatewayAppSettings" yaml:"kernelGatewayAppSettings"`
	// The security group IDs for the Amazon VPC that the space uses for communication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-defaultspacesettings.html#cfn-sagemaker-domain-defaultspacesettings-securitygroups
	//
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Default storage settings for a space.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-defaultspacesettings.html#cfn-sagemaker-domain-defaultspacesettings-spacestoragesettings
	//
	SpaceStorageSettings interface{} `field:"optional" json:"spaceStorageSettings" yaml:"spaceStorageSettings"`
}

