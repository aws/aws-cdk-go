package awsbedrockagentcorealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awsbedrockagentcorealpha/v2/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsbedrockagentcore"
	"github.com/aws/constructs-go/constructs/v10"
)

// L2 construct for `AWS::BedrockAgentCore::OAuth2CredentialProvider`.
//
// Prefer the static factories (for example {@link OAuth2CredentialProvider.usingSlack}) so you only pass
// the OAuth2 settings that apply to that vendor. To attach the identity to a gateway target, use
// {@link GatewayCredentialProvider.fromOauthIdentity } with this construct, or
// {@link OAuth2CredentialProvider.bindForGatewayOAuthTarget} with {@link GatewayCredentialProvider.fromOauthIdentityArn }.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//
//   oAuth2CredentialProvider := bedrock_agentcore_alpha.NewOAuth2CredentialProvider(this, jsii.String("MyOAuth2CredentialProvider"), &OAuth2CredentialProviderProps{
//   	CredentialProviderVendor: jsii.String("credentialProviderVendor"),
//   	Oauth2ProviderConfigInput: &Oauth2ProviderConfigInputProperty{
//   		AtlassianOauth2ProviderConfig: &AtlassianOauth2ProviderConfigInputProperty{
//   			ClientId: jsii.String("clientId"),
//   			ClientSecret: jsii.String("clientSecret"),
//   		},
//   		CustomOauth2ProviderConfig: &CustomOauth2ProviderConfigInputProperty{
//   			OauthDiscovery: &Oauth2DiscoveryProperty{
//   				AuthorizationServerMetadata: &Oauth2AuthorizationServerMetadataProperty{
//   					AuthorizationEndpoint: jsii.String("authorizationEndpoint"),
//   					Issuer: jsii.String("issuer"),
//   					TokenEndpoint: jsii.String("tokenEndpoint"),
//
//   					// the properties below are optional
//   					ResponseTypes: []*string{
//   						jsii.String("responseTypes"),
//   					},
//   				},
//   				DiscoveryUrl: jsii.String("discoveryUrl"),
//   			},
//
//   			// the properties below are optional
//   			ClientId: jsii.String("clientId"),
//   			ClientSecret: jsii.String("clientSecret"),
//   			OnBehalfOfTokenExchangeConfig: &OnBehalfOfTokenExchangeConfigProperty{
//   				GrantType: jsii.String("grantType"),
//
//   				// the properties below are optional
//   				TokenExchangeGrantTypeConfig: &TokenExchangeGrantTypeConfigProperty{
//   					ActorTokenContent: jsii.String("actorTokenContent"),
//
//   					// the properties below are optional
//   					ActorTokenScopes: []*string{
//   						jsii.String("actorTokenScopes"),
//   					},
//   				},
//   			},
//   		},
//   		GithubOauth2ProviderConfig: &GithubOauth2ProviderConfigInputProperty{
//   			ClientId: jsii.String("clientId"),
//   			ClientSecret: jsii.String("clientSecret"),
//   		},
//   		GoogleOauth2ProviderConfig: &GoogleOauth2ProviderConfigInputProperty{
//   			ClientId: jsii.String("clientId"),
//   			ClientSecret: jsii.String("clientSecret"),
//   		},
//   		IncludedOauth2ProviderConfig: &IncludedOauth2ProviderConfigInputProperty{
//   			ClientId: jsii.String("clientId"),
//   			ClientSecret: jsii.String("clientSecret"),
//
//   			// the properties below are optional
//   			AuthorizationEndpoint: jsii.String("authorizationEndpoint"),
//   			Issuer: jsii.String("issuer"),
//   			TokenEndpoint: jsii.String("tokenEndpoint"),
//   		},
//   		LinkedinOauth2ProviderConfig: &LinkedinOauth2ProviderConfigInputProperty{
//   			ClientId: jsii.String("clientId"),
//   			ClientSecret: jsii.String("clientSecret"),
//   		},
//   		MicrosoftOauth2ProviderConfig: &MicrosoftOauth2ProviderConfigInputProperty{
//   			ClientId: jsii.String("clientId"),
//   			ClientSecret: jsii.String("clientSecret"),
//
//   			// the properties below are optional
//   			TenantId: jsii.String("tenantId"),
//   		},
//   		SalesforceOauth2ProviderConfig: &SalesforceOauth2ProviderConfigInputProperty{
//   			ClientId: jsii.String("clientId"),
//   			ClientSecret: jsii.String("clientSecret"),
//   		},
//   		SlackOauth2ProviderConfig: &SlackOauth2ProviderConfigInputProperty{
//   			ClientId: jsii.String("clientId"),
//   			ClientSecret: jsii.String("clientSecret"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	OAuth2CredentialProviderName: jsii.String("oAuth2CredentialProviderName"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   })
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-oauth2credentialprovider.html
//
// Experimental.
type OAuth2CredentialProvider interface {
	awscdk.Resource
	IOAuth2CredentialProvider
	// Callback URL for the OAuth2 authorization flow.
	// Experimental.
	CallbackUrl() *string
	// The ARN of the Secrets Manager secret for the OAuth2 client credentials.
	//
	// May be undefined for resources imported without this attribute.
	// Experimental.
	ClientSecretArn() *string
	// Timestamp when the credential provider was created.
	// Experimental.
	CreatedTime() *string
	// The ARN of this credential provider.
	// Experimental.
	CredentialProviderArn() *string
	// OAuth2 vendor string passed to CloudFormation.
	// Experimental.
	CredentialProviderVendor() *string
	// The environment this resource belongs to.
	//
	// For resources that are created and managed in a Stack (those created by
	// creating new class instances like `new Role()`, `new Bucket()`, etc.), this
	// is always the same as the environment of the stack they belong to.
	//
	// For referenced resources (those obtained from referencing methods like
	// `Role.fromRoleArn()`, `Bucket.fromBucketName()`, etc.), they might be
	// different than the stack they were imported into.
	// Experimental.
	Env() *interfaces.ResourceEnvironment
	// The principal to grant permissions to.
	// Experimental.
	GrantPrincipal() awsiam.IPrincipal
	// Timestamp when the credential provider was last updated.
	// Experimental.
	LastUpdatedTime() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// The name of this OAuth2 credential provider.
	// Experimental.
	OAuth2CredentialProviderName() *string
	// A reference to a OAuth2CredentialProvider resource.
	// Experimental.
	OAuth2CredentialProviderRef() *interfacesawsbedrockagentcore.OAuth2CredentialProviderReference
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//   cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// Override the cross-stack reference strength for this resource.
	//
	// When set, any cross-stack reference to this resource will use the specified
	// mechanism instead of the global default determined by the
	// `@aws-cdk/core:defaultCrossStackReferences` context key. This is useful for
	// selectively weakening specific references to avoid the "deadly embrace" problem
	// without changing the app-wide default.
	// Experimental.
	ApplyCrossStackReferenceStrength(strength awscdk.ReferenceStrength)
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// ARNs and OAuth scopes for {@link GatewayCredentialProvider.fromOauthIdentity } / {@link GatewayCredentialProvider.fromOauthIdentityArn }.
	// Experimental.
	BindForGatewayOAuthTarget(scopes *[]*string, customParameters *map[string]*string) *GatewayOAuth2IdentityBinding
	// Experimental.
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	// Experimental.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	// Experimental.
	GetResourceNameAttribute(nameAttr *string) *string
	// [disable-awslint:no-grants].
	// Experimental.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// [disable-awslint:no-grants].
	// Experimental.
	GrantAdmin(grantee awsiam.IGrantable) awsiam.Grant
	// [disable-awslint:no-grants].
	// Experimental.
	GrantFullAccess(grantee awsiam.IGrantable) awsiam.Grant
	// [disable-awslint:no-grants].
	// Experimental.
	GrantRead(grantee awsiam.IGrantable) awsiam.Grant
	// [disable-awslint:no-grants].
	// Experimental.
	GrantUse(grantee awsiam.IGrantable) awsiam.Grant
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	// Experimental.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for OAuth2CredentialProvider
type jsiiProxy_OAuth2CredentialProvider struct {
	internal.Type__awscdkResource
	jsiiProxy_IOAuth2CredentialProvider
}

func (j *jsiiProxy_OAuth2CredentialProvider) CallbackUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"callbackUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OAuth2CredentialProvider) ClientSecretArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OAuth2CredentialProvider) CreatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OAuth2CredentialProvider) CredentialProviderArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialProviderArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OAuth2CredentialProvider) CredentialProviderVendor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialProviderVendor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OAuth2CredentialProvider) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OAuth2CredentialProvider) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OAuth2CredentialProvider) LastUpdatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastUpdatedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OAuth2CredentialProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OAuth2CredentialProvider) OAuth2CredentialProviderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oAuth2CredentialProviderName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OAuth2CredentialProvider) OAuth2CredentialProviderRef() *interfacesawsbedrockagentcore.OAuth2CredentialProviderReference {
	var returns *interfacesawsbedrockagentcore.OAuth2CredentialProviderReference
	_jsii_.Get(
		j,
		"oAuth2CredentialProviderRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OAuth2CredentialProvider) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OAuth2CredentialProvider) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewOAuth2CredentialProvider(scope constructs.Construct, id *string, props *OAuth2CredentialProviderProps) OAuth2CredentialProvider {
	_init_.Initialize()

	if err := validateNewOAuth2CredentialProviderParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_OAuth2CredentialProvider{}

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-agentcore-alpha.OAuth2CredentialProvider",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewOAuth2CredentialProvider_Override(o OAuth2CredentialProvider, scope constructs.Construct, id *string, props *OAuth2CredentialProviderProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-agentcore-alpha.OAuth2CredentialProvider",
		[]interface{}{scope, id, props},
		o,
	)
}

