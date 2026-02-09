package previewawswafv2mixins


// A single rule, which you can use in a `WebACL` or `RuleGroup` to identify web requests that you want to manage in some way.
//
// Each rule includes one top-level `Statement` that AWS WAF uses to identify matching web requests, and parameters that govern how AWS WAF handles them.
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
//   ruleProperty := &RuleProperty{
//   	Action: &RuleActionProperty{
//   		Allow: allow,
//   		Block: block,
//   		Captcha: captcha,
//   		Challenge: challenge,
//   		Count: count,
//   	},
//   	CaptchaConfig: &CaptchaConfigProperty{
//   		ImmunityTimeProperty: &ImmunityTimePropertyProperty{
//   			ImmunityTime: jsii.Number(123),
//   		},
//   	},
//   	ChallengeConfig: &ChallengeConfigProperty{
//   		ImmunityTimeProperty: &ImmunityTimePropertyProperty{
//   			ImmunityTime: jsii.Number(123),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Priority: jsii.Number(123),
//   	RuleLabels: []interface{}{
//   		&LabelProperty{
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	Statement: &StatementProperty{
//   		AndStatement: &AndStatementProperty{
//   			Statements: []interface{}{
//   				statementProperty_,
//   			},
//   		},
//   		AsnMatchStatement: &AsnMatchStatementProperty{
//   			AsnList: []interface{}{
//   				jsii.Number(123),
//   			},
//   			ForwardedIpConfig: &ForwardedIPConfigurationProperty{
//   				FallbackBehavior: jsii.String("fallbackBehavior"),
//   				HeaderName: jsii.String("headerName"),
//   			},
//   		},
//   		ByteMatchStatement: &ByteMatchStatementProperty{
//   			FieldToMatch: &FieldToMatchProperty{
//   				AllQueryArguments: allQueryArguments,
//   				Body: &BodyProperty{
//   					OversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				Cookies: &CookiesProperty{
//   					MatchPattern: &CookieMatchPatternProperty{
//   						All: all,
//   						ExcludedCookies: []*string{
//   							jsii.String("excludedCookies"),
//   						},
//   						IncludedCookies: []*string{
//   							jsii.String("includedCookies"),
//   						},
//   					},
//   					MatchScope: jsii.String("matchScope"),
//   					OversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				HeaderOrder: &HeaderOrderProperty{
//   					OversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				Headers: &HeadersProperty{
//   					MatchPattern: &HeaderMatchPatternProperty{
//   						All: all,
//   						ExcludedHeaders: []*string{
//   							jsii.String("excludedHeaders"),
//   						},
//   						IncludedHeaders: []*string{
//   							jsii.String("includedHeaders"),
//   						},
//   					},
//   					MatchScope: jsii.String("matchScope"),
//   					OversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				Ja3Fingerprint: &JA3FingerprintProperty{
//   					FallbackBehavior: jsii.String("fallbackBehavior"),
//   				},
//   				Ja4Fingerprint: &JA4FingerprintProperty{
//   					FallbackBehavior: jsii.String("fallbackBehavior"),
//   				},
//   				JsonBody: &JsonBodyProperty{
//   					InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   					MatchPattern: &JsonMatchPatternProperty{
//   						All: all,
//   						IncludedPaths: []*string{
//   							jsii.String("includedPaths"),
//   						},
//   					},
//   					MatchScope: jsii.String("matchScope"),
//   					OversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				Method: method,
//   				QueryString: queryString,
//   				SingleHeader: singleHeader,
//   				SingleQueryArgument: singleQueryArgument,
//   				UriFragment: &UriFragmentProperty{
//   					FallbackBehavior: jsii.String("fallbackBehavior"),
//   				},
//   				UriPath: uriPath,
//   			},
//   			PositionalConstraint: jsii.String("positionalConstraint"),
//   			SearchString: jsii.String("searchString"),
//   			SearchStringBase64: jsii.String("searchStringBase64"),
//   			TextTransformations: []interface{}{
//   				&TextTransformationProperty{
//   					Priority: jsii.Number(123),
//   					Type: jsii.String("type"),
//   				},
//   			},
//   		},
//   		GeoMatchStatement: &GeoMatchStatementProperty{
//   			CountryCodes: []*string{
//   				jsii.String("countryCodes"),
//   			},
//   			ForwardedIpConfig: &ForwardedIPConfigurationProperty{
//   				FallbackBehavior: jsii.String("fallbackBehavior"),
//   				HeaderName: jsii.String("headerName"),
//   			},
//   		},
//   		IpSetReferenceStatement: map[string]*string{
//   			"arn": jsii.String("arn"),
//   			"ipSetForwardedIpConfig": map[string]*string{
//   				"fallbackBehavior": jsii.String("fallbackBehavior"),
//   				"headerName": jsii.String("headerName"),
//   				"position": jsii.String("position"),
//   			},
//   		},
//   		LabelMatchStatement: &LabelMatchStatementProperty{
//   			Key: jsii.String("key"),
//   			Scope: jsii.String("scope"),
//   		},
//   		NotStatement: &NotStatementProperty{
//   			Statement: statementProperty_,
//   		},
//   		OrStatement: &OrStatementProperty{
//   			Statements: []interface{}{
//   				statementProperty_,
//   			},
//   		},
//   		RateBasedStatement: &RateBasedStatementProperty{
//   			AggregateKeyType: jsii.String("aggregateKeyType"),
//   			CustomKeys: []interface{}{
//   				&RateBasedStatementCustomKeyProperty{
//   					Asn: asn,
//   					Cookie: &RateLimitCookieProperty{
//   						Name: jsii.String("name"),
//   						TextTransformations: []interface{}{
//   							&TextTransformationProperty{
//   								Priority: jsii.Number(123),
//   								Type: jsii.String("type"),
//   							},
//   						},
//   					},
//   					ForwardedIp: forwardedIp,
//   					Header: &RateLimitHeaderProperty{
//   						Name: jsii.String("name"),
//   						TextTransformations: []interface{}{
//   							&TextTransformationProperty{
//   								Priority: jsii.Number(123),
//   								Type: jsii.String("type"),
//   							},
//   						},
//   					},
//   					HttpMethod: httpMethod,
//   					Ip: ip,
//   					Ja3Fingerprint: &RateLimitJA3FingerprintProperty{
//   						FallbackBehavior: jsii.String("fallbackBehavior"),
//   					},
//   					Ja4Fingerprint: &RateLimitJA4FingerprintProperty{
//   						FallbackBehavior: jsii.String("fallbackBehavior"),
//   					},
//   					LabelNamespace: &RateLimitLabelNamespaceProperty{
//   						Namespace: jsii.String("namespace"),
//   					},
//   					QueryArgument: &RateLimitQueryArgumentProperty{
//   						Name: jsii.String("name"),
//   						TextTransformations: []interface{}{
//   							&TextTransformationProperty{
//   								Priority: jsii.Number(123),
//   								Type: jsii.String("type"),
//   							},
//   						},
//   					},
//   					QueryString: &RateLimitQueryStringProperty{
//   						TextTransformations: []interface{}{
//   							&TextTransformationProperty{
//   								Priority: jsii.Number(123),
//   								Type: jsii.String("type"),
//   							},
//   						},
//   					},
//   					UriPath: &RateLimitUriPathProperty{
//   						TextTransformations: []interface{}{
//   							&TextTransformationProperty{
//   								Priority: jsii.Number(123),
//   								Type: jsii.String("type"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   			EvaluationWindowSec: jsii.Number(123),
//   			ForwardedIpConfig: &ForwardedIPConfigurationProperty{
//   				FallbackBehavior: jsii.String("fallbackBehavior"),
//   				HeaderName: jsii.String("headerName"),
//   			},
//   			Limit: jsii.Number(123),
//   			ScopeDownStatement: statementProperty_,
//   		},
//   		RegexMatchStatement: &RegexMatchStatementProperty{
//   			FieldToMatch: &FieldToMatchProperty{
//   				AllQueryArguments: allQueryArguments,
//   				Body: &BodyProperty{
//   					OversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				Cookies: &CookiesProperty{
//   					MatchPattern: &CookieMatchPatternProperty{
//   						All: all,
//   						ExcludedCookies: []*string{
//   							jsii.String("excludedCookies"),
//   						},
//   						IncludedCookies: []*string{
//   							jsii.String("includedCookies"),
//   						},
//   					},
//   					MatchScope: jsii.String("matchScope"),
//   					OversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				HeaderOrder: &HeaderOrderProperty{
//   					OversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				Headers: &HeadersProperty{
//   					MatchPattern: &HeaderMatchPatternProperty{
//   						All: all,
//   						ExcludedHeaders: []*string{
//   							jsii.String("excludedHeaders"),
//   						},
//   						IncludedHeaders: []*string{
//   							jsii.String("includedHeaders"),
//   						},
//   					},
//   					MatchScope: jsii.String("matchScope"),
//   					OversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				Ja3Fingerprint: &JA3FingerprintProperty{
//   					FallbackBehavior: jsii.String("fallbackBehavior"),
//   				},
//   				Ja4Fingerprint: &JA4FingerprintProperty{
//   					FallbackBehavior: jsii.String("fallbackBehavior"),
//   				},
//   				JsonBody: &JsonBodyProperty{
//   					InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   					MatchPattern: &JsonMatchPatternProperty{
//   						All: all,
//   						IncludedPaths: []*string{
//   							jsii.String("includedPaths"),
//   						},
//   					},
//   					MatchScope: jsii.String("matchScope"),
//   					OversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				Method: method,
//   				QueryString: queryString,
//   				SingleHeader: singleHeader,
//   				SingleQueryArgument: singleQueryArgument,
//   				UriFragment: &UriFragmentProperty{
//   					FallbackBehavior: jsii.String("fallbackBehavior"),
//   				},
//   				UriPath: uriPath,
//   			},
//   			RegexString: jsii.String("regexString"),
//   			TextTransformations: []interface{}{
//   				&TextTransformationProperty{
//   					Priority: jsii.Number(123),
//   					Type: jsii.String("type"),
//   				},
//   			},
//   		},
//   		RegexPatternSetReferenceStatement: &RegexPatternSetReferenceStatementProperty{
//   			Arn: jsii.String("arn"),
//   			FieldToMatch: &FieldToMatchProperty{
//   				AllQueryArguments: allQueryArguments,
//   				Body: &BodyProperty{
//   					OversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				Cookies: &CookiesProperty{
//   					MatchPattern: &CookieMatchPatternProperty{
//   						All: all,
//   						ExcludedCookies: []*string{
//   							jsii.String("excludedCookies"),
//   						},
//   						IncludedCookies: []*string{
//   							jsii.String("includedCookies"),
//   						},
//   					},
//   					MatchScope: jsii.String("matchScope"),
//   					OversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				HeaderOrder: &HeaderOrderProperty{
//   					OversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				Headers: &HeadersProperty{
//   					MatchPattern: &HeaderMatchPatternProperty{
//   						All: all,
//   						ExcludedHeaders: []*string{
//   							jsii.String("excludedHeaders"),
//   						},
//   						IncludedHeaders: []*string{
//   							jsii.String("includedHeaders"),
//   						},
//   					},
//   					MatchScope: jsii.String("matchScope"),
//   					OversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				Ja3Fingerprint: &JA3FingerprintProperty{
//   					FallbackBehavior: jsii.String("fallbackBehavior"),
//   				},
//   				Ja4Fingerprint: &JA4FingerprintProperty{
//   					FallbackBehavior: jsii.String("fallbackBehavior"),
//   				},
//   				JsonBody: &JsonBodyProperty{
//   					InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   					MatchPattern: &JsonMatchPatternProperty{
//   						All: all,
//   						IncludedPaths: []*string{
//   							jsii.String("includedPaths"),
//   						},
//   					},
//   					MatchScope: jsii.String("matchScope"),
//   					OversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				Method: method,
//   				QueryString: queryString,
//   				SingleHeader: singleHeader,
//   				SingleQueryArgument: singleQueryArgument,
//   				UriFragment: &UriFragmentProperty{
//   					FallbackBehavior: jsii.String("fallbackBehavior"),
//   				},
//   				UriPath: uriPath,
//   			},
//   			TextTransformations: []interface{}{
//   				&TextTransformationProperty{
//   					Priority: jsii.Number(123),
//   					Type: jsii.String("type"),
//   				},
//   			},
//   		},
//   		SizeConstraintStatement: &SizeConstraintStatementProperty{
//   			ComparisonOperator: jsii.String("comparisonOperator"),
//   			FieldToMatch: &FieldToMatchProperty{
//   				AllQueryArguments: allQueryArguments,
//   				Body: &BodyProperty{
//   					OversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				Cookies: &CookiesProperty{
//   					MatchPattern: &CookieMatchPatternProperty{
//   						All: all,
//   						ExcludedCookies: []*string{
//   							jsii.String("excludedCookies"),
//   						},
//   						IncludedCookies: []*string{
//   							jsii.String("includedCookies"),
//   						},
//   					},
//   					MatchScope: jsii.String("matchScope"),
//   					OversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				HeaderOrder: &HeaderOrderProperty{
//   					OversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				Headers: &HeadersProperty{
//   					MatchPattern: &HeaderMatchPatternProperty{
//   						All: all,
//   						ExcludedHeaders: []*string{
//   							jsii.String("excludedHeaders"),
//   						},
//   						IncludedHeaders: []*string{
//   							jsii.String("includedHeaders"),
//   						},
//   					},
//   					MatchScope: jsii.String("matchScope"),
//   					OversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				Ja3Fingerprint: &JA3FingerprintProperty{
//   					FallbackBehavior: jsii.String("fallbackBehavior"),
//   				},
//   				Ja4Fingerprint: &JA4FingerprintProperty{
//   					FallbackBehavior: jsii.String("fallbackBehavior"),
//   				},
//   				JsonBody: &JsonBodyProperty{
//   					InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   					MatchPattern: &JsonMatchPatternProperty{
//   						All: all,
//   						IncludedPaths: []*string{
//   							jsii.String("includedPaths"),
//   						},
//   					},
//   					MatchScope: jsii.String("matchScope"),
//   					OversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				Method: method,
//   				QueryString: queryString,
//   				SingleHeader: singleHeader,
//   				SingleQueryArgument: singleQueryArgument,
//   				UriFragment: &UriFragmentProperty{
//   					FallbackBehavior: jsii.String("fallbackBehavior"),
//   				},
//   				UriPath: uriPath,
//   			},
//   			Size: jsii.Number(123),
//   			TextTransformations: []interface{}{
//   				&TextTransformationProperty{
//   					Priority: jsii.Number(123),
//   					Type: jsii.String("type"),
//   				},
//   			},
//   		},
//   		SqliMatchStatement: &SqliMatchStatementProperty{
//   			FieldToMatch: &FieldToMatchProperty{
//   				AllQueryArguments: allQueryArguments,
//   				Body: &BodyProperty{
//   					OversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				Cookies: &CookiesProperty{
//   					MatchPattern: &CookieMatchPatternProperty{
//   						All: all,
//   						ExcludedCookies: []*string{
//   							jsii.String("excludedCookies"),
//   						},
//   						IncludedCookies: []*string{
//   							jsii.String("includedCookies"),
//   						},
//   					},
//   					MatchScope: jsii.String("matchScope"),
//   					OversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				HeaderOrder: &HeaderOrderProperty{
//   					OversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				Headers: &HeadersProperty{
//   					MatchPattern: &HeaderMatchPatternProperty{
//   						All: all,
//   						ExcludedHeaders: []*string{
//   							jsii.String("excludedHeaders"),
//   						},
//   						IncludedHeaders: []*string{
//   							jsii.String("includedHeaders"),
//   						},
//   					},
//   					MatchScope: jsii.String("matchScope"),
//   					OversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				Ja3Fingerprint: &JA3FingerprintProperty{
//   					FallbackBehavior: jsii.String("fallbackBehavior"),
//   				},
//   				Ja4Fingerprint: &JA4FingerprintProperty{
//   					FallbackBehavior: jsii.String("fallbackBehavior"),
//   				},
//   				JsonBody: &JsonBodyProperty{
//   					InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   					MatchPattern: &JsonMatchPatternProperty{
//   						All: all,
//   						IncludedPaths: []*string{
//   							jsii.String("includedPaths"),
//   						},
//   					},
//   					MatchScope: jsii.String("matchScope"),
//   					OversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				Method: method,
//   				QueryString: queryString,
//   				SingleHeader: singleHeader,
//   				SingleQueryArgument: singleQueryArgument,
//   				UriFragment: &UriFragmentProperty{
//   					FallbackBehavior: jsii.String("fallbackBehavior"),
//   				},
//   				UriPath: uriPath,
//   			},
//   			SensitivityLevel: jsii.String("sensitivityLevel"),
//   			TextTransformations: []interface{}{
//   				&TextTransformationProperty{
//   					Priority: jsii.Number(123),
//   					Type: jsii.String("type"),
//   				},
//   			},
//   		},
//   		XssMatchStatement: &XssMatchStatementProperty{
//   			FieldToMatch: &FieldToMatchProperty{
//   				AllQueryArguments: allQueryArguments,
//   				Body: &BodyProperty{
//   					OversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				Cookies: &CookiesProperty{
//   					MatchPattern: &CookieMatchPatternProperty{
//   						All: all,
//   						ExcludedCookies: []*string{
//   							jsii.String("excludedCookies"),
//   						},
//   						IncludedCookies: []*string{
//   							jsii.String("includedCookies"),
//   						},
//   					},
//   					MatchScope: jsii.String("matchScope"),
//   					OversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				HeaderOrder: &HeaderOrderProperty{
//   					OversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				Headers: &HeadersProperty{
//   					MatchPattern: &HeaderMatchPatternProperty{
//   						All: all,
//   						ExcludedHeaders: []*string{
//   							jsii.String("excludedHeaders"),
//   						},
//   						IncludedHeaders: []*string{
//   							jsii.String("includedHeaders"),
//   						},
//   					},
//   					MatchScope: jsii.String("matchScope"),
//   					OversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				Ja3Fingerprint: &JA3FingerprintProperty{
//   					FallbackBehavior: jsii.String("fallbackBehavior"),
//   				},
//   				Ja4Fingerprint: &JA4FingerprintProperty{
//   					FallbackBehavior: jsii.String("fallbackBehavior"),
//   				},
//   				JsonBody: &JsonBodyProperty{
//   					InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   					MatchPattern: &JsonMatchPatternProperty{
//   						All: all,
//   						IncludedPaths: []*string{
//   							jsii.String("includedPaths"),
//   						},
//   					},
//   					MatchScope: jsii.String("matchScope"),
//   					OversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				Method: method,
//   				QueryString: queryString,
//   				SingleHeader: singleHeader,
//   				SingleQueryArgument: singleQueryArgument,
//   				UriFragment: &UriFragmentProperty{
//   					FallbackBehavior: jsii.String("fallbackBehavior"),
//   				},
//   				UriPath: uriPath,
//   			},
//   			TextTransformations: []interface{}{
//   				&TextTransformationProperty{
//   					Priority: jsii.Number(123),
//   					Type: jsii.String("type"),
//   				},
//   			},
//   		},
//   	},
//   	VisibilityConfig: &VisibilityConfigProperty{
//   		CloudWatchMetricsEnabled: jsii.Boolean(false),
//   		MetricName: jsii.String("metricName"),
//   		SampledRequestsEnabled: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-rule.html
//
type CfnRuleGroupPropsMixin_RuleProperty struct {
	// The action that AWS WAF should take on a web request when it matches the rule statement.
	//
	// Settings at the web ACL level can override the rule action setting.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-rule.html#cfn-wafv2-rulegroup-rule-action
	//
	Action interface{} `field:"optional" json:"action" yaml:"action"`
	// Specifies how AWS WAF should handle `CAPTCHA` evaluations.
	//
	// If you don't specify this, AWS WAF uses the `CAPTCHA` configuration that's defined for the web ACL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-rule.html#cfn-wafv2-rulegroup-rule-captchaconfig
	//
	CaptchaConfig interface{} `field:"optional" json:"captchaConfig" yaml:"captchaConfig"`
	// Specifies how AWS WAF should handle `Challenge` evaluations.
	//
	// If you don't specify this, AWS WAF uses the challenge configuration that's defined for the web ACL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-rule.html#cfn-wafv2-rulegroup-rule-challengeconfig
	//
	ChallengeConfig interface{} `field:"optional" json:"challengeConfig" yaml:"challengeConfig"`
	// The name of the rule.
	//
	// If you change the name of a `Rule` after you create it and you want the rule's metric name to reflect the change, update the metric name in the rule's `VisibilityConfig` settings. AWS WAF doesn't automatically update the metric name when you update the rule name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-rule.html#cfn-wafv2-rulegroup-rule-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// If you define more than one `Rule` in a `WebACL` , AWS WAF evaluates each request against the `Rules` in order based on the value of `Priority` .
	//
	// AWS WAF processes rules with lower priority first. The priorities don't need to be consecutive, but they must all be different.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-rule.html#cfn-wafv2-rulegroup-rule-priority
	//
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Labels to apply to web requests that match the rule match statement.
	//
	// AWS WAF applies fully qualified labels to matching web requests. A fully qualified label is the concatenation of a label namespace and a rule label. The rule's rule group or web ACL defines the label namespace.
	//
	// > Any rule that isn't a rule group reference statement or managed rule group statement can add labels to matching web requests.
	//
	// Rules that run after this rule in the web ACL can match against these labels using a `LabelMatchStatement` .
	//
	// For each label, provide a case-sensitive string containing optional namespaces and a label name, according to the following guidelines:
	//
	// - Separate each component of the label with a colon.
	// - Each namespace or name can have up to 128 characters.
	// - You can specify up to 5 namespaces in a label.
	// - Don't use the following reserved words in your label specification: `aws` , `waf` , `managed` , `rulegroup` , `webacl` , `regexpatternset` , or `ipset` .
	//
	// For example, `myLabelName` or `nameSpace1:nameSpace2:myLabelName` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-rule.html#cfn-wafv2-rulegroup-rule-rulelabels
	//
	RuleLabels interface{} `field:"optional" json:"ruleLabels" yaml:"ruleLabels"`
	// The AWS WAF processing statement for the rule, for example `ByteMatchStatement` or `SizeConstraintStatement` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-rule.html#cfn-wafv2-rulegroup-rule-statement
	//
	Statement interface{} `field:"optional" json:"statement" yaml:"statement"`
	// Defines and enables Amazon CloudWatch metrics and web request sample collection.
	//
	// If you change the name of a `Rule` after you create it and you want the rule's metric name to reflect the change, update the metric name as well. AWS WAF doesn't automatically update the metric name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-rule.html#cfn-wafv2-rulegroup-rule-visibilityconfig
	//
	VisibilityConfig interface{} `field:"optional" json:"visibilityConfig" yaml:"visibilityConfig"`
}

