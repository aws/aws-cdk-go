package previewawscustomerprofilesmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawscustomerprofiles/previewawscustomerprofilesmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A calculated attribute definition for Customer Profiles.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCalculatedAttributeDefinitionPropsMixin := awscdkmixinspreview.Mixins.NewCfnCalculatedAttributeDefinitionPropsMixin(&CfnCalculatedAttributeDefinitionMixinProps{
//   	AttributeDetails: &AttributeDetailsProperty{
//   		Attributes: []interface{}{
//   			&AttributeItemProperty{
//   				Name: jsii.String("name"),
//   			},
//   		},
//   		Expression: jsii.String("expression"),
//   	},
//   	CalculatedAttributeName: jsii.String("calculatedAttributeName"),
//   	Conditions: &ConditionsProperty{
//   		ObjectCount: jsii.Number(123),
//   		Range: &RangeProperty{
//   			TimestampFormat: jsii.String("timestampFormat"),
//   			TimestampSource: jsii.String("timestampSource"),
//   			Unit: jsii.String("unit"),
//   			Value: jsii.Number(123),
//   			ValueRange: &ValueRangeProperty{
//   				End: jsii.Number(123),
//   				Start: jsii.Number(123),
//   			},
//   		},
//   		Threshold: &ThresholdProperty{
//   			Operator: jsii.String("operator"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	DisplayName: jsii.String("displayName"),
//   	DomainName: jsii.String("domainName"),
//   	Statistic: jsii.String("statistic"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	UseHistoricalData: jsii.Boolean(false),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-calculatedattributedefinition.html
//
type CfnCalculatedAttributeDefinitionPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnCalculatedAttributeDefinitionMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnCalculatedAttributeDefinitionPropsMixin
type jsiiProxy_CfnCalculatedAttributeDefinitionPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnCalculatedAttributeDefinitionPropsMixin) Props() *CfnCalculatedAttributeDefinitionMixinProps {
	var returns *CfnCalculatedAttributeDefinitionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCalculatedAttributeDefinitionPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CustomerProfiles::CalculatedAttributeDefinition`.
func NewCfnCalculatedAttributeDefinitionPropsMixin(props *CfnCalculatedAttributeDefinitionMixinProps, options *mixins.CfnPropertyMixinOptions) CfnCalculatedAttributeDefinitionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnCalculatedAttributeDefinitionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCalculatedAttributeDefinitionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_customerprofiles.mixins.CfnCalculatedAttributeDefinitionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CustomerProfiles::CalculatedAttributeDefinition`.
func NewCfnCalculatedAttributeDefinitionPropsMixin_Override(c CfnCalculatedAttributeDefinitionPropsMixin, props *CfnCalculatedAttributeDefinitionMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_customerprofiles.mixins.CfnCalculatedAttributeDefinitionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnCalculatedAttributeDefinitionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCalculatedAttributeDefinitionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_customerprofiles.mixins.CfnCalculatedAttributeDefinitionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCalculatedAttributeDefinitionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_customerprofiles.mixins.CfnCalculatedAttributeDefinitionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCalculatedAttributeDefinitionPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnCalculatedAttributeDefinitionPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