// Import an existing OAuth2 credential provider.
// Experimental.
func OAuth2CredentialProvider_FromOAuth2CredentialProviderAttributes(scope constructs.Construct, id *string, attrs *OAuth2CredentialProviderAttributes) IOAuth2CredentialProvider {
	_init_.Initialize()

	if err := validateOAuth2CredentialProvider_FromOAuth2CredentialProviderAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IOAuth2CredentialProvider

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.OAuth2CredentialProvider",
		"fromOAuth2CredentialProviderAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func OAuth2CredentialProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOAuth2CredentialProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.OAuth2CredentialProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
// Experimental.
func OAuth2CredentialProvider_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateOAuth2CredentialProvider_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.OAuth2CredentialProvider",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func OAuth2CredentialProvider_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateOAuth2CredentialProvider_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.OAuth2CredentialProvider",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Create a credential provider for Atlassian OAuth2.
// Experimental.
func OAuth2CredentialProvider_UsingAtlassian(scope constructs.Construct, id *string, props *AtlassianOAuth2CredentialProviderProps) OAuth2CredentialProvider {
	_init_.Initialize()

	if err := validateOAuth2CredentialProvider_UsingAtlassianParameters(scope, id, props); err != nil {
		panic(err)
	}
	var returns OAuth2CredentialProvider

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.OAuth2CredentialProvider",
		"usingAtlassian",
		[]interface{}{scope, id, props},
		&returns,
	)

	return returns
}

