package previewawselasticloadbalancingv2mixins


// Properties for CfnListenerRulePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnListenerRuleMixinProps := &CfnListenerRuleMixinProps{
//   	Actions: []interface{}{
//   		&ActionProperty{
//   			AuthenticateCognitoConfig: &AuthenticateCognitoConfigProperty{
//   				AuthenticationRequestExtraParams: map[string]*string{
//   					"authenticationRequestExtraParamsKey": jsii.String("authenticationRequestExtraParams"),
//   				},
//   				OnUnauthenticatedRequest: jsii.String("onUnauthenticatedRequest"),
//   				Scope: jsii.String("scope"),
//   				SessionCookieName: jsii.String("sessionCookieName"),
//   				SessionTimeout: jsii.Number(123),
//   				UserPoolArn: jsii.String("userPoolArn"),
//   				UserPoolClientId: jsii.String("userPoolClientId"),
//   				UserPoolDomain: jsii.String("userPoolDomain"),
//   			},
//   			AuthenticateOidcConfig: &AuthenticateOidcConfigProperty{
//   				AuthenticationRequestExtraParams: map[string]*string{
//   					"authenticationRequestExtraParamsKey": jsii.String("authenticationRequestExtraParams"),
//   				},
//   				AuthorizationEndpoint: jsii.String("authorizationEndpoint"),
//   				ClientId: jsii.String("clientId"),
//   				ClientSecret: jsii.String("clientSecret"),
//   				Issuer: jsii.String("issuer"),
//   				OnUnauthenticatedRequest: jsii.String("onUnauthenticatedRequest"),
//   				Scope: jsii.String("scope"),
//   				SessionCookieName: jsii.String("sessionCookieName"),
//   				SessionTimeout: jsii.Number(123),
//   				TokenEndpoint: jsii.String("tokenEndpoint"),
//   				UseExistingClientSecret: jsii.Boolean(false),
//   				UserInfoEndpoint: jsii.String("userInfoEndpoint"),
//   			},
//   			FixedResponseConfig: &FixedResponseConfigProperty{
//   				ContentType: jsii.String("contentType"),
//   				MessageBody: jsii.String("messageBody"),
//   				StatusCode: jsii.String("statusCode"),
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
//   			JwtValidationConfig: &JwtValidationConfigProperty{
//   				AdditionalClaims: []interface{}{
//   					&JwtValidationActionAdditionalClaimProperty{
//   						Format: jsii.String("format"),
//   						Name: jsii.String("name"),
//   						Values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   				},
//   				Issuer: jsii.String("issuer"),
//   				JwksEndpoint: jsii.String("jwksEndpoint"),
//   			},
//   			Order: jsii.Number(123),
//   			RedirectConfig: &RedirectConfigProperty{
//   				Host: jsii.String("host"),
//   				Path: jsii.String("path"),
//   				Port: jsii.String("port"),
//   				Protocol: jsii.String("protocol"),
//   				Query: jsii.String("query"),
//   				StatusCode: jsii.String("statusCode"),
//   			},
//   			TargetGroupArn: jsii.String("targetGroupArn"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	Conditions: []interface{}{
//   		&RuleConditionProperty{
//   			Field: jsii.String("field"),
//   			HostHeaderConfig: &HostHeaderConfigProperty{
//   				RegexValues: []*string{
//   					jsii.String("regexValues"),
//   				},
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			HttpHeaderConfig: &HttpHeaderConfigProperty{
//   				HttpHeaderName: jsii.String("httpHeaderName"),
//   				RegexValues: []*string{
//   					jsii.String("regexValues"),
//   				},
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
//   				RegexValues: []*string{
//   					jsii.String("regexValues"),
//   				},
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
//   			RegexValues: []*string{
//   				jsii.String("regexValues"),
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
//   	ListenerArn: jsii.String("listenerArn"),
//   	Priority: jsii.Number(123),
//   	Transforms: []interface{}{
//   		&TransformProperty{
//   			HostHeaderRewriteConfig: &RewriteConfigObjectProperty{
//   				Rewrites: []interface{}{
//   					&RewriteConfigProperty{
//   						Regex: jsii.String("regex"),
//   						Replace: jsii.String("replace"),
//   					},
//   				},
//   			},
//   			Type: jsii.String("type"),
//   			UrlRewriteConfig: &RewriteConfigObjectProperty{
//   				Rewrites: []interface{}{
//   					&RewriteConfigProperty{
//   						Regex: jsii.String("regex"),
//   						Replace: jsii.String("replace"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listenerrule.html
//
type CfnListenerRuleMixinProps struct {
	// The actions.
	//
	// The rule must include exactly one of the following types of actions: `forward` , `fixed-response` , or `redirect` , and it must be the last action to be performed. If the rule is for an HTTPS listener, it can also optionally include an authentication action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listenerrule.html#cfn-elasticloadbalancingv2-listenerrule-actions
	//
	Actions interface{} `field:"optional" json:"actions" yaml:"actions"`
	// The conditions.
	//
	// The rule can optionally include up to one of each of the following conditions: `http-request-method` , `host-header` , `path-pattern` , and `source-ip` . A rule can also optionally include one or more of each of the following conditions: `http-header` and `query-string` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listenerrule.html#cfn-elasticloadbalancingv2-listenerrule-conditions
	//
	Conditions interface{} `field:"optional" json:"conditions" yaml:"conditions"`
	// The Amazon Resource Name (ARN) of the listener.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listenerrule.html#cfn-elasticloadbalancingv2-listenerrule-listenerarn
	//
	ListenerArn *string `field:"optional" json:"listenerArn" yaml:"listenerArn"`
	// The rule priority. A listener can't have multiple rules with the same priority.
	//
	// If you try to reorder rules by updating their priorities, do not specify a new priority if an existing rule already uses this priority, as this can cause an error. If you need to reuse a priority with a different rule, you must remove it as a priority first, and then specify it in a subsequent update.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listenerrule.html#cfn-elasticloadbalancingv2-listenerrule-priority
	//
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listenerrule.html#cfn-elasticloadbalancingv2-listenerrule-transforms
	//
	Transforms interface{} `field:"optional" json:"transforms" yaml:"transforms"`
}

