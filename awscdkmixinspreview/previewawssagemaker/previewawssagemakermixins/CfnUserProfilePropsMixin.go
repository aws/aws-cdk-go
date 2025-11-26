package previewawssagemakermixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawssagemaker/previewawssagemakermixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a user profile.
//
// A user profile represents a single user within a domain, and is the main way to reference a "person" for the purposes of sharing, reporting, and other user-oriented features. This entity is created when a user onboards to Amazon SageMaker Studio. If an administrator invites a person by email or imports them from IAM Identity Center , a user profile is automatically created. A user profile is the primary holder of settings for an individual user and has a reference to the user's private Amazon Elastic File System (EFS) home directory.
//
// > If you're using IAM Identity Center authentication, a user in IAM Identity Center , or a group in IAM Identity Center containing that user, must be assigned to the Amazon SageMaker Studio application from the IAM Identity Center Console to create a user profile. For more information about application assignment, see [Assign user access](https://docs.aws.amazon.com/singlesignon/latest/userguide/assignuserstoapp.html) . After assignment is complete, a user profile can be created for that user in IAM Identity Center with AWS CloudFormation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnUserProfilePropsMixin := awscdkmixinspreview.Mixins.NewCfnUserProfilePropsMixin(&CfnUserProfileMixinProps{
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
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-userprofile.html
//
type CfnUserProfilePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnUserProfileMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnUserProfilePropsMixin
type jsiiProxy_CfnUserProfilePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnUserProfilePropsMixin) Props() *CfnUserProfileMixinProps {
	var returns *CfnUserProfileMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserProfilePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SageMaker::UserProfile`.
func NewCfnUserProfilePropsMixin(props *CfnUserProfileMixinProps, options *mixins.CfnPropertyMixinOptions) CfnUserProfilePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnUserProfilePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnUserProfilePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnUserProfilePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SageMaker::UserProfile`.
func NewCfnUserProfilePropsMixin_Override(c CfnUserProfilePropsMixin, props *CfnUserProfileMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnUserProfilePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnUserProfilePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnUserProfilePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnUserProfilePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnUserProfilePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnUserProfilePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnUserProfilePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnUserProfilePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

