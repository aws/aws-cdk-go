package awsaccessanalyzer

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsaccessanalyzer/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates an archive rule for the specified analyzer.
//
// Archive rules automatically archive new findings that meet the criteria you define when you create the rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnArchiveRulePropsMixin := awscdkcfnpropertymixins.Aws_accessanalyzer.NewCfnArchiveRulePropsMixin(&CfnArchiveRuleMixinProps{
//   	AnalyzerName: jsii.String("analyzerName"),
//   	Filter: map[string]interface{}{
//   		"filterKey": &FilterItemsProperty{
//   			"contains": []*string{
//   				jsii.String("contains"),
//   			},
//   			"eq": []*string{
//   				jsii.String("eq"),
//   			},
//   			"exists": jsii.Boolean(false),
//   			"neq": []*string{
//   				jsii.String("neq"),
//   			},
//   		},
//   	},
//   	RuleName: jsii.String("ruleName"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-accessanalyzer-archiverule.html
//
type CfnArchiveRulePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnArchiveRuleMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnArchiveRulePropsMixin
type jsiiProxy_CfnArchiveRulePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnArchiveRulePropsMixin) Props() *CfnArchiveRuleMixinProps {
	var returns *CfnArchiveRuleMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnArchiveRulePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::AccessAnalyzer::ArchiveRule`.
func NewCfnArchiveRulePropsMixin(props *CfnArchiveRuleMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnArchiveRulePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnArchiveRulePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnArchiveRulePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_accessanalyzer.CfnArchiveRulePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::AccessAnalyzer::ArchiveRule`.
func NewCfnArchiveRulePropsMixin_Override(c CfnArchiveRulePropsMixin, props *CfnArchiveRuleMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_accessanalyzer.CfnArchiveRulePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnArchiveRulePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnArchiveRulePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_accessanalyzer.CfnArchiveRulePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnArchiveRulePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_accessanalyzer.CfnArchiveRulePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnArchiveRulePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnArchiveRulePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

