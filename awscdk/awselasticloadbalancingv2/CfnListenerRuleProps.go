package awselasticloadbalancingv2


// Properties for defining a `CfnListenerRule`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnListenerRuleProps := &CfnListenerRuleProps{
//   	Actions: []interface{}{
//   		&ActionProperty{
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			AuthenticateCognitoConfig: &AuthenticateCognitoConfigProperty{
//   				UserPoolArn: jsii.String("userPoolArn"),
//   				UserPoolClientId: jsii.String("userPoolClientId"),
//   				UserPoolDomain: jsii.String("userPoolDomain"),
//
//   				// the properties below are optional
//   				AuthenticationRequestExtraParams: map[string]*string{
//   					"authenticationRequestExtraParamsKey": jsii.String("authenticationRequestExtraParams"),
//   				},
//   				OnUnauthenticatedRequest: jsii.String("onUnauthenticatedRequest"),
//   				Scope: jsii.String("scope"),
//   				SessionCookieName: jsii.String("sessionCookieName"),
//   				SessionTimeout: jsii.Number(123),
//   			},
//   			AuthenticateOidcConfig: &AuthenticateOidcConfigProperty{
//   				AuthorizationEndpoint: jsii.String("authorizationEndpoint"),
//   				ClientId: jsii.String("clientId"),
//   				Issuer: jsii.String("issuer"),
//   				TokenEndpoint: jsii.String("tokenEndpoint"),
//   				UserInfoEndpoint: jsii.String("userInfoEndpoint"),
//
//   				// the properties below are optional
//   				AuthenticationRequestExtraParams: map[string]*string{
//   					"authenticationRequestExtraParamsKey": jsii.String("authenticationRequestExtraParams"),
//   				},
//   				ClientSecret: jsii.String("clientSecret"),
//   				OnUnauthenticatedRequest: jsii.String("onUnauthenticatedRequest"),
//   				Scope: jsii.String("scope"),
//   				SessionCookieName: jsii.String("sessionCookieName"),
//   				SessionTimeout: jsii.Number(123),
//   				UseExistingClientSecret: jsii.Boolean(false),
//   			},
//   			FixedResponseConfig: &FixedResponseConfigProperty{
//   				StatusCode: jsii.String("statusCode"),
//
//   				// the properties below are optional
//   				ContentType: jsii.String("contentType"),
//   				MessageBody: jsii.String("messageBody"),
//   			},
//   			ForwardConfig: &ForwardConfigProperty{
//   				TargetGroups: []interface{}{
//   					&TargetGroupTupleProperty{
//   						TargetGroupArn: jsii.String("targetGroupArn"),
//   						Weight: jsii.Number(123),
//   					},
//   				},
//   				TargetGroupStickinessConfig: &TargetGroupStickinessConfigProperty{
//   					DurationSeconds: jsii.Number(123),
//   					Enabled: jsii.Boolean(false),
//   				},
//   			},
//   			Order: jsii.Number(123),
//   			RedirectConfig: &RedirectConfigProperty{
//   				StatusCode: jsii.String("statusCode"),
//
//   				// the properties below are optional
//   				Host: jsii.String("host"),
//   				Path: jsii.String("path"),
//   				Port: jsii.String("port"),
//   				Protocol: jsii.String("protocol"),
//   				Query: jsii.String("query"),
//   			},
//   			TargetGroupArn: jsii.String("targetGroupArn"),
//   		},
//   	},
//   	Conditions: []interface{}{
//   		&RuleConditionProperty{
//   			Field: jsii.String("field"),
//   			HostHeaderConfig: &HostHeaderConfigProperty{
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			HttpHeaderConfig: &HttpHeaderConfigProperty{
//   				HttpHeaderName: jsii.String("httpHeaderName"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			HttpRequestMethodConfig: &HttpRequestMethodConfigProperty{
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			PathPatternConfig: &PathPatternConfigProperty{
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			QueryStringConfig: &QueryStringConfigProperty{
//   				Values: []interface{}{
//   					&QueryStringKeyValueProperty{
//   						Key: jsii.String("key"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   			},
//   			SourceIpConfig: &SourceIpConfigProperty{
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   	Priority: jsii.Number(123),
//
//   	// the properties below are optional
//   	ListenerArn: jsii.String("listenerArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listenerrule.html
//
type CfnListenerRuleProps struct {
	// The actions.
	//
	// The rule must include exactly one of the following types of actions: `forward` , `fixed-response` , or `redirect` , and it must be the last action to be performed. If the rule is for an HTTPS listener, it can also optionally include an authentication action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listenerrule.html#cfn-elasticloadbalancingv2-listenerrule-actions
	//
	Actions interface{} `field:"required" json:"actions" yaml:"actions"`
	// The conditions.
	//
	// The rule can optionally include up to one of each of the following conditions: `http-request-method` , `host-header` , `path-pattern` , and `source-ip` . A rule can also optionally include one or more of each of the following conditions: `http-header` and `query-string` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listenerrule.html#cfn-elasticloadbalancingv2-listenerrule-conditions
	//
	Conditions interface{} `field:"required" json:"conditions" yaml:"conditions"`
	// The rule priority. A listener can't have multiple rules with the same priority.
	//
	// If you try to reorder rules by updating their priorities, do not specify a new priority if an existing rule already uses this priority, as this can cause an error. If you need to reuse a priority with a different rule, you must remove it as a priority first, and then specify it in a subsequent update.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listenerrule.html#cfn-elasticloadbalancingv2-listenerrule-priority
	//
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// The Amazon Resource Name (ARN) of the listener.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listenerrule.html#cfn-elasticloadbalancingv2-listenerrule-listenerarn
	//
	ListenerArn *string `field:"optional" json:"listenerArn" yaml:"listenerArn"`
}