// Create a credential provider for Auth0 OAuth2 (`IncludedOauth2ProviderConfig`).
// See: https://docs.aws.amazon.com/bedrock-agentcore/latest/devguide/identity-idp-auth0.html
//
// Experimental.
func OAuth2CredentialProvider_UsingAuth0(scope constructs.Construct, id *string, props *IncludedOauth2TenantCredentialProviderProps) OAuth2CredentialProvider {
	_init_.Initialize()

	if err := validateOAuth2CredentialProvider_UsingAuth0Parameters(scope, id, props); err != nil {
		panic(err)
	}
	var returns OAuth2CredentialProvider

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.OAuth2CredentialProvider",
		"usingAuth0",
		[]interface{}{scope, id, props},
		&returns,
	)

	return returns
}

// Create a credential provider for Amazon Cognito OAuth2 (`IncludedOauth2ProviderConfig`).
// See: https://docs.aws.amazon.com/bedrock-agentcore/latest/devguide/identity-idp-cognito.html
//
// Experimental.
func OAuth2CredentialProvider_UsingCognito(scope constructs.Construct, id *string, props *IncludedOauth2TenantCredentialProviderProps) OAuth2CredentialProvider {
	_init_.Initialize()

	if err := validateOAuth2CredentialProvider_UsingCognitoParameters(scope, id, props); err != nil {
		panic(err)
	}
	var returns OAuth2CredentialProvider

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.OAuth2CredentialProvider",
		"usingCognito",
		[]interface{}{scope, id, props},
		&returns,
	)

	return returns
}

