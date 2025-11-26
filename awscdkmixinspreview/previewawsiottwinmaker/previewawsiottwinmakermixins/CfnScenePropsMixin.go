package previewawsiottwinmakermixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsiottwinmaker/previewawsiottwinmakermixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Use the `AWS::IoTTwinMaker::Scene` resource to declare a scene.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnScenePropsMixin := awscdkmixinspreview.Mixins.NewCfnScenePropsMixin(&CfnSceneMixinProps{
//   	Capabilities: []*string{
//   		jsii.String("capabilities"),
//   	},
//   	ContentLocation: jsii.String("contentLocation"),
//   	Description: jsii.String("description"),
//   	SceneId: jsii.String("sceneId"),
//   	SceneMetadata: map[string]*string{
//   		"sceneMetadataKey": jsii.String("sceneMetadata"),
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	WorkspaceId: jsii.String("workspaceId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iottwinmaker-scene.html
//
type CfnScenePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnSceneMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnScenePropsMixin
type jsiiProxy_CfnScenePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnScenePropsMixin) Props() *CfnSceneMixinProps {
	var returns *CfnSceneMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScenePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::IoTTwinMaker::Scene`.
func NewCfnScenePropsMixin(props *CfnSceneMixinProps, options *mixins.CfnPropertyMixinOptions) CfnScenePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnScenePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnScenePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iottwinmaker.mixins.CfnScenePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::IoTTwinMaker::Scene`.
func NewCfnScenePropsMixin_Override(c CfnScenePropsMixin, props *CfnSceneMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iottwinmaker.mixins.CfnScenePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnScenePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnScenePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_iottwinmaker.mixins.CfnScenePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnScenePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_iottwinmaker.mixins.CfnScenePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnScenePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnScenePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

