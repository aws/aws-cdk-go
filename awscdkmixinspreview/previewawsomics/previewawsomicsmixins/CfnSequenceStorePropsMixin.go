package previewawsomicsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsomics/previewawsomicsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a sequence store.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var s3AccessPolicy interface{}
//
//   cfnSequenceStorePropsMixin := awscdkmixinspreview.Mixins.NewCfnSequenceStorePropsMixin(&CfnSequenceStoreMixinProps{
//   	AccessLogLocation: jsii.String("accessLogLocation"),
//   	Description: jsii.String("description"),
//   	ETagAlgorithmFamily: jsii.String("eTagAlgorithmFamily"),
//   	FallbackLocation: jsii.String("fallbackLocation"),
//   	Name: jsii.String("name"),
//   	PropagatedSetLevelTags: []*string{
//   		jsii.String("propagatedSetLevelTags"),
//   	},
//   	S3AccessPolicy: s3AccessPolicy,
//   	SseConfig: &SseConfigProperty{
//   		KeyArn: jsii.String("keyArn"),
//   		Type: jsii.String("type"),
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-sequencestore.html
//
type CfnSequenceStorePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnSequenceStoreMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnSequenceStorePropsMixin
type jsiiProxy_CfnSequenceStorePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnSequenceStorePropsMixin) Props() *CfnSequenceStoreMixinProps {
	var returns *CfnSequenceStoreMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSequenceStorePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Omics::SequenceStore`.
func NewCfnSequenceStorePropsMixin(props *CfnSequenceStoreMixinProps, options *mixins.CfnPropertyMixinOptions) CfnSequenceStorePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnSequenceStorePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnSequenceStorePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_omics.mixins.CfnSequenceStorePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Omics::SequenceStore`.
func NewCfnSequenceStorePropsMixin_Override(c CfnSequenceStorePropsMixin, props *CfnSequenceStoreMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_omics.mixins.CfnSequenceStorePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnSequenceStorePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSequenceStorePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_omics.mixins.CfnSequenceStorePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSequenceStorePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_omics.mixins.CfnSequenceStorePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSequenceStorePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnSequenceStorePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