// Create a credential provider for a custom OAuth2 authorization server (discovery document or metadata).
// Experimental.
func OAuth2CredentialProvider_UsingCustom(scope constructs.Construct, id *string, props *CustomOAuth2CredentialProviderProps) OAuth2CredentialProvider {
	_init_.Initialize()

	if err := validateOAuth2CredentialProvider_UsingCustomParameters(scope, id, props); err != nil {
		panic(err)
	}
	var returns OAuth2CredentialProvider

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.OAuth2CredentialProvider",
		"usingCustom",
		[]interface{}{scope, id, props},
		&returns,
	)

	return returns
}

// Create a credential provider for CyberArk OAuth2 (`IncludedOauth2ProviderConfig`).
// See: https://docs.aws.amazon.com/bedrock-agentcore/latest/devguide/identity-idp-cyberark.html
//
// Experimental.
func OAuth2CredentialProvider_UsingCyberArk(scope constructs.Construct, id *string, props *IncludedOauth2TenantCredentialProviderProps) OAuth2CredentialProvider {
	_init_.Initialize()

	if err := validateOAuth2CredentialProvider_UsingCyberArkParameters(scope, id, props); err != nil {
		panic(err)
	}
	var returns OAuth2CredentialProvider

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.OAuth2CredentialProvider",
		"usingCyberArk",
		[]interface{}{scope, id, props},
		&returns,
	)

	return returns
}

// Create a credential provider for Dropbox OAuth2 (`IncludedOauth2ProviderConfig`).
// See: https://docs.aws.amazon.com/bedrock-agentcore/latest/devguide/identity-idp-dropbox.html
//
// Experimental.
func OAuth2CredentialProvider_UsingDropbox(scope constructs.Construct, id *string, props *DropboxOAuth2CredentialProviderProps) OAuth2CredentialProvider {
	_init_.Initialize()

	if err := validateOAuth2CredentialProvider_UsingDropboxParameters(scope, id, props); err != nil {
		panic(err)
	}
	var returns OAuth2CredentialProvider

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.OAuth2CredentialProvider",
		"usingDropbox",
		[]interface{}{scope, id, props},
		&returns,
	)

	return returns
}

// Create a credential provider for Facebook OAuth2 (`IncludedOauth2ProviderConfig`).
// See: https://docs.aws.amazon.com/bedrock-agentcore/latest/devguide/identity-idp-facebook.html
//
// Experimental.
func OAuth2CredentialProvider_UsingFacebook(scope constructs.Construct, id *string, props *FacebookOAuth2CredentialProviderProps) OAuth2CredentialProvider {
	_init_.Initialize()

	if err := validateOAuth2CredentialProvider_UsingFacebookParameters(scope, id, props); err != nil {
		panic(err)
	}
	var returns OAuth2CredentialProvider

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.OAuth2CredentialProvider",
		"usingFacebook",
		[]interface{}{scope, id, props},
		&returns,
	)

	return returns
}

