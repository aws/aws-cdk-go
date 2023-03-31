package awselasticloadbalancingv2actions

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscognito"
	"github.com/aws/aws-cdk-go/awscdk/awselasticloadbalancingv2"
)

// Properties for AuthenticateCognitoAction.
//
// Example:
//   import cognito "github.com/aws/aws-cdk-go/awscdk"
//   import ec2 "github.com/aws/aws-cdk-go/awscdk"
//   import elbv2 "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/constructs-go/constructs"
//   import actions "github.com/aws/aws-cdk-go/awscdk"
//
//   cognitoStack struct {
//   stack
//   }
//
//   lb := elbv2.NewApplicationLoadBalancer(this, jsii.String("LB"), &applicationLoadBalancerProps{
//   	vpc: vpc,
//   	internetFacing: jsii.Boolean(true),
//   })
//
//   userPool := cognito.NewUserPool(this, jsii.String("UserPool"))
//   userPoolClient := cognito.NewUserPoolClient(this, jsii.String("Client"), &userPoolClientProps{
//   	userPool: userPool,
//
//   	// Required minimal configuration for use with an ELB
//   	generateSecret: jsii.Boolean(true),
//   	authFlows: &authFlow{
//   		userPassword: jsii.Boolean(true),
//   	},
//   	oAuth: &oAuthSettings{
//   		flows: &oAuthFlows{
//   			authorizationCodeGrant: jsii.Boolean(true),
//   		},
//   		scopes: []oAuthScope{
//   			cognito.*oAuthScope_EMAIL(),
//   		},
//   		callbackUrls: []*string{
//   			fmt.Sprintf("https://%v/oauth2/idpresponse", lb.loadBalancerDnsName),
//   		},
//   	},
//   })
//   cfnClient := userPoolClient.node.defaultChild.(cfnUserPoolClient)
//   cfnClient.addPropertyOverride(jsii.String("RefreshTokenValidity"), jsii.Number(1))
//   cfnClient.addPropertyOverride(jsii.String("SupportedIdentityProviders"), []interface{}{
//   	jsii.String("COGNITO"),
//   })
//
//   userPoolDomain := cognito.NewUserPoolDomain(this, jsii.String("Domain"), &userPoolDomainProps{
//   	userPool: userPool,
//   	cognitoDomain: &cognitoDomainOptions{
//   		domainPrefix: jsii.String("test-cdk-prefix"),
//   	},
//   })
//
//   lb.addListener(jsii.String("Listener"), &baseApplicationListenerProps{
//   	port: jsii.Number(443),
//   	certificates: []iListenerCertificate{
//   		certificate,
//   	},
//   	defaultAction: actions.NewAuthenticateCognitoAction(&authenticateCognitoActionProps{
//   		userPool: userPool,
//   		userPoolClient: userPoolClient,
//   		userPoolDomain: userPoolDomain,
//   		next: elbv2.listenerAction.fixedResponse(jsii.Number(200), &fixedResponseOptions{
//   			contentType: jsii.String("text/plain"),
//   			messageBody: jsii.String("Authenticated"),
//   		}),
//   	}),
//   })
//
//   awscdk.NewCfnOutput(this, jsii.String("DNS"), &cfnOutputProps{
//   	value: lb.loadBalancerDnsName,
//   })
//
//   app := awscdk.NewApp()
//   NewCognitoStack(app, jsii.String("integ-cognito"))
//   app.synth()
//
// Experimental.
type AuthenticateCognitoActionProps struct {
	// What action to execute next.
	//
	// Multiple actions form a linked chain; the chain must always terminate in a
	// (weighted)forward, fixedResponse or redirect action.
	// Experimental.
	Next awselasticloadbalancingv2.ListenerAction `field:"required" json:"next" yaml:"next"`
	// The Amazon Cognito user pool.
	// Experimental.
	UserPool awscognito.IUserPool `field:"required" json:"userPool" yaml:"userPool"`
	// The Amazon Cognito user pool client.
	// Experimental.
	UserPoolClient awscognito.IUserPoolClient `field:"required" json:"userPoolClient" yaml:"userPoolClient"`
	// The domain prefix or fully-qualified domain name of the Amazon Cognito user pool.
	// Experimental.
	UserPoolDomain awscognito.IUserPoolDomain `field:"required" json:"userPoolDomain" yaml:"userPoolDomain"`
	// The query parameters (up to 10) to include in the redirect request to the authorization endpoint.
	// Experimental.
	AuthenticationRequestExtraParams *map[string]*string `field:"optional" json:"authenticationRequestExtraParams" yaml:"authenticationRequestExtraParams"`
	// The behavior if the user is not authenticated.
	// Experimental.
	OnUnauthenticatedRequest awselasticloadbalancingv2.UnauthenticatedAction `field:"optional" json:"onUnauthenticatedRequest" yaml:"onUnauthenticatedRequest"`
	// The set of user claims to be requested from the IdP.
	//
	// To verify which scope values your IdP supports and how to separate multiple values, see the documentation for your IdP.
	// Experimental.
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// The name of the cookie used to maintain session information.
	// Experimental.
	SessionCookieName *string `field:"optional" json:"sessionCookieName" yaml:"sessionCookieName"`
	// The maximum duration of the authentication session.
	// Experimental.
	SessionTimeout awscdk.Duration `field:"optional" json:"sessionTimeout" yaml:"sessionTimeout"`
}

