package mixinsawswafv2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnRuleGroupPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var all interface{}
//   var allow interface{}
//   var allQueryArguments interface{}
//   var asn interface{}
//   var block interface{}
//   var captcha interface{}
//   var challenge interface{}
//   var count interface{}
//   var forwardedIp interface{}
//   var httpMethod interface{}
//   var ip interface{}
//   var method interface{}
//   var queryString interface{}
//   var singleHeader interface{}
//   var singleQueryArgument interface{}
//   var statementProperty_ StatementProperty
//   var uriPath interface{}
//
//   cfnRuleGroupMixinProps := &CfnRuleGroupMixinProps{
//   	AvailableLabels: []interface{}{
//   		&LabelSummaryProperty{
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	Capacity: jsii.Number(123),
//   	ConsumedLabels: []interface{}{
//   		&LabelSummaryProperty{
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	CustomResponseBodies: map[string]interface{}{
//   		"customResponseBodiesKey": &CustomResponseBodyProperty{
//   			"content": jsii.String("content"),
//   			"contentType": jsii.String("contentType"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	Rules: []interface{}{
//   		&RuleProperty{
//   			Action: &RuleActionProperty{
//   				Allow: allow,
//   				Block: block,
//   				Captcha: captcha,
//   				Challenge: challenge,
//   				Count: count,
//   			},
//   			CaptchaConfig: &CaptchaConfigProperty{
//   				ImmunityTimeProperty: &ImmunityTimePropertyProperty{
//   					ImmunityTime: jsii.Number(123),
//   				},
//   			},
//   			ChallengeConfig: &ChallengeConfigProperty{
//   				ImmunityTimeProperty: &ImmunityTimePropertyProperty{
//   					ImmunityTime: jsii.Number(123),
//   				},
//   			},
//   			Name: jsii.String("name"),
//   			Priority: jsii.Number(123),
//   			RuleLabels: []interface{}{
//   				&LabelProperty{
//   					Name: jsii.String("name"),
//   				},
//   			},
//   			Statement: &StatementProperty{
//   				AndStatement: &AndStatementProperty{
//   					Statements: []interface{}{
//   						statementProperty_,
//   					},
//   				},
//   				AsnMatchStatement: &AsnMatchStatementProperty{
//   					AsnList: []interface{}{
//   						jsii.Number(123),
//   					},
//   					ForwardedIpConfig: &ForwardedIPConfigurationProperty{
//   						FallbackBehavior: jsii.String("fallbackBehavior"),
//   						HeaderName: jsii.String("headerName"),
//   					},
//   				},
//   				ByteMatchStatement: &ByteMatchStatementProperty{
//   					FieldToMatch: &FieldToMatchProperty{
//   						AllQueryArguments: allQueryArguments,
//   						Body: &BodyProperty{
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						Cookies: &CookiesProperty{
//   							MatchPattern: &CookieMatchPatternProperty{
//   								All: all,
//   								ExcludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								IncludedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							MatchScope: jsii.String("matchScope"),
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						Headers: &HeadersProperty{
//   							MatchPattern: &HeaderMatchPatternProperty{
//   								All: all,
//   								ExcludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								IncludedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							MatchScope: jsii.String("matchScope"),
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						Ja3Fingerprint: &JA3FingerprintProperty{
//   							FallbackBehavior: jsii.String("fallbackBehavior"),
//   						},
//   						Ja4Fingerprint: &JA4FingerprintProperty{
//   							FallbackBehavior: jsii.String("fallbackBehavior"),
//   						},
//   						JsonBody: &JsonBodyProperty{
//   							InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							MatchPattern: &JsonMatchPatternProperty{
//   								All: all,
//   								IncludedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							MatchScope: jsii.String("matchScope"),
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						Method: method,
//   						QueryString: queryString,
//   						SingleHeader: singleHeader,
//   						SingleQueryArgument: singleQueryArgument,
//   						UriFragment: &UriFragmentProperty{
//   							FallbackBehavior: jsii.String("fallbackBehavior"),
//   						},
//   						UriPath: uriPath,
//   					},
//   					PositionalConstraint: jsii.String("positionalConstraint"),
//   					SearchString: jsii.String("searchString"),
//   					SearchStringBase64: jsii.String("searchStringBase64"),
//   					TextTransformations: []interface{}{
//   						&TextTransformationProperty{
//   							Priority: jsii.Number(123),
//   							Type: jsii.String("type"),
//   						},
//   					},
//   				},
//   				GeoMatchStatement: &GeoMatchStatementProperty{
//   					CountryCodes: []*string{
//   						jsii.String("countryCodes"),
//   					},
//   					ForwardedIpConfig: &ForwardedIPConfigurationProperty{
//   						FallbackBehavior: jsii.String("fallbackBehavior"),
//   						HeaderName: jsii.String("headerName"),
//   					},
//   				},
//   				IpSetReferenceStatement: map[string]*string{
//   					"arn": jsii.String("arn"),
//   					"ipSetForwardedIpConfig": map[string]*string{
//   						"fallbackBehavior": jsii.String("fallbackBehavior"),
//   						"headerName": jsii.String("headerName"),
//   						"position": jsii.String("position"),
//   					},
//   				},
//   				LabelMatchStatement: &LabelMatchStatementProperty{
//   					Key: jsii.String("key"),
//   					Scope: jsii.String("scope"),
//   				},
//   				NotStatement: &NotStatementProperty{
//   					Statement: statementProperty_,
//   				},
//   				OrStatement: &OrStatementProperty{
//   					Statements: []interface{}{
//   						statementProperty_,
//   					},
//   				},
//   				RateBasedStatement: &RateBasedStatementProperty{
//   					AggregateKeyType: jsii.String("aggregateKeyType"),
//   					CustomKeys: []interface{}{
//   						&RateBasedStatementCustomKeyProperty{
//   							Asn: asn,
//   							Cookie: &RateLimitCookieProperty{
//   								Name: jsii.String("name"),
//   								TextTransformations: []interface{}{
//   									&TextTransformationProperty{
//   										Priority: jsii.Number(123),
//   										Type: jsii.String("type"),
//   									},
//   								},
//   							},
//   							ForwardedIp: forwardedIp,
//   							Header: &RateLimitHeaderProperty{
//   								Name: jsii.String("name"),
//   								TextTransformations: []interface{}{
//   									&TextTransformationProperty{
//   										Priority: jsii.Number(123),
//   										Type: jsii.String("type"),
//   									},
//   								},
//   							},
//   							HttpMethod: httpMethod,
//   							Ip: ip,
//   							Ja3Fingerprint: &RateLimitJA3FingerprintProperty{
//   								FallbackBehavior: jsii.String("fallbackBehavior"),
//   							},
//   							Ja4Fingerprint: &RateLimitJA4FingerprintProperty{
//   								FallbackBehavior: jsii.String("fallbackBehavior"),
//   							},
//   							LabelNamespace: &RateLimitLabelNamespaceProperty{
//   								Namespace: jsii.String("namespace"),
//   							},
//   							QueryArgument: &RateLimitQueryArgumentProperty{
//   								Name: jsii.String("name"),
//   								TextTransformations: []interface{}{
//   									&TextTransformationProperty{
//   										Priority: jsii.Number(123),
//   										Type: jsii.String("type"),
//   									},
//   								},
//   							},
//   							QueryString: &RateLimitQueryStringProperty{
//   								TextTransformations: []interface{}{
//   									&TextTransformationProperty{
//   										Priority: jsii.Number(123),
//   										Type: jsii.String("type"),
//   									},
//   								},
//   							},
//   							UriPath: &RateLimitUriPathProperty{
//   								TextTransformations: []interface{}{
//   									&TextTransformationProperty{
//   										Priority: jsii.Number(123),
//   										Type: jsii.String("type"),
//   									},
//   								},
//   							},
//   						},
//   					},
//   					EvaluationWindowSec: jsii.Number(123),
//   					ForwardedIpConfig: &ForwardedIPConfigurationProperty{
//   						FallbackBehavior: jsii.String("fallbackBehavior"),
//   						HeaderName: jsii.String("headerName"),
//   					},
//   					Limit: jsii.Number(123),
//   					ScopeDownStatement: statementProperty_,
//   				},
//   				RegexMatchStatement: &RegexMatchStatementProperty{
//   					FieldToMatch: &FieldToMatchProperty{
//   						AllQueryArguments: allQueryArguments,
//   						Body: &BodyProperty{
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						Cookies: &CookiesProperty{
//   							MatchPattern: &CookieMatchPatternProperty{
//   								All: all,
//   								ExcludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								IncludedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							MatchScope: jsii.String("matchScope"),
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						Headers: &HeadersProperty{
//   							MatchPattern: &HeaderMatchPatternProperty{
//   								All: all,
//   								ExcludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								IncludedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							MatchScope: jsii.String("matchScope"),
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						Ja3Fingerprint: &JA3FingerprintProperty{
//   							FallbackBehavior: jsii.String("fallbackBehavior"),
//   						},
//   						Ja4Fingerprint: &JA4FingerprintProperty{
//   							FallbackBehavior: jsii.String("fallbackBehavior"),
//   						},
//   						JsonBody: &JsonBodyProperty{
//   							InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							MatchPattern: &JsonMatchPatternProperty{
//   								All: all,
//   								IncludedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							MatchScope: jsii.String("matchScope"),
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						Method: method,
//   						QueryString: queryString,
//   						SingleHeader: singleHeader,
//   						SingleQueryArgument: singleQueryArgument,
//   						UriFragment: &UriFragmentProperty{
//   							FallbackBehavior: jsii.String("fallbackBehavior"),
//   						},
//   						UriPath: uriPath,
//   					},
//   					RegexString: jsii.String("regexString"),
//   					TextTransformations: []interface{}{
//   						&TextTransformationProperty{
//   							Priority: jsii.Number(123),
//   							Type: jsii.String("type"),
//   						},
//   					},
//   				},
//   				RegexPatternSetReferenceStatement: &RegexPatternSetReferenceStatementProperty{
//   					Arn: jsii.String("arn"),
//   					FieldToMatch: &FieldToMatchProperty{
//   						AllQueryArguments: allQueryArguments,
//   						Body: &BodyProperty{
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						Cookies: &CookiesProperty{
//   							MatchPattern: &CookieMatchPatternProperty{
//   								All: all,
//   								ExcludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								IncludedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							MatchScope: jsii.String("matchScope"),
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						Headers: &HeadersProperty{
//   							MatchPattern: &HeaderMatchPatternProperty{
//   								All: all,
//   								ExcludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								IncludedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							MatchScope: jsii.String("matchScope"),
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						Ja3Fingerprint: &JA3FingerprintProperty{
//   							FallbackBehavior: jsii.String("fallbackBehavior"),
//   						},
//   						Ja4Fingerprint: &JA4FingerprintProperty{
//   							FallbackBehavior: jsii.String("fallbackBehavior"),
//   						},
//   						JsonBody: &JsonBodyProperty{
//   							InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							MatchPattern: &JsonMatchPatternProperty{
//   								All: all,
//   								IncludedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							MatchScope: jsii.String("matchScope"),
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						Method: method,
//   						QueryString: queryString,
//   						SingleHeader: singleHeader,
//   						SingleQueryArgument: singleQueryArgument,
//   						UriFragment: &UriFragmentProperty{
//   							FallbackBehavior: jsii.String("fallbackBehavior"),
//   						},
//   						UriPath: uriPath,
//   					},
//   					TextTransformations: []interface{}{
//   						&TextTransformationProperty{
//   							Priority: jsii.Number(123),
//   							Type: jsii.String("type"),
//   						},
//   					},
//   				},
//   				SizeConstraintStatement: &SizeConstraintStatementProperty{
//   					ComparisonOperator: jsii.String("comparisonOperator"),
//   					FieldToMatch: &FieldToMatchProperty{
//   						AllQueryArguments: allQueryArguments,
//   						Body: &BodyProperty{
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						Cookies: &CookiesProperty{
//   							MatchPattern: &CookieMatchPatternProperty{
//   								All: all,
//   								ExcludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								IncludedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							MatchScope: jsii.String("matchScope"),
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						Headers: &HeadersProperty{
//   							MatchPattern: &HeaderMatchPatternProperty{
//   								All: all,
//   								ExcludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								IncludedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							MatchScope: jsii.String("matchScope"),
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						Ja3Fingerprint: &JA3FingerprintProperty{
//   							FallbackBehavior: jsii.String("fallbackBehavior"),
//   						},
//   						Ja4Fingerprint: &JA4FingerprintProperty{
//   							FallbackBehavior: jsii.String("fallbackBehavior"),
//   						},
//   						JsonBody: &JsonBodyProperty{
//   							InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							MatchPattern: &JsonMatchPatternProperty{
//   								All: all,
//   								IncludedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							MatchScope: jsii.String("matchScope"),
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						Method: method,
//   						QueryString: queryString,
//   						SingleHeader: singleHeader,
//   						SingleQueryArgument: singleQueryArgument,
//   						UriFragment: &UriFragmentProperty{
//   							FallbackBehavior: jsii.String("fallbackBehavior"),
//   						},
//   						UriPath: uriPath,
//   					},
//   					Size: jsii.Number(123),
//   					TextTransformations: []interface{}{
//   						&TextTransformationProperty{
//   							Priority: jsii.Number(123),
//   							Type: jsii.String("type"),
//   						},
//   					},
//   				},
//   				SqliMatchStatement: &SqliMatchStatementProperty{
//   					FieldToMatch: &FieldToMatchProperty{
//   						AllQueryArguments: allQueryArguments,
//   						Body: &BodyProperty{
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						Cookies: &CookiesProperty{
//   							MatchPattern: &CookieMatchPatternProperty{
//   								All: all,
//   								ExcludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								IncludedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							MatchScope: jsii.String("matchScope"),
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						Headers: &HeadersProperty{
//   							MatchPattern: &HeaderMatchPatternProperty{
//   								All: all,
//   								ExcludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								IncludedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							MatchScope: jsii.String("matchScope"),
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						Ja3Fingerprint: &JA3FingerprintProperty{
//   							FallbackBehavior: jsii.String("fallbackBehavior"),
//   						},
//   						Ja4Fingerprint: &JA4FingerprintProperty{
//   							FallbackBehavior: jsii.String("fallbackBehavior"),
//   						},
//   						JsonBody: &JsonBodyProperty{
//   							InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							MatchPattern: &JsonMatchPatternProperty{
//   								All: all,
//   								IncludedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							MatchScope: jsii.String("matchScope"),
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						Method: method,
//   						QueryString: queryString,
//   						SingleHeader: singleHeader,
//   						SingleQueryArgument: singleQueryArgument,
//   						UriFragment: &UriFragmentProperty{
//   							FallbackBehavior: jsii.String("fallbackBehavior"),
//   						},
//   						UriPath: uriPath,
//   					},
//   					SensitivityLevel: jsii.String("sensitivityLevel"),
//   					TextTransformations: []interface{}{
//   						&TextTransformationProperty{
//   							Priority: jsii.Number(123),
//   							Type: jsii.String("type"),
//   						},
//   					},
//   				},
//   				XssMatchStatement: &XssMatchStatementProperty{
//   					FieldToMatch: &FieldToMatchProperty{
//   						AllQueryArguments: allQueryArguments,
//   						Body: &BodyProperty{
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						Cookies: &CookiesProperty{
//   							MatchPattern: &CookieMatchPatternProperty{
//   								All: all,
//   								ExcludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								IncludedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							MatchScope: jsii.String("matchScope"),
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						Headers: &HeadersProperty{
//   							MatchPattern: &HeaderMatchPatternProperty{
//   								All: all,
//   								ExcludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								IncludedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							MatchScope: jsii.String("matchScope"),
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						Ja3Fingerprint: &JA3FingerprintProperty{
//   							FallbackBehavior: jsii.String("fallbackBehavior"),
//   						},
//   						Ja4Fingerprint: &JA4FingerprintProperty{
//   							FallbackBehavior: jsii.String("fallbackBehavior"),
//   						},
//   						JsonBody: &JsonBodyProperty{
//   							InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							MatchPattern: &JsonMatchPatternProperty{
//   								All: all,
//   								IncludedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							MatchScope: jsii.String("matchScope"),
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						Method: method,
//   						QueryString: queryString,
//   						SingleHeader: singleHeader,
//   						SingleQueryArgument: singleQueryArgument,
//   						UriFragment: &UriFragmentProperty{
//   							FallbackBehavior: jsii.String("fallbackBehavior"),
//   						},
//   						UriPath: uriPath,
//   					},
//   					TextTransformations: []interface{}{
//   						&TextTransformationProperty{
//   							Priority: jsii.Number(123),
//   							Type: jsii.String("type"),
//   						},
//   					},
//   				},
//   			},
//   			VisibilityConfig: &VisibilityConfigProperty{
//   				CloudWatchMetricsEnabled: jsii.Boolean(false),
//   				MetricName: jsii.String("metricName"),
//   				SampledRequestsEnabled: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	Scope: jsii.String("scope"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VisibilityConfig: &VisibilityConfigProperty{
//   		CloudWatchMetricsEnabled: jsii.Boolean(false),
//   		MetricName: jsii.String("metricName"),
//   		SampledRequestsEnabled: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-rulegroup.html
//
type CfnRuleGroupMixinProps struct {
	// The labels that one or more rules in this rule group add to matching web requests.
	//
	// These labels are defined in the `RuleLabels` for a `Rule` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-rulegroup.html#cfn-wafv2-rulegroup-availablelabels
	//
	AvailableLabels interface{} `field:"optional" json:"availableLabels" yaml:"availableLabels"`
	// The web ACL capacity units (WCUs) required for this rule group.
	//
	// When you create your own rule group, you define this, and you cannot change it after creation. When you add or modify the rules in a rule group, AWS WAF enforces this limit.
	//
	// AWS WAF uses WCUs to calculate and control the operating resources that are used to run your rules, rule groups, and web ACLs. AWS WAF calculates capacity differently for each rule type, to reflect the relative cost of each rule. Simple rules that cost little to run use fewer WCUs than more complex rules that use more processing power. Rule group capacity is fixed at creation, which helps users plan their web ACL WCU usage when they use a rule group. The WCU limit for web ACLs is 1,500.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-rulegroup.html#cfn-wafv2-rulegroup-capacity
	//
	Capacity *float64 `field:"optional" json:"capacity" yaml:"capacity"`
	// The labels that one or more rules in this rule group match against in label match statements.
	//
	// These labels are defined in a `LabelMatchStatement` specification, in the `Statement` definition of a rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-rulegroup.html#cfn-wafv2-rulegroup-consumedlabels
	//
	ConsumedLabels interface{} `field:"optional" json:"consumedLabels" yaml:"consumedLabels"`
	// A map of custom response keys and content bodies.
	//
	// When you create a rule with a block action, you can send a custom response to the web request. You define these for the rule group, and then use them in the rules that you define in the rule group.
	//
	// For information about customizing web requests and responses, see [Customizing web requests and responses in AWS WAF](https://docs.aws.amazon.com/waf/latest/developerguide/waf-custom-request-response.html) in the *AWS WAF Developer Guide* .
	//
	// For information about the limits on count and size for custom request and response settings, see [AWS WAF quotas](https://docs.aws.amazon.com/waf/latest/developerguide/limits.html) in the *AWS WAF Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-rulegroup.html#cfn-wafv2-rulegroup-customresponsebodies
	//
	CustomResponseBodies interface{} `field:"optional" json:"customResponseBodies" yaml:"customResponseBodies"`
	// A description of the rule group that helps with identification.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-rulegroup.html#cfn-wafv2-rulegroup-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the rule group.
	//
	// You cannot change the name of a rule group after you create it.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-rulegroup.html#cfn-wafv2-rulegroup-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The rule statements used to identify the web requests that you want to allow, block, or count.
	//
	// Each rule includes one top-level statement that AWS WAF uses to identify matching web requests, and parameters that govern how AWS WAF handles them.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-rulegroup.html#cfn-wafv2-rulegroup-rules
	//
	Rules interface{} `field:"optional" json:"rules" yaml:"rules"`
	// Specifies whether this is for an Amazon CloudFront distribution or for a regional application.
	//
	// For an AWS Amplify application, use `CLOUDFRONT` . A regional application can be an Application Load Balancer (ALB), an  REST API, an AWS AppSync GraphQL API, an Amazon Cognito user pool, an AWS App Runner service, or an AWS Verified Access instance. Valid Values are `CLOUDFRONT` and `REGIONAL` .
	//
	// > For `CLOUDFRONT` , you must create your WAFv2 resources in the US East (N. Virginia) Region, `us-east-1` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-rulegroup.html#cfn-wafv2-rulegroup-scope
	//
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// Key:value pairs associated with an AWS resource.
	//
	// The key:value pair can be anything you define. Typically, the tag key represents a category (such as "environment") and the tag value represents a specific value within that category (such as "test," "development," or "production"). You can add up to 50 tags to each AWS resource.
	//
	// > To modify tags on existing resources, use the AWS WAF APIs or command line interface. With AWS CloudFormation , you can only add tags to AWS WAF resources during resource creation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-rulegroup.html#cfn-wafv2-rulegroup-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Defines and enables Amazon CloudWatch metrics and web request sample collection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-rulegroup.html#cfn-wafv2-rulegroup-visibilityconfig
	//
	VisibilityConfig interface{} `field:"optional" json:"visibilityConfig" yaml:"visibilityConfig"`
}