// Create a credential provider for FusionAuth OAuth2 (`IncludedOauth2ProviderConfig`).
// See: https://docs.aws.amazon.com/bedrock-agentcore/latest/devguide/identity-idp-fusionauth.html
//
// Experimental.
func OAuth2CredentialProvider_UsingFusionAuth(scope constructs.Construct, id *string, props *IncludedOauth2TenantCredentialProviderProps) OAuth2CredentialProvider {
	_init_.Initialize()

	if err := validateOAuth2CredentialProvider_UsingFusionAuthParameters(scope, id, props); err != nil {
		panic(err)
	}
	var returns OAuth2CredentialProvider

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.OAuth2CredentialProvider",
		"usingFusionAuth",
		[]interface{}{scope, id, props},
		&returns,
	)

	return returns
}

// Create a credential provider for GitHub OAuth2.
// Experimental.
func OAuth2CredentialProvider_UsingGithub(scope constructs.Construct, id *string, props *GithubOAuth2CredentialProviderProps) OAuth2CredentialProvider {
	_init_.Initialize()

	if err := validateOAuth2CredentialProvider_UsingGithubParameters(scope, id, props); err != nil {
		panic(err)
	}
	var returns OAuth2CredentialProvider

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.OAuth2CredentialProvider",
		"usingGithub",
		[]interface{}{scope, id, props},
		&returns,
	)

	return returns
}

// Create a credential provider for Google OAuth2.
// Experimental.
func OAuth2CredentialProvider_UsingGoogle(scope constructs.Construct, id *string, props *GoogleOAuth2CredentialProviderProps) OAuth2CredentialProvider {
	_init_.Initialize()

	if err := validateOAuth2CredentialProvider_UsingGoogleParameters(scope, id, props); err != nil {
		panic(err)
	}
	var returns OAuth2CredentialProvider

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.OAuth2CredentialProvider",
		"usingGoogle",
		[]interface{}{scope, id, props},
		&returns,
	)

	return returns
}

// Create a credential provider for HubSpot OAuth2 (`IncludedOauth2ProviderConfig`).
// See: https://docs.aws.amazon.com/bedrock-agentcore/latest/devguide/identity-idp-hubspot.html
//
// Experimental.
func OAuth2CredentialProvider_UsingHubspot(scope constructs.Construct, id *string, props *HubspotOAuth2CredentialProviderProps) OAuth2CredentialProvider {
	_init_.Initialize()

	if err := validateOAuth2CredentialProvider_UsingHubspotParameters(scope, id, props); err != nil {
		panic(err)
	}
	var returns OAuth2CredentialProvider

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.OAuth2CredentialProvider",
		"usingHubspot",
		[]interface{}{scope, id, props},
		&returns,
	)

	return returns
}

// Create a credential provider for LinkedIn OAuth2.
// Experimental.
func OAuth2CredentialProvider_UsingLinkedin(scope constructs.Construct, id *string, props *LinkedinOAuth2CredentialProviderProps) OAuth2CredentialProvider {
	_init_.Initialize()

	if err := validateOAuth2CredentialProvider_UsingLinkedinParameters(scope, id, props); err != nil {
		panic(err)
	}
	var returns OAuth2CredentialProvider

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.OAuth2CredentialProvider",
		"usingLinkedin",
		[]interface{}{scope, id, props},
		&returns,
	)

	return returns
}

