package awsappfabric

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsappfabric/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates an app authorization within an app bundle, which allows AppFabric to connect to an application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnAppAuthorizationPropsMixin := awscdkcfnpropertymixins.Aws_appfabric.NewCfnAppAuthorizationPropsMixin(&CfnAppAuthorizationMixinProps{
//   	App: jsii.String("app"),
//   	AppBundleIdentifier: jsii.String("appBundleIdentifier"),
//   	AuthType: jsii.String("authType"),
//   	Credential: &CredentialProperty{
//   		ApiKeyCredential: &ApiKeyCredentialProperty{
//   			ApiKey: jsii.String("apiKey"),
//   		},
//   		Oauth2Credential: &Oauth2CredentialProperty{
//   			ClientId: jsii.String("clientId"),
//   			ClientSecret: jsii.String("clientSecret"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Tenant: &TenantProperty{
//   		TenantDisplayName: jsii.String("tenantDisplayName"),
//   		TenantIdentifier: jsii.String("tenantIdentifier"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appfabric-appauthorization.html
//
type CfnAppAuthorizationPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnAppAuthorizationMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAppAuthorizationPropsMixin
type jsiiProxy_CfnAppAuthorizationPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnAppAuthorizationPropsMixin) Props() *CfnAppAuthorizationMixinProps {
	var returns *CfnAppAuthorizationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAppAuthorizationPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::AppFabric::AppAuthorization`.
func NewCfnAppAuthorizationPropsMixin(props *CfnAppAuthorizationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnAppAuthorizationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAppAuthorizationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAppAuthorizationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_appfabric.CfnAppAuthorizationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::AppFabric::AppAuthorization`.
func NewCfnAppAuthorizationPropsMixin_Override(c CfnAppAuthorizationPropsMixin, props *CfnAppAuthorizationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_appfabric.CfnAppAuthorizationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnAppAuthorizationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAppAuthorizationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_appfabric.CfnAppAuthorizationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAppAuthorizationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_appfabric.CfnAppAuthorizationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAppAuthorizationPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnAppAuthorizationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

