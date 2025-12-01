package previewawselasticloadbalancingv2mixins


// Specifies an action for a listener rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   actionProperty := &ActionProperty{
//   	AuthenticateCognitoConfig: &AuthenticateCognitoConfigProperty{
//   		AuthenticationRequestExtraParams: map[string]*string{
//   			"authenticationRequestExtraParamsKey": jsii.String("authenticationRequestExtraParams"),
//   		},
//   		OnUnauthenticatedRequest: jsii.String("onUnauthenticatedRequest"),
//   		Scope: jsii.String("scope"),
//   		SessionCookieName: jsii.String("sessionCookieName"),
//   		SessionTimeout: jsii.Number(123),
//   		UserPoolArn: jsii.String("userPoolArn"),
//   		UserPoolClientId: jsii.String("userPoolClientId"),
//   		UserPoolDomain: jsii.String("userPoolDomain"),
//   	},
//   	AuthenticateOidcConfig: &AuthenticateOidcConfigProperty{
//   		AuthenticationRequestExtraParams: map[string]*string{
//   			"authenticationRequestExtraParamsKey": jsii.String("authenticationRequestExtraParams"),
//   		},
//   		AuthorizationEndpoint: jsii.String("authorizationEndpoint"),
//   		ClientId: jsii.String("clientId"),
//   		ClientSecret: jsii.String("clientSecret"),
//   		Issuer: jsii.String("issuer"),
//   		OnUnauthenticatedRequest: jsii.String("onUnauthenticatedRequest"),
//   		Scope: jsii.String("scope"),
//   		SessionCookieName: jsii.String("sessionCookieName"),
//   		SessionTimeout: jsii.Number(123),
//   		TokenEndpoint: jsii.String("tokenEndpoint"),
//   		UseExistingClientSecret: jsii.Boolean(false),
//   		UserInfoEndpoint: jsii.String("userInfoEndpoint"),
//   	},
//   	FixedResponseConfig: &FixedResponseConfigProperty{
//   		ContentType: jsii.String("contentType"),
//   		MessageBody: jsii.String("messageBody"),
//   		StatusCode: jsii.String("statusCode"),
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
//   	JwtValidationConfig: &JwtValidationConfigProperty{
//   		AdditionalClaims: []interface{}{
//   			&JwtValidationActionAdditionalClaimProperty{
//   				Format: jsii.String("format"),
//   				Name: jsii.String("name"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   		Issuer: jsii.String("issuer"),
//   		JwksEndpoint: jsii.String("jwksEndpoint"),
//   	},
//   	Order: jsii.Number(123),
//   	RedirectConfig: &RedirectConfigProperty{
//   		Host: jsii.String("host"),
//   		Path: jsii.String("path"),
//   		Port: jsii.String("port"),
//   		Protocol: jsii.String("protocol"),
//   		Query: jsii.String("query"),
//   		StatusCode: jsii.String("statusCode"),
//   	},
//   	TargetGroupArn: jsii.String("targetGroupArn"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-action.html
//
type CfnListenerRulePropsMixin_ActionProperty struct {
	// [HTTPS listeners] Information for using Amazon Cognito to authenticate users.
	//
	// Specify only when `Type` is `authenticate-cognito` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-action.html#cfn-elasticloadbalancingv2-listenerrule-action-authenticatecognitoconfig
	//
	AuthenticateCognitoConfig interface{} `field:"optional" json:"authenticateCognitoConfig" yaml:"authenticateCognitoConfig"`
	// [HTTPS listeners] Information about an identity provider that is compliant with OpenID Connect (OIDC).
	//
	// Specify only when `Type` is `authenticate-oidc` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-action.html#cfn-elasticloadbalancingv2-listenerrule-action-authenticateoidcconfig
	//
	AuthenticateOidcConfig interface{} `field:"optional" json:"authenticateOidcConfig" yaml:"authenticateOidcConfig"`
	// [Application Load Balancer] Information for creating an action that returns a custom HTTP response.
	//
	// Specify only when `Type` is `fixed-response` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-action.html#cfn-elasticloadbalancingv2-listenerrule-action-fixedresponseconfig
	//
	FixedResponseConfig interface{} `field:"optional" json:"fixedResponseConfig" yaml:"fixedResponseConfig"`
	// Information for creating an action that distributes requests among multiple target groups. Specify only when `Type` is `forward` .
	//
	// If you specify both `ForwardConfig` and `TargetGroupArn` , you can specify only one target group using `ForwardConfig` and it must be the same target group specified in `TargetGroupArn` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-action.html#cfn-elasticloadbalancingv2-listenerrule-action-forwardconfig
	//
	ForwardConfig interface{} `field:"optional" json:"forwardConfig" yaml:"forwardConfig"`
	// [HTTPS listeners] Information for validating JWT access tokens in client requests.
	//
	// Specify only when `Type` is `jwt-validation` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-action.html#cfn-elasticloadbalancingv2-listenerrule-action-jwtvalidationconfig
	//
	JwtValidationConfig interface{} `field:"optional" json:"jwtValidationConfig" yaml:"jwtValidationConfig"`
	// The order for the action.
	//
	// This value is required for rules with multiple actions. The action with the lowest value for order is performed first.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-action.html#cfn-elasticloadbalancingv2-listenerrule-action-order
	//
	Order *float64 `field:"optional" json:"order" yaml:"order"`
	// [Application Load Balancer] Information for creating a redirect action.
	//
	// Specify only when `Type` is `redirect` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-action.html#cfn-elasticloadbalancingv2-listenerrule-action-redirectconfig
	//
	RedirectConfig interface{} `field:"optional" json:"redirectConfig" yaml:"redirectConfig"`
	// The Amazon Resource Name (ARN) of the target group.
	//
	// Specify only when `Type` is `forward` and you want to route to a single target group. To route to multiple target groups, you must use `ForwardConfig` instead.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-action.html#cfn-elasticloadbalancingv2-listenerrule-action-targetgrouparn
	//
	TargetGroupArn *string `field:"optional" json:"targetGroupArn" yaml:"targetGroupArn"`
	// The type of action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-action.html#cfn-elasticloadbalancingv2-listenerrule-action-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

