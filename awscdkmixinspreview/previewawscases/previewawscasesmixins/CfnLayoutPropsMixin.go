package previewawscasesmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawscases/previewawscasesmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a layout in the Cases domain.
//
// Layouts define the following configuration in the top section and More Info tab of the Cases user interface:
//
// - Fields to display to the users
// - Field ordering
//
// > Title and Status fields cannot be part of layouts since they are not configurable.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLayoutPropsMixin := awscdkmixinspreview.Mixins.NewCfnLayoutPropsMixin(&CfnLayoutMixinProps{
//   	Content: &LayoutContentProperty{
//   		Basic: &BasicLayoutProperty{
//   			MoreInfo: &LayoutSectionsProperty{
//   				Sections: []interface{}{
//   					&SectionProperty{
//   						FieldGroup: &FieldGroupProperty{
//   							Fields: []interface{}{
//   								&FieldItemProperty{
//   									Id: jsii.String("id"),
//   								},
//   							},
//   							Name: jsii.String("name"),
//   						},
//   					},
//   				},
//   			},
//   			TopPanel: &LayoutSectionsProperty{
//   				Sections: []interface{}{
//   					&SectionProperty{
//   						FieldGroup: &FieldGroupProperty{
//   							Fields: []interface{}{
//   								&FieldItemProperty{
//   									Id: jsii.String("id"),
//   								},
//   							},
//   							Name: jsii.String("name"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	DomainId: jsii.String("domainId"),
//   	Name: jsii.String("name"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cases-layout.html
//
type CfnLayoutPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnLayoutMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnLayoutPropsMixin
type jsiiProxy_CfnLayoutPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnLayoutPropsMixin) Props() *CfnLayoutMixinProps {
	var returns *CfnLayoutMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayoutPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Cases::Layout`.
func NewCfnLayoutPropsMixin(props *CfnLayoutMixinProps, options *mixins.CfnPropertyMixinOptions) CfnLayoutPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnLayoutPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnLayoutPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cases.mixins.CfnLayoutPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Cases::Layout`.
func NewCfnLayoutPropsMixin_Override(c CfnLayoutPropsMixin, props *CfnLayoutMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cases.mixins.CfnLayoutPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnLayoutPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnLayoutPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_cases.mixins.CfnLayoutPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLayoutPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_cases.mixins.CfnLayoutPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnLayoutPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnLayoutPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

