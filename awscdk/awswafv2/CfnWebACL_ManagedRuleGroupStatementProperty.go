package awswafv2


// A rule statement used to run the rules that are defined in a managed rule group.
//
// To use this, provide the vendor name and the name of the rule group in this statement.
//
// You cannot nest a `ManagedRuleGroupStatement` , for example for use inside a `NotStatement` or `OrStatement` . It can only be referenced as a top-level statement within a rule.
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
//   managedRuleGroupStatementProperty := &ManagedRuleGroupStatementProperty{
//   	Name: jsii.String("name"),
//   	VendorName: jsii.String("vendorName"),
//
//   	// the properties below are optional
//   	ExcludedRules: []interface{}{
//   		&ExcludedRuleProperty{
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	ManagedRuleGroupConfigs: []interface{}{
//   		&ManagedRuleGroupConfigProperty{
//   			AwsManagedRulesAtpRuleSet: &AWSManagedRulesATPRuleSetProperty{
//   				LoginPath: jsii.String("loginPath"),
//
//   				// the properties below are optional
//   				RequestInspection: &RequestInspectionProperty{
//   					PasswordField: &FieldIdentifierProperty{
//   						Identifier: jsii.String("identifier"),
//   					},
//   					PayloadType: jsii.String("payloadType"),
//   					UsernameField: &FieldIdentifierProperty{
//   						Identifier: jsii.String("identifier"),
//   					},
//   				},
//   				ResponseInspection: &ResponseInspectionProperty{
//   					BodyContains: &ResponseInspectionBodyContainsProperty{
//   						FailureStrings: []*string{
//   							jsii.String("failureStrings"),
//   						},
//   						SuccessStrings: []*string{
//   							jsii.String("successStrings"),
//   						},
//   					},
//   					Header: &ResponseInspectionHeaderProperty{
//   						FailureValues: []*string{
//   							jsii.String("failureValues"),
//   						},
//   						Name: jsii.String("name"),
//   						SuccessValues: []*string{
//   							jsii.String("successValues"),
//   						},
//   					},
//   					Json: &ResponseInspectionJsonProperty{
//   						FailureValues: []*string{
//   							jsii.String("failureValues"),
//   						},
//   						Identifier: jsii.String("identifier"),
//   						SuccessValues: []*string{
//   							jsii.String("successValues"),
//   						},
//   					},
//   					StatusCode: &ResponseInspectionStatusCodeProperty{
//   						FailureCodes: []interface{}{
//   							jsii.Number(123),
//   						},
//   						SuccessCodes: []interface{}{
//   							jsii.Number(123),
//   						},
//   					},
//   				},
//   			},
//   			AwsManagedRulesBotControlRuleSet: &AWSManagedRulesBotControlRuleSetProperty{
//   				InspectionLevel: jsii.String("inspectionLevel"),
//   			},
//   			LoginPath: jsii.String("loginPath"),
//   			PasswordField: &FieldIdentifierProperty{
//   				Identifier: jsii.String("identifier"),
//   			},
//   			PayloadType: jsii.String("payloadType"),
//   			UsernameField: &FieldIdentifierProperty{
//   				Identifier: jsii.String("identifier"),
//   			},
//   		},
//   	},
//   	RuleActionOverrides: []interface{}{
//   		&RuleActionOverrideProperty{
//   			ActionToUse: &RuleActionProperty{
//   				Allow: &AllowActionProperty{
//   					CustomRequestHandling: &CustomRequestHandlingProperty{
//   						InsertHeaders: []interface{}{
//   							&CustomHTTPHeaderProperty{
//   								Name: jsii.String("name"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   					},
//   				},
//   				Block: &BlockActionProperty{
//   					CustomResponse: &CustomResponseProperty{
//   						ResponseCode: jsii.Number(123),
//
//   						// the properties below are optional
//   						CustomResponseBodyKey: jsii.String("customResponseBodyKey"),
//   						ResponseHeaders: []interface{}{
//   							&CustomHTTPHeaderProperty{
//   								Name: jsii.String("name"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   					},
//   				},
//   				Captcha: &CaptchaActionProperty{
//   					CustomRequestHandling: &CustomRequestHandlingProperty{
//   						InsertHeaders: []interface{}{
//   							&CustomHTTPHeaderProperty{
//   								Name: jsii.String("name"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   					},
//   				},
//   				Challenge: &ChallengeActionProperty{
//   					CustomRequestHandling: &CustomRequestHandlingProperty{
//   						InsertHeaders: []interface{}{
//   							&CustomHTTPHeaderProperty{
//   								Name: jsii.String("name"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   					},
//   				},
//   				Count: &CountActionProperty{
//   					CustomRequestHandling: &CustomRequestHandlingProperty{
//   						InsertHeaders: []interface{}{
//   							&CustomHTTPHeaderProperty{
//   								Name: jsii.String("name"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   			Name: jsii.String("name"),
//   		},
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
//   	Version: jsii.String("version"),
//   }
//
type CfnWebACL_ManagedRuleGroupStatementProperty struct {
	// The name of the managed rule group.
	//
	// You use this, along with the vendor name, to identify the rule group.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The name of the managed rule group vendor.
	//
	// You use this, along with the rule group name, to identify the rule group.
	VendorName *string `field:"required" json:"vendorName" yaml:"vendorName"`
	// Rules in the referenced rule group whose actions are set to `Count` .
	//
	// > Instead of this option, use `RuleActionOverrides` . It accepts any valid action setting, including `Count` .
	ExcludedRules interface{} `field:"optional" json:"excludedRules" yaml:"excludedRules"`
	// Additional information that's used by a managed rule group. Many managed rule groups don't require this.
	//
	// Use the `AWSManagedRulesATPRuleSet` configuration object for the account takeover prevention managed rule group, to provide information such as the sign-in page of your application and the type of content to accept or reject from the client.
	//
	// Use the `AWSManagedRulesBotControlRuleSet` configuration object to configure the protection level that you want the Bot Control rule group to use.
	ManagedRuleGroupConfigs interface{} `field:"optional" json:"managedRuleGroupConfigs" yaml:"managedRuleGroupConfigs"`
	// Action settings to use in the place of the rule actions that are configured inside the rule group.
	//
	// You specify one override for each rule whose action you want to change.
	//
	// You can use overrides for testing, for example you can override all of rule actions to `Count` and then monitor the resulting count metrics to understand how the rule group would handle your web traffic. You can also permanently override some or all actions, to modify how the rule group manages your web traffic.
	RuleActionOverrides interface{} `field:"optional" json:"ruleActionOverrides" yaml:"ruleActionOverrides"`
	// An optional nested statement that narrows the scope of the web requests that are evaluated by the managed rule group.
	//
	// Requests are only evaluated by the rule group if they match the scope-down statement. You can use any nestable `Statement` in the scope-down statement, and you can nest statements at any level, the same as you can for a rule statement.
	ScopeDownStatement interface{} `field:"optional" json:"scopeDownStatement" yaml:"scopeDownStatement"`
	// The version of the managed rule group to use.
	//
	// If you specify this, the version setting is fixed until you change it. If you don't specify this, AWS WAF uses the vendor's default version, and then keeps the version at the vendor's default when the vendor updates the managed rule group settings.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