// Create a credential provider for Microsoft (Entra ID) OAuth2.
// Experimental.
func OAuth2CredentialProvider_UsingMicrosoft(scope constructs.Construct, id *string, props *MicrosoftOAuth2CredentialProviderProps) OAuth2CredentialProvider {
	_init_.Initialize()

	if err := validateOAuth2CredentialProvider_UsingMicrosoftParameters(scope, id, props); err != nil {
		panic(err)
	}
	var returns OAuth2CredentialProvider

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.OAuth2CredentialProvider",
		"usingMicrosoft",
		[]interface{}{scope, id, props},
		&returns,
	)

	return returns
}

// Create a credential provider for Notion OAuth2 (`IncludedOauth2ProviderConfig`).
// See: https://docs.aws.amazon.com/bedrock-agentcore/latest/devguide/identity-idp-notion.html
//
// Experimental.
func OAuth2CredentialProvider_UsingNotion(scope constructs.Construct, id *string, props *NotionOAuth2CredentialProviderProps) OAuth2CredentialProvider {
	_init_.Initialize()

	if err := validateOAuth2CredentialProvider_UsingNotionParameters(scope, id, props); err != nil {
		panic(err)
	}
	var returns OAuth2CredentialProvider

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.OAuth2CredentialProvider",
		"usingNotion",
		[]interface{}{scope, id, props},
		&returns,
	)

	return returns
}

// Create a credential provider for Okta OAuth2 (`IncludedOauth2ProviderConfig`).
// See: https://docs.aws.amazon.com/bedrock-agentcore/latest/devguide/identity-idp-okta.html
//
// Experimental.
func OAuth2CredentialProvider_UsingOkta(scope constructs.Construct, id *string, props *IncludedOauth2TenantCredentialProviderProps) OAuth2CredentialProvider {
	_init_.Initialize()

	if err := validateOAuth2CredentialProvider_UsingOktaParameters(scope, id, props); err != nil {
		panic(err)
	}
	var returns OAuth2CredentialProvider

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.OAuth2CredentialProvider",
		"usingOkta",
		[]interface{}{scope, id, props},
		&returns,
	)

	return returns
}

// Create a credential provider for OneLogin OAuth2 (`IncludedOauth2ProviderConfig`).
// See: https://docs.aws.amazon.com/bedrock-agentcore/latest/devguide/identity-idp-onelogin.html
//
// Experimental.
func OAuth2CredentialProvider_UsingOneLogin(scope constructs.Construct, id *string, props *IncludedOauth2TenantCredentialProviderProps) OAuth2CredentialProvider {
	_init_.Initialize()

	if err := validateOAuth2CredentialProvider_UsingOneLoginParameters(scope, id, props); err != nil {
		panic(err)
	}
	var returns OAuth2CredentialProvider

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.OAuth2CredentialProvider",
		"usingOneLogin",
		[]interface{}{scope, id, props},
		&returns,
	)

	return returns
}

// Create a credential provider for PingOne OAuth2 (`IncludedOauth2ProviderConfig`).
// See: https://docs.aws.amazon.com/bedrock-agentcore/latest/devguide/identity-idp-pingidentity.html
//
// Experimental.
func OAuth2CredentialProvider_UsingPingOne(scope constructs.Construct, id *string, props *IncludedOauth2TenantCredentialProviderProps) OAuth2CredentialProvider {
	_init_.Initialize()

	if err := validateOAuth2CredentialProvider_UsingPingOneParameters(scope, id, props); err != nil {
		panic(err)
	}
	var returns OAuth2CredentialProvider

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.OAuth2CredentialProvider",
		"usingPingOne",
		[]interface{}{scope, id, props},
		&returns,
	)

	return returns
}

