package awselasticloadbalancingv2actions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
)

// Properties for AuthenticateCognitoAction.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var listenerAction listenerAction
//   var userPool userPool
//   var userPoolClient userPoolClient
//   var userPoolDomain userPoolDomain
//
//   authenticateCognitoActionProps := &AuthenticateCognitoActionProps{
//   	Next: listenerAction,
//   	UserPool: userPool,
//   	UserPoolClient: userPoolClient,
//   	UserPoolDomain: userPoolDomain,
//
//   	// the properties below are optional
//   	AuthenticationRequestExtraParams: map[string]*string{
//   		"authenticationRequestExtraParamsKey": jsii.String("authenticationRequestExtraParams"),
//   	},
//   	OnUnauthenticatedRequest: awscdk.Aws_elasticloadbalancingv2.UnauthenticatedAction_DENY,
//   	Scope: jsii.String("scope"),
//   	SessionCookieName: jsii.String("sessionCookieName"),
//   	SessionTimeout: cdk.Duration_Minutes(jsii.Number(30)),
//   }
//
type AuthenticateCognitoActionProps struct {
	// What action to execute next.
	//
	// Multiple actions form a linked chain; the chain must always terminate in a
	// (weighted)forward, fixedResponse or redirect action.
	Next awselasticloadbalancingv2.ListenerAction `field:"required" json:"next" yaml:"next"`
	// The Amazon Cognito user pool.
	UserPool awscognito.IUserPool `field:"required" json:"userPool" yaml:"userPool"`
	// The Amazon Cognito user pool client.
	UserPoolClient awscognito.IUserPoolClient `field:"required" json:"userPoolClient" yaml:"userPoolClient"`
	// The domain prefix or fully-qualified domain name of the Amazon Cognito user pool.
	UserPoolDomain awscognito.IUserPoolDomain `field:"required" json:"userPoolDomain" yaml:"userPoolDomain"`
	// The query parameters (up to 10) to include in the redirect request to the authorization endpoint.
	AuthenticationRequestExtraParams *map[string]*string `field:"optional" json:"authenticationRequestExtraParams" yaml:"authenticationRequestExtraParams"`
	// The behavior if the user is not authenticated.
	OnUnauthenticatedRequest awselasticloadbalancingv2.UnauthenticatedAction `field:"optional" json:"onUnauthenticatedRequest" yaml:"onUnauthenticatedRequest"`
	// The set of user claims to be requested from the IdP.
	//
	// To verify which scope values your IdP supports and how to separate multiple values, see the documentation for your IdP.
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// The name of the cookie used to maintain session information.
	SessionCookieName *string `field:"optional" json:"sessionCookieName" yaml:"sessionCookieName"`
	// The maximum duration of the authentication session.
	SessionTimeout awscdk.Duration `field:"optional" json:"sessionTimeout" yaml:"sessionTimeout"`
}

