package previewawsbedrockmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsbedrock/previewawsbedrockmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a new version of an existing Automated Reasoning policy.
//
// This allows you to iterate on your policy rules while maintaining previous versions for rollback or comparison purposes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAutomatedReasoningPolicyVersionPropsMixin := awscdkmixinspreview.Mixins.NewCfnAutomatedReasoningPolicyVersionPropsMixin(&CfnAutomatedReasoningPolicyVersionMixinProps{
//   	LastUpdatedDefinitionHash: jsii.String("lastUpdatedDefinitionHash"),
//   	PolicyArn: jsii.String("policyArn"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-automatedreasoningpolicyversion.html
//
type CfnAutomatedReasoningPolicyVersionPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnAutomatedReasoningPolicyVersionMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAutomatedReasoningPolicyVersionPropsMixin
type jsiiProxy_CfnAutomatedReasoningPolicyVersionPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnAutomatedReasoningPolicyVersionPropsMixin) Props() *CfnAutomatedReasoningPolicyVersionMixinProps {
	var returns *CfnAutomatedReasoningPolicyVersionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutomatedReasoningPolicyVersionPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Bedrock::AutomatedReasoningPolicyVersion`.
func NewCfnAutomatedReasoningPolicyVersionPropsMixin(props *CfnAutomatedReasoningPolicyVersionMixinProps, options *mixins.CfnPropertyMixinOptions) CfnAutomatedReasoningPolicyVersionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAutomatedReasoningPolicyVersionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAutomatedReasoningPolicyVersionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnAutomatedReasoningPolicyVersionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Bedrock::AutomatedReasoningPolicyVersion`.
func NewCfnAutomatedReasoningPolicyVersionPropsMixin_Override(c CfnAutomatedReasoningPolicyVersionPropsMixin, props *CfnAutomatedReasoningPolicyVersionMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnAutomatedReasoningPolicyVersionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnAutomatedReasoningPolicyVersionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAutomatedReasoningPolicyVersionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnAutomatedReasoningPolicyVersionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAutomatedReasoningPolicyVersionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnAutomatedReasoningPolicyVersionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAutomatedReasoningPolicyVersionPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnAutomatedReasoningPolicyVersionPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

