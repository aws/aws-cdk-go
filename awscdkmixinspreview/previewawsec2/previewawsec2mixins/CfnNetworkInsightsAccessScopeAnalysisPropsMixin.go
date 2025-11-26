package previewawsec2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsec2/previewawsec2mixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Describes a Network Access Scope analysis.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnNetworkInsightsAccessScopeAnalysisPropsMixin := awscdkmixinspreview.Mixins.NewCfnNetworkInsightsAccessScopeAnalysisPropsMixin(&CfnNetworkInsightsAccessScopeAnalysisMixinProps{
//   	NetworkInsightsAccessScopeId: jsii.String("networkInsightsAccessScopeId"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinsightsaccessscopeanalysis.html
//
type CfnNetworkInsightsAccessScopeAnalysisPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnNetworkInsightsAccessScopeAnalysisMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnNetworkInsightsAccessScopeAnalysisPropsMixin
type jsiiProxy_CfnNetworkInsightsAccessScopeAnalysisPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnNetworkInsightsAccessScopeAnalysisPropsMixin) Props() *CfnNetworkInsightsAccessScopeAnalysisMixinProps {
	var returns *CfnNetworkInsightsAccessScopeAnalysisMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNetworkInsightsAccessScopeAnalysisPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EC2::NetworkInsightsAccessScopeAnalysis`.
func NewCfnNetworkInsightsAccessScopeAnalysisPropsMixin(props *CfnNetworkInsightsAccessScopeAnalysisMixinProps, options *mixins.CfnPropertyMixinOptions) CfnNetworkInsightsAccessScopeAnalysisPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnNetworkInsightsAccessScopeAnalysisPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnNetworkInsightsAccessScopeAnalysisPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnNetworkInsightsAccessScopeAnalysisPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EC2::NetworkInsightsAccessScopeAnalysis`.
func NewCfnNetworkInsightsAccessScopeAnalysisPropsMixin_Override(c CfnNetworkInsightsAccessScopeAnalysisPropsMixin, props *CfnNetworkInsightsAccessScopeAnalysisMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnNetworkInsightsAccessScopeAnalysisPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnNetworkInsightsAccessScopeAnalysisPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnNetworkInsightsAccessScopeAnalysisPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnNetworkInsightsAccessScopeAnalysisPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnNetworkInsightsAccessScopeAnalysisPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnNetworkInsightsAccessScopeAnalysisPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnNetworkInsightsAccessScopeAnalysisPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnNetworkInsightsAccessScopeAnalysisPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

