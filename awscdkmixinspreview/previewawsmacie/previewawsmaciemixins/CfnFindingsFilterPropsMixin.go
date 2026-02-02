package previewawsmaciemixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsmacie/previewawsmaciemixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::Macie::FindingsFilter` resource specifies a findings filter.
//
// In Amazon Macie , a *findings filter* , also referred to as a *filter rule* , is a set of custom criteria that specifies which findings to include or exclude from the results of a query for findings. The criteria can help you identify and focus on findings that have specific characteristics, such as severity, type, or the name of an affected AWS resource. You can also configure a findings filter to suppress (automatically archive) findings that match the filter's criteria. For more information, see [Filtering Macie findings](https://docs.aws.amazon.com/macie/latest/user/findings-filter-overview.html) in the *Amazon Macie User Guide* .
//
// An `AWS::Macie::Session` resource must exist for an AWS account before you can create an `AWS::Macie::FindingsFilter` resource for the account. Use a [DependsOn attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html) to ensure that an `AWS::Macie::Session` resource is created before other Macie resources are created for an account. For example, `"DependsOn": "Session"` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnFindingsFilterPropsMixin := awscdkmixinspreview.Mixins.NewCfnFindingsFilterPropsMixin(&CfnFindingsFilterMixinProps{
//   	Action: jsii.String("action"),
//   	Description: jsii.String("description"),
//   	FindingCriteria: &FindingCriteriaProperty{
//   		Criterion: map[string]interface{}{
//   			"criterionKey": &CriterionAdditionalPropertiesProperty{
//   				"eq": []*string{
//   					jsii.String("eq"),
//   				},
//   				"gt": jsii.Number(123),
//   				"gte": jsii.Number(123),
//   				"lt": jsii.Number(123),
//   				"lte": jsii.Number(123),
//   				"neq": []*string{
//   					jsii.String("neq"),
//   				},
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Position: jsii.Number(123),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-findingsfilter.html
//
type CfnFindingsFilterPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnFindingsFilterMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnFindingsFilterPropsMixin
type jsiiProxy_CfnFindingsFilterPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnFindingsFilterPropsMixin) Props() *CfnFindingsFilterMixinProps {
	var returns *CfnFindingsFilterMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFindingsFilterPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Macie::FindingsFilter`.
func NewCfnFindingsFilterPropsMixin(props *CfnFindingsFilterMixinProps, options *mixins.CfnPropertyMixinOptions) CfnFindingsFilterPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnFindingsFilterPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnFindingsFilterPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_macie.mixins.CfnFindingsFilterPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Macie::FindingsFilter`.
func NewCfnFindingsFilterPropsMixin_Override(c CfnFindingsFilterPropsMixin, props *CfnFindingsFilterMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_macie.mixins.CfnFindingsFilterPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnFindingsFilterPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnFindingsFilterPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_macie.mixins.CfnFindingsFilterPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFindingsFilterPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_macie.mixins.CfnFindingsFilterPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnFindingsFilterPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnFindingsFilterPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

