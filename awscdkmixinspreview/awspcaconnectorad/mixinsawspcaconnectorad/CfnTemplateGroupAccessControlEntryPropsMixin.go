package mixinsawspcaconnectorad

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awspcaconnectorad/mixinsawspcaconnectorad/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// Create a group access control entry.
//
// Allow or deny Active Directory groups from enrolling and/or autoenrolling with the template based on the group security identifiers (SIDs).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnTemplateGroupAccessControlEntryPropsMixin := awscdkmixinspreview.Mixins.NewCfnTemplateGroupAccessControlEntryPropsMixin(&CfnTemplateGroupAccessControlEntryMixinProps{
//   	AccessRights: &AccessRightsProperty{
//   		AutoEnroll: jsii.String("autoEnroll"),
//   		Enroll: jsii.String("enroll"),
//   	},
//   	GroupDisplayName: jsii.String("groupDisplayName"),
//   	GroupSecurityIdentifier: jsii.String("groupSecurityIdentifier"),
//   	TemplateArn: jsii.String("templateArn"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcaconnectorad-templategroupaccesscontrolentry.html
//
type CfnTemplateGroupAccessControlEntryPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnTemplateGroupAccessControlEntryMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnTemplateGroupAccessControlEntryPropsMixin
type jsiiProxy_CfnTemplateGroupAccessControlEntryPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
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

func (j *jsiiProxy_CfnTemplateGroupAccessControlEntryPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::PCAConnectorAD::TemplateGroupAccessControlEntry`.
func NewCfnTemplateGroupAccessControlEntryPropsMixin(props *CfnTemplateGroupAccessControlEntryMixinProps, options *mixins.CfnPropertyMixinOptions) CfnTemplateGroupAccessControlEntryPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnTemplateGroupAccessControlEntryPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnTemplateGroupAccessControlEntryPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_pcaconnectorad.mixins.CfnTemplateGroupAccessControlEntryPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::PCAConnectorAD::TemplateGroupAccessControlEntry`.
func NewCfnTemplateGroupAccessControlEntryPropsMixin_Override(c CfnTemplateGroupAccessControlEntryPropsMixin, props *CfnTemplateGroupAccessControlEntryMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_pcaconnectorad.mixins.CfnTemplateGroupAccessControlEntryPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnTemplateGroupAccessControlEntryPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTemplateGroupAccessControlEntryPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_pcaconnectorad.mixins.CfnTemplateGroupAccessControlEntryPropsMixin",
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
		"@aws-cdk/mixins-preview.aws_pcaconnectorad.mixins.CfnTemplateGroupAccessControlEntryPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTemplateGroupAccessControlEntryPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

