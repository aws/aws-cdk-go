package previewawscognitomixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawscognito/previewawscognitomixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a new set of branding settings for a user pool style and associates it with an app client.
//
// This operation is the programmatic option for the creation of a new style in the branding designer.
//
// Provides values for UI customization in a `Settings` JSON object and image files in an `Assets` array. To send the JSON object `Document` type parameter in `Settings` , you might need to update to the most recent version of your AWS SDK.
//
// This operation has a 2-megabyte request-size limit and include the CSS settings and image assets for your app client. Your branding settings might exceed 2MB in size. Amazon Cognito doesn't require that you pass all parameters in one request and preserves existing style settings that you don't specify. If your request is larger than 2MB, separate it into multiple requests, each with a size smaller than the limit.
//
// As a best practice, modify the output of [DescribeManagedLoginBrandingByClient](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_DescribeManagedLoginBrandingByClient.html) into the request parameters for this operation. To get all settings, set `ReturnMergedResources` to `true` . For more information, see [API and SDK operations for managed login branding](https://docs.aws.amazon.com/cognito/latest/developerguide/managed-login-brandingdesigner.html#branding-designer-api)
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var settings interface{}
//
//   cfnManagedLoginBrandingPropsMixin := awscdkmixinspreview.Mixins.NewCfnManagedLoginBrandingPropsMixin(&CfnManagedLoginBrandingMixinProps{
//   	Assets: []interface{}{
//   		&AssetTypeProperty{
//   			Bytes: jsii.String("bytes"),
//   			Category: jsii.String("category"),
//   			ColorMode: jsii.String("colorMode"),
//   			Extension: jsii.String("extension"),
//   			ResourceId: jsii.String("resourceId"),
//   		},
//   	},
//   	ClientId: jsii.String("clientId"),
//   	ReturnMergedResources: jsii.Boolean(false),
//   	Settings: settings,
//   	UseCognitoProvidedValues: jsii.Boolean(false),
//   	UserPoolId: jsii.String("userPoolId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-managedloginbranding.html
//
type CfnManagedLoginBrandingPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnManagedLoginBrandingMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnManagedLoginBrandingPropsMixin
type jsiiProxy_CfnManagedLoginBrandingPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnManagedLoginBrandingPropsMixin) Props() *CfnManagedLoginBrandingMixinProps {
	var returns *CfnManagedLoginBrandingMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnManagedLoginBrandingPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Cognito::ManagedLoginBranding`.
func NewCfnManagedLoginBrandingPropsMixin(props *CfnManagedLoginBrandingMixinProps, options *mixins.CfnPropertyMixinOptions) CfnManagedLoginBrandingPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnManagedLoginBrandingPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnManagedLoginBrandingPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnManagedLoginBrandingPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Cognito::ManagedLoginBranding`.
func NewCfnManagedLoginBrandingPropsMixin_Override(c CfnManagedLoginBrandingPropsMixin, props *CfnManagedLoginBrandingMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnManagedLoginBrandingPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnManagedLoginBrandingPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnManagedLoginBrandingPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnManagedLoginBrandingPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnManagedLoginBrandingPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnManagedLoginBrandingPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnManagedLoginBrandingPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnManagedLoginBrandingPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

