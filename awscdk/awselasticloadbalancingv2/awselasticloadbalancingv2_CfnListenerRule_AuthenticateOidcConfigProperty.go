package awselasticloadbalancingv2


// Specifies information required using an identity provide (IdP) that is compliant with OpenID Connect (OIDC) to authenticate users.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   authenticateOidcConfigProperty := &authenticateOidcConfigProperty{
//   	authorizationEndpoint: jsii.String("authorizationEndpoint"),
//   	clientId: jsii.String("clientId"),
//   	issuer: jsii.String("issuer"),
//   	tokenEndpoint: jsii.String("tokenEndpoint"),
//   	userInfoEndpoint: jsii.String("userInfoEndpoint"),
//
//   	// the properties below are optional
//   	authenticationRequestExtraParams: map[string]*string{
//   		"authenticationRequestExtraParamsKey": jsii.String("authenticationRequestExtraParams"),
//   	},
//   	clientSecret: jsii.String("clientSecret"),
//   	onUnauthenticatedRequest: jsii.String("onUnauthenticatedRequest"),
//   	scope: jsii.String("scope"),
//   	sessionCookieName: jsii.String("sessionCookieName"),
//   	sessionTimeout: jsii.Number(123),
//   	useExistingClientSecret: jsii.Boolean(false),
//   }
//
type CfnListenerRule_AuthenticateOidcConfigProperty struct {
	// The authorization endpoint of the IdP.
	//
	// This must be a full URL, including the HTTPS protocol, the domain, and the path.
	AuthorizationEndpoint *string `field:"required" json:"authorizationEndpoint" yaml:"authorizationEndpoint"`
	// The OAuth 2.0 client identifier.
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The OIDC issuer identifier of the IdP.
	//
	// This must be a full URL, including the HTTPS protocol, the domain, and the path.
	Issuer *string `field:"required" json:"issuer" yaml:"issuer"`
	// The token endpoint of the IdP.
	//
	// This must be a full URL, including the HTTPS protocol, the domain, and the path.
	TokenEndpoint *string `field:"required" json:"tokenEndpoint" yaml:"tokenEndpoint"`
	// The user info endpoint of the IdP.
	//
	// This must be a full URL, including the HTTPS protocol, the domain, and the path.
	UserInfoEndpoint *string `field:"required" json:"userInfoEndpoint" yaml:"userInfoEndpoint"`
	// The query parameters (up to 10) to include in the redirect request to the authorization endpoint.
	AuthenticationRequestExtraParams interface{} `field:"optional" json:"authenticationRequestExtraParams" yaml:"authenticationRequestExtraParams"`
	// The OAuth 2.0 client secret. This parameter is required if you are creating a rule. If you are modifying a rule, you can omit this parameter if you set `UseExistingClientSecret` to true.
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// The behavior if the user is not authenticated. The following are possible values:.
	//
	// - deny `` - Return an HTTP 401 Unauthorized error.
	// - allow `` - Allow the request to be forwarded to the target.
	// - authenticate `` - Redirect the request to the IdP authorization endpoint. This is the default value.
	OnUnauthenticatedRequest *string `field:"optional" json:"onUnauthenticatedRequest" yaml:"onUnauthenticatedRequest"`
	// The set of user claims to be requested from the IdP. The default is `openid` .
	//
	// To verify which scope values your IdP supports and how to separate multiple values, see the documentation for your IdP.
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// The name of the cookie used to maintain session information.
	//
	// The default is AWSELBAuthSessionCookie.
	SessionCookieName *string `field:"optional" json:"sessionCookieName" yaml:"sessionCookieName"`
	// The maximum duration of the authentication session, in seconds.
	//
	// The default is 604800 seconds (7 days).
	SessionTimeout *float64 `field:"optional" json:"sessionTimeout" yaml:"sessionTimeout"`
	// Indicates whether to use the existing client secret when modifying a rule.
	//
	// If you are creating a rule, you can omit this parameter or set it to false.
	UseExistingClientSecret interface{} `field:"optional" json:"useExistingClientSecret" yaml:"useExistingClientSecret"`
}

