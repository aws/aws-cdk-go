package previewawssagemakermixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawssagemaker/previewawssagemakermixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a `Domain` .
//
// A domain consists of an associated Amazon Elastic File System volume, a list of authorized users, and a variety of security, application, policy, and Amazon Virtual Private Cloud (VPC) configurations. Users within a domain can share notebook files and other artifacts with each other.
//
// *EFS storage*
//
// When a domain is created, an EFS volume is created for use by all of the users within the domain. Each user receives a private home directory within the EFS volume for notebooks, Git repositories, and data files.
//
// SageMaker AI uses the AWS Key Management Service ( AWS KMS) to encrypt the EFS volume attached to the domain with an AWS managed key by default. For more control, you can specify a customer managed key. For more information, see [Protect Data at Rest Using Encryption](https://docs.aws.amazon.com/sagemaker/latest/dg/encryption-at-rest.html) .
//
// *VPC configuration*
//
// All traffic between the domain and the Amazon EFS volume is through the specified VPC and subnets. For other traffic, you can specify the `AppNetworkAccessType` parameter. `AppNetworkAccessType` corresponds to the network access type that you choose when you onboard to the domain. The following options are available:
//
// - `PublicInternetOnly` - Non-EFS traffic goes through a VPC managed by Amazon SageMaker AI, which allows internet access. This is the default value.
// - `VpcOnly` - All traffic is through the specified VPC and subnets. Internet access is disabled by default. To allow internet access, you must specify a NAT gateway.
//
// When internet access is disabled, you won't be able to run a Amazon SageMaker AI Studio notebook or to train or host models unless your VPC has an interface endpoint to the SageMaker AI API and runtime or a NAT gateway and your security groups allow outbound connections.
//
// > NFS traffic over TCP on port 2049 needs to be allowed in both inbound and outbound rules in order to launch a Amazon SageMaker AI Studio app successfully.
//
// For more information, see [Connect Amazon SageMaker AI Studio Notebooks to Resources in a VPC](https://docs.aws.amazon.com/sagemaker/latest/dg/studio-notebooks-and-internet-access.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDomainPropsMixin := awscdkmixinspreview.Mixins.NewCfnDomainPropsMixin(&CfnDomainMixinProps{
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
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-domain.html
//
type CfnDomainPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnDomainMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDomainPropsMixin
type jsiiProxy_CfnDomainPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnDomainPropsMixin) Props() *CfnDomainMixinProps {
	var returns *CfnDomainMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SageMaker::Domain`.
func NewCfnDomainPropsMixin(props *CfnDomainMixinProps, options *mixins.CfnPropertyMixinOptions) CfnDomainPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDomainPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDomainPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnDomainPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SageMaker::Domain`.
func NewCfnDomainPropsMixin_Override(c CfnDomainPropsMixin, props *CfnDomainMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnDomainPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnDomainPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDomainPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnDomainPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDomainPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnDomainPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDomainPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"applyTo",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDomainPropsMixin) Supports(construct constructs.IConstruct) *bool {
	if err := c.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

