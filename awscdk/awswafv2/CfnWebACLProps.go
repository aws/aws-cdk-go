package awswafv2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnWebACL`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var all interface{}
//   var allQueryArguments interface{}
//   var count interface{}
//   var method interface{}
//   var none interface{}
//   var queryString interface{}
//   var singleHeader interface{}
//   var singleQueryArgument interface{}
//   var statementProperty_ statementProperty
//   var uriPath interface{}
//
//   cfnWebACLProps := &CfnWebACLProps{
//   	DefaultAction: &DefaultActionProperty{
//   		Allow: &AllowActionProperty{
//   			CustomRequestHandling: &CustomRequestHandlingProperty{
//   				InsertHeaders: []interface{}{
//   					&CustomHTTPHeaderProperty{
//   						Name: jsii.String("name"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   		Block: &BlockActionProperty{
//   			CustomResponse: &CustomResponseProperty{
//   				ResponseCode: jsii.Number(123),
//
//   				// the properties below are optional
//   				CustomResponseBodyKey: jsii.String("customResponseBodyKey"),
//   				ResponseHeaders: []interface{}{
//   					&CustomHTTPHeaderProperty{
//   						Name: jsii.String("name"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	Scope: jsii.String("scope"),
//   	VisibilityConfig: &VisibilityConfigProperty{
//   		CloudWatchMetricsEnabled: jsii.Boolean(false),
//   		MetricName: jsii.String("metricName"),
//   		SampledRequestsEnabled: jsii.Boolean(false),
//   	},
//
//   	// the properties below are optional
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
//   			Name: jsii.String("name"),
//   			Priority: jsii.Number(123),
//   			Statement: &statementProperty{
//   				AndStatement: &AndStatementProperty{
//   					Statements: []interface{}{
//   						statementProperty_,
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
//   						JsonBody: &JsonBodyProperty{
//   							MatchPattern: &JsonMatchPatternProperty{
//   								All: all,
//   								IncludedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							MatchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						Method: method,
//   						QueryString: queryString,
//   						SingleHeader: singleHeader,
//   						SingleQueryArgument: singleQueryArgument,
//   						UriPath: uriPath,
//   					},
//   					PositionalConstraint: jsii.String("positionalConstraint"),
//   					TextTransformations: []interface{}{
//   						&TextTransformationProperty{
//   							Priority: jsii.Number(123),
//   							Type: jsii.String("type"),
//   						},
//   					},
//
//   					// the properties below are optional
//   					SearchString: jsii.String("searchString"),
//   					SearchStringBase64: jsii.String("searchStringBase64"),
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
//   				IpSetReferenceStatement: map[string]interface{}{
//   					"arn": jsii.String("arn"),
//
//   					// the properties below are optional
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
//   				ManagedRuleGroupStatement: &ManagedRuleGroupStatementProperty{
//   					Name: jsii.String("name"),
//   					VendorName: jsii.String("vendorName"),
//
//   					// the properties below are optional
//   					ExcludedRules: []interface{}{
//   						&ExcludedRuleProperty{
//   							Name: jsii.String("name"),
//   						},
//   					},
//   					ManagedRuleGroupConfigs: []interface{}{
//   						&ManagedRuleGroupConfigProperty{
//   							AwsManagedRulesAtpRuleSet: &AWSManagedRulesATPRuleSetProperty{
//   								LoginPath: jsii.String("loginPath"),
//
//   								// the properties below are optional
//   								RequestInspection: &RequestInspectionProperty{
//   									PasswordField: &FieldIdentifierProperty{
//   										Identifier: jsii.String("identifier"),
//   									},
//   									PayloadType: jsii.String("payloadType"),
//   									UsernameField: &FieldIdentifierProperty{
//   										Identifier: jsii.String("identifier"),
//   									},
//   								},
//   								ResponseInspection: &ResponseInspectionProperty{
//   									BodyContains: &ResponseInspectionBodyContainsProperty{
//   										FailureStrings: []*string{
//   											jsii.String("failureStrings"),
//   										},
//   										SuccessStrings: []*string{
//   											jsii.String("successStrings"),
//   										},
//   									},
//   									Header: &ResponseInspectionHeaderProperty{
//   										FailureValues: []*string{
//   											jsii.String("failureValues"),
//   										},
//   										Name: jsii.String("name"),
//   										SuccessValues: []*string{
//   											jsii.String("successValues"),
//   										},
//   									},
//   									Json: &ResponseInspectionJsonProperty{
//   										FailureValues: []*string{
//   											jsii.String("failureValues"),
//   										},
//   										Identifier: jsii.String("identifier"),
//   										SuccessValues: []*string{
//   											jsii.String("successValues"),
//   										},
//   									},
//   									StatusCode: &ResponseInspectionStatusCodeProperty{
//   										FailureCodes: []interface{}{
//   											jsii.Number(123),
//   										},
//   										SuccessCodes: []interface{}{
//   											jsii.Number(123),
//   										},
//   									},
//   								},
//   							},
//   							AwsManagedRulesBotControlRuleSet: &AWSManagedRulesBotControlRuleSetProperty{
//   								InspectionLevel: jsii.String("inspectionLevel"),
//   							},
//   							LoginPath: jsii.String("loginPath"),
//   							PasswordField: &FieldIdentifierProperty{
//   								Identifier: jsii.String("identifier"),
//   							},
//   							PayloadType: jsii.String("payloadType"),
//   							UsernameField: &FieldIdentifierProperty{
//   								Identifier: jsii.String("identifier"),
//   							},
//   						},
//   					},
//   					RuleActionOverrides: []interface{}{
//   						&RuleActionOverrideProperty{
//   							ActionToUse: &RuleActionProperty{
//   								Allow: &AllowActionProperty{
//   									CustomRequestHandling: &CustomRequestHandlingProperty{
//   										InsertHeaders: []interface{}{
//   											&CustomHTTPHeaderProperty{
//   												Name: jsii.String("name"),
//   												Value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   								Block: &BlockActionProperty{
//   									CustomResponse: &CustomResponseProperty{
//   										ResponseCode: jsii.Number(123),
//
//   										// the properties below are optional
//   										CustomResponseBodyKey: jsii.String("customResponseBodyKey"),
//   										ResponseHeaders: []interface{}{
//   											&CustomHTTPHeaderProperty{
//   												Name: jsii.String("name"),
//   												Value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   								Captcha: &CaptchaActionProperty{
//   									CustomRequestHandling: &CustomRequestHandlingProperty{
//   										InsertHeaders: []interface{}{
//   											&CustomHTTPHeaderProperty{
//   												Name: jsii.String("name"),
//   												Value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   								Challenge: &ChallengeActionProperty{
//   									CustomRequestHandling: &CustomRequestHandlingProperty{
//   										InsertHeaders: []interface{}{
//   											&CustomHTTPHeaderProperty{
//   												Name: jsii.String("name"),
//   												Value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   								Count: &CountActionProperty{
//   									CustomRequestHandling: &CustomRequestHandlingProperty{
//   										InsertHeaders: []interface{}{
//   											&CustomHTTPHeaderProperty{
//   												Name: jsii.String("name"),
//   												Value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   							},
//   							Name: jsii.String("name"),
//   						},
//   					},
//   					ScopeDownStatement: statementProperty_,
//   					Version: jsii.String("version"),
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
//   					Limit: jsii.Number(123),
//
//   					// the properties below are optional
//   					ForwardedIpConfig: &ForwardedIPConfigurationProperty{
//   						FallbackBehavior: jsii.String("fallbackBehavior"),
//   						HeaderName: jsii.String("headerName"),
//   					},
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
//   						JsonBody: &JsonBodyProperty{
//   							MatchPattern: &JsonMatchPatternProperty{
//   								All: all,
//   								IncludedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							MatchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						Method: method,
//   						QueryString: queryString,
//   						SingleHeader: singleHeader,
//   						SingleQueryArgument: singleQueryArgument,
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
//   						JsonBody: &JsonBodyProperty{
//   							MatchPattern: &JsonMatchPatternProperty{
//   								All: all,
//   								IncludedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							MatchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						Method: method,
//   						QueryString: queryString,
//   						SingleHeader: singleHeader,
//   						SingleQueryArgument: singleQueryArgument,
//   						UriPath: uriPath,
//   					},
//   					TextTransformations: []interface{}{
//   						&TextTransformationProperty{
//   							Priority: jsii.Number(123),
//   							Type: jsii.String("type"),
//   						},
//   					},
//   				},
//   				RuleGroupReferenceStatement: &RuleGroupReferenceStatementProperty{
//   					Arn: jsii.String("arn"),
//
//   					// the properties below are optional
//   					ExcludedRules: []interface{}{
//   						&ExcludedRuleProperty{
//   							Name: jsii.String("name"),
//   						},
//   					},
//   					RuleActionOverrides: []interface{}{
//   						&RuleActionOverrideProperty{
//   							ActionToUse: &RuleActionProperty{
//   								Allow: &AllowActionProperty{
//   									CustomRequestHandling: &CustomRequestHandlingProperty{
//   										InsertHeaders: []interface{}{
//   											&CustomHTTPHeaderProperty{
//   												Name: jsii.String("name"),
//   												Value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   								Block: &BlockActionProperty{
//   									CustomResponse: &CustomResponseProperty{
//   										ResponseCode: jsii.Number(123),
//
//   										// the properties below are optional
//   										CustomResponseBodyKey: jsii.String("customResponseBodyKey"),
//   										ResponseHeaders: []interface{}{
//   											&CustomHTTPHeaderProperty{
//   												Name: jsii.String("name"),
//   												Value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   								Captcha: &CaptchaActionProperty{
//   									CustomRequestHandling: &CustomRequestHandlingProperty{
//   										InsertHeaders: []interface{}{
//   											&CustomHTTPHeaderProperty{
//   												Name: jsii.String("name"),
//   												Value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   								Challenge: &ChallengeActionProperty{
//   									CustomRequestHandling: &CustomRequestHandlingProperty{
//   										InsertHeaders: []interface{}{
//   											&CustomHTTPHeaderProperty{
//   												Name: jsii.String("name"),
//   												Value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   								Count: &CountActionProperty{
//   									CustomRequestHandling: &CustomRequestHandlingProperty{
//   										InsertHeaders: []interface{}{
//   											&CustomHTTPHeaderProperty{
//   												Name: jsii.String("name"),
//   												Value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   							},
//   							Name: jsii.String("name"),
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
//   						JsonBody: &JsonBodyProperty{
//   							MatchPattern: &JsonMatchPatternProperty{
//   								All: all,
//   								IncludedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							MatchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						Method: method,
//   						QueryString: queryString,
//   						SingleHeader: singleHeader,
//   						SingleQueryArgument: singleQueryArgument,
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
//   						JsonBody: &JsonBodyProperty{
//   							MatchPattern: &JsonMatchPatternProperty{
//   								All: all,
//   								IncludedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							MatchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						Method: method,
//   						QueryString: queryString,
//   						SingleHeader: singleHeader,
//   						SingleQueryArgument: singleQueryArgument,
//   						UriPath: uriPath,
//   					},
//   					TextTransformations: []interface{}{
//   						&TextTransformationProperty{
//   							Priority: jsii.Number(123),
//   							Type: jsii.String("type"),
//   						},
//   					},
//
//   					// the properties below are optional
//   					SensitivityLevel: jsii.String("sensitivityLevel"),
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
//   						JsonBody: &JsonBodyProperty{
//   							MatchPattern: &JsonMatchPatternProperty{
//   								All: all,
//   								IncludedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							MatchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						Method: method,
//   						QueryString: queryString,
//   						SingleHeader: singleHeader,
//   						SingleQueryArgument: singleQueryArgument,
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
//
//   			// the properties below are optional
//   			Action: &RuleActionProperty{
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
//   			OverrideAction: &OverrideActionProperty{
//   				Count: count,
//   				None: none,
//   			},
//   			RuleLabels: []interface{}{
//   				&LabelProperty{
//   					Name: jsii.String("name"),
//   				},
//   			},
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TokenDomains: []*string{
//   		jsii.String("tokenDomains"),
//   	},
//   }
//
type CfnWebACLProps struct {
	// The action to perform if none of the `Rules` contained in the `WebACL` match.
	DefaultAction interface{} `field:"required" json:"defaultAction" yaml:"defaultAction"`
	// Specifies whether this is for an Amazon CloudFront distribution or for a regional application.
	//
	// A regional application can be an Application Load Balancer (ALB), an Amazon API Gateway REST API, an AWS AppSync GraphQL API, an Amazon Cognito user pool, or an AWS App Runner service. Valid Values are `CLOUDFRONT` and `REGIONAL` .
	//
	// > For `CLOUDFRONT` , you must create your WAFv2 resources in the US East (N. Virginia) Region, `us-east-1` .
	//
	// For information about how to define the association of the web ACL with your resource, see `WebACLAssociation` .
	Scope *string `field:"required" json:"scope" yaml:"scope"`
	// Defines and enables Amazon CloudWatch metrics and web request sample collection.
	VisibilityConfig interface{} `field:"required" json:"visibilityConfig" yaml:"visibilityConfig"`
	// Specifies how AWS WAF should handle `CAPTCHA` evaluations for rules that don't have their own `CaptchaConfig` settings.
	//
	// If you don't specify this, AWS WAF uses its default settings for `CaptchaConfig` .
	CaptchaConfig interface{} `field:"optional" json:"captchaConfig" yaml:"captchaConfig"`
	// Specifies how AWS WAF should handle challenge evaluations for rules that don't have their own `ChallengeConfig` settings.
	//
	// If you don't specify this, AWS WAF uses its default settings for `ChallengeConfig` .
	ChallengeConfig interface{} `field:"optional" json:"challengeConfig" yaml:"challengeConfig"`
	// A map of custom response keys and content bodies.
	//
	// When you create a rule with a block action, you can send a custom response to the web request. You define these for the web ACL, and then use them in the rules and default actions that you define in the web ACL.
	//
	// For information about customizing web requests and responses, see [Customizing web requests and responses in AWS WAF](https://docs.aws.amazon.com/waf/latest/developerguide/waf-custom-request-response.html) in the *AWS WAF Developer Guide* .
	//
	// For information about the limits on count and size for custom request and response settings, see [AWS WAF quotas](https://docs.aws.amazon.com/waf/latest/developerguide/limits.html) in the *AWS WAF Developer Guide* .
	CustomResponseBodies interface{} `field:"optional" json:"customResponseBodies" yaml:"customResponseBodies"`
	// A description of the web ACL that helps with identification.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the web ACL.
	//
	// You cannot change the name of a web ACL after you create it.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The rule statements used to identify the web requests that you want to allow, block, or count.
	//
	// Each rule includes one top-level statement that AWS WAF uses to identify matching web requests, and parameters that govern how AWS WAF handles them.
	Rules interface{} `field:"optional" json:"rules" yaml:"rules"`
	// Key:value pairs associated with an AWS resource.
	//
	// The key:value pair can be anything you define. Typically, the tag key represents a category (such as "environment") and the tag value represents a specific value within that category (such as "test," "development," or "production"). You can add up to 50 tags to each AWS resource.
	//
	// > To modify tags on existing resources, use the AWS WAF APIs or command line interface. With AWS CloudFormation , you can only add tags to AWS WAF resources during resource creation.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Specifies the domains that AWS WAF should accept in a web request token.
	//
	// This enables the use of tokens across multiple protected websites. When AWS WAF provides a token, it uses the domain of the AWS resource that the web ACL is protecting. If you don't specify a list of token domains, AWS WAF accepts tokens only for the domain of the protected resource. With a token domain list, AWS WAF accepts the resource's host domain plus all domains in the token domain list, including their prefixed subdomains.
	TokenDomains *[]*string `field:"optional" json:"tokenDomains" yaml:"tokenDomains"`
}

