package previewalexaaskmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewalexaask/previewalexaaskmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `Alexa::ASK::Skill` resource creates an Alexa skill that enables customers to access new abilities.
//
// For more information about developing a skill, see the  .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var manifest interface{}
//
//   cfnSkillPropsMixin := awscdkmixinspreview.Mixins.NewCfnSkillPropsMixin(&CfnSkillMixinProps{
//   	AuthenticationConfiguration: &AuthenticationConfigurationProperty{
//   		ClientId: jsii.String("clientId"),
//   		ClientSecret: jsii.String("clientSecret"),
//   		RefreshToken: jsii.String("refreshToken"),
//   	},
//   	SkillPackage: &SkillPackageProperty{
//   		Overrides: &OverridesProperty{
//   			Manifest: manifest,
//   		},
//   		S3Bucket: jsii.String("s3Bucket"),
//   		S3BucketRole: jsii.String("s3BucketRole"),
//   		S3Key: jsii.String("s3Key"),
//   		S3ObjectVersion: jsii.String("s3ObjectVersion"),
//   	},
//   	VendorId: jsii.String("vendorId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ask-skill.html
//
type CfnSkillPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnSkillMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnSkillPropsMixin
type jsiiProxy_CfnSkillPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnSkillPropsMixin) Props() *CfnSkillMixinProps {
	var returns *CfnSkillMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSkillPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `Alexa::ASK::Skill`.
func NewCfnSkillPropsMixin(props *CfnSkillMixinProps, options *mixins.CfnPropertyMixinOptions) CfnSkillPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnSkillPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnSkillPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.alexa_ask.mixins.CfnSkillPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `Alexa::ASK::Skill`.
func NewCfnSkillPropsMixin_Override(c CfnSkillPropsMixin, props *CfnSkillMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.alexa_ask.mixins.CfnSkillPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnSkillPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSkillPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.alexa_ask.mixins.CfnSkillPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSkillPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.alexa_ask.mixins.CfnSkillPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSkillPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnSkillPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

