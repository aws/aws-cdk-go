package awswafv2


// A single rule, which you can use in a `WebACL` or `RuleGroup` to identify web requests that you want to allow, block, or count.
//
// Each rule includes one top-level `Statement` that AWS WAF uses to identify matching web requests, and parameters that govern how AWS WAF handles them.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var all interface{}
//   var allow interface{}
//   var allQueryArguments interface{}
//   var block interface{}
//   var captcha interface{}
//   var challenge interface{}
//   var count interface{}
//   var method interface{}
//   var queryString interface{}
//   var singleHeader interface{}
//   var singleQueryArgument interface{}
//   var statementProperty_ statementProperty
//   var uriPath interface{}
//
//   ruleProperty := &RuleProperty{
//   	Name: jsii.String("name"),
//   	Priority: jsii.Number(123),
//   	Statement: &statementProperty{
//   		AndStatement: &AndStatementProperty{
//   			Statements: []interface{}{
//   				statementProperty_,
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
//   				JsonBody: &JsonBodyProperty{
//   					MatchPattern: &JsonMatchPatternProperty{
//   						All: all,
//   						IncludedPaths: []*string{
//   							jsii.String("includedPaths"),
//   						},
//   					},
//   					MatchScope: jsii.String("matchScope"),
//
//   					// the properties below are optional
//   					InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   					OversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				Method: method,
//   				QueryString: queryString,
//   				SingleHeader: singleHeader,
//   				SingleQueryArgument: singleQueryArgument,
//   				UriPath: uriPath,
//   			},
//   			PositionalConstraint: jsii.String("positionalConstraint"),
//   			TextTransformations: []interface{}{
//   				&TextTransformationProperty{
//   					Priority: jsii.Number(123),
//   					Type: jsii.String("type"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			SearchString: jsii.String("searchString"),
//   			SearchStringBase64: jsii.String("searchStringBase64"),
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
//   		IpSetReferenceStatement: map[string]interface{}{
//   			"arn": jsii.String("arn"),
//
//   			// the properties below are optional
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
//   			Limit: jsii.Number(123),
//
//   			// the properties below are optional
//   			ForwardedIpConfig: &ForwardedIPConfigurationProperty{
//   				FallbackBehavior: jsii.String("fallbackBehavior"),
//   				HeaderName: jsii.String("headerName"),
//   			},
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
//   				JsonBody: &JsonBodyProperty{
//   					MatchPattern: &JsonMatchPatternProperty{
//   						All: all,
//   						IncludedPaths: []*string{
//   							jsii.String("includedPaths"),
//   						},
//   					},
//   					MatchScope: jsii.String("matchScope"),
//
//   					// the properties below are optional
//   					InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   					OversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				Method: method,
//   				QueryString: queryString,
//   				SingleHeader: singleHeader,
//   				SingleQueryArgument: singleQueryArgument,
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
//   				JsonBody: &JsonBodyProperty{
//   					MatchPattern: &JsonMatchPatternProperty{
//   						All: all,
//   						IncludedPaths: []*string{
//   							jsii.String("includedPaths"),
//   						},
//   					},
//   					MatchScope: jsii.String("matchScope"),
//
//   					// the properties below are optional
//   					InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   					OversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				Method: method,
//   				QueryString: queryString,
//   				SingleHeader: singleHeader,
//   				SingleQueryArgument: singleQueryArgument,
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
//   				JsonBody: &JsonBodyProperty{
//   					MatchPattern: &JsonMatchPatternProperty{
//   						All: all,
//   						IncludedPaths: []*string{
//   							jsii.String("includedPaths"),
//   						},
//   					},
//   					MatchScope: jsii.String("matchScope"),
//
//   					// the properties below are optional
//   					InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   					OversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				Method: method,
//   				QueryString: queryString,
//   				SingleHeader: singleHeader,
//   				SingleQueryArgument: singleQueryArgument,
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
//   				JsonBody: &JsonBodyProperty{
//   					MatchPattern: &JsonMatchPatternProperty{
//   						All: all,
//   						IncludedPaths: []*string{
//   							jsii.String("includedPaths"),
//   						},
//   					},
//   					MatchScope: jsii.String("matchScope"),
//
//   					// the properties below are optional
//   					InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   					OversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				Method: method,
//   				QueryString: queryString,
//   				SingleHeader: singleHeader,
//   				SingleQueryArgument: singleQueryArgument,
//   				UriPath: uriPath,
//   			},
//   			TextTransformations: []interface{}{
//   				&TextTransformationProperty{
//   					Priority: jsii.Number(123),
//   					Type: jsii.String("type"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			SensitivityLevel: jsii.String("sensitivityLevel"),
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
//   				JsonBody: &JsonBodyProperty{
//   					MatchPattern: &JsonMatchPatternProperty{
//   						All: all,
//   						IncludedPaths: []*string{
//   							jsii.String("includedPaths"),
//   						},
//   					},
//   					MatchScope: jsii.String("matchScope"),
//
//   					// the properties below are optional
//   					InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   					OversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				Method: method,
//   				QueryString: queryString,
//   				SingleHeader: singleHeader,
//   				SingleQueryArgument: singleQueryArgument,
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
//
//   	// the properties below are optional
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
//   	RuleLabels: []interface{}{
//   		&LabelProperty{
//   			Name: jsii.String("name"),
//   		},
//   	},
//   }
//
type CfnRuleGroup_RuleProperty struct {
	// The name of the rule.
	//
	// You can't change the name of a `Rule` after you create it.
	Name *string `field:"required" json:"name" yaml:"name"`
	// If you define more than one `Rule` in a `WebACL` , AWS WAF evaluates each request against the `Rules` in order based on the value of `Priority` .
	//
	// AWS WAF processes rules with lower priority first. The priorities don't need to be consecutive, but they must all be different.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// The AWS WAF processing statement for the rule, for example `ByteMatchStatement` or `SizeConstraintStatement` .
	Statement interface{} `field:"required" json:"statement" yaml:"statement"`
	// Defines and enables Amazon CloudWatch metrics and web request sample collection.
	VisibilityConfig interface{} `field:"required" json:"visibilityConfig" yaml:"visibilityConfig"`
	// The action that AWS WAF should take on a web request when it matches the rule statement.
	//
	// Settings at the web ACL level can override the rule action setting.
	Action interface{} `field:"optional" json:"action" yaml:"action"`
	// Specifies how AWS WAF should handle `CAPTCHA` evaluations.
	//
	// If you don't specify this, AWS WAF uses the `CAPTCHA` configuration that's defined for the web ACL.
	CaptchaConfig interface{} `field:"optional" json:"captchaConfig" yaml:"captchaConfig"`
	// Specifies how AWS WAF should handle `Challenge` evaluations.
	//
	// If you don't specify this, AWS WAF uses the challenge configuration that's defined for the web ACL.
	ChallengeConfig interface{} `field:"optional" json:"challengeConfig" yaml:"challengeConfig"`
	// Labels to apply to web requests that match the rule match statement.
	//
	// AWS WAF applies fully qualified labels to matching web requests. A fully qualified label is the concatenation of a label namespace and a rule label. The rule's rule group or web ACL defines the label namespace.
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
	RuleLabels interface{} `field:"optional" json:"ruleLabels" yaml:"ruleLabels"`
}

