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
//   listener.addAction(jsii.String("DefaultAction"), &addApplicationActionProps{
//   	action: elbv2.listenerAction.authenticateOidc(&authenticateOidcOptions{
//   		authorizationEndpoint: jsii.String("https://example.com/openid"),
//   		// Other OIDC properties here
//   		clientId: jsii.String("..."),
//   		clientSecret: awscdk.SecretValue.secretsManager(jsii.String("...")),
//   		issuer: jsii.String("..."),
//   		tokenEndpoint: jsii.String("..."),
//   		userInfoEndpoint: jsii.String("..."),
//
//   		// Next
//   		next: elbv2.*listenerAction.forward([]iApplicationTargetGroup{
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
	// The query parameters (up to 10) to include in the redirect request to the authorization endpoint.
	AuthenticationRequestExtraParams *map[string]*string `field:"optional" json:"authenticationRequestExtraParams" yaml:"authenticationRequestExtraParams"`
	// The behavior if the user is not authenticated.
	OnUnauthenticatedRequest UnauthenticatedAction `field:"optional" json:"onUnauthenticatedRequest" yaml:"onUnauthenticatedRequest"`
	// The set of user claims to be requested from the IdP.
	//
	// To verify which scope values your IdP supports and how to separate multiple values, see the documentation for your IdP.
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// The name of the cookie used to maintain session information.
	SessionCookieName *string `field:"optional" json:"sessionCookieName" yaml:"sessionCookieName"`
	// The maximum duration of the authentication session.
	SessionTimeout awscdk.Duration `field:"optional" json:"sessionTimeout" yaml:"sessionTimeout"`
}

