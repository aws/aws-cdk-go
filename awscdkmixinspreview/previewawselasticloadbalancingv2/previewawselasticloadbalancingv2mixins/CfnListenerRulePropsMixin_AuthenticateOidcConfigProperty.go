package previewawselasticloadbalancingv2mixins


// Specifies information required using an identity provide (IdP) that is compliant with OpenID Connect (OIDC) to authenticate users.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   authenticateOidcConfigProperty := &AuthenticateOidcConfigProperty{
//   	AuthenticationRequestExtraParams: map[string]*string{
//   		"authenticationRequestExtraParamsKey": jsii.String("authenticationRequestExtraParams"),
//   	},
//   	AuthorizationEndpoint: jsii.String("authorizationEndpoint"),
//   	ClientId: jsii.String("clientId"),
//   	ClientSecret: jsii.String("clientSecret"),
//   	Issuer: jsii.String("issuer"),
//   	OnUnauthenticatedRequest: jsii.String("onUnauthenticatedRequest"),
//   	Scope: jsii.String("scope"),
//   	SessionCookieName: jsii.String("sessionCookieName"),
//   	SessionTimeout: jsii.Number(123),
//   	TokenEndpoint: jsii.String("tokenEndpoint"),
//   	UseExistingClientSecret: jsii.Boolean(false),
//   	UserInfoEndpoint: jsii.String("userInfoEndpoint"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-authenticateoidcconfig.html
//
type CfnListenerRulePropsMixin_AuthenticateOidcConfigProperty struct {
	// The query parameters (up to 10) to include in the redirect request to the authorization endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-authenticateoidcconfig.html#cfn-elasticloadbalancingv2-listenerrule-authenticateoidcconfig-authenticationrequestextraparams
	//
	AuthenticationRequestExtraParams interface{} `field:"optional" json:"authenticationRequestExtraParams" yaml:"authenticationRequestExtraParams"`
	// The authorization endpoint of the IdP.
	//
	// This must be a full URL, including the HTTPS protocol, the domain, and the path.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-authenticateoidcconfig.html#cfn-elasticloadbalancingv2-listenerrule-authenticateoidcconfig-authorizationendpoint
	//
	AuthorizationEndpoint *string `field:"optional" json:"authorizationEndpoint" yaml:"authorizationEndpoint"`
	// The OAuth 2.0 client identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-authenticateoidcconfig.html#cfn-elasticloadbalancingv2-listenerrule-authenticateoidcconfig-clientid
	//
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// The OAuth 2.0 client secret. This parameter is required if you are creating a rule. If you are modifying a rule, you can omit this parameter if you set `UseExistingClientSecret` to true.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-authenticateoidcconfig.html#cfn-elasticloadbalancingv2-listenerrule-authenticateoidcconfig-clientsecret
	//
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// The OIDC issuer identifier of the IdP.
	//
	// This must be a full URL, including the HTTPS protocol, the domain, and the path.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-authenticateoidcconfig.html#cfn-elasticloadbalancingv2-listenerrule-authenticateoidcconfig-issuer
	//
	Issuer *string `field:"optional" json:"issuer" yaml:"issuer"`
	// The behavior if the user is not authenticated. The following are possible values:.
	//
	// - deny `` - Return an HTTP 401 Unauthorized error.
	// - allow `` - Allow the request to be forwarded to the target.
	// - authenticate `` - Redirect the request to the IdP authorization endpoint. This is the default value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-authenticateoidcconfig.html#cfn-elasticloadbalancingv2-listenerrule-authenticateoidcconfig-onunauthenticatedrequest
	//
	OnUnauthenticatedRequest *string `field:"optional" json:"onUnauthenticatedRequest" yaml:"onUnauthenticatedRequest"`
	// The set of user claims to be requested from the IdP. The default is `openid` .
	//
	// To verify which scope values your IdP supports and how to separate multiple values, see the documentation for your IdP.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-authenticateoidcconfig.html#cfn-elasticloadbalancingv2-listenerrule-authenticateoidcconfig-scope
	//
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// The name of the cookie used to maintain session information.
	//
	// The default is AWSELBAuthSessionCookie.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-authenticateoidcconfig.html#cfn-elasticloadbalancingv2-listenerrule-authenticateoidcconfig-sessioncookiename
	//
	SessionCookieName *string `field:"optional" json:"sessionCookieName" yaml:"sessionCookieName"`
	// The maximum duration of the authentication session, in seconds.
	//
	// The default is 604800 seconds (7 days).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-authenticateoidcconfig.html#cfn-elasticloadbalancingv2-listenerrule-authenticateoidcconfig-sessiontimeout
	//
	SessionTimeout *float64 `field:"optional" json:"sessionTimeout" yaml:"sessionTimeout"`
	// The token endpoint of the IdP.
	//
	// This must be a full URL, including the HTTPS protocol, the domain, and the path.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-authenticateoidcconfig.html#cfn-elasticloadbalancingv2-listenerrule-authenticateoidcconfig-tokenendpoint
	//
	TokenEndpoint *string `field:"optional" json:"tokenEndpoint" yaml:"tokenEndpoint"`
	// Indicates whether to use the existing client secret when modifying a rule.
	//
	// If you are creating a rule, you can omit this parameter or set it to false.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-authenticateoidcconfig.html#cfn-elasticloadbalancingv2-listenerrule-authenticateoidcconfig-useexistingclientsecret
	//
	UseExistingClientSecret interface{} `field:"optional" json:"useExistingClientSecret" yaml:"useExistingClientSecret"`
	// The user info endpoint of the IdP.
	//
	// This must be a full URL, including the HTTPS protocol, the domain, and the path.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-authenticateoidcconfig.html#cfn-elasticloadbalancingv2-listenerrule-authenticateoidcconfig-userinfoendpoint
	//
	UserInfoEndpoint *string `field:"optional" json:"userInfoEndpoint" yaml:"userInfoEndpoint"`
}

