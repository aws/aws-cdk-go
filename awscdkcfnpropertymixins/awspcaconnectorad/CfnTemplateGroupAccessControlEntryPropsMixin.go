package awspcaconnectorad

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awspcaconnectorad/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Create a group access control entry.
//
// Allow or deny Active Directory groups from enrolling and/or autoenrolling with the template based on the group security identifiers (SIDs).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnTemplateGroupAccessControlEntryPropsMixin := awscdkcfnpropertymixins.Aws_pcaconnectorad.NewCfnTemplateGroupAccessControlEntryPropsMixin(&CfnTemplateGroupAccessControlEntryMixinProps{
//   	AccessRights: &AccessRightsProperty{
//   		AutoEnroll: jsii.String("autoEnroll"),
//   		Enroll: jsii.String("enroll"),
//   	},
//   	GroupDisplayName: jsii.String("groupDisplayName"),
//   	GroupSecurityIdentifier: jsii.String("groupSecurityIdentifier"),
//   	TemplateArn: jsii.String("templateArn"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcaconnectorad-templategroupaccesscontrolentry.html
//
type CfnTemplateGroupAccessControlEntryPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnTemplateGroupAccessControlEntryMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnTemplateGroupAccessControlEntryPropsMixin
type jsiiProxy_CfnTemplateGroupAccessControlEntryPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnTemplateGroupAccessControlEntryPropsMixin) Props() *CfnTemplateGroupAccessControlEntryMixinProps {
	var returns *CfnTemplateGroupAccessControlEntryMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTemplateGroupAccessControlEntryPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::PCAConnectorAD::TemplateGroupAccessControlEntry`.
func NewCfnTemplateGroupAccessControlEntryPropsMixin(props *CfnTemplateGroupAccessControlEntryMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnTemplateGroupAccessControlEntryPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnTemplateGroupAccessControlEntryPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnTemplateGroupAccessControlEntryPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_pcaconnectorad.CfnTemplateGroupAccessControlEntryPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::PCAConnectorAD::TemplateGroupAccessControlEntry`.
func NewCfnTemplateGroupAccessControlEntryPropsMixin_Override(c CfnTemplateGroupAccessControlEntryPropsMixin, props *CfnTemplateGroupAccessControlEntryMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_pcaconnectorad.CfnTemplateGroupAccessControlEntryPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnTemplateGroupAccessControlEntryPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTemplateGroupAccessControlEntryPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_pcaconnectorad.CfnTemplateGroupAccessControlEntryPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTemplateGroupAccessControlEntryPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_pcaconnectorad.CfnTemplateGroupAccessControlEntryPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTemplateGroupAccessControlEntryPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnTemplateGroupAccessControlEntryPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

