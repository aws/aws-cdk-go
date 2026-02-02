package previewawsxraymixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsxray/previewawsxraymixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Use the `AWS::XRay::SamplingRule` resource to specify a sampling rule, which controls sampling behavior for instrumented applications.
//
// Include a `SamplingRule` entity to create or update a sampling rule.
//
// > `SamplingRule.Version` can only be set when creating a sampling rule. Updating the version will cause the update to fail.
//
// Services retrieve rules with [GetSamplingRules](https://docs.aws.amazon.com//xray/latest/api/API_GetSamplingRules.html) , and evaluate each rule in ascending order of *priority* for each request. If a rule matches, the service records a trace, borrowing it from the reservoir size. After 10 seconds, the service reports back to X-Ray with [GetSamplingTargets](https://docs.aws.amazon.com//xray/latest/api/API_GetSamplingTargets.html) to get updated versions of each in-use rule. The updated rule contains a trace quota that the service can use instead of borrowing from the reservoir.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnSamplingRulePropsMixin := awscdkmixinspreview.Mixins.NewCfnSamplingRulePropsMixin(&CfnSamplingRuleMixinProps{
//   	RuleName: jsii.String("ruleName"),
//   	SamplingRule: &SamplingRuleProperty{
//   		Attributes: map[string]*string{
//   			"attributesKey": jsii.String("attributes"),
//   		},
//   		FixedRate: jsii.Number(123),
//   		Host: jsii.String("host"),
//   		HttpMethod: jsii.String("httpMethod"),
//   		Priority: jsii.Number(123),
//   		ReservoirSize: jsii.Number(123),
//   		ResourceArn: jsii.String("resourceArn"),
//   		RuleArn: jsii.String("ruleArn"),
//   		RuleName: jsii.String("ruleName"),
//   		ServiceName: jsii.String("serviceName"),
//   		ServiceType: jsii.String("serviceType"),
//   		UrlPath: jsii.String("urlPath"),
//   		Version: jsii.Number(123),
//   	},
//   	SamplingRuleRecord: &SamplingRuleRecordProperty{
//   		CreatedAt: jsii.String("createdAt"),
//   		ModifiedAt: jsii.String("modifiedAt"),
//   		SamplingRule: &SamplingRuleProperty{
//   			Attributes: map[string]*string{
//   				"attributesKey": jsii.String("attributes"),
//   			},
//   			FixedRate: jsii.Number(123),
//   			Host: jsii.String("host"),
//   			HttpMethod: jsii.String("httpMethod"),
//   			Priority: jsii.Number(123),
//   			ReservoirSize: jsii.Number(123),
//   			ResourceArn: jsii.String("resourceArn"),
//   			RuleArn: jsii.String("ruleArn"),
//   			RuleName: jsii.String("ruleName"),
//   			ServiceName: jsii.String("serviceName"),
//   			ServiceType: jsii.String("serviceType"),
//   			UrlPath: jsii.String("urlPath"),
//   			Version: jsii.Number(123),
//   		},
//   	},
//   	SamplingRuleUpdate: &SamplingRuleUpdateProperty{
//   		Attributes: map[string]*string{
//   			"attributesKey": jsii.String("attributes"),
//   		},
//   		FixedRate: jsii.Number(123),
//   		Host: jsii.String("host"),
//   		HttpMethod: jsii.String("httpMethod"),
//   		Priority: jsii.Number(123),
//   		ReservoirSize: jsii.Number(123),
//   		ResourceArn: jsii.String("resourceArn"),
//   		RuleArn: jsii.String("ruleArn"),
//   		RuleName: jsii.String("ruleName"),
//   		ServiceName: jsii.String("serviceName"),
//   		ServiceType: jsii.String("serviceType"),
//   		UrlPath: jsii.String("urlPath"),
//   	},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-xray-samplingrule.html
//
type CfnSamplingRulePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnSamplingRuleMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnSamplingRulePropsMixin
type jsiiProxy_CfnSamplingRulePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnSamplingRulePropsMixin) Props() *CfnSamplingRuleMixinProps {
	var returns *CfnSamplingRuleMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSamplingRulePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::XRay::SamplingRule`.
func NewCfnSamplingRulePropsMixin(props *CfnSamplingRuleMixinProps, options *mixins.CfnPropertyMixinOptions) CfnSamplingRulePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnSamplingRulePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnSamplingRulePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_xray.mixins.CfnSamplingRulePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::XRay::SamplingRule`.
func NewCfnSamplingRulePropsMixin_Override(c CfnSamplingRulePropsMixin, props *CfnSamplingRuleMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_xray.mixins.CfnSamplingRulePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnSamplingRulePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSamplingRulePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_xray.mixins.CfnSamplingRulePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSamplingRulePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_xray.mixins.CfnSamplingRulePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSamplingRulePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnSamplingRulePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

