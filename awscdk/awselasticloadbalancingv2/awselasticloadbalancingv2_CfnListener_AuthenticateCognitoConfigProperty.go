package awselasticloadbalancingv2


// Specifies information required when integrating with Amazon Cognito to authenticate users.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   authenticateCognitoConfigProperty := &authenticateCognitoConfigProperty{
//   	userPoolArn: jsii.String("userPoolArn"),
//   	userPoolClientId: jsii.String("userPoolClientId"),
//   	userPoolDomain: jsii.String("userPoolDomain"),
//
//   	// the properties below are optional
//   	authenticationRequestExtraParams: map[string]*string{
//   		"authenticationRequestExtraParamsKey": jsii.String("authenticationRequestExtraParams"),
//   	},
//   	onUnauthenticatedRequest: jsii.String("onUnauthenticatedRequest"),
//   	scope: jsii.String("scope"),
//   	sessionCookieName: jsii.String("sessionCookieName"),
//   	sessionTimeout: jsii.String("sessionTimeout"),
//   }
//
type CfnListener_AuthenticateCognitoConfigProperty struct {
	// The Amazon Resource Name (ARN) of the Amazon Cognito user pool.
	UserPoolArn *string `field:"required" json:"userPoolArn" yaml:"userPoolArn"`
	// The ID of the Amazon Cognito user pool client.
	UserPoolClientId *string `field:"required" json:"userPoolClientId" yaml:"userPoolClientId"`
	// The domain prefix or fully-qualified domain name of the Amazon Cognito user pool.
	UserPoolDomain *string `field:"required" json:"userPoolDomain" yaml:"userPoolDomain"`
	// The query parameters (up to 10) to include in the redirect request to the authorization endpoint.
	AuthenticationRequestExtraParams interface{} `field:"optional" json:"authenticationRequestExtraParams" yaml:"authenticationRequestExtraParams"`
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
	SessionTimeout *string `field:"optional" json:"sessionTimeout" yaml:"sessionTimeout"`
}

