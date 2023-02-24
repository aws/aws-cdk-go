package awswafv2


// A logical rule statement used to combine other rule statements with OR logic.
//
// You provide more than one `Statement` within the `OrStatement` .
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
//   orStatementProperty := &OrStatementProperty{
//   	Statements: []interface{}{
//   		&statementProperty{
//   			AndStatement: &AndStatementProperty{
//   				Statements: []interface{}{
//   					statementProperty_,
//   				},
//   			},
//   			ByteMatchStatement: &ByteMatchStatementProperty{
//   				FieldToMatch: &FieldToMatchProperty{
//   					AllQueryArguments: allQueryArguments,
//   					Body: &BodyProperty{
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Cookies: &CookiesProperty{
//   						MatchPattern: &CookieMatchPatternProperty{
//   							All: all,
//   							ExcludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							IncludedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Headers: &HeadersProperty{
//   						MatchPattern: &HeaderMatchPatternProperty{
//   							All: all,
//   							ExcludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							IncludedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					JsonBody: &JsonBodyProperty{
//   						MatchPattern: &JsonMatchPatternProperty{
//   							All: all,
//   							IncludedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Method: method,
//   					QueryString: queryString,
//   					SingleHeader: singleHeader,
//   					SingleQueryArgument: singleQueryArgument,
//   					UriPath: uriPath,
//   				},
//   				PositionalConstraint: jsii.String("positionalConstraint"),
//   				TextTransformations: []interface{}{
//   					&TextTransformationProperty{
//   						Priority: jsii.Number(123),
//   						Type: jsii.String("type"),
//   					},
//   				},
//
//   				// the properties below are optional
//   				SearchString: jsii.String("searchString"),
//   				SearchStringBase64: jsii.String("searchStringBase64"),
//   			},
//   			GeoMatchStatement: &GeoMatchStatementProperty{
//   				CountryCodes: []*string{
//   					jsii.String("countryCodes"),
//   				},
//   				ForwardedIpConfig: &ForwardedIPConfigurationProperty{
//   					FallbackBehavior: jsii.String("fallbackBehavior"),
//   					HeaderName: jsii.String("headerName"),
//   				},
//   			},
//   			IpSetReferenceStatement: map[string]interface{}{
//   				"arn": jsii.String("arn"),
//
//   				// the properties below are optional
//   				"ipSetForwardedIpConfig": map[string]*string{
//   					"fallbackBehavior": jsii.String("fallbackBehavior"),
//   					"headerName": jsii.String("headerName"),
//   					"position": jsii.String("position"),
//   				},
//   			},
//   			LabelMatchStatement: &LabelMatchStatementProperty{
//   				Key: jsii.String("key"),
//   				Scope: jsii.String("scope"),
//   			},
//   			ManagedRuleGroupStatement: &ManagedRuleGroupStatementProperty{
//   				Name: jsii.String("name"),
//   				VendorName: jsii.String("vendorName"),
//
//   				// the properties below are optional
//   				ExcludedRules: []interface{}{
//   					&ExcludedRuleProperty{
//   						Name: jsii.String("name"),
//   					},
//   				},
//   				ManagedRuleGroupConfigs: []interface{}{
//   					&ManagedRuleGroupConfigProperty{
//   						AwsManagedRulesBotControlRuleSet: &AWSManagedRulesBotControlRuleSetProperty{
//   							InspectionLevel: jsii.String("inspectionLevel"),
//   						},
//   						LoginPath: jsii.String("loginPath"),
//   						PasswordField: &FieldIdentifierProperty{
//   							Identifier: jsii.String("identifier"),
//   						},
//   						PayloadType: jsii.String("payloadType"),
//   						UsernameField: &FieldIdentifierProperty{
//   							Identifier: jsii.String("identifier"),
//   						},
//   					},
//   				},
//   				RuleActionOverrides: []interface{}{
//   					&RuleActionOverrideProperty{
//   						ActionToUse: &RuleActionProperty{
//   							Allow: &AllowActionProperty{
//   								CustomRequestHandling: &CustomRequestHandlingProperty{
//   									InsertHeaders: []interface{}{
//   										&CustomHTTPHeaderProperty{
//   											Name: jsii.String("name"),
//   											Value: jsii.String("value"),
//   										},
//   									},
//   								},
//   							},
//   							Block: &BlockActionProperty{
//   								CustomResponse: &CustomResponseProperty{
//   									ResponseCode: jsii.Number(123),
//
//   									// the properties below are optional
//   									CustomResponseBodyKey: jsii.String("customResponseBodyKey"),
//   									ResponseHeaders: []interface{}{
//   										&CustomHTTPHeaderProperty{
//   											Name: jsii.String("name"),
//   											Value: jsii.String("value"),
//   										},
//   									},
//   								},
//   							},
//   							Captcha: &CaptchaActionProperty{
//   								CustomRequestHandling: &CustomRequestHandlingProperty{
//   									InsertHeaders: []interface{}{
//   										&CustomHTTPHeaderProperty{
//   											Name: jsii.String("name"),
//   											Value: jsii.String("value"),
//   										},
//   									},
//   								},
//   							},
//   							Challenge: &ChallengeActionProperty{
//   								CustomRequestHandling: &CustomRequestHandlingProperty{
//   									InsertHeaders: []interface{}{
//   										&CustomHTTPHeaderProperty{
//   											Name: jsii.String("name"),
//   											Value: jsii.String("value"),
//   										},
//   									},
//   								},
//   							},
//   							Count: &CountActionProperty{
//   								CustomRequestHandling: &CustomRequestHandlingProperty{
//   									InsertHeaders: []interface{}{
//   										&CustomHTTPHeaderProperty{
//   											Name: jsii.String("name"),
//   											Value: jsii.String("value"),
//   										},
//   									},
//   								},
//   							},
//   						},
//   						Name: jsii.String("name"),
//   					},
//   				},
//   				ScopeDownStatement: statementProperty_,
//   				Version: jsii.String("version"),
//   			},
//   			NotStatement: &NotStatementProperty{
//   				Statement: statementProperty_,
//   			},
//   			OrStatement: &OrStatementProperty{
//   				Statements: []interface{}{
//   					statementProperty_,
//   				},
//   			},
//   			RateBasedStatement: &RateBasedStatementProperty{
//   				AggregateKeyType: jsii.String("aggregateKeyType"),
//   				Limit: jsii.Number(123),
//
//   				// the properties below are optional
//   				ForwardedIpConfig: &ForwardedIPConfigurationProperty{
//   					FallbackBehavior: jsii.String("fallbackBehavior"),
//   					HeaderName: jsii.String("headerName"),
//   				},
//   				ScopeDownStatement: statementProperty_,
//   			},
//   			RegexMatchStatement: &RegexMatchStatementProperty{
//   				FieldToMatch: &FieldToMatchProperty{
//   					AllQueryArguments: allQueryArguments,
//   					Body: &BodyProperty{
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Cookies: &CookiesProperty{
//   						MatchPattern: &CookieMatchPatternProperty{
//   							All: all,
//   							ExcludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							IncludedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Headers: &HeadersProperty{
//   						MatchPattern: &HeaderMatchPatternProperty{
//   							All: all,
//   							ExcludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							IncludedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					JsonBody: &JsonBodyProperty{
//   						MatchPattern: &JsonMatchPatternProperty{
//   							All: all,
//   							IncludedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Method: method,
//   					QueryString: queryString,
//   					SingleHeader: singleHeader,
//   					SingleQueryArgument: singleQueryArgument,
//   					UriPath: uriPath,
//   				},
//   				RegexString: jsii.String("regexString"),
//   				TextTransformations: []interface{}{
//   					&TextTransformationProperty{
//   						Priority: jsii.Number(123),
//   						Type: jsii.String("type"),
//   					},
//   				},
//   			},
//   			RegexPatternSetReferenceStatement: &RegexPatternSetReferenceStatementProperty{
//   				Arn: jsii.String("arn"),
//   				FieldToMatch: &FieldToMatchProperty{
//   					AllQueryArguments: allQueryArguments,
//   					Body: &BodyProperty{
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Cookies: &CookiesProperty{
//   						MatchPattern: &CookieMatchPatternProperty{
//   							All: all,
//   							ExcludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							IncludedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Headers: &HeadersProperty{
//   						MatchPattern: &HeaderMatchPatternProperty{
//   							All: all,
//   							ExcludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							IncludedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					JsonBody: &JsonBodyProperty{
//   						MatchPattern: &JsonMatchPatternProperty{
//   							All: all,
//   							IncludedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Method: method,
//   					QueryString: queryString,
//   					SingleHeader: singleHeader,
//   					SingleQueryArgument: singleQueryArgument,
//   					UriPath: uriPath,
//   				},
//   				TextTransformations: []interface{}{
//   					&TextTransformationProperty{
//   						Priority: jsii.Number(123),
//   						Type: jsii.String("type"),
//   					},
//   				},
//   			},
//   			RuleGroupReferenceStatement: &RuleGroupReferenceStatementProperty{
//   				Arn: jsii.String("arn"),
//
//   				// the properties below are optional
//   				ExcludedRules: []interface{}{
//   					&ExcludedRuleProperty{
//   						Name: jsii.String("name"),
//   					},
//   				},
//   				RuleActionOverrides: []interface{}{
//   					&RuleActionOverrideProperty{
//   						ActionToUse: &RuleActionProperty{
//   							Allow: &AllowActionProperty{
//   								CustomRequestHandling: &CustomRequestHandlingProperty{
//   									InsertHeaders: []interface{}{
//   										&CustomHTTPHeaderProperty{
//   											Name: jsii.String("name"),
//   											Value: jsii.String("value"),
//   										},
//   									},
//   								},
//   							},
//   							Block: &BlockActionProperty{
//   								CustomResponse: &CustomResponseProperty{
//   									ResponseCode: jsii.Number(123),
//
//   									// the properties below are optional
//   									CustomResponseBodyKey: jsii.String("customResponseBodyKey"),
//   									ResponseHeaders: []interface{}{
//   										&CustomHTTPHeaderProperty{
//   											Name: jsii.String("name"),
//   											Value: jsii.String("value"),
//   										},
//   									},
//   								},
//   							},
//   							Captcha: &CaptchaActionProperty{
//   								CustomRequestHandling: &CustomRequestHandlingProperty{
//   									InsertHeaders: []interface{}{
//   										&CustomHTTPHeaderProperty{
//   											Name: jsii.String("name"),
//   											Value: jsii.String("value"),
//   										},
//   									},
//   								},
//   							},
//   							Challenge: &ChallengeActionProperty{
//   								CustomRequestHandling: &CustomRequestHandlingProperty{
//   									InsertHeaders: []interface{}{
//   										&CustomHTTPHeaderProperty{
//   											Name: jsii.String("name"),
//   											Value: jsii.String("value"),
//   										},
//   									},
//   								},
//   							},
//   							Count: &CountActionProperty{
//   								CustomRequestHandling: &CustomRequestHandlingProperty{
//   									InsertHeaders: []interface{}{
//   										&CustomHTTPHeaderProperty{
//   											Name: jsii.String("name"),
//   											Value: jsii.String("value"),
//   										},
//   									},
//   								},
//   							},
//   						},
//   						Name: jsii.String("name"),
//   					},
//   				},
//   			},
//   			SizeConstraintStatement: &SizeConstraintStatementProperty{
//   				ComparisonOperator: jsii.String("comparisonOperator"),
//   				FieldToMatch: &FieldToMatchProperty{
//   					AllQueryArguments: allQueryArguments,
//   					Body: &BodyProperty{
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Cookies: &CookiesProperty{
//   						MatchPattern: &CookieMatchPatternProperty{
//   							All: all,
//   							ExcludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							IncludedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Headers: &HeadersProperty{
//   						MatchPattern: &HeaderMatchPatternProperty{
//   							All: all,
//   							ExcludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							IncludedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					JsonBody: &JsonBodyProperty{
//   						MatchPattern: &JsonMatchPatternProperty{
//   							All: all,
//   							IncludedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Method: method,
//   					QueryString: queryString,
//   					SingleHeader: singleHeader,
//   					SingleQueryArgument: singleQueryArgument,
//   					UriPath: uriPath,
//   				},
//   				Size: jsii.Number(123),
//   				TextTransformations: []interface{}{
//   					&TextTransformationProperty{
//   						Priority: jsii.Number(123),
//   						Type: jsii.String("type"),
//   					},
//   				},
//   			},
//   			SqliMatchStatement: &SqliMatchStatementProperty{
//   				FieldToMatch: &FieldToMatchProperty{
//   					AllQueryArguments: allQueryArguments,
//   					Body: &BodyProperty{
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Cookies: &CookiesProperty{
//   						MatchPattern: &CookieMatchPatternProperty{
//   							All: all,
//   							ExcludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							IncludedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Headers: &HeadersProperty{
//   						MatchPattern: &HeaderMatchPatternProperty{
//   							All: all,
//   							ExcludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							IncludedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					JsonBody: &JsonBodyProperty{
//   						MatchPattern: &JsonMatchPatternProperty{
//   							All: all,
//   							IncludedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Method: method,
//   					QueryString: queryString,
//   					SingleHeader: singleHeader,
//   					SingleQueryArgument: singleQueryArgument,
//   					UriPath: uriPath,
//   				},
//   				TextTransformations: []interface{}{
//   					&TextTransformationProperty{
//   						Priority: jsii.Number(123),
//   						Type: jsii.String("type"),
//   					},
//   				},
//
//   				// the properties below are optional
//   				SensitivityLevel: jsii.String("sensitivityLevel"),
//   			},
//   			XssMatchStatement: &XssMatchStatementProperty{
//   				FieldToMatch: &FieldToMatchProperty{
//   					AllQueryArguments: allQueryArguments,
//   					Body: &BodyProperty{
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Cookies: &CookiesProperty{
//   						MatchPattern: &CookieMatchPatternProperty{
//   							All: all,
//   							ExcludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							IncludedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Headers: &HeadersProperty{
//   						MatchPattern: &HeaderMatchPatternProperty{
//   							All: all,
//   							ExcludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							IncludedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					JsonBody: &JsonBodyProperty{
//   						MatchPattern: &JsonMatchPatternProperty{
//   							All: all,
//   							IncludedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Method: method,
//   					QueryString: queryString,
//   					SingleHeader: singleHeader,
//   					SingleQueryArgument: singleQueryArgument,
//   					UriPath: uriPath,
//   				},
//   				TextTransformations: []interface{}{
//   					&TextTransformationProperty{
//   						Priority: jsii.Number(123),
//   						Type: jsii.String("type"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnWebACL_OrStatementProperty struct {
	// The statements to combine with OR logic.
	//
	// You can use any statements that can be nested.
	Statements interface{} `field:"required" json:"statements" yaml:"statements"`
}

