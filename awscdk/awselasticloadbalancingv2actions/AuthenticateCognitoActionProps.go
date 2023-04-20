package awselasticloadbalancingv2actions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
)

// Properties for AuthenticateCognitoAction.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var vpc vpc
//   var certificate certificate
//
//
//   lb := elbv2.NewApplicationLoadBalancer(this, jsii.String("LB"), &ApplicationLoadBalancerProps{
//   	Vpc: Vpc,
//   	InternetFacing: jsii.Boolean(true),
//   })
//
//   userPool := awscdk.Aws_cognito.NewUserPool(this, jsii.String("UserPool"))
//   userPoolClient := awscdk.Aws_cognito.NewUserPoolClient(this, jsii.String("Client"), &UserPoolClientProps{
//   	UserPool: UserPool,
//
//   	// Required minimal configuration for use with an ELB
//   	GenerateSecret: jsii.Boolean(true),
//   	AuthFlows: &AuthFlow{
//   		UserPassword: jsii.Boolean(true),
//   	},
//   	OAuth: &OAuthSettings{
//   		Flows: &OAuthFlows{
//   			AuthorizationCodeGrant: jsii.Boolean(true),
//   		},
//   		Scopes: []oAuthScope{
//   			awscdk.*Aws_cognito.*oAuthScope_EMAIL(),
//   		},
//   		CallbackUrls: []*string{
//   			fmt.Sprintf("https://%v/oauth2/idpresponse", lb.LoadBalancerDnsName),
//   		},
//   	},
//   })
//   cfnClient := userPoolClient.Node.defaultChild.(cfnUserPoolClient)
//   cfnClient.AddPropertyOverride(jsii.String("RefreshTokenValidity"), jsii.Number(1))
//   cfnClient.AddPropertyOverride(jsii.String("SupportedIdentityProviders"), []interface{}{
//   	jsii.String("COGNITO"),
//   })
//
//   userPoolDomain := awscdk.Aws_cognito.NewUserPoolDomain(this, jsii.String("Domain"), &UserPoolDomainProps{
//   	UserPool: UserPool,
//   	CognitoDomain: &CognitoDomainOptions{
//   		DomainPrefix: jsii.String("test-cdk-prefix"),
//   	},
//   })
//
//   lb.AddListener(jsii.String("Listener"), &BaseApplicationListenerProps{
//   	Port: jsii.Number(443),
//   	Certificates: []iListenerCertificate{
//   		certificate,
//   	},
//   	DefaultAction: actions.NewAuthenticateCognitoAction(&AuthenticateCognitoActionProps{
//   		UserPool: *UserPool,
//   		UserPoolClient: *UserPoolClient,
//   		UserPoolDomain: *UserPoolDomain,
//   		Next: elbv2.ListenerAction_FixedResponse(jsii.Number(200), &FixedResponseOptions{
//   			ContentType: jsii.String("text/plain"),
//   			MessageBody: jsii.String("Authenticated"),
//   		}),
//   	}),
//   })
//
//   awscdk.NewCfnOutput(this, jsii.String("DNS"), &CfnOutputProps{
//   	Value: lb.*LoadBalancerDnsName,
//   })
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

