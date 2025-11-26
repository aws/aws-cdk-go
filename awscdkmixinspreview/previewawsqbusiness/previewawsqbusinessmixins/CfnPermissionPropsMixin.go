package previewawsqbusinessmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsqbusiness/previewawsqbusinessmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Adds or updates a permission policy for a Amazon Q Business application, allowing cross-account access for an ISV.
//
// This operation creates a new policy statement for the specified Amazon Q Business application. The policy statement defines the IAM actions that the ISV is allowed to perform on the Amazon Q Business application's resources.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnPermissionPropsMixin := awscdkmixinspreview.Mixins.NewCfnPermissionPropsMixin(&CfnPermissionMixinProps{
//   	Actions: []*string{
//   		jsii.String("actions"),
//   	},
//   	ApplicationId: jsii.String("applicationId"),
//   	Conditions: []interface{}{
//   		&ConditionProperty{
//   			ConditionKey: jsii.String("conditionKey"),
//   			ConditionOperator: jsii.String("conditionOperator"),
//   			ConditionValues: []*string{
//   				jsii.String("conditionValues"),
//   			},
//   		},
//   	},
//   	Principal: jsii.String("principal"),
//   	StatementId: jsii.String("statementId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-permission.html
//
type CfnPermissionPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnPermissionMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnPermissionPropsMixin
type jsiiProxy_CfnPermissionPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnPermissionPropsMixin) Props() *CfnPermissionMixinProps {
	var returns *CfnPermissionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPermissionPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::QBusiness::Permission`.
func NewCfnPermissionPropsMixin(props *CfnPermissionMixinProps, options *mixins.CfnPropertyMixinOptions) CfnPermissionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnPermissionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnPermissionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnPermissionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::QBusiness::Permission`.
func NewCfnPermissionPropsMixin_Override(c CfnPermissionPropsMixin, props *CfnPermissionMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnPermissionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnPermissionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnPermissionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnPermissionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPermissionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnPermissionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPermissionPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnPermissionPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

