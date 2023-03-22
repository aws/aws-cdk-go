package awselasticloadbalancingv2


// Specifies an action for a listener rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   actionProperty := &actionProperty{
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	authenticateCognitoConfig: &authenticateCognitoConfigProperty{
//   		userPoolArn: jsii.String("userPoolArn"),
//   		userPoolClientId: jsii.String("userPoolClientId"),
//   		userPoolDomain: jsii.String("userPoolDomain"),
//
//   		// the properties below are optional
//   		authenticationRequestExtraParams: map[string]*string{
//   			"authenticationRequestExtraParamsKey": jsii.String("authenticationRequestExtraParams"),
//   		},
//   		onUnauthenticatedRequest: jsii.String("onUnauthenticatedRequest"),
//   		scope: jsii.String("scope"),
//   		sessionCookieName: jsii.String("sessionCookieName"),
//   		sessionTimeout: jsii.Number(123),
//   	},
//   	authenticateOidcConfig: &authenticateOidcConfigProperty{
//   		authorizationEndpoint: jsii.String("authorizationEndpoint"),
//   		clientId: jsii.String("clientId"),
//   		issuer: jsii.String("issuer"),
//   		tokenEndpoint: jsii.String("tokenEndpoint"),
//   		userInfoEndpoint: jsii.String("userInfoEndpoint"),
//
//   		// the properties below are optional
//   		authenticationRequestExtraParams: map[string]*string{
//   			"authenticationRequestExtraParamsKey": jsii.String("authenticationRequestExtraParams"),
//   		},
//   		clientSecret: jsii.String("clientSecret"),
//   		onUnauthenticatedRequest: jsii.String("onUnauthenticatedRequest"),
//   		scope: jsii.String("scope"),
//   		sessionCookieName: jsii.String("sessionCookieName"),
//   		sessionTimeout: jsii.Number(123),
//   		useExistingClientSecret: jsii.Boolean(false),
//   	},
//   	fixedResponseConfig: &fixedResponseConfigProperty{
//   		statusCode: jsii.String("statusCode"),
//
//   		// the properties below are optional
//   		contentType: jsii.String("contentType"),
//   		messageBody: jsii.String("messageBody"),
//   	},
//   	forwardConfig: &forwardConfigProperty{
//   		targetGroups: []interface{}{
//   			&targetGroupTupleProperty{
//   				targetGroupArn: jsii.String("targetGroupArn"),
//   				weight: jsii.Number(123),
//   			},
//   		},
//   		targetGroupStickinessConfig: &targetGroupStickinessConfigProperty{
//   			durationSeconds: jsii.Number(123),
//   			enabled: jsii.Boolean(false),
//   		},
//   	},
//   	order: jsii.Number(123),
//   	redirectConfig: &redirectConfigProperty{
//   		statusCode: jsii.String("statusCode"),
//
//   		// the properties below are optional
//   		host: jsii.String("host"),
//   		path: jsii.String("path"),
//   		port: jsii.String("port"),
//   		protocol: jsii.String("protocol"),
//   		query: jsii.String("query"),
//   	},
//   	targetGroupArn: jsii.String("targetGroupArn"),
//   }
//
type CfnListenerRule_ActionProperty struct {
	// The type of action.
	Type *string `field:"required" json:"type" yaml:"type"`
	// [HTTPS listeners] Information for using Amazon Cognito to authenticate users.
	//
	// Specify only when `Type` is `authenticate-cognito` .
	AuthenticateCognitoConfig interface{} `field:"optional" json:"authenticateCognitoConfig" yaml:"authenticateCognitoConfig"`
	// [HTTPS listeners] Information about an identity provider that is compliant with OpenID Connect (OIDC).
	//
	// Specify only when `Type` is `authenticate-oidc` .
	AuthenticateOidcConfig interface{} `field:"optional" json:"authenticateOidcConfig" yaml:"authenticateOidcConfig"`
	// [Application Load Balancer] Information for creating an action that returns a custom HTTP response.
	//
	// Specify only when `Type` is `fixed-response` .
	FixedResponseConfig interface{} `field:"optional" json:"fixedResponseConfig" yaml:"fixedResponseConfig"`
	// Information for creating an action that distributes requests among one or more target groups.
	//
	// For Network Load Balancers, you can specify a single target group. Specify only when `Type` is `forward` . If you specify both `ForwardConfig` and `TargetGroupArn` , you can specify only one target group using `ForwardConfig` and it must be the same target group specified in `TargetGroupArn` .
	ForwardConfig interface{} `field:"optional" json:"forwardConfig" yaml:"forwardConfig"`
	// The order for the action.
	//
	// This value is required for rules with multiple actions. The action with the lowest value for order is performed first.
	Order *float64 `field:"optional" json:"order" yaml:"order"`
	// [Application Load Balancer] Information for creating a redirect action.
	//
	// Specify only when `Type` is `redirect` .
	RedirectConfig interface{} `field:"optional" json:"redirectConfig" yaml:"redirectConfig"`
	// The Amazon Resource Name (ARN) of the target group.
	//
	// Specify only when `Type` is `forward` and you want to route to a single target group. To route to one or more target groups, use `ForwardConfig` instead.
	TargetGroupArn *string `field:"optional" json:"targetGroupArn" yaml:"targetGroupArn"`
}

