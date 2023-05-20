package awswafv2


// A rate-based rule counts incoming requests and rate limits requests when they are coming at too fast a rate.
//
// The rule categorizes requests according to your aggregation criteria, collects them into aggregation instances, and counts and rate limits the requests for each instance.
//
// You can specify individual aggregation keys, like IP address or HTTP method. You can also specify aggregation key combinations, like IP address and HTTP method, or HTTP method, query argument, and cookie.
//
// Each unique set of values for the aggregation keys that you specify is a separate aggregation instance, with the value from each key contributing to the aggregation instance definition.
//
// For example, assume the rule evaluates web requests with the following IP address and HTTP method values:
//
// - IP address 10.1.1.1, HTTP method POST
// - IP address 10.1.1.1, HTTP method GET
// - IP address 127.0.0.0, HTTP method POST
// - IP address 10.1.1.1, HTTP method GET
//
// The rule would create different aggregation instances according to your aggregation criteria, for example:
//
// - If the aggregation criteria is just the IP address, then each individual address is an aggregation instance, and AWS WAF counts requests separately for each. The aggregation instances and request counts for our example would be the following:
//
// - IP address 10.1.1.1: count 3
// - IP address 127.0.0.0: count 1
// - If the aggregation criteria is HTTP method, then each individual HTTP method is an aggregation instance. The aggregation instances and request counts for our example would be the following:
//
// - HTTP method POST: count 2
// - HTTP method GET: count 2
// - If the aggregation criteria is IP address and HTTP method, then each IP address and each HTTP method would contribute to the combined aggregation instance. The aggregation instances and request counts for our example would be the following:
//
// - IP address 10.1.1.1, HTTP method POST: count 1
// - IP address 10.1.1.1, HTTP method GET: count 2
// - IP address 127.0.0.0, HTTP method POST: count 1
//
// For any n-tuple of aggregation keys, each unique combination of values for the keys defines a separate aggregation instance, which AWS WAF counts and rate-limits individually.
//
// You can optionally nest another statement inside the rate-based statement, to narrow the scope of the rule so that it only counts and rate limits requests that match the nested statement. You can use this nested scope-down statement in conjunction with your aggregation key specifications or you can just count and rate limit all requests that match the scope-down statement, without additional aggregation. When you choose to just manage all requests that match a scope-down statement, the aggregation instance is singular for the rule.
//
// You cannot nest a `RateBasedStatement` inside another statement, for example inside a `NotStatement` or `OrStatement` . You can define a `RateBasedStatement` inside a web ACL and inside a rule group.
//
// For additional information about the options, see [Rate limiting web requests using rate-based rules](https://docs.aws.amazon.com/waf/latest/developerguide/waf-rate-based-rules.html) in the *AWS WAF Developer Guide* .
//
// If you only aggregate on the individual IP address or forwarded IP address, you can retrieve the list of IP addresses that AWS WAF is currently rate limiting for a rule through the API call `GetRateBasedStatementManagedKeys` . This option is not available for other aggregation configurations.
//
// AWS WAF tracks and manages web requests separately for each instance of a rate-based rule that you use. For example, if you provide the same rate-based rule settings in two web ACLs, each of the two rule statements represents a separate instance of the rate-based rule and gets its own tracking and management by AWS WAF . If you define a rate-based rule inside a rule group, and then use that rule group in multiple places, each use creates a separate instance of the rate-based rule that gets its own tracking and management by AWS WAF .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var all interface{}
//   var allQueryArguments interface{}
//   var method interface{}
//   var queryString interface{}
//   var singleHeader interface{}
//   var singleQueryArgument interface{}
//   var statementProperty_ statementProperty
//   var uriPath interface{}
//
//   rateBasedStatementProperty := &RateBasedStatementProperty{
//   	AggregateKeyType: jsii.String("aggregateKeyType"),
//   	Limit: jsii.Number(123),
//
//   	// the properties below are optional
//   	ForwardedIpConfig: &ForwardedIPConfigurationProperty{
//   		FallbackBehavior: jsii.String("fallbackBehavior"),
//   		HeaderName: jsii.String("headerName"),
//   	},
//   	ScopeDownStatement: &statementProperty{
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
//   		ManagedRuleGroupStatement: &ManagedRuleGroupStatementProperty{
//   			Name: jsii.String("name"),
//   			VendorName: jsii.String("vendorName"),
//
//   			// the properties below are optional
//   			ExcludedRules: []interface{}{
//   				&ExcludedRuleProperty{
//   					Name: jsii.String("name"),
//   				},
//   			},
//   			ManagedRuleGroupConfigs: []interface{}{
//   				&ManagedRuleGroupConfigProperty{
//   					AwsManagedRulesAtpRuleSet: &AWSManagedRulesATPRuleSetProperty{
//   						LoginPath: jsii.String("loginPath"),
//
//   						// the properties below are optional
//   						RequestInspection: &RequestInspectionProperty{
//   							PasswordField: &FieldIdentifierProperty{
//   								Identifier: jsii.String("identifier"),
//   							},
//   							PayloadType: jsii.String("payloadType"),
//   							UsernameField: &FieldIdentifierProperty{
//   								Identifier: jsii.String("identifier"),
//   							},
//   						},
//   						ResponseInspection: &ResponseInspectionProperty{
//   							BodyContains: &ResponseInspectionBodyContainsProperty{
//   								FailureStrings: []*string{
//   									jsii.String("failureStrings"),
//   								},
//   								SuccessStrings: []*string{
//   									jsii.String("successStrings"),
//   								},
//   							},
//   							Header: &ResponseInspectionHeaderProperty{
//   								FailureValues: []*string{
//   									jsii.String("failureValues"),
//   								},
//   								Name: jsii.String("name"),
//   								SuccessValues: []*string{
//   									jsii.String("successValues"),
//   								},
//   							},
//   							Json: &ResponseInspectionJsonProperty{
//   								FailureValues: []*string{
//   									jsii.String("failureValues"),
//   								},
//   								Identifier: jsii.String("identifier"),
//   								SuccessValues: []*string{
//   									jsii.String("successValues"),
//   								},
//   							},
//   							StatusCode: &ResponseInspectionStatusCodeProperty{
//   								FailureCodes: []interface{}{
//   									jsii.Number(123),
//   								},
//   								SuccessCodes: []interface{}{
//   									jsii.Number(123),
//   								},
//   							},
//   						},
//   					},
//   					AwsManagedRulesBotControlRuleSet: &AWSManagedRulesBotControlRuleSetProperty{
//   						InspectionLevel: jsii.String("inspectionLevel"),
//   					},
//   					LoginPath: jsii.String("loginPath"),
//   					PasswordField: &FieldIdentifierProperty{
//   						Identifier: jsii.String("identifier"),
//   					},
//   					PayloadType: jsii.String("payloadType"),
//   					UsernameField: &FieldIdentifierProperty{
//   						Identifier: jsii.String("identifier"),
//   					},
//   				},
//   			},
//   			RuleActionOverrides: []interface{}{
//   				&RuleActionOverrideProperty{
//   					ActionToUse: &RuleActionProperty{
//   						Allow: &AllowActionProperty{
//   							CustomRequestHandling: &CustomRequestHandlingProperty{
//   								InsertHeaders: []interface{}{
//   									&CustomHTTPHeaderProperty{
//   										Name: jsii.String("name"),
//   										Value: jsii.String("value"),
//   									},
//   								},
//   							},
//   						},
//   						Block: &BlockActionProperty{
//   							CustomResponse: &CustomResponseProperty{
//   								ResponseCode: jsii.Number(123),
//
//   								// the properties below are optional
//   								CustomResponseBodyKey: jsii.String("customResponseBodyKey"),
//   								ResponseHeaders: []interface{}{
//   									&CustomHTTPHeaderProperty{
//   										Name: jsii.String("name"),
//   										Value: jsii.String("value"),
//   									},
//   								},
//   							},
//   						},
//   						Captcha: &CaptchaActionProperty{
//   							CustomRequestHandling: &CustomRequestHandlingProperty{
//   								InsertHeaders: []interface{}{
//   									&CustomHTTPHeaderProperty{
//   										Name: jsii.String("name"),
//   										Value: jsii.String("value"),
//   									},
//   								},
//   							},
//   						},
//   						Challenge: &ChallengeActionProperty{
//   							CustomRequestHandling: &CustomRequestHandlingProperty{
//   								InsertHeaders: []interface{}{
//   									&CustomHTTPHeaderProperty{
//   										Name: jsii.String("name"),
//   										Value: jsii.String("value"),
//   									},
//   								},
//   							},
//   						},
//   						Count: &CountActionProperty{
//   							CustomRequestHandling: &CustomRequestHandlingProperty{
//   								InsertHeaders: []interface{}{
//   									&CustomHTTPHeaderProperty{
//   										Name: jsii.String("name"),
//   										Value: jsii.String("value"),
//   									},
//   								},
//   							},
//   						},
//   					},
//   					Name: jsii.String("name"),
//   				},
//   			},
//   			ScopeDownStatement: statementProperty_,
//   			Version: jsii.String("version"),
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
//   		RuleGroupReferenceStatement: &RuleGroupReferenceStatementProperty{
//   			Arn: jsii.String("arn"),
//
//   			// the properties below are optional
//   			ExcludedRules: []interface{}{
//   				&ExcludedRuleProperty{
//   					Name: jsii.String("name"),
//   				},
//   			},
//   			RuleActionOverrides: []interface{}{
//   				&RuleActionOverrideProperty{
//   					ActionToUse: &RuleActionProperty{
//   						Allow: &AllowActionProperty{
//   							CustomRequestHandling: &CustomRequestHandlingProperty{
//   								InsertHeaders: []interface{}{
//   									&CustomHTTPHeaderProperty{
//   										Name: jsii.String("name"),
//   										Value: jsii.String("value"),
//   									},
//   								},
//   							},
//   						},
//   						Block: &BlockActionProperty{
//   							CustomResponse: &CustomResponseProperty{
//   								ResponseCode: jsii.Number(123),
//
//   								// the properties below are optional
//   								CustomResponseBodyKey: jsii.String("customResponseBodyKey"),
//   								ResponseHeaders: []interface{}{
//   									&CustomHTTPHeaderProperty{
//   										Name: jsii.String("name"),
//   										Value: jsii.String("value"),
//   									},
//   								},
//   							},
//   						},
//   						Captcha: &CaptchaActionProperty{
//   							CustomRequestHandling: &CustomRequestHandlingProperty{
//   								InsertHeaders: []interface{}{
//   									&CustomHTTPHeaderProperty{
//   										Name: jsii.String("name"),
//   										Value: jsii.String("value"),
//   									},
//   								},
//   							},
//   						},
//   						Challenge: &ChallengeActionProperty{
//   							CustomRequestHandling: &CustomRequestHandlingProperty{
//   								InsertHeaders: []interface{}{
//   									&CustomHTTPHeaderProperty{
//   										Name: jsii.String("name"),
//   										Value: jsii.String("value"),
//   									},
//   								},
//   							},
//   						},
//   						Count: &CountActionProperty{
//   							CustomRequestHandling: &CustomRequestHandlingProperty{
//   								InsertHeaders: []interface{}{
//   									&CustomHTTPHeaderProperty{
//   										Name: jsii.String("name"),
//   										Value: jsii.String("value"),
//   									},
//   								},
//   							},
//   						},
//   					},
//   					Name: jsii.String("name"),
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
//   }
//
type CfnWebACL_RateBasedStatementProperty struct {
	// Setting that indicates how to aggregate the request counts.
	//
	// > Web requests that are missing any of the components specified in the aggregation keys are omitted from the rate-based rule evaluation and handling.
	//
	// - `CONSTANT` - Count and limit the requests that match the rate-based rule's scope-down statement. With this option, the counted requests aren't further aggregated. The scope-down statement is the only specification used. When the count of all requests that satisfy the scope-down statement goes over the limit, AWS WAF applies the rule action to all requests that satisfy the scope-down statement.
	//
	// With this option, you must configure the `ScopeDownStatement` property.
	// - `CUSTOM_KEYS` - Aggregate the request counts using one or more web request components as the aggregate keys.
	//
	// With this option, you must specify the aggregate keys in the `CustomKeys` property.
	//
	// To aggregate on only the IP address or only the forwarded IP address, don't use custom keys. Instead, set the aggregate key type to `IP` or `FORWARDED_IP` .
	// - `FORWARDED_IP` - Aggregate the request counts on the first IP address in an HTTP header.
	//
	// With this option, you must specify the header to use in the `ForwardedIPConfig` property.
	//
	// To aggregate on a combination of the forwarded IP address with other aggregate keys, use `CUSTOM_KEYS` .
	// - `IP` - Aggregate the request counts on the IP address from the web request origin.
	//
	// To aggregate on a combination of the IP address with other aggregate keys, use `CUSTOM_KEYS` .
	AggregateKeyType *string `field:"required" json:"aggregateKeyType" yaml:"aggregateKeyType"`
	// The limit on requests per 5-minute period for a single aggregation instance for the rate-based rule.
	//
	// If the rate-based statement includes a `ScopeDownStatement` , this limit is applied only to the requests that match the statement.
	//
	// Examples:
	//
	// - If you aggregate on just the IP address, this is the limit on requests from any single IP address.
	// - If you aggregate on the HTTP method and the query argument name "city", then this is the limit on requests for any single method, city pair.
	Limit *float64 `field:"required" json:"limit" yaml:"limit"`
	// The configuration for inspecting IP addresses in an HTTP header that you specify, instead of using the IP address that's reported by the web request origin.
	//
	// Commonly, this is the X-Forwarded-For (XFF) header, but you can specify any header name.
	//
	// > If the specified header isn't present in the request, AWS WAF doesn't apply the rule to the web request at all.
	//
	// This is required if you specify a forwarded IP in the rule's aggregate key settings.
	ForwardedIpConfig interface{} `field:"optional" json:"forwardedIpConfig" yaml:"forwardedIpConfig"`
	// An optional nested statement that narrows the scope of the web requests that are evaluated by the rate-based statement.
	//
	// Requests are only tracked by the rate-based statement if they match the scope-down statement. You can use any nestable `Statement` in the scope-down statement, and you can nest statements at any level, the same as you can for a rule statement.
	ScopeDownStatement interface{} `field:"optional" json:"scopeDownStatement" yaml:"scopeDownStatement"`
}

