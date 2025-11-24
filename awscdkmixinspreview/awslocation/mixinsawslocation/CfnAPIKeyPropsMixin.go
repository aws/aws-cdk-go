package mixinsawslocation

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awslocation/mixinsawslocation/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// The API key resource in your AWS account, which lets you grant actions for Amazon Location resources to the API key bearer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAPIKeyPropsMixin := awscdkmixinspreview.Mixins.NewCfnAPIKeyPropsMixin(&CfnAPIKeyMixinProps{
//   	Description: jsii.String("description"),
//   	ExpireTime: jsii.String("expireTime"),
//   	ForceDelete: jsii.Boolean(false),
//   	ForceUpdate: jsii.Boolean(false),
//   	KeyName: jsii.String("keyName"),
//   	NoExpiry: jsii.Boolean(false),
//   	Restrictions: &ApiKeyRestrictionsProperty{
//   		AllowActions: []*string{
//   			jsii.String("allowActions"),
//   		},
//   		AllowAndroidApps: []interface{}{
//   			&AndroidAppProperty{
//   				CertificateFingerprint: jsii.String("certificateFingerprint"),
//   				Package: jsii.String("package"),
//   			},
//   		},
//   		AllowAppleApps: []interface{}{
//   			&AppleAppProperty{
//   				BundleId: jsii.String("bundleId"),
//   			},
//   		},
//   		AllowReferers: []*string{
//   			jsii.String("allowReferers"),
//   		},
//   		AllowResources: []*string{
//   			jsii.String("allowResources"),
//   		},
//   	},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-apikey.html
//
type CfnAPIKeyPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnAPIKeyMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAPIKeyPropsMixin
type jsiiProxy_CfnAPIKeyPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnAPIKeyPropsMixin) Props() *CfnAPIKeyMixinProps {
	var returns *CfnAPIKeyMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPIKeyPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Location::APIKey`.
func NewCfnAPIKeyPropsMixin(props *CfnAPIKeyMixinProps, options *mixins.CfnPropertyMixinOptions) CfnAPIKeyPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAPIKeyPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAPIKeyPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_location.mixins.CfnAPIKeyPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Location::APIKey`.
func NewCfnAPIKeyPropsMixin_Override(c CfnAPIKeyPropsMixin, props *CfnAPIKeyMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_location.mixins.CfnAPIKeyPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnAPIKeyPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAPIKeyPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_location.mixins.CfnAPIKeyPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAPIKeyPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_location.mixins.CfnAPIKeyPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAPIKeyPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnAPIKeyPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

