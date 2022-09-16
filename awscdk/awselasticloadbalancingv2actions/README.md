# Actions for AWS Elastic Load Balancing V2

This package contains integration actions for ELBv2. See the README of the `@aws-cdk/aws-elasticloadbalancingv2` library.

## Cognito

ELB allows for requests to be authenticated against a Cognito user pool using
the `AuthenticateCognitoAction`. For details on the setup's requirements,
read [Prepare to use Amazon
Cognito](https://docs.aws.amazon.com/elasticloadbalancing/latest/application/listener-authenticate-users.html#cognito-requirements).
Here's an example:

```go
// Example automatically generated from non-compiling source. May contain errors.
lb := elbv2.NewApplicationLoadBalancer(this, jsii.String("LB"), map[string]interface{}{
	"vpc": vpc,
	"internetFacing": jsii.Boolean(true),
})

userPool := cognito.NewUserPool(this, jsii.String("UserPool"))
userPoolClient := cognito.NewUserPoolClient(this, jsii.String("Client"), map[string]interface{}{
	"userPool": userPool,

	// Required minimal configuration for use with an ELB
	"generateSecret": jsii.Boolean(true),
	"authFlows": map[string]*bool{
		"userPassword": jsii.Boolean(true),
	},
	"oAuth": map[string]interface{}{
		"flows": map[string]*bool{
			"authorizationCodeGrant": jsii.Boolean(true),
		},
		"scopes": []interface{}{
			cognito.OAuthScope_EMAIL,
		},
		"callbackUrls": []*string{
			fmt.Sprintf("https://%v/oauth2/idpresponse", lb.loadBalancerDnsName),
		},
	},
})
cfnClient := userPoolClient.node.defaultChild.(cognito.CfnUserPoolClient)
cfnClient.addPropertyOverride(jsii.String("RefreshTokenValidity"), jsii.Number(1))
cfnClient.addPropertyOverride(jsii.String("SupportedIdentityProviders"), []interface{}{
	jsii.String("COGNITO"),
})

userPoolDomain := cognito.NewUserPoolDomain(this, jsii.String("Domain"), map[string]interface{}{
	"userPool": userPool,
	"cognitoDomain": map[string]*string{
		"domainPrefix": jsii.String("test-cdk-prefix"),
	},
})

lb.addListener(jsii.String("Listener"), map[string]interface{}{
	"port": jsii.Number(443),
	"certificates": []interface{}{
		certificate,
	},
	"defaultAction": actions.NewAuthenticateCognitoAction(map[string]interface{}{
		"userPool": userPool,
		"userPoolClient": userPoolClient,
		"userPoolDomain": userPoolDomain,
		"next": elbv2.ListenerAction_fixedResponse(jsii.Number(200), map[string]*string{
			"contentType": jsii.String("text/plain"),
			"messageBody": jsii.String("Authenticated"),
		}),
	}),
})

awscdk.NewCfnOutput(this, jsii.String("DNS"), &cfnOutputProps{
	value: lb.loadBalancerDnsName,
})
```
