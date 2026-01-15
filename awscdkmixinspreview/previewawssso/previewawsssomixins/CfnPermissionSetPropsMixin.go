package previewawsssomixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawssso/previewawsssomixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies a permission set within a specified  instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var inlinePolicy interface{}
//
//   cfnPermissionSetPropsMixin := awscdkmixinspreview.Mixins.NewCfnPermissionSetPropsMixin(&CfnPermissionSetMixinProps{
//   	CustomerManagedPolicyReferences: []interface{}{
//   		&CustomerManagedPolicyReferenceProperty{
//   			Name: jsii.String("name"),
//   			Path: jsii.String("path"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	InlinePolicy: inlinePolicy,
//   	InstanceArn: jsii.String("instanceArn"),
//   	ManagedPolicies: []*string{
//   		jsii.String("managedPolicies"),
//   	},
//   	Name: jsii.String("name"),
//   	PermissionsBoundary: &PermissionsBoundaryProperty{
//   		CustomerManagedPolicyReference: &CustomerManagedPolicyReferenceProperty{
//   			Name: jsii.String("name"),
//   			Path: jsii.String("path"),
//   		},
//   		ManagedPolicyArn: jsii.String("managedPolicyArn"),
//   	},
//   	RelayStateType: jsii.String("relayStateType"),
//   	SessionDuration: jsii.String("sessionDuration"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-permissionset.html
//
type CfnPermissionSetPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnPermissionSetMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnPermissionSetPropsMixin
type jsiiProxy_CfnPermissionSetPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnPermissionSetPropsMixin) Props() *CfnPermissionSetMixinProps {
	var returns *CfnPermissionSetMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPermissionSetPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SSO::PermissionSet`.
func NewCfnPermissionSetPropsMixin(props *CfnPermissionSetMixinProps, options *mixins.CfnPropertyMixinOptions) CfnPermissionSetPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnPermissionSetPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnPermissionSetPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sso.mixins.CfnPermissionSetPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SSO::PermissionSet`.
func NewCfnPermissionSetPropsMixin_Override(c CfnPermissionSetPropsMixin, props *CfnPermissionSetMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sso.mixins.CfnPermissionSetPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnPermissionSetPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnPermissionSetPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_sso.mixins.CfnPermissionSetPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPermissionSetPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_sso.mixins.CfnPermissionSetPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPermissionSetPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnPermissionSetPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

