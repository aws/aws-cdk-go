package awscognito

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awscognito/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::Cognito::IdentityPoolRoleAttachment` resource manages the role configuration for an Amazon Cognito identity pool.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//   var roles interface{}
//
//   cfnIdentityPoolRoleAttachmentPropsMixin := awscdkcfnpropertymixins.Aws_cognito.NewCfnIdentityPoolRoleAttachmentPropsMixin(&CfnIdentityPoolRoleAttachmentMixinProps{
//   	IdentityPoolId: jsii.String("identityPoolId"),
//   	RoleMappings: map[string]interface{}{
//   		"roleMappingsKey": &RoleMappingProperty{
//   			"ambiguousRoleResolution": jsii.String("ambiguousRoleResolution"),
//   			"identityProvider": jsii.String("identityProvider"),
//   			"rulesConfiguration": &RulesConfigurationTypeProperty{
//   				"rules": []interface{}{
//   					&MappingRuleProperty{
//   						"claim": jsii.String("claim"),
//   						"matchType": jsii.String("matchType"),
//   						"roleArn": jsii.String("roleArn"),
//   						"value": jsii.String("value"),
//   					},
//   				},
//   			},
//   			"type": jsii.String("type"),
//   		},
//   	},
//   	Roles: roles,
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypoolroleattachment.html
//
type CfnIdentityPoolRoleAttachmentPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnIdentityPoolRoleAttachmentMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnIdentityPoolRoleAttachmentPropsMixin
type jsiiProxy_CfnIdentityPoolRoleAttachmentPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnIdentityPoolRoleAttachmentPropsMixin) Props() *CfnIdentityPoolRoleAttachmentMixinProps {
	var returns *CfnIdentityPoolRoleAttachmentMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentityPoolRoleAttachmentPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Cognito::IdentityPoolRoleAttachment`.
func NewCfnIdentityPoolRoleAttachmentPropsMixin(props *CfnIdentityPoolRoleAttachmentMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnIdentityPoolRoleAttachmentPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnIdentityPoolRoleAttachmentPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnIdentityPoolRoleAttachmentPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnIdentityPoolRoleAttachmentPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Cognito::IdentityPoolRoleAttachment`.
func NewCfnIdentityPoolRoleAttachmentPropsMixin_Override(c CfnIdentityPoolRoleAttachmentPropsMixin, props *CfnIdentityPoolRoleAttachmentMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnIdentityPoolRoleAttachmentPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnIdentityPoolRoleAttachmentPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnIdentityPoolRoleAttachmentPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnIdentityPoolRoleAttachmentPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnIdentityPoolRoleAttachmentPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnIdentityPoolRoleAttachmentPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnIdentityPoolRoleAttachmentPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnIdentityPoolRoleAttachmentPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

