package previewawsvpclatticemixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsvpclattice/previewawsvpclatticemixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a listener rule.
//
// Each listener has a default rule for checking connection requests, but you can define additional rules. Each rule consists of a priority, one or more actions, and one or more conditions. For more information, see [Listener rules](https://docs.aws.amazon.com/vpc-lattice/latest/ug/listeners.html#listener-rules) in the *Amazon VPC Lattice User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRulePropsMixin := awscdkmixinspreview.Mixins.NewCfnRulePropsMixin(&CfnRuleMixinProps{
//   	Action: &ActionProperty{
//   		FixedResponse: &FixedResponseProperty{
//   			StatusCode: jsii.Number(123),
//   		},
//   		Forward: &ForwardProperty{
//   			TargetGroups: []interface{}{
//   				&WeightedTargetGroupProperty{
//   					TargetGroupIdentifier: jsii.String("targetGroupIdentifier"),
//   					Weight: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   	ListenerIdentifier: jsii.String("listenerIdentifier"),
//   	Match: &MatchProperty{
//   		HttpMatch: &HttpMatchProperty{
//   			HeaderMatches: []interface{}{
//   				&HeaderMatchProperty{
//   					CaseSensitive: jsii.Boolean(false),
//   					Match: &HeaderMatchTypeProperty{
//   						Contains: jsii.String("contains"),
//   						Exact: jsii.String("exact"),
//   						Prefix: jsii.String("prefix"),
//   					},
//   					Name: jsii.String("name"),
//   				},
//   			},
//   			Method: jsii.String("method"),
//   			PathMatch: &PathMatchProperty{
//   				CaseSensitive: jsii.Boolean(false),
//   				Match: &PathMatchTypeProperty{
//   					Exact: jsii.String("exact"),
//   					Prefix: jsii.String("prefix"),
//   				},
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Priority: jsii.Number(123),
//   	ServiceIdentifier: jsii.String("serviceIdentifier"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-rule.html
//
type CfnRulePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnRuleMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnRulePropsMixin
type jsiiProxy_CfnRulePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnRulePropsMixin) Props() *CfnRuleMixinProps {
	var returns *CfnRuleMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRulePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::VpcLattice::Rule`.
func NewCfnRulePropsMixin(props *CfnRuleMixinProps, options *mixins.CfnPropertyMixinOptions) CfnRulePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnRulePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnRulePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_vpclattice.mixins.CfnRulePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::VpcLattice::Rule`.
func NewCfnRulePropsMixin_Override(c CfnRulePropsMixin, props *CfnRuleMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_vpclattice.mixins.CfnRulePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnRulePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnRulePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_vpclattice.mixins.CfnRulePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRulePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_vpclattice.mixins.CfnRulePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnRulePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnRulePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

