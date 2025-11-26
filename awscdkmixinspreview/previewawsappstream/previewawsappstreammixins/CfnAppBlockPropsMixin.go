package previewawsappstreammixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsappstream/previewawsappstreammixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// This resource creates an app block.
//
// App blocks store details about the virtual hard disk that contains the files for the application in an S3 bucket. It also stores the setup script with details about how to mount the virtual hard disk. App blocks are only supported for Elastic fleets.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAppBlockPropsMixin := awscdkmixinspreview.Mixins.NewCfnAppBlockPropsMixin(&CfnAppBlockMixinProps{
//   	Description: jsii.String("description"),
//   	DisplayName: jsii.String("displayName"),
//   	Name: jsii.String("name"),
//   	PackagingType: jsii.String("packagingType"),
//   	PostSetupScriptDetails: &ScriptDetailsProperty{
//   		ExecutableParameters: jsii.String("executableParameters"),
//   		ExecutablePath: jsii.String("executablePath"),
//   		ScriptS3Location: &S3LocationProperty{
//   			S3Bucket: jsii.String("s3Bucket"),
//   			S3Key: jsii.String("s3Key"),
//   		},
//   		TimeoutInSeconds: jsii.Number(123),
//   	},
//   	SetupScriptDetails: &ScriptDetailsProperty{
//   		ExecutableParameters: jsii.String("executableParameters"),
//   		ExecutablePath: jsii.String("executablePath"),
//   		ScriptS3Location: &S3LocationProperty{
//   			S3Bucket: jsii.String("s3Bucket"),
//   			S3Key: jsii.String("s3Key"),
//   		},
//   		TimeoutInSeconds: jsii.Number(123),
//   	},
//   	SourceS3Location: &S3LocationProperty{
//   		S3Bucket: jsii.String("s3Bucket"),
//   		S3Key: jsii.String("s3Key"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-appblock.html
//
type CfnAppBlockPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnAppBlockMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAppBlockPropsMixin
type jsiiProxy_CfnAppBlockPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnAppBlockPropsMixin) Props() *CfnAppBlockMixinProps {
	var returns *CfnAppBlockMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAppBlockPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::AppStream::AppBlock`.
func NewCfnAppBlockPropsMixin(props *CfnAppBlockMixinProps, options *mixins.CfnPropertyMixinOptions) CfnAppBlockPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAppBlockPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAppBlockPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_appstream.mixins.CfnAppBlockPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::AppStream::AppBlock`.
func NewCfnAppBlockPropsMixin_Override(c CfnAppBlockPropsMixin, props *CfnAppBlockMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_appstream.mixins.CfnAppBlockPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnAppBlockPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAppBlockPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_appstream.mixins.CfnAppBlockPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAppBlockPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_appstream.mixins.CfnAppBlockPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAppBlockPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnAppBlockPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

