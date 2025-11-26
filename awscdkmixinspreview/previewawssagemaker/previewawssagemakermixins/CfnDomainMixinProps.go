package previewawssagemakermixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnDomainPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDomainMixinProps := &CfnDomainMixinProps{
//   	AppNetworkAccessType: jsii.String("appNetworkAccessType"),
//   	AppSecurityGroupManagement: jsii.String("appSecurityGroupManagement"),
//   	AuthMode: jsii.String("authMode"),
//   	DefaultSpaceSettings: &DefaultSpaceSettingsProperty{
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
//   		SecurityGroups: []*string{
//   			jsii.String("securityGroups"),
//   		},
//   		SpaceStorageSettings: &DefaultSpaceStorageSettingsProperty{
//   			DefaultEbsStorageSettings: &DefaultEbsStorageSettingsProperty{
//   				DefaultEbsVolumeSizeInGb: jsii.Number(123),
//   				MaximumEbsVolumeSizeInGb: jsii.Number(123),
//   			},
//   		},
//   	},
//   	DefaultUserSettings: &UserSettingsProperty{
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
//   		RSessionAppSettings: &RSessionAppSettingsProperty{
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
//   	DomainName: jsii.String("domainName"),
//   	DomainSettings: &DomainSettingsProperty{
//   		DockerSettings: &DockerSettingsProperty{
//   			EnableDockerAccess: jsii.String("enableDockerAccess"),
//   			VpcOnlyTrustedAccounts: []*string{
//   				jsii.String("vpcOnlyTrustedAccounts"),
//   			},
//   		},
//   		ExecutionRoleIdentityConfig: jsii.String("executionRoleIdentityConfig"),
//   		IpAddressType: jsii.String("ipAddressType"),
//   		RStudioServerProDomainSettings: &RStudioServerProDomainSettingsProperty{
//   			DefaultResourceSpec: &ResourceSpecProperty{
//   				InstanceType: jsii.String("instanceType"),
//   				LifecycleConfigArn: jsii.String("lifecycleConfigArn"),
//   				SageMakerImageArn: jsii.String("sageMakerImageArn"),
//   				SageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   			},
//   			DomainExecutionRoleArn: jsii.String("domainExecutionRoleArn"),
//   			RStudioConnectUrl: jsii.String("rStudioConnectUrl"),
//   			RStudioPackageManagerUrl: jsii.String("rStudioPackageManagerUrl"),
//   		},
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		UnifiedStudioSettings: &UnifiedStudioSettingsProperty{
//   			DomainAccountId: jsii.String("domainAccountId"),
//   			DomainId: jsii.String("domainId"),
//   			DomainRegion: jsii.String("domainRegion"),
//   			EnvironmentId: jsii.String("environmentId"),
//   			ProjectId: jsii.String("projectId"),
//   			ProjectS3Path: jsii.String("projectS3Path"),
//   			StudioWebPortalAccess: jsii.String("studioWebPortalAccess"),
//   		},
//   	},
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	SubnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   	TagPropagation: jsii.String("tagPropagation"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VpcId: jsii.String("vpcId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-domain.html
//
type CfnDomainMixinProps struct {
	// Specifies the VPC used for non-EFS traffic. The default value is `PublicInternetOnly` .
	//
	// - `PublicInternetOnly` - Non-EFS traffic is through a VPC managed by Amazon SageMaker AI , which allows direct internet access
	// - `VpcOnly` - All Studio traffic is through the specified VPC and subnets
	//
	// *Valid Values* : `PublicInternetOnly | VpcOnly`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-domain.html#cfn-sagemaker-domain-appnetworkaccesstype
	//
	AppNetworkAccessType *string `field:"optional" json:"appNetworkAccessType" yaml:"appNetworkAccessType"`
	// The entity that creates and manages the required security groups for inter-app communication in `VpcOnly` mode.
	//
	// Required when `CreateDomain.AppNetworkAccessType` is `VpcOnly` and `DomainSettings.RStudioServerProDomainSettings.DomainExecutionRoleArn` is provided. If setting up the domain for use with RStudio, this value must be set to `Service` .
	//
	// *Allowed Values* : `Service` | `Customer`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-domain.html#cfn-sagemaker-domain-appsecuritygroupmanagement
	//
	AppSecurityGroupManagement *string `field:"optional" json:"appSecurityGroupManagement" yaml:"appSecurityGroupManagement"`
	// The mode of authentication that members use to access the Domain.
	//
	// *Valid Values* : `SSO | IAM`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-domain.html#cfn-sagemaker-domain-authmode
	//
	AuthMode *string `field:"optional" json:"authMode" yaml:"authMode"`
	// The default settings for shared spaces that users create in the domain.
	//
	// SageMaker applies these settings only to shared spaces. It doesn't apply them to private spaces.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-domain.html#cfn-sagemaker-domain-defaultspacesettings
	//
	DefaultSpaceSettings interface{} `field:"optional" json:"defaultSpaceSettings" yaml:"defaultSpaceSettings"`
	// The default user settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-domain.html#cfn-sagemaker-domain-defaultusersettings
	//
	DefaultUserSettings interface{} `field:"optional" json:"defaultUserSettings" yaml:"defaultUserSettings"`
	// The domain name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-domain.html#cfn-sagemaker-domain-domainname
	//
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// A collection of settings that apply to the `SageMaker Domain` .
	//
	// These settings are specified through the `CreateDomain` API call.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-domain.html#cfn-sagemaker-domain-domainsettings
	//
	DomainSettings interface{} `field:"optional" json:"domainSettings" yaml:"domainSettings"`
	// SageMaker uses AWS KMS to encrypt the EFS volume attached to the Domain with an AWS managed customer master key (CMK) by default.
	//
	// For more control, specify a customer managed CMK.
	//
	// *Length Constraints* : Maximum length of 2048.
	//
	// *Pattern* : `.*`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-domain.html#cfn-sagemaker-domain-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The VPC subnets that Studio uses for communication.
	//
	// *Length Constraints* : Maximum length of 32.
	//
	// *Array members* : Minimum number of 1 item. Maximum number of 16 items.
	//
	// *Pattern* : `[-0-9a-zA-Z]+`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-domain.html#cfn-sagemaker-domain-subnetids
	//
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
	// Indicates whether the tags added to Domain, User Profile and Space entity is propagated to all SageMaker resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-domain.html#cfn-sagemaker-domain-tagpropagation
	//
	TagPropagation *string `field:"optional" json:"tagPropagation" yaml:"tagPropagation"`
	// Tags to associated with the Domain.
	//
	// Each tag consists of a key and an optional value. Tag keys must be unique per resource. Tags are searchable using the Search API.
	//
	// Tags that you specify for the Domain are also added to all apps that are launched in the Domain.
	//
	// *Array members* : Minimum number of 0 items. Maximum number of 50 items.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-domain.html#cfn-sagemaker-domain-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The ID of the Amazon Virtual Private Cloud (Amazon VPC) that Studio uses for communication.
	//
	// *Length Constraints* : Maximum length of 32.
	//
	// *Pattern* : `[-0-9a-zA-Z]+`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-domain.html#cfn-sagemaker-domain-vpcid
	//
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

