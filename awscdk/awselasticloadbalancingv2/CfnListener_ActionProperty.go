package awselasticloadbalancingv2


// Specifies an action for a listener rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   actionProperty := &ActionProperty{
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	AuthenticateCognitoConfig: &AuthenticateCognitoConfigProperty{
//   		UserPoolArn: jsii.String("userPoolArn"),
//   		UserPoolClientId: jsii.String("userPoolClientId"),
//   		UserPoolDomain: jsii.String("userPoolDomain"),
//
//   		// the properties below are optional
//   		AuthenticationRequestExtraParams: map[string]*string{
//   			"authenticationRequestExtraParamsKey": jsii.String("authenticationRequestExtraParams"),
//   		},
//   		OnUnauthenticatedRequest: jsii.String("onUnauthenticatedRequest"),
//   		Scope: jsii.String("scope"),
//   		SessionCookieName: jsii.String("sessionCookieName"),
//   		SessionTimeout: jsii.String("sessionTimeout"),
//   	},
//   	AuthenticateOidcConfig: &AuthenticateOidcConfigProperty{
//   		AuthorizationEndpoint: jsii.String("authorizationEndpoint"),
//   		ClientId: jsii.String("clientId"),
//   		Issuer: jsii.String("issuer"),
//   		TokenEndpoint: jsii.String("tokenEndpoint"),
//   		UserInfoEndpoint: jsii.String("userInfoEndpoint"),
//
//   		// the properties below are optional
//   		AuthenticationRequestExtraParams: map[string]*string{
//   			"authenticationRequestExtraParamsKey": jsii.String("authenticationRequestExtraParams"),
//   		},
//   		ClientSecret: jsii.String("clientSecret"),
//   		OnUnauthenticatedRequest: jsii.String("onUnauthenticatedRequest"),
//   		Scope: jsii.String("scope"),
//   		SessionCookieName: jsii.String("sessionCookieName"),
//   		SessionTimeout: jsii.String("sessionTimeout"),
//   		UseExistingClientSecret: jsii.Boolean(false),
//   	},
//   	FixedResponseConfig: &FixedResponseConfigProperty{
//   		StatusCode: jsii.String("statusCode"),
//
//   		// the properties below are optional
//   		ContentType: jsii.String("contentType"),
//   		MessageBody: jsii.String("messageBody"),
//   	},
//   	ForwardConfig: &ForwardConfigProperty{
//   		TargetGroups: []interface{}{
//   			&TargetGroupTupleProperty{
//   				TargetGroupArn: jsii.String("targetGroupArn"),
//   				Weight: jsii.Number(123),
//   			},
//   		},
//   		TargetGroupStickinessConfig: &TargetGroupStickinessConfigProperty{
//   			DurationSeconds: jsii.Number(123),
//   			Enabled: jsii.Boolean(false),
//   		},
//   	},
//   	Order: jsii.Number(123),
//   	RedirectConfig: &RedirectConfigProperty{
//   		StatusCode: jsii.String("statusCode"),
//
//   		// the properties below are optional
//   		Host: jsii.String("host"),
//   		Path: jsii.String("path"),
//   		Port: jsii.String("port"),
//   		Protocol: jsii.String("protocol"),
//   		Query: jsii.String("query"),
//   	},
//   	TargetGroupArn: jsii.String("targetGroupArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-action.html
//
type CfnListener_ActionProperty struct {
	// The type of action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-action.html#cfn-elasticloadbalancingv2-listener-action-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// [HTTPS listeners] Information for using Amazon Cognito to authenticate users.
	//
	// Specify only when `Type` is `authenticate-cognito` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-action.html#cfn-elasticloadbalancingv2-listener-action-authenticatecognitoconfig
	//
	AuthenticateCognitoConfig interface{} `field:"optional" json:"authenticateCognitoConfig" yaml:"authenticateCognitoConfig"`
	// [HTTPS listeners] Information about an identity provider that is compliant with OpenID Connect (OIDC).
	//
	// Specify only when `Type` is `authenticate-oidc` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-action.html#cfn-elasticloadbalancingv2-listener-action-authenticateoidcconfig
	//
	AuthenticateOidcConfig interface{} `field:"optional" json:"authenticateOidcConfig" yaml:"authenticateOidcConfig"`
	// [Application Load Balancer] Information for creating an action that returns a custom HTTP response.
	//
	// Specify only when `Type` is `fixed-response` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-action.html#cfn-elasticloadbalancingv2-listener-action-fixedresponseconfig
	//
	FixedResponseConfig interface{} `field:"optional" json:"fixedResponseConfig" yaml:"fixedResponseConfig"`
	// Information for creating an action that distributes requests among one or more target groups.
	//
	// For Network Load Balancers, you can specify a single target group. Specify only when `Type` is `forward` . If you specify both `ForwardConfig` and `TargetGroupArn` , you can specify only one target group using `ForwardConfig` and it must be the same target group specified in `TargetGroupArn` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-action.html#cfn-elasticloadbalancingv2-listener-action-forwardconfig
	//
	ForwardConfig interface{} `field:"optional" json:"forwardConfig" yaml:"forwardConfig"`
	// The order for the action.
	//
	// This value is required for rules with multiple actions. The action with the lowest value for order is performed first.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-action.html#cfn-elasticloadbalancingv2-listener-action-order
	//
	Order *float64 `field:"optional" json:"order" yaml:"order"`
	// [Application Load Balancer] Information for creating a redirect action.
	//
	// Specify only when `Type` is `redirect` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-action.html#cfn-elasticloadbalancingv2-listener-action-redirectconfig
	//
	RedirectConfig interface{} `field:"optional" json:"redirectConfig" yaml:"redirectConfig"`
	// The Amazon Resource Name (ARN) of the target group.
	//
	// Specify only when `Type` is `forward` and you want to route to a single target group. To route to one or more target groups, use `ForwardConfig` instead.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-action.html#cfn-elasticloadbalancingv2-listener-action-targetgrouparn
	//
	TargetGroupArn *string `field:"optional" json:"targetGroupArn" yaml:"targetGroupArn"`
}

