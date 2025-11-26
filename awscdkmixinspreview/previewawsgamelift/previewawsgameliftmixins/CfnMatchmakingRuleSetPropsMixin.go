package previewawsgameliftmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsgamelift/previewawsgameliftmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a new rule set for FlexMatch matchmaking.
//
// A rule set describes the type of match to create, such as the number and size of teams. It also sets the parameters for acceptable player matches, such as minimum skill level or character type.
//
// To create a matchmaking rule set, provide unique rule set name and the rule set body in JSON format. Rule sets must be defined in the same Region as the matchmaking configuration they are used with.
//
// Since matchmaking rule sets cannot be edited, it is a good idea to check the rule set syntax.
//
// *Learn more*
//
// - [Build a rule set](https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/match-rulesets.html)
// - [Design a matchmaker](https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/match-configuration.html)
// - [Matchmaking with FlexMatch](https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/match-intro.html)
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnMatchmakingRuleSetPropsMixin := awscdkmixinspreview.Mixins.NewCfnMatchmakingRuleSetPropsMixin(&CfnMatchmakingRuleSetMixinProps{
//   	Name: jsii.String("name"),
//   	RuleSetBody: jsii.String("ruleSetBody"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-matchmakingruleset.html
//
type CfnMatchmakingRuleSetPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnMatchmakingRuleSetMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnMatchmakingRuleSetPropsMixin
type jsiiProxy_CfnMatchmakingRuleSetPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnMatchmakingRuleSetPropsMixin) Props() *CfnMatchmakingRuleSetMixinProps {
	var returns *CfnMatchmakingRuleSetMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMatchmakingRuleSetPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::GameLift::MatchmakingRuleSet`.
func NewCfnMatchmakingRuleSetPropsMixin(props *CfnMatchmakingRuleSetMixinProps, options *mixins.CfnPropertyMixinOptions) CfnMatchmakingRuleSetPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnMatchmakingRuleSetPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnMatchmakingRuleSetPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnMatchmakingRuleSetPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::GameLift::MatchmakingRuleSet`.
func NewCfnMatchmakingRuleSetPropsMixin_Override(c CfnMatchmakingRuleSetPropsMixin, props *CfnMatchmakingRuleSetMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnMatchmakingRuleSetPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnMatchmakingRuleSetPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnMatchmakingRuleSetPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnMatchmakingRuleSetPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnMatchmakingRuleSetPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnMatchmakingRuleSetPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnMatchmakingRuleSetPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnMatchmakingRuleSetPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

