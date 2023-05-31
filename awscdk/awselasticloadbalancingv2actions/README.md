# Actions for AWS Elastic Load Balancing V2

This package contains integration actions for ELBv2. See the README of the `@aws-cdk/aws-elasticloadbalancingv2` library.

## Cognito

ELB allows for requests to be authenticated against a Cognito user pool using
the `AuthenticateCognitoAction`. For details on the setup's requirements,
read [Prepare to use Amazon
Cognito](https://docs.aws.amazon.com/elasticloadbalancing/latest/application/listener-authenticate-users.html#cognito-requirements).
Here's an example:

```go
import "github.com/aws/aws-cdk-go/awscdk"
import ec2 "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/constructs-go/constructs"
import actions "github.com/aws/aws-cdk-go/awscdk"

cognitoStack struct {
stack
}

lb := elbv2.NewApplicationLoadBalancer(this, jsii.String("LB"), &ApplicationLoadBalancerProps{
	Vpc: Vpc,
	InternetFacing: jsii.Boolean(true),
})

userPool := cognito.NewUserPool(this, jsii.String("UserPool"))
userPoolClient := cognito.NewUserPoolClient(this, jsii.String("Client"), &UserPoolClientProps{
	UserPool: UserPool,

	// Required minimal configuration for use with an ELB
	GenerateSecret: jsii.Boolean(true),
	AuthFlows: &AuthFlow{
		UserPassword: jsii.Boolean(true),
	},
	OAuth: &OAuthSettings{
		Flows: &OAuthFlows{
			AuthorizationCodeGrant: jsii.Boolean(true),
		},
		Scopes: []oAuthScope{
			cognito.*oAuthScope_EMAIL(),
		},
		CallbackUrls: []*string{
			fmt.Sprintf("https://%v/oauth2/idpresponse", lb.LoadBalancerDnsName),
		},
	},
})
cfnClient := userPoolClient.Node.defaultChild.(cfnUserPoolClient)
cfnClient.AddPropertyOverride(jsii.String("RefreshTokenValidity"), jsii.Number(1))
cfnClient.AddPropertyOverride(jsii.String("SupportedIdentityProviders"), []interface{}{
	jsii.String("COGNITO"),
})

userPoolDomain := cognito.NewUserPoolDomain(this, jsii.String("Domain"), &UserPoolDomainProps{
	UserPool: UserPool,
	CognitoDomain: &CognitoDomainOptions{
		DomainPrefix: jsii.String("test-cdk-prefix"),
	},
})

lb.AddListener(jsii.String("Listener"), &BaseApplicationListenerProps{
	Port: jsii.Number(443),
	Certificates: []iListenerCertificate{
		certificate,
	},
	DefaultAction: actions.NewAuthenticateCognitoAction(&AuthenticateCognitoActionProps{
		UserPool: *UserPool,
		UserPoolClient: *UserPoolClient,
		UserPoolDomain: *UserPoolDomain,
		Next: elbv2.ListenerAction_FixedResponse(jsii.Number(200), &FixedResponseOptions{
			ContentType: jsii.String("text/plain"),
			MessageBody: jsii.String("Authenticated"),
		}),
	}),
})

awscdk.NewCfnOutput(this, jsii.String("DNS"), &CfnOutputProps{
	Value: lb.*LoadBalancerDnsName,
})

app := awscdk.NewApp()
NewCognitoStack(app, jsii.String("integ-cognito"))
app.Synth()
```

> NOTE: this example seems incomplete, I was not able to get the redirect back to the
> Load Balancer after authentication working. Would love some pointers on what a full working
> setup actually looks like!
