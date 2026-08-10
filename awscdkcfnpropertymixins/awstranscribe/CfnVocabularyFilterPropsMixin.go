package awstranscribe

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awstranscribe/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a custom vocabulary filter that you can use to mask, delete, or flag specific words from your transcript.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnVocabularyFilterPropsMixin := awscdkcfnpropertymixins.Aws_transcribe.NewCfnVocabularyFilterPropsMixin(&CfnVocabularyFilterMixinProps{
//   	DataAccessRoleArn: jsii.String("dataAccessRoleArn"),
//   	LanguageCode: jsii.String("languageCode"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VocabularyFilterFileUri: jsii.String("vocabularyFilterFileUri"),
//   	VocabularyFilterName: jsii.String("vocabularyFilterName"),
//   	Words: []*string{
//   		jsii.String("words"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transcribe-vocabularyfilter.html
//
type CfnVocabularyFilterPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnVocabularyFilterMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnVocabularyFilterPropsMixin
type jsiiProxy_CfnVocabularyFilterPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnVocabularyFilterPropsMixin) Props() *CfnVocabularyFilterMixinProps {
	var returns *CfnVocabularyFilterMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVocabularyFilterPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Transcribe::VocabularyFilter`.
func NewCfnVocabularyFilterPropsMixin(props *CfnVocabularyFilterMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnVocabularyFilterPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnVocabularyFilterPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnVocabularyFilterPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_transcribe.CfnVocabularyFilterPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Transcribe::VocabularyFilter`.
func NewCfnVocabularyFilterPropsMixin_Override(c CfnVocabularyFilterPropsMixin, props *CfnVocabularyFilterMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_transcribe.CfnVocabularyFilterPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnVocabularyFilterPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnVocabularyFilterPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_transcribe.CfnVocabularyFilterPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnVocabularyFilterPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_transcribe.CfnVocabularyFilterPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnVocabularyFilterPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnVocabularyFilterPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

