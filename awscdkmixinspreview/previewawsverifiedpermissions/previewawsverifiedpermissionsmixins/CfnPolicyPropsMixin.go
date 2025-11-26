package previewawsverifiedpermissionsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsverifiedpermissions/previewawsverifiedpermissionsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates or updates a Cedar policy and saves it in the specified policy store.
//
// You can create either a static policy or a policy linked to a policy template.
//
// You can directly update only static policies. To update a template-linked policy, you must update its linked policy template instead.
//
// - To create a static policy, in the `Definition` include a `Static` element that includes the Cedar policy text in the `Statement` element.
// - To create a policy that is dynamically linked to a policy template, in the `Definition` include a `Templatelinked` element that specifies the policy template ID and the principal and resource to associate with this policy. If the policy template is ever updated, any policies linked to the policy template automatically use the updated template.
//
// > - If policy validation is enabled in the policy store, then updating a static policy causes Verified Permissions to validate the policy against the schema in the policy store. If the updated static policy doesn't pass validation, the operation fails and the update isn't stored.
// > - When you edit a static policy, You can change only certain elements of a static policy:
// >
// > - The action referenced by the policy.
// > - A condition clause, such as when and unless.
// >
// > You can't change these elements of a static policy:
// >
// > - Changing a policy from a static policy to a template-linked policy.
// > - Changing the effect of a static policy from permit or forbid.
// > - The principal referenced by a static policy.
// > - The resource referenced by a static policy.
// > - To update a template-linked policy, you must update the template instead.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnPolicyPropsMixin := awscdkmixinspreview.Mixins.NewCfnPolicyPropsMixin(&CfnPolicyMixinProps{
//   	Definition: &PolicyDefinitionProperty{
//   		Static: &StaticPolicyDefinitionProperty{
//   			Description: jsii.String("description"),
//   			Statement: jsii.String("statement"),
//   		},
//   		TemplateLinked: &TemplateLinkedPolicyDefinitionProperty{
//   			PolicyTemplateId: jsii.String("policyTemplateId"),
//   			Principal: &EntityIdentifierProperty{
//   				EntityId: jsii.String("entityId"),
//   				EntityType: jsii.String("entityType"),
//   			},
//   			Resource: &EntityIdentifierProperty{
//   				EntityId: jsii.String("entityId"),
//   				EntityType: jsii.String("entityType"),
//   			},
//   		},
//   	},
//   	PolicyStoreId: jsii.String("policyStoreId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-verifiedpermissions-policy.html
//
type CfnPolicyPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnPolicyMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnPolicyPropsMixin
type jsiiProxy_CfnPolicyPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnPolicyPropsMixin) Props() *CfnPolicyMixinProps {
	var returns *CfnPolicyMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPolicyPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::VerifiedPermissions::Policy`.
func NewCfnPolicyPropsMixin(props *CfnPolicyMixinProps, options *mixins.CfnPropertyMixinOptions) CfnPolicyPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnPolicyPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnPolicyPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_verifiedpermissions.mixins.CfnPolicyPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::VerifiedPermissions::Policy`.
func NewCfnPolicyPropsMixin_Override(c CfnPolicyPropsMixin, props *CfnPolicyMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_verifiedpermissions.mixins.CfnPolicyPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnPolicyPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnPolicyPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_verifiedpermissions.mixins.CfnPolicyPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPolicyPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_verifiedpermissions.mixins.CfnPolicyPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPolicyPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnPolicyPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

