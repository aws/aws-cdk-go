package previewawssagemakermixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawssagemaker/previewawssagemakermixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a private space or a space used for real time collaboration in a domain.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnSpacePropsMixin := awscdkmixinspreview.Mixins.NewCfnSpacePropsMixin(&CfnSpaceMixinProps{
//   	DomainId: jsii.String("domainId"),
//   	OwnershipSettings: &OwnershipSettingsProperty{
//   		OwnerUserProfileName: jsii.String("ownerUserProfileName"),
//   	},
//   	SpaceDisplayName: jsii.String("spaceDisplayName"),
//   	SpaceName: jsii.String("spaceName"),
//   	SpaceSettings: &SpaceSettingsProperty{
//   		AppType: jsii.String("appType"),
//   		CodeEditorAppSettings: &SpaceCodeEditorAppSettingsProperty{
//   			AppLifecycleManagement: &SpaceAppLifecycleManagementProperty{
//   				IdleSettings: &SpaceIdleSettingsProperty{
//   					IdleTimeoutInMinutes: jsii.Number(123),
//   				},
//   			},
//   			DefaultResourceSpec: &ResourceSpecProperty{
//   				InstanceType: jsii.String("instanceType"),
//   				LifecycleConfigArn: jsii.String("lifecycleConfigArn"),
//   				SageMakerImageArn: jsii.String("sageMakerImageArn"),
//   				SageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   			},
//   		},
//   		CustomFileSystems: []interface{}{
//   			&CustomFileSystemProperty{
//   				EfsFileSystem: &EFSFileSystemProperty{
//   					FileSystemId: jsii.String("fileSystemId"),
//   				},
//   				FSxLustreFileSystem: &FSxLustreFileSystemProperty{
//   					FileSystemId: jsii.String("fileSystemId"),
//   				},
//   				S3FileSystem: &S3FileSystemProperty{
//   					S3Uri: jsii.String("s3Uri"),
//   				},
//   			},
//   		},
//   		JupyterLabAppSettings: &SpaceJupyterLabAppSettingsProperty{
//   			AppLifecycleManagement: &SpaceAppLifecycleManagementProperty{
//   				IdleSettings: &SpaceIdleSettingsProperty{
//   					IdleTimeoutInMinutes: jsii.Number(123),
//   				},
//   			},
//   			CodeRepositories: []interface{}{
//   				&CodeRepositoryProperty{
//   					RepositoryUrl: jsii.String("repositoryUrl"),
//   				},
//   			},
//   			DefaultResourceSpec: &ResourceSpecProperty{
//   				InstanceType: jsii.String("instanceType"),
//   				LifecycleConfigArn: jsii.String("lifecycleConfigArn"),
//   				SageMakerImageArn: jsii.String("sageMakerImageArn"),
//   				SageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
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
//   		RemoteAccess: jsii.String("remoteAccess"),
//   		SpaceManagedResources: jsii.String("spaceManagedResources"),
//   		SpaceStorageSettings: &SpaceStorageSettingsProperty{
//   			EbsStorageSettings: &EbsStorageSettingsProperty{
//   				EbsVolumeSizeInGb: jsii.Number(123),
//   			},
//   		},
//   	},
//   	SpaceSharingSettings: &SpaceSharingSettingsProperty{
//   		SharingType: jsii.String("sharingType"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-space.html
//
type CfnSpacePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnSpaceMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnSpacePropsMixin
type jsiiProxy_CfnSpacePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnSpacePropsMixin) Props() *CfnSpaceMixinProps {
	var returns *CfnSpaceMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSpacePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SageMaker::Space`.
func NewCfnSpacePropsMixin(props *CfnSpaceMixinProps, options *mixins.CfnPropertyMixinOptions) CfnSpacePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnSpacePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnSpacePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnSpacePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SageMaker::Space`.
func NewCfnSpacePropsMixin_Override(c CfnSpacePropsMixin, props *CfnSpaceMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnSpacePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnSpacePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSpacePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnSpacePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSpacePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnSpacePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSpacePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnSpacePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

