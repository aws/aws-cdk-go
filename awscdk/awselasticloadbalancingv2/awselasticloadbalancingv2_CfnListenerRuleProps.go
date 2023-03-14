package awselasticloadbalancingv2


// Properties for defining a `CfnListenerRule`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnListenerRuleProps := &cfnListenerRuleProps{
//   	actions: []interface{}{
//   		&actionProperty{
//   			type: jsii.String("type"),
//
//   			// the properties below are optional
//   			authenticateCognitoConfig: &authenticateCognitoConfigProperty{
//   				userPoolArn: jsii.String("userPoolArn"),
//   				userPoolClientId: jsii.String("userPoolClientId"),
//   				userPoolDomain: jsii.String("userPoolDomain"),
//
//   				// the properties below are optional
//   				authenticationRequestExtraParams: map[string]*string{
//   					"authenticationRequestExtraParamsKey": jsii.String("authenticationRequestExtraParams"),
//   				},
//   				onUnauthenticatedRequest: jsii.String("onUnauthenticatedRequest"),
//   				scope: jsii.String("scope"),
//   				sessionCookieName: jsii.String("sessionCookieName"),
//   				sessionTimeout: jsii.Number(123),
//   			},
//   			authenticateOidcConfig: &authenticateOidcConfigProperty{
//   				authorizationEndpoint: jsii.String("authorizationEndpoint"),
//   				clientId: jsii.String("clientId"),
//   				issuer: jsii.String("issuer"),
//   				tokenEndpoint: jsii.String("tokenEndpoint"),
//   				userInfoEndpoint: jsii.String("userInfoEndpoint"),
//
//   				// the properties below are optional
//   				authenticationRequestExtraParams: map[string]*string{
//   					"authenticationRequestExtraParamsKey": jsii.String("authenticationRequestExtraParams"),
//   				},
//   				clientSecret: jsii.String("clientSecret"),
//   				onUnauthenticatedRequest: jsii.String("onUnauthenticatedRequest"),
//   				scope: jsii.String("scope"),
//   				sessionCookieName: jsii.String("sessionCookieName"),
//   				sessionTimeout: jsii.Number(123),
//   				useExistingClientSecret: jsii.Boolean(false),
//   			},
//   			fixedResponseConfig: &fixedResponseConfigProperty{
//   				statusCode: jsii.String("statusCode"),
//
//   				// the properties below are optional
//   				contentType: jsii.String("contentType"),
//   				messageBody: jsii.String("messageBody"),
//   			},
//   			forwardConfig: &forwardConfigProperty{
//   				targetGroups: []interface{}{
//   					&targetGroupTupleProperty{
//   						targetGroupArn: jsii.String("targetGroupArn"),
//   						weight: jsii.Number(123),
//   					},
//   				},
//   				targetGroupStickinessConfig: &targetGroupStickinessConfigProperty{
//   					durationSeconds: jsii.Number(123),
//   					enabled: jsii.Boolean(false),
//   				},
//   			},
//   			order: jsii.Number(123),
//   			redirectConfig: &redirectConfigProperty{
//   				statusCode: jsii.String("statusCode"),
//
//   				// the properties below are optional
//   				host: jsii.String("host"),
//   				path: jsii.String("path"),
//   				port: jsii.String("port"),
//   				protocol: jsii.String("protocol"),
//   				query: jsii.String("query"),
//   			},
//   			targetGroupArn: jsii.String("targetGroupArn"),
//   		},
//   	},
//   	conditions: []interface{}{
//   		&ruleConditionProperty{
//   			field: jsii.String("field"),
//   			hostHeaderConfig: &hostHeaderConfigProperty{
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			httpHeaderConfig: &httpHeaderConfigProperty{
//   				httpHeaderName: jsii.String("httpHeaderName"),
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			httpRequestMethodConfig: &httpRequestMethodConfigProperty{
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			pathPatternConfig: &pathPatternConfigProperty{
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			queryStringConfig: &queryStringConfigProperty{
//   				values: []interface{}{
//   					&queryStringKeyValueProperty{
//   						key: jsii.String("key"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   			},
//   			sourceIpConfig: &sourceIpConfigProperty{
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   	listenerArn: jsii.String("listenerArn"),
//   	priority: jsii.Number(123),
//   }
//
type CfnListenerRuleProps struct {
	// The actions.
	//
	// The rule must include exactly one of the following types of actions: `forward` , `fixed-response` , or `redirect` , and it must be the last action to be performed. If the rule is for an HTTPS listener, it can also optionally include an authentication action.
	Actions interface{} `field:"required" json:"actions" yaml:"actions"`
	// The conditions.
	//
	// The rule can optionally include up to one of each of the following conditions: `http-request-method` , `host-header` , `path-pattern` , and `source-ip` . A rule can also optionally include one or more of each of the following conditions: `http-header` and `query-string` .
	Conditions interface{} `field:"required" json:"conditions" yaml:"conditions"`
	// The Amazon Resource Name (ARN) of the listener.
	ListenerArn *string `field:"required" json:"listenerArn" yaml:"listenerArn"`
	// The rule priority. A listener can't have multiple rules with the same priority.
	//
	// If you try to reorder rules by updating their priorities, do not specify a new priority if an existing rule already uses this priority, as this can cause an error. If you need to reuse a priority with a different rule, you must remove it as a priority first, and then specify it in a subsequent update.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
}

