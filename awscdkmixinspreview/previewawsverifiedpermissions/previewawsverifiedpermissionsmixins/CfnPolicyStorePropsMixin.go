package previewawsverifiedpermissionsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsverifiedpermissions/previewawsverifiedpermissionsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a policy store.
//
// A policy store is a container for policy resources. You can create a separate policy store for each of your applications.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnPolicyStorePropsMixin := awscdkmixinspreview.Mixins.NewCfnPolicyStorePropsMixin(&CfnPolicyStoreMixinProps{
//   	DeletionProtection: &DeletionProtectionProperty{
//   		Mode: jsii.String("mode"),
//   	},
//   	Description: jsii.String("description"),
//   	Schema: &SchemaDefinitionProperty{
//   		CedarFormat: jsii.String("cedarFormat"),
//   		CedarJson: jsii.String("cedarJson"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ValidationSettings: &ValidationSettingsProperty{
//   		Mode: jsii.String("mode"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-verifiedpermissions-policystore.html
//
type CfnPolicyStorePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnPolicyStoreMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnPolicyStorePropsMixin
type jsiiProxy_CfnPolicyStorePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnPolicyStorePropsMixin) Props() *CfnPolicyStoreMixinProps {
	var returns *CfnPolicyStoreMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPolicyStorePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::VerifiedPermissions::PolicyStore`.
func NewCfnPolicyStorePropsMixin(props *CfnPolicyStoreMixinProps, options *mixins.CfnPropertyMixinOptions) CfnPolicyStorePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnPolicyStorePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnPolicyStorePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_verifiedpermissions.mixins.CfnPolicyStorePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::VerifiedPermissions::PolicyStore`.
func NewCfnPolicyStorePropsMixin_Override(c CfnPolicyStorePropsMixin, props *CfnPolicyStoreMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_verifiedpermissions.mixins.CfnPolicyStorePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnPolicyStorePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnPolicyStorePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_verifiedpermissions.mixins.CfnPolicyStorePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPolicyStorePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_verifiedpermissions.mixins.CfnPolicyStorePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPolicyStorePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnPolicyStorePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

