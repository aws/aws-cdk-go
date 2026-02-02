package previewawscleanroomsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawscleanrooms/previewawscleanroomsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// An object that defines the privacy budget template.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnPrivacyBudgetTemplatePropsMixin := awscdkmixinspreview.Mixins.NewCfnPrivacyBudgetTemplatePropsMixin(&CfnPrivacyBudgetTemplateMixinProps{
//   	AutoRefresh: jsii.String("autoRefresh"),
//   	MembershipIdentifier: jsii.String("membershipIdentifier"),
//   	Parameters: &ParametersProperty{
//   		BudgetParameters: []interface{}{
//   			&BudgetParameterProperty{
//   				AutoRefresh: jsii.String("autoRefresh"),
//   				Budget: jsii.Number(123),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   		Epsilon: jsii.Number(123),
//   		ResourceArn: jsii.String("resourceArn"),
//   		UsersNoisePerQuery: jsii.Number(123),
//   	},
//   	PrivacyBudgetType: jsii.String("privacyBudgetType"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-privacybudgettemplate.html
//
type CfnPrivacyBudgetTemplatePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnPrivacyBudgetTemplateMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnPrivacyBudgetTemplatePropsMixin
type jsiiProxy_CfnPrivacyBudgetTemplatePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnPrivacyBudgetTemplatePropsMixin) Props() *CfnPrivacyBudgetTemplateMixinProps {
	var returns *CfnPrivacyBudgetTemplateMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPrivacyBudgetTemplatePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CleanRooms::PrivacyBudgetTemplate`.
func NewCfnPrivacyBudgetTemplatePropsMixin(props *CfnPrivacyBudgetTemplateMixinProps, options *mixins.CfnPropertyMixinOptions) CfnPrivacyBudgetTemplatePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnPrivacyBudgetTemplatePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnPrivacyBudgetTemplatePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnPrivacyBudgetTemplatePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CleanRooms::PrivacyBudgetTemplate`.
func NewCfnPrivacyBudgetTemplatePropsMixin_Override(c CfnPrivacyBudgetTemplatePropsMixin, props *CfnPrivacyBudgetTemplateMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnPrivacyBudgetTemplatePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnPrivacyBudgetTemplatePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnPrivacyBudgetTemplatePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnPrivacyBudgetTemplatePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPrivacyBudgetTemplatePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnPrivacyBudgetTemplatePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPrivacyBudgetTemplatePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnPrivacyBudgetTemplatePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

