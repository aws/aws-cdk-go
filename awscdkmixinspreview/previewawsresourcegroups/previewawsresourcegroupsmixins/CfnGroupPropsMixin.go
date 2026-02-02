package previewawsresourcegroupsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsresourcegroups/previewawsresourcegroupsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a resource group with the specified name and description.
//
// You can optionally include either a resource query or a service configuration. For more information about constructing a resource query, see [Build queries and groups in Resource Groups](https://docs.aws.amazon.com//ARG/latest/userguide/getting_started-query.html) in the *Resource Groups User Guide* . For more information about service-linked groups and service configurations, see [Service configurations for Resource Groups](https://docs.aws.amazon.com//ARG/latest/APIReference/about-slg.html) .
//
// *Minimum permissions*
//
// To run this command, you must have the following permissions:
//
// - `resource-groups:CreateGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnGroupPropsMixin := awscdkmixinspreview.Mixins.NewCfnGroupPropsMixin(&CfnGroupMixinProps{
//   	Configuration: []interface{}{
//   		&ConfigurationItemProperty{
//   			Parameters: []interface{}{
//   				&ConfigurationParameterProperty{
//   					Name: jsii.String("name"),
//   					Values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   			},
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	ResourceQuery: &ResourceQueryProperty{
//   		Query: &QueryProperty{
//   			ResourceTypeFilters: []*string{
//   				jsii.String("resourceTypeFilters"),
//   			},
//   			StackIdentifier: jsii.String("stackIdentifier"),
//   			TagFilters: []interface{}{
//   				&TagFilterProperty{
//   					Key: jsii.String("key"),
//   					Values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   			},
//   		},
//   		Type: jsii.String("type"),
//   	},
//   	Resources: []*string{
//   		jsii.String("resources"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resourcegroups-group.html
//
type CfnGroupPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnGroupMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnGroupPropsMixin
type jsiiProxy_CfnGroupPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnGroupPropsMixin) Props() *CfnGroupMixinProps {
	var returns *CfnGroupMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGroupPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ResourceGroups::Group`.
func NewCfnGroupPropsMixin(props *CfnGroupMixinProps, options *mixins.CfnPropertyMixinOptions) CfnGroupPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnGroupPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnGroupPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_resourcegroups.mixins.CfnGroupPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ResourceGroups::Group`.
func NewCfnGroupPropsMixin_Override(c CfnGroupPropsMixin, props *CfnGroupMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_resourcegroups.mixins.CfnGroupPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnGroupPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnGroupPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_resourcegroups.mixins.CfnGroupPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnGroupPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_resourcegroups.mixins.CfnGroupPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnGroupPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnGroupPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

