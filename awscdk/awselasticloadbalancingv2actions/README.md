# Actions for AWS Elastic Load Balancing V2

This package contains integration actions for ELBv2. See the README of the `@aws-cdk/aws-elasticloadbalancingv2` library.

## Cognito

ELB allows for requests to be authenticated against a Cognito user pool using
the `AuthenticateCognitoAction`. For details on the setup's requirements,
read [Prepare to use Amazon
Cognito](https://docs.aws.amazon.com/elasticloadbalancing/latest/application/listener-authenticate-users.html#cognito-requirements).
Here's an example:

```go
import cognito "github.com/aws/aws-cdk-go/awscdk"
import ec2 "github.com/aws/aws-cdk-go/awscdk"
import elbv2 "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/constructs-go/constructs"
import actions "github.com/aws/aws-cdk-go/awscdk"

cognitoStack struct {
stack
}

lb := elbv2.NewApplicationLoadBalancer(this, jsii.String("LB"), &applicationLoadBalancerProps{
	vpc: vpc,
	internetFacing: jsii.Boolean(true),
})

userPool := cognito.NewUserPool(this, jsii.String("UserPool"))
userPoolClient := cognito.NewUserPoolClient(this, jsii.String("Client"), &userPoolClientProps{
	userPool: userPool,

	// Required minimal configuration for use with an ELB
	generateSecret: jsii.Boolean(true),
	authFlows: &authFlow{
		userPassword: jsii.Boolean(true),
	},
	oAuth: &oAuthSettings{
		flows: &oAuthFlows{
			authorizationCodeGrant: jsii.Boolean(true),
		},
		scopes: []oAuthScope{
			cognito.*oAuthScope_EMAIL(),
		},
		callbackUrls: []*string{
			fmt.Sprintf("https://%v/oauth2/idpresponse", lb.loadBalancerDnsName),
		},
	},
})
cfnClient := userPoolClient.node.defaultChild.(cfnUserPoolClient)
cfnClient.addPropertyOverride(jsii.String("RefreshTokenValidity"), jsii.Number(1))
cfnClient.addPropertyOverride(jsii.String("SupportedIdentityProviders"), []interface{}{
	jsii.String("COGNITO"),
})

userPoolDomain := cognito.NewUserPoolDomain(this, jsii.String("Domain"), &userPoolDomainProps{
	userPool: userPool,
	cognitoDomain: &cognitoDomainOptions{
		domainPrefix: jsii.String("test-cdk-prefix"),
	},
})

lb.addListener(jsii.String("Listener"), &baseApplicationListenerProps{
	port: jsii.Number(443),
	certificates: []iListenerCertificate{
		certificate,
	},
	defaultAction: actions.NewAuthenticateCognitoAction(&authenticateCognitoActionProps{
		userPool: userPool,
		userPoolClient: userPoolClient,
		userPoolDomain: userPoolDomain,
		next: elbv2.listenerAction.fixedResponse(jsii.Number(200), &fixedResponseOptions{
			contentType: jsii.String("text/plain"),
			messageBody: jsii.String("Authenticated"),
		}),
	}),
})

awscdk.NewCfnOutput(this, jsii.String("DNS"), &cfnOutputProps{
	value: lb.loadBalancerDnsName,
})

app := awscdk.NewApp()
NewCognitoStack(app, jsii.String("integ-cognito"))
app.synth()
```

> NOTE: this example seems incomplete, I was not able to get the redirect back to the
> Load Balancer after authentication working. Would love some pointers on what a full working
> setup actually looks like!
