package awsbedrockagentcore

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsbedrockagentcore/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::BedrockAgentCore::OAuth2CredentialProvider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnOAuth2CredentialProviderPropsMixin := awscdkcfnpropertymixins.Aws_bedrockagentcore.NewCfnOAuth2CredentialProviderPropsMixin(&CfnOAuth2CredentialProviderMixinProps{
//   	CredentialProviderVendor: jsii.String("credentialProviderVendor"),
//   	Name: jsii.String("name"),
//   	Oauth2ProviderConfigInput: &Oauth2ProviderConfigInputProperty{
//   		AtlassianOauth2ProviderConfig: &AtlassianOauth2ProviderConfigInputProperty{
//   			ClientId: jsii.String("clientId"),
//   			ClientSecret: jsii.String("clientSecret"),
//   			ClientSecretConfig: &SecretReferenceProperty{
//   				JsonKey: jsii.String("jsonKey"),
//   				SecretId: jsii.String("secretId"),
//   			},
//   			ClientSecretSource: jsii.String("clientSecretSource"),
//   		},
//   		CustomOauth2ProviderConfig: &CustomOauth2ProviderConfigInputProperty{
//   			ClientAuthenticationMethod: jsii.String("clientAuthenticationMethod"),
//   			ClientId: jsii.String("clientId"),
//   			ClientSecret: jsii.String("clientSecret"),
//   			ClientSecretConfig: &SecretReferenceProperty{
//   				JsonKey: jsii.String("jsonKey"),
//   				SecretId: jsii.String("secretId"),
//   			},
//   			ClientSecretSource: jsii.String("clientSecretSource"),
//   			OauthDiscovery: &Oauth2DiscoveryProperty{
//   				AuthorizationServerMetadata: &Oauth2AuthorizationServerMetadataProperty{
//   					AuthorizationEndpoint: jsii.String("authorizationEndpoint"),
//   					Issuer: jsii.String("issuer"),
//   					ResponseTypes: []*string{
//   						jsii.String("responseTypes"),
//   					},
//   					TokenEndpoint: jsii.String("tokenEndpoint"),
//   				},
//   				DiscoveryUrl: jsii.String("discoveryUrl"),
//   			},
//   			OnBehalfOfTokenExchangeConfig: &OnBehalfOfTokenExchangeConfigProperty{
//   				GrantType: jsii.String("grantType"),
//   				TokenExchangeGrantTypeConfig: &TokenExchangeGrantTypeConfigProperty{
//   					ActorTokenContent: jsii.String("actorTokenContent"),
//   					ActorTokenScopes: []*string{
//   						jsii.String("actorTokenScopes"),
//   					},
//   				},
//   			},
//   		},
//   		GithubOauth2ProviderConfig: &GithubOauth2ProviderConfigInputProperty{
//   			ClientId: jsii.String("clientId"),
//   			ClientSecret: jsii.String("clientSecret"),
//   			ClientSecretConfig: &SecretReferenceProperty{
//   				JsonKey: jsii.String("jsonKey"),
//   				SecretId: jsii.String("secretId"),
//   			},
//   			ClientSecretSource: jsii.String("clientSecretSource"),
//   		},
//   		GoogleOauth2ProviderConfig: &GoogleOauth2ProviderConfigInputProperty{
//   			ClientId: jsii.String("clientId"),
//   			ClientSecret: jsii.String("clientSecret"),
//   			ClientSecretConfig: &SecretReferenceProperty{
//   				JsonKey: jsii.String("jsonKey"),
//   				SecretId: jsii.String("secretId"),
//   			},
//   			ClientSecretSource: jsii.String("clientSecretSource"),
//   		},
//   		IncludedOauth2ProviderConfig: &IncludedOauth2ProviderConfigInputProperty{
//   			AuthorizationEndpoint: jsii.String("authorizationEndpoint"),
//   			ClientId: jsii.String("clientId"),
//   			ClientSecret: jsii.String("clientSecret"),
//   			ClientSecretConfig: &SecretReferenceProperty{
//   				JsonKey: jsii.String("jsonKey"),
//   				SecretId: jsii.String("secretId"),
//   			},
//   			ClientSecretSource: jsii.String("clientSecretSource"),
//   			Issuer: jsii.String("issuer"),
//   			TokenEndpoint: jsii.String("tokenEndpoint"),
//   		},
//   		LinkedinOauth2ProviderConfig: &LinkedinOauth2ProviderConfigInputProperty{
//   			ClientId: jsii.String("clientId"),
//   			ClientSecret: jsii.String("clientSecret"),
//   			ClientSecretConfig: &SecretReferenceProperty{
//   				JsonKey: jsii.String("jsonKey"),
//   				SecretId: jsii.String("secretId"),
//   			},
//   			ClientSecretSource: jsii.String("clientSecretSource"),
//   		},
//   		MicrosoftOauth2ProviderConfig: &MicrosoftOauth2ProviderConfigInputProperty{
//   			ClientId: jsii.String("clientId"),
//   			ClientSecret: jsii.String("clientSecret"),
//   			ClientSecretConfig: &SecretReferenceProperty{
//   				JsonKey: jsii.String("jsonKey"),
//   				SecretId: jsii.String("secretId"),
//   			},
//   			ClientSecretSource: jsii.String("clientSecretSource"),
//   			TenantId: jsii.String("tenantId"),
//   		},
//   		SalesforceOauth2ProviderConfig: &SalesforceOauth2ProviderConfigInputProperty{
//   			ClientId: jsii.String("clientId"),
//   			ClientSecret: jsii.String("clientSecret"),
//   			ClientSecretConfig: &SecretReferenceProperty{
//   				JsonKey: jsii.String("jsonKey"),
//   				SecretId: jsii.String("secretId"),
//   			},
//   			ClientSecretSource: jsii.String("clientSecretSource"),
//   		},
//   		SlackOauth2ProviderConfig: &SlackOauth2ProviderConfigInputProperty{
//   			ClientId: jsii.String("clientId"),
//   			ClientSecret: jsii.String("clientSecret"),
//   			ClientSecretConfig: &SecretReferenceProperty{
//   				JsonKey: jsii.String("jsonKey"),
//   				SecretId: jsii.String("secretId"),
//   			},
//   			ClientSecretSource: jsii.String("clientSecretSource"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-oauth2credentialprovider.html
//
type CfnOAuth2CredentialProviderPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnOAuth2CredentialProviderMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnOAuth2CredentialProviderPropsMixin
type jsiiProxy_CfnOAuth2CredentialProviderPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnOAuth2CredentialProviderPropsMixin) Props() *CfnOAuth2CredentialProviderMixinProps {
	var returns *CfnOAuth2CredentialProviderMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOAuth2CredentialProviderPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::BedrockAgentCore::OAuth2CredentialProvider`.
func NewCfnOAuth2CredentialProviderPropsMixin(props *CfnOAuth2CredentialProviderMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnOAuth2CredentialProviderPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnOAuth2CredentialProviderPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnOAuth2CredentialProviderPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnOAuth2CredentialProviderPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::BedrockAgentCore::OAuth2CredentialProvider`.
func NewCfnOAuth2CredentialProviderPropsMixin_Override(c CfnOAuth2CredentialProviderPropsMixin, props *CfnOAuth2CredentialProviderMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnOAuth2CredentialProviderPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnOAuth2CredentialProviderPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnOAuth2CredentialProviderPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnOAuth2CredentialProviderPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnOAuth2CredentialProviderPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnOAuth2CredentialProviderPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnOAuth2CredentialProviderPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnOAuth2CredentialProviderPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

