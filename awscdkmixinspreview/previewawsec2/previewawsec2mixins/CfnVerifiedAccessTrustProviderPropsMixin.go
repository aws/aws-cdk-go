package previewawsec2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsec2/previewawsec2mixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A trust provider is a third-party entity that creates, maintains, and manages identity information for users and devices.
//
// When an application request is made, the identity information sent by the trust provider is evaluated by Verified Access before allowing or denying the application request.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnVerifiedAccessTrustProviderPropsMixin := awscdkmixinspreview.Mixins.NewCfnVerifiedAccessTrustProviderPropsMixin(&CfnVerifiedAccessTrustProviderMixinProps{
//   	Description: jsii.String("description"),
//   	DeviceOptions: &DeviceOptionsProperty{
//   		PublicSigningKeyUrl: jsii.String("publicSigningKeyUrl"),
//   		TenantId: jsii.String("tenantId"),
//   	},
//   	DeviceTrustProviderType: jsii.String("deviceTrustProviderType"),
//   	NativeApplicationOidcOptions: &NativeApplicationOidcOptionsProperty{
//   		AuthorizationEndpoint: jsii.String("authorizationEndpoint"),
//   		ClientId: jsii.String("clientId"),
//   		ClientSecret: jsii.String("clientSecret"),
//   		Issuer: jsii.String("issuer"),
//   		PublicSigningKeyEndpoint: jsii.String("publicSigningKeyEndpoint"),
//   		Scope: jsii.String("scope"),
//   		TokenEndpoint: jsii.String("tokenEndpoint"),
//   		UserInfoEndpoint: jsii.String("userInfoEndpoint"),
//   	},
//   	OidcOptions: &OidcOptionsProperty{
//   		AuthorizationEndpoint: jsii.String("authorizationEndpoint"),
//   		ClientId: jsii.String("clientId"),
//   		ClientSecret: jsii.String("clientSecret"),
//   		Issuer: jsii.String("issuer"),
//   		Scope: jsii.String("scope"),
//   		TokenEndpoint: jsii.String("tokenEndpoint"),
//   		UserInfoEndpoint: jsii.String("userInfoEndpoint"),
//   	},
//   	PolicyReferenceName: jsii.String("policyReferenceName"),
//   	SseSpecification: &SseSpecificationProperty{
//   		CustomerManagedKeyEnabled: jsii.Boolean(false),
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TrustProviderType: jsii.String("trustProviderType"),
//   	UserTrustProviderType: jsii.String("userTrustProviderType"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-verifiedaccesstrustprovider.html
//
type CfnVerifiedAccessTrustProviderPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnVerifiedAccessTrustProviderMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnVerifiedAccessTrustProviderPropsMixin
type jsiiProxy_CfnVerifiedAccessTrustProviderPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnVerifiedAccessTrustProviderPropsMixin) Props() *CfnVerifiedAccessTrustProviderMixinProps {
	var returns *CfnVerifiedAccessTrustProviderMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVerifiedAccessTrustProviderPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EC2::VerifiedAccessTrustProvider`.
func NewCfnVerifiedAccessTrustProviderPropsMixin(props *CfnVerifiedAccessTrustProviderMixinProps, options *mixins.CfnPropertyMixinOptions) CfnVerifiedAccessTrustProviderPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnVerifiedAccessTrustProviderPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnVerifiedAccessTrustProviderPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVerifiedAccessTrustProviderPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EC2::VerifiedAccessTrustProvider`.
func NewCfnVerifiedAccessTrustProviderPropsMixin_Override(c CfnVerifiedAccessTrustProviderPropsMixin, props *CfnVerifiedAccessTrustProviderMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVerifiedAccessTrustProviderPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnVerifiedAccessTrustProviderPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnVerifiedAccessTrustProviderPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVerifiedAccessTrustProviderPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnVerifiedAccessTrustProviderPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVerifiedAccessTrustProviderPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnVerifiedAccessTrustProviderPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnVerifiedAccessTrustProviderPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

