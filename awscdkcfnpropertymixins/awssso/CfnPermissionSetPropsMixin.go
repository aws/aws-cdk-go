package awssso

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awssso/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies a permission set within a specified  instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var inlinePolicy interface{}
//   var mergeStrategy IMergeStrategy
//
//   cfnPermissionSetPropsMixin := awscdkcfnpropertymixins.Aws_sso.NewCfnPermissionSetPropsMixin(&CfnPermissionSetMixinProps{
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
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-permissionset.html
//
type CfnPermissionSetPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnPermissionSetMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnPermissionSetPropsMixin
type jsiiProxy_CfnPermissionSetPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
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

func (j *jsiiProxy_CfnPermissionSetPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SSO::PermissionSet`.
func NewCfnPermissionSetPropsMixin(props *CfnPermissionSetMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnPermissionSetPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnPermissionSetPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnPermissionSetPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_sso.CfnPermissionSetPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SSO::PermissionSet`.
func NewCfnPermissionSetPropsMixin_Override(c CfnPermissionSetPropsMixin, props *CfnPermissionSetMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_sso.CfnPermissionSetPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnPermissionSetPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnPermissionSetPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_sso.CfnPermissionSetPropsMixin",
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
		"@aws-cdk/cfn-property-mixins.aws_sso.CfnPermissionSetPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPermissionSetPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
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

