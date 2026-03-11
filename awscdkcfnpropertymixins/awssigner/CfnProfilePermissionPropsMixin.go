package awssigner

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awssigner/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Adds cross-account permissions to a signing profile.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnProfilePermissionPropsMixin := awscdkcfnpropertymixins.Aws_signer.NewCfnProfilePermissionPropsMixin(&CfnProfilePermissionMixinProps{
//   	Action: jsii.String("action"),
//   	Principal: jsii.String("principal"),
//   	ProfileName: jsii.String("profileName"),
//   	ProfileVersion: jsii.String("profileVersion"),
//   	StatementId: jsii.String("statementId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-signer-profilepermission.html
//
type CfnProfilePermissionPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnProfilePermissionMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnProfilePermissionPropsMixin
type jsiiProxy_CfnProfilePermissionPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnProfilePermissionPropsMixin) Props() *CfnProfilePermissionMixinProps {
	var returns *CfnProfilePermissionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProfilePermissionPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Signer::ProfilePermission`.
func NewCfnProfilePermissionPropsMixin(props *CfnProfilePermissionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnProfilePermissionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnProfilePermissionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnProfilePermissionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_signer.CfnProfilePermissionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Signer::ProfilePermission`.
func NewCfnProfilePermissionPropsMixin_Override(c CfnProfilePermissionPropsMixin, props *CfnProfilePermissionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_signer.CfnProfilePermissionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnProfilePermissionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnProfilePermissionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_signer.CfnProfilePermissionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnProfilePermissionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_signer.CfnProfilePermissionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnProfilePermissionPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnProfilePermissionPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

