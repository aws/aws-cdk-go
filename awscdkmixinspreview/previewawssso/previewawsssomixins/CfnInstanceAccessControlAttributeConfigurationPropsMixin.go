package previewawsssomixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawssso/previewawsssomixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Enables the attribute-based access control (ABAC) feature for the specified  instance.
//
// You can also specify new attributes to add to your ABAC configuration during the enabling process. For more information about ABAC, see [Attribute-Based Access Control](https://docs.aws.amazon.com//singlesignon/latest/userguide/abac.html) in the *User Guide* .
//
// > The `InstanceAccessControlAttributeConfiguration` property has been deprecated but is still supported for backwards compatibility purposes. We recommend that you use the `AccessControlAttributes` property instead.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnInstanceAccessControlAttributeConfigurationPropsMixin := awscdkmixinspreview.Mixins.NewCfnInstanceAccessControlAttributeConfigurationPropsMixin(&CfnInstanceAccessControlAttributeConfigurationMixinProps{
//   	AccessControlAttributes: []interface{}{
//   		&AccessControlAttributeProperty{
//   			Key: jsii.String("key"),
//   			Value: &AccessControlAttributeValueProperty{
//   				Source: []*string{
//   					jsii.String("source"),
//   				},
//   			},
//   		},
//   	},
//   	InstanceAccessControlAttributeConfiguration: &InstanceAccessControlAttributeConfigurationProperty{
//   		AccessControlAttributes: []interface{}{
//   			&AccessControlAttributeProperty{
//   				Key: jsii.String("key"),
//   				Value: &AccessControlAttributeValueProperty{
//   					Source: []*string{
//   						jsii.String("source"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	InstanceArn: jsii.String("instanceArn"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-instanceaccesscontrolattributeconfiguration.html
//
type CfnInstanceAccessControlAttributeConfigurationPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnInstanceAccessControlAttributeConfigurationMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnInstanceAccessControlAttributeConfigurationPropsMixin
type jsiiProxy_CfnInstanceAccessControlAttributeConfigurationPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnInstanceAccessControlAttributeConfigurationPropsMixin) Props() *CfnInstanceAccessControlAttributeConfigurationMixinProps {
	var returns *CfnInstanceAccessControlAttributeConfigurationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceAccessControlAttributeConfigurationPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SSO::InstanceAccessControlAttributeConfiguration`.
func NewCfnInstanceAccessControlAttributeConfigurationPropsMixin(props *CfnInstanceAccessControlAttributeConfigurationMixinProps, options *mixins.CfnPropertyMixinOptions) CfnInstanceAccessControlAttributeConfigurationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnInstanceAccessControlAttributeConfigurationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnInstanceAccessControlAttributeConfigurationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sso.mixins.CfnInstanceAccessControlAttributeConfigurationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SSO::InstanceAccessControlAttributeConfiguration`.
func NewCfnInstanceAccessControlAttributeConfigurationPropsMixin_Override(c CfnInstanceAccessControlAttributeConfigurationPropsMixin, props *CfnInstanceAccessControlAttributeConfigurationMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sso.mixins.CfnInstanceAccessControlAttributeConfigurationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnInstanceAccessControlAttributeConfigurationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnInstanceAccessControlAttributeConfigurationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_sso.mixins.CfnInstanceAccessControlAttributeConfigurationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnInstanceAccessControlAttributeConfigurationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_sso.mixins.CfnInstanceAccessControlAttributeConfigurationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnInstanceAccessControlAttributeConfigurationPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnInstanceAccessControlAttributeConfigurationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

