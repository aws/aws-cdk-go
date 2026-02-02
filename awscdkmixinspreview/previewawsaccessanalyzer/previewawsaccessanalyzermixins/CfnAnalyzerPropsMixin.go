package previewawsaccessanalyzermixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsaccessanalyzer/previewawsaccessanalyzermixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::AccessAnalyzer::Analyzer` resource specifies a new analyzer.
//
// The analyzer is an object that represents the IAM Access Analyzer feature. An analyzer is required for Access Analyzer to become operational.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAnalyzerPropsMixin := awscdkmixinspreview.Mixins.NewCfnAnalyzerPropsMixin(&CfnAnalyzerMixinProps{
//   	AnalyzerConfiguration: &AnalyzerConfigurationProperty{
//   		InternalAccessConfiguration: &InternalAccessConfigurationProperty{
//   			InternalAccessAnalysisRule: &InternalAccessAnalysisRuleProperty{
//   				Inclusions: []interface{}{
//   					&InternalAccessAnalysisRuleCriteriaProperty{
//   						AccountIds: []*string{
//   							jsii.String("accountIds"),
//   						},
//   						ResourceArns: []*string{
//   							jsii.String("resourceArns"),
//   						},
//   						ResourceTypes: []*string{
//   							jsii.String("resourceTypes"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		UnusedAccessConfiguration: &UnusedAccessConfigurationProperty{
//   			AnalysisRule: &AnalysisRuleProperty{
//   				Exclusions: []interface{}{
//   					&AnalysisRuleCriteriaProperty{
//   						AccountIds: []*string{
//   							jsii.String("accountIds"),
//   						},
//   						ResourceTags: []interface{}{
//   							[]interface{}{
//   								&CfnTag{
//   									Key: jsii.String("key"),
//   									Value: jsii.String("value"),
//   								},
//   							},
//   						},
//   					},
//   				},
//   			},
//   			UnusedAccessAge: jsii.Number(123),
//   		},
//   	},
//   	AnalyzerName: jsii.String("analyzerName"),
//   	ArchiveRules: []interface{}{
//   		&ArchiveRuleProperty{
//   			Filter: []interface{}{
//   				&FilterProperty{
//   					Contains: []*string{
//   						jsii.String("contains"),
//   					},
//   					Eq: []*string{
//   						jsii.String("eq"),
//   					},
//   					Exists: jsii.Boolean(false),
//   					Neq: []*string{
//   						jsii.String("neq"),
//   					},
//   					Property: jsii.String("property"),
//   				},
//   			},
//   			RuleName: jsii.String("ruleName"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-accessanalyzer-analyzer.html
//
type CfnAnalyzerPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnAnalyzerMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAnalyzerPropsMixin
type jsiiProxy_CfnAnalyzerPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnAnalyzerPropsMixin) Props() *CfnAnalyzerMixinProps {
	var returns *CfnAnalyzerMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnalyzerPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::AccessAnalyzer::Analyzer`.
func NewCfnAnalyzerPropsMixin(props *CfnAnalyzerMixinProps, options *mixins.CfnPropertyMixinOptions) CfnAnalyzerPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAnalyzerPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAnalyzerPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_accessanalyzer.mixins.CfnAnalyzerPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::AccessAnalyzer::Analyzer`.
func NewCfnAnalyzerPropsMixin_Override(c CfnAnalyzerPropsMixin, props *CfnAnalyzerMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_accessanalyzer.mixins.CfnAnalyzerPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnAnalyzerPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAnalyzerPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_accessanalyzer.mixins.CfnAnalyzerPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAnalyzerPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_accessanalyzer.mixins.CfnAnalyzerPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAnalyzerPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnAnalyzerPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

