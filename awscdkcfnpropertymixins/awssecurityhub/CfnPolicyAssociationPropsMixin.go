package awssecurityhub

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awssecurityhub/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::SecurityHub::PolicyAssociation` resource specifies associations for a configuration policy or a self-managed configuration.
//
// You can associate a AWS Security Hub CSPM configuration policy or self-managed configuration with the organization root, organizational units (OUs), or AWS accounts . After a successful association, the configuration policy takes effect in the specified targets. For more information, see [Creating and associating Security Hub CSPM configuration policies](https://docs.aws.amazon.com/securityhub/latest/userguide/create-associate-policy.html) in the *AWS Security Hub CSPM User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnPolicyAssociationPropsMixin := awscdkcfnpropertymixins.Aws_securityhub.NewCfnPolicyAssociationPropsMixin(&CfnPolicyAssociationMixinProps{
//   	ConfigurationPolicyId: jsii.String("configurationPolicyId"),
//   	TargetId: jsii.String("targetId"),
//   	TargetType: jsii.String("targetType"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-policyassociation.html
//
type CfnPolicyAssociationPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnPolicyAssociationMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnPolicyAssociationPropsMixin
type jsiiProxy_CfnPolicyAssociationPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnPolicyAssociationPropsMixin) Props() *CfnPolicyAssociationMixinProps {
	var returns *CfnPolicyAssociationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPolicyAssociationPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SecurityHub::PolicyAssociation`.
func NewCfnPolicyAssociationPropsMixin(props *CfnPolicyAssociationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnPolicyAssociationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnPolicyAssociationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnPolicyAssociationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnPolicyAssociationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SecurityHub::PolicyAssociation`.
func NewCfnPolicyAssociationPropsMixin_Override(c CfnPolicyAssociationPropsMixin, props *CfnPolicyAssociationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnPolicyAssociationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnPolicyAssociationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnPolicyAssociationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnPolicyAssociationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPolicyAssociationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnPolicyAssociationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPolicyAssociationPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnPolicyAssociationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

