package previewawssagemakermixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawssagemaker/previewawssagemakermixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a configuration for running a SageMaker AI image as a KernelGateway app.
//
// The configuration specifies the Amazon Elastic File System storage volume on the image, and a list of the kernels in the image.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAppImageConfigPropsMixin := awscdkmixinspreview.Mixins.NewCfnAppImageConfigPropsMixin(&CfnAppImageConfigMixinProps{
//   	AppImageConfigName: jsii.String("appImageConfigName"),
//   	CodeEditorAppImageConfig: &CodeEditorAppImageConfigProperty{
//   		ContainerConfig: &ContainerConfigProperty{
//   			ContainerArguments: []*string{
//   				jsii.String("containerArguments"),
//   			},
//   			ContainerEntrypoint: []*string{
//   				jsii.String("containerEntrypoint"),
//   			},
//   			ContainerEnvironmentVariables: []interface{}{
//   				&CustomImageContainerEnvironmentVariableProperty{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	JupyterLabAppImageConfig: &JupyterLabAppImageConfigProperty{
//   		ContainerConfig: &ContainerConfigProperty{
//   			ContainerArguments: []*string{
//   				jsii.String("containerArguments"),
//   			},
//   			ContainerEntrypoint: []*string{
//   				jsii.String("containerEntrypoint"),
//   			},
//   			ContainerEnvironmentVariables: []interface{}{
//   				&CustomImageContainerEnvironmentVariableProperty{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	KernelGatewayImageConfig: &KernelGatewayImageConfigProperty{
//   		FileSystemConfig: &FileSystemConfigProperty{
//   			DefaultGid: jsii.Number(123),
//   			DefaultUid: jsii.Number(123),
//   			MountPath: jsii.String("mountPath"),
//   		},
//   		KernelSpecs: []interface{}{
//   			&KernelSpecProperty{
//   				DisplayName: jsii.String("displayName"),
//   				Name: jsii.String("name"),
//   			},
//   		},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-appimageconfig.html
//
type CfnAppImageConfigPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnAppImageConfigMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAppImageConfigPropsMixin
type jsiiProxy_CfnAppImageConfigPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnAppImageConfigPropsMixin) Props() *CfnAppImageConfigMixinProps {
	var returns *CfnAppImageConfigMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAppImageConfigPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SageMaker::AppImageConfig`.
func NewCfnAppImageConfigPropsMixin(props *CfnAppImageConfigMixinProps, options *mixins.CfnPropertyMixinOptions) CfnAppImageConfigPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAppImageConfigPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAppImageConfigPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnAppImageConfigPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SageMaker::AppImageConfig`.
func NewCfnAppImageConfigPropsMixin_Override(c CfnAppImageConfigPropsMixin, props *CfnAppImageConfigMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnAppImageConfigPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnAppImageConfigPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAppImageConfigPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnAppImageConfigPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAppImageConfigPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnAppImageConfigPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAppImageConfigPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnAppImageConfigPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

