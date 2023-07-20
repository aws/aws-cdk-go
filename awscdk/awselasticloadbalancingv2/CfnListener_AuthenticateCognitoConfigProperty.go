package awselasticloadbalancingv2


// Specifies information required when integrating with Amazon Cognito to authenticate users.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   authenticateCognitoConfigProperty := &AuthenticateCognitoConfigProperty{
//   	UserPoolArn: jsii.String("userPoolArn"),
//   	UserPoolClientId: jsii.String("userPoolClientId"),
//   	UserPoolDomain: jsii.String("userPoolDomain"),
//
//   	// the properties below are optional
//   	AuthenticationRequestExtraParams: map[string]*string{
//   		"authenticationRequestExtraParamsKey": jsii.String("authenticationRequestExtraParams"),
//   	},
//   	OnUnauthenticatedRequest: jsii.String("onUnauthenticatedRequest"),
//   	Scope: jsii.String("scope"),
//   	SessionCookieName: jsii.String("sessionCookieName"),
//   	SessionTimeout: jsii.String("sessionTimeout"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-authenticatecognitoconfig.html
//
type CfnListener_AuthenticateCognitoConfigProperty struct {
	// The Amazon Resource Name (ARN) of the Amazon Cognito user pool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-authenticatecognitoconfig.html#cfn-elasticloadbalancingv2-listener-authenticatecognitoconfig-userpoolarn
	//
	UserPoolArn *string `field:"required" json:"userPoolArn" yaml:"userPoolArn"`
	// The ID of the Amazon Cognito user pool client.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-authenticatecognitoconfig.html#cfn-elasticloadbalancingv2-listener-authenticatecognitoconfig-userpoolclientid
	//
	UserPoolClientId *string `field:"required" json:"userPoolClientId" yaml:"userPoolClientId"`
	// The domain prefix or fully-qualified domain name of the Amazon Cognito user pool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-authenticatecognitoconfig.html#cfn-elasticloadbalancingv2-listener-authenticatecognitoconfig-userpooldomain
	//
	UserPoolDomain *string `field:"required" json:"userPoolDomain" yaml:"userPoolDomain"`
	// The query parameters (up to 10) to include in the redirect request to the authorization endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-authenticatecognitoconfig.html#cfn-elasticloadbalancingv2-listener-authenticatecognitoconfig-authenticationrequestextraparams
	//
	AuthenticationRequestExtraParams interface{} `field:"optional" json:"authenticationRequestExtraParams" yaml:"authenticationRequestExtraParams"`
	// The behavior if the user is not authenticated. The following are possible values:.
	//
	// - deny `` - Return an HTTP 401 Unauthorized error.
	// - allow `` - Allow the request to be forwarded to the target.
	// - authenticate `` - Redirect the request to the IdP authorization endpoint. This is the default value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-authenticatecognitoconfig.html#cfn-elasticloadbalancingv2-listener-authenticatecognitoconfig-onunauthenticatedrequest
	//
	OnUnauthenticatedRequest *string `field:"optional" json:"onUnauthenticatedRequest" yaml:"onUnauthenticatedRequest"`
	// The set of user claims to be requested from the IdP. The default is `openid` .
	//
	// To verify which scope values your IdP supports and how to separate multiple values, see the documentation for your IdP.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-authenticatecognitoconfig.html#cfn-elasticloadbalancingv2-listener-authenticatecognitoconfig-scope
	//
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// The name of the cookie used to maintain session information.
	//
	// The default is AWSELBAuthSessionCookie.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-authenticatecognitoconfig.html#cfn-elasticloadbalancingv2-listener-authenticatecognitoconfig-sessioncookiename
	//
	SessionCookieName *string `field:"optional" json:"sessionCookieName" yaml:"sessionCookieName"`
	// The maximum duration of the authentication session, in seconds.
	//
	// The default is 604800 seconds (7 days).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-authenticatecognitoconfig.html#cfn-elasticloadbalancingv2-listener-authenticatecognitoconfig-sessiontimeout
	//
	SessionTimeout *string `field:"optional" json:"sessionTimeout" yaml:"sessionTimeout"`
}

