package awselasticloadbalancingv2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Options for `ListenerAction.authenciateOidc()`.
//
// Example:
//   var listener applicationListener
//   var myTargetGroup applicationTargetGroup
//
//
//   listener.AddAction(jsii.String("DefaultAction"), &AddApplicationActionProps{
//   	Action: elbv2.ListenerAction_AuthenticateOidc(&AuthenticateOidcOptions{
//   		AuthorizationEndpoint: jsii.String("https://example.com/openid"),
//   		// Other OIDC properties here
//   		ClientId: jsii.String("..."),
//   		ClientSecret: awscdk.SecretValue_SecretsManager(jsii.String("...")),
//   		Issuer: jsii.String("..."),
//   		TokenEndpoint: jsii.String("..."),
//   		UserInfoEndpoint: jsii.String("..."),
//
//   		// Next
//   		Next: elbv2.ListenerAction_Forward([]iApplicationTargetGroup{
//   			myTargetGroup,
//   		}),
//   	}),
//   })
//
type AuthenticateOidcOptions struct {
	// The authorization endpoint of the IdP.
	//
	// This must be a full URL, including the HTTPS protocol, the domain, and the path.
	AuthorizationEndpoint *string `field:"required" json:"authorizationEndpoint" yaml:"authorizationEndpoint"`
	// The OAuth 2.0 client identifier.
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The OAuth 2.0 client secret.
	ClientSecret awscdk.SecretValue `field:"required" json:"clientSecret" yaml:"clientSecret"`
	// The OIDC issuer identifier of the IdP.
	//
	// This must be a full URL, including the HTTPS protocol, the domain, and the path.
	Issuer *string `field:"required" json:"issuer" yaml:"issuer"`
	// What action to execute next.
	Next ListenerAction `field:"required" json:"next" yaml:"next"`
	// The token endpoint of the IdP.
	//
	// This must be a full URL, including the HTTPS protocol, the domain, and the path.
	TokenEndpoint *string `field:"required" json:"tokenEndpoint" yaml:"tokenEndpoint"`
	// The user info endpoint of the IdP.
	//
	// This must be a full URL, including the HTTPS protocol, the domain, and the path.
	UserInfoEndpoint *string `field:"required" json:"userInfoEndpoint" yaml:"userInfoEndpoint"`
	// Allow HTTPS outbound traffic to communicate with the IdP.
	//
	// Set this property to false if the IP address used for the IdP endpoint is identifiable
	// and you want to control outbound traffic.
	// Then allow HTTPS outbound traffic to the IdP's IP address using the listener's `connections` property.
	// See: https://repost.aws/knowledge-center/elb-configure-authentication-alb
	//
	// Default: true.
	//
	AllowHttpsOutbound *bool `field:"optional" json:"allowHttpsOutbound" yaml:"allowHttpsOutbound"`
	// The query parameters (up to 10) to include in the redirect request to the authorization endpoint.
	// Default: - No extra parameters.
	//
	AuthenticationRequestExtraParams *map[string]*string `field:"optional" json:"authenticationRequestExtraParams" yaml:"authenticationRequestExtraParams"`
	// The behavior if the user is not authenticated.
	// Default: UnauthenticatedAction.AUTHENTICATE
	//
	OnUnauthenticatedRequest UnauthenticatedAction `field:"optional" json:"onUnauthenticatedRequest" yaml:"onUnauthenticatedRequest"`
	// The set of user claims to be requested from the IdP.
	//
	// To verify which scope values your IdP supports and how to separate multiple values, see the documentation for your IdP.
	// Default: "openid".
	//
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// The name of the cookie used to maintain session information.
	// Default: "AWSELBAuthSessionCookie".
	//
	SessionCookieName *string `field:"optional" json:"sessionCookieName" yaml:"sessionCookieName"`
	// The maximum duration of the authentication session.
	// Default: Duration.days(7)
	//
	SessionTimeout awscdk.Duration `field:"optional" json:"sessionTimeout" yaml:"sessionTimeout"`
}