// Create a credential provider for Reddit OAuth2 (`IncludedOauth2ProviderConfig`).
// See: https://docs.aws.amazon.com/bedrock-agentcore/latest/devguide/identity-idp-reddit.html
//
// Experimental.
func OAuth2CredentialProvider_UsingReddit(scope constructs.Construct, id *string, props *RedditOAuth2CredentialProviderProps) OAuth2CredentialProvider {
	_init_.Initialize()

	if err := validateOAuth2CredentialProvider_UsingRedditParameters(scope, id, props); err != nil {
		panic(err)
	}
	var returns OAuth2CredentialProvider

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.OAuth2CredentialProvider",
		"usingReddit",
		[]interface{}{scope, id, props},
		&returns,
	)

	return returns
}

// Create a credential provider for Salesforce OAuth2.
// Experimental.
func OAuth2CredentialProvider_UsingSalesforce(scope constructs.Construct, id *string, props *SalesforceOAuth2CredentialProviderProps) OAuth2CredentialProvider {
	_init_.Initialize()

	if err := validateOAuth2CredentialProvider_UsingSalesforceParameters(scope, id, props); err != nil {
		panic(err)
	}
	var returns OAuth2CredentialProvider

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.OAuth2CredentialProvider",
		"usingSalesforce",
		[]interface{}{scope, id, props},
		&returns,
	)

	return returns
}

// Create a credential provider for Slack OAuth2.
// Experimental.
func OAuth2CredentialProvider_UsingSlack(scope constructs.Construct, id *string, props *SlackOAuth2CredentialProviderProps) OAuth2CredentialProvider {
	_init_.Initialize()

	if err := validateOAuth2CredentialProvider_UsingSlackParameters(scope, id, props); err != nil {
		panic(err)
	}
	var returns OAuth2CredentialProvider

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.OAuth2CredentialProvider",
		"usingSlack",
		[]interface{}{scope, id, props},
		&returns,
	)

	return returns
}

// Create a credential provider for Spotify OAuth2 (`IncludedOauth2ProviderConfig`).
// See: https://docs.aws.amazon.com/bedrock-agentcore/latest/devguide/identity-idp-spotify.html
//
// Experimental.
func OAuth2CredentialProvider_UsingSpotify(scope constructs.Construct, id *string, props *SpotifyOAuth2CredentialProviderProps) OAuth2CredentialProvider {
	_init_.Initialize()

	if err := validateOAuth2CredentialProvider_UsingSpotifyParameters(scope, id, props); err != nil {
		panic(err)
	}
	var returns OAuth2CredentialProvider

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.OAuth2CredentialProvider",
		"usingSpotify",
		[]interface{}{scope, id, props},
		&returns,
	)

	return returns
}

// Create a credential provider for Twitch OAuth2 (`IncludedOauth2ProviderConfig`).
// See: https://docs.aws.amazon.com/bedrock-agentcore/latest/devguide/identity-idp-twitch.html
//
// Experimental.
func OAuth2CredentialProvider_UsingTwitch(scope constructs.Construct, id *string, props *TwitchOAuth2CredentialProviderProps) OAuth2CredentialProvider {
	_init_.Initialize()

	if err := validateOAuth2CredentialProvider_UsingTwitchParameters(scope, id, props); err != nil {
		panic(err)
	}
	var returns OAuth2CredentialProvider

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.OAuth2CredentialProvider",
		"usingTwitch",
		[]interface{}{scope, id, props},
		&returns,
	)

	return returns
}

// Create a credential provider for X (Twitter) OAuth2 (`IncludedOauth2ProviderConfig`).
// See: https://docs.aws.amazon.com/bedrock-agentcore/latest/devguide/identity-idp-x.html
//
// Experimental.
func OAuth2CredentialProvider_UsingX(scope constructs.Construct, id *string, props *XOAuth2CredentialProviderProps) OAuth2CredentialProvider {
	_init_.Initialize()

	if err := validateOAuth2CredentialProvider_UsingXParameters(scope, id, props); err != nil {
		panic(err)
	}
	var returns OAuth2CredentialProvider

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.OAuth2CredentialProvider",
		"usingX",
		[]interface{}{scope, id, props},
		&returns,
	)

	return returns
}

// Create a credential provider for Yandex OAuth2 (`IncludedOauth2ProviderConfig`).
// See: https://docs.aws.amazon.com/bedrock-agentcore/latest/devguide/identity-idp-yandex.html
//
// Experimental.
func OAuth2CredentialProvider_UsingYandex(scope constructs.Construct, id *string, props *YandexOAuth2CredentialProviderProps) OAuth2CredentialProvider {
	_init_.Initialize()

	if err := validateOAuth2CredentialProvider_UsingYandexParameters(scope, id, props); err != nil {
		panic(err)
	}
	var returns OAuth2CredentialProvider

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.OAuth2CredentialProvider",
		"usingYandex",
		[]interface{}{scope, id, props},
		&returns,
	)

	return returns
}

// Create a credential provider for Zoom OAuth2 (`IncludedOauth2ProviderConfig`).
// See: https://docs.aws.amazon.com/bedrock-agentcore/latest/devguide/identity-idp-zoom.html
//
// Experimental.
func OAuth2CredentialProvider_UsingZoom(scope constructs.Construct, id *string, props *ZoomOAuth2CredentialProviderProps) OAuth2CredentialProvider {
	_init_.Initialize()

	if err := validateOAuth2CredentialProvider_UsingZoomParameters(scope, id, props); err != nil {
		panic(err)
	}
	var returns OAuth2CredentialProvider

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.OAuth2CredentialProvider",
		"usingZoom",
		[]interface{}{scope, id, props},
		&returns,
	)

	return returns
}

func OAuth2CredentialProvider_PROPERTY_INJECTION_ID() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-agentcore-alpha.OAuth2CredentialProvider",
		"PROPERTY_INJECTION_ID",
		&returns,
	)
	return returns
}

func (o *jsiiProxy_OAuth2CredentialProvider) ApplyCrossStackReferenceStrength(strength awscdk.ReferenceStrength) {
	if err := o.validateApplyCrossStackReferenceStrengthParameters(strength); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"applyCrossStackReferenceStrength",
		[]interface{}{strength},
	)
}

func (o *jsiiProxy_OAuth2CredentialProvider) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := o.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (o *jsiiProxy_OAuth2CredentialProvider) BindForGatewayOAuthTarget(scopes *[]*string, customParameters *map[string]*string) *GatewayOAuth2IdentityBinding {
	if err := o.validateBindForGatewayOAuthTargetParameters(scopes); err != nil {
		panic(err)
	}
	var returns *GatewayOAuth2IdentityBinding

	_jsii_.Invoke(
		o,
		"bindForGatewayOAuthTarget",
		[]interface{}{scopes, customParameters},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OAuth2CredentialProvider) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OAuth2CredentialProvider) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := o.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OAuth2CredentialProvider) GetResourceNameAttribute(nameAttr *string) *string {
	if err := o.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OAuth2CredentialProvider) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
	if err := o.validateGrantParameters(grantee); err != nil {
		panic(err)
	}
	args := []interface{}{grantee}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		o,
		"grant",
		args,
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OAuth2CredentialProvider) GrantAdmin(grantee awsiam.IGrantable) awsiam.Grant {
	if err := o.validateGrantAdminParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		o,
		"grantAdmin",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OAuth2CredentialProvider) GrantFullAccess(grantee awsiam.IGrantable) awsiam.Grant {
	if err := o.validateGrantFullAccessParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		o,
		"grantFullAccess",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OAuth2CredentialProvider) GrantRead(grantee awsiam.IGrantable) awsiam.Grant {
	if err := o.validateGrantReadParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		o,
		"grantRead",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OAuth2CredentialProvider) GrantUse(grantee awsiam.IGrantable) awsiam.Grant {
	if err := o.validateGrantUseParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		o,
		"grantUse",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OAuth2CredentialProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OAuth2CredentialProvider) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		o,
		"with",
		args,
		&returns,
	)

	return returns
}

