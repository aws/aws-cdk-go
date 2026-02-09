package awswafv2


// A logical rule statement used to combine other rule statements with AND logic.
//
// You provide more than one `Statement` within the `AndStatement` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var all interface{}
//   var allQueryArguments interface{}
//   var andStatementProperty_ AndStatementProperty
//   var asn interface{}
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
//   andStatementProperty := &AndStatementProperty{
//   	Statements: []interface{}{
//   		&StatementProperty{
//   			AndStatement: andStatementProperty_,
//   			AsnMatchStatement: &AsnMatchStatementProperty{
//   				AsnList: []interface{}{
//   					jsii.Number(123),
//   				},
//   				ForwardedIpConfig: &ForwardedIPConfigurationProperty{
//   					FallbackBehavior: jsii.String("fallbackBehavior"),
//   					HeaderName: jsii.String("headerName"),
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
//   					HeaderOrder: &HeaderOrderProperty{
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
//   					Ja3Fingerprint: &JA3FingerprintProperty{
//   						FallbackBehavior: jsii.String("fallbackBehavior"),
//   					},
//   					Ja4Fingerprint: &JA4FingerprintProperty{
//   						FallbackBehavior: jsii.String("fallbackBehavior"),
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
//   					UriFragment: &UriFragmentProperty{
//   						FallbackBehavior: jsii.String("fallbackBehavior"),
//   					},
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
//   				CustomKeys: []interface{}{
//   					&RateBasedStatementCustomKeyProperty{
//   						Asn: asn,
//   						Cookie: &RateLimitCookieProperty{
//   							Name: jsii.String("name"),
//   							TextTransformations: []interface{}{
//   								&TextTransformationProperty{
//   									Priority: jsii.Number(123),
//   									Type: jsii.String("type"),
//   								},
//   							},
//   						},
//   						ForwardedIp: forwardedIp,
//   						Header: &RateLimitHeaderProperty{
//   							Name: jsii.String("name"),
//   							TextTransformations: []interface{}{
//   								&TextTransformationProperty{
//   									Priority: jsii.Number(123),
//   									Type: jsii.String("type"),
//   								},
//   							},
//   						},
//   						HttpMethod: httpMethod,
//   						Ip: ip,
//   						Ja3Fingerprint: &RateLimitJA3FingerprintProperty{
//   							FallbackBehavior: jsii.String("fallbackBehavior"),
//   						},
//   						Ja4Fingerprint: &RateLimitJA4FingerprintProperty{
//   							FallbackBehavior: jsii.String("fallbackBehavior"),
//   						},
//   						LabelNamespace: &RateLimitLabelNamespaceProperty{
//   							Namespace: jsii.String("namespace"),
//   						},
//   						QueryArgument: &RateLimitQueryArgumentProperty{
//   							Name: jsii.String("name"),
//   							TextTransformations: []interface{}{
//   								&TextTransformationProperty{
//   									Priority: jsii.Number(123),
//   									Type: jsii.String("type"),
//   								},
//   							},
//   						},
//   						QueryString: &RateLimitQueryStringProperty{
//   							TextTransformations: []interface{}{
//   								&TextTransformationProperty{
//   									Priority: jsii.Number(123),
//   									Type: jsii.String("type"),
//   								},
//   							},
//   						},
//   						UriPath: &RateLimitUriPathProperty{
//   							TextTransformations: []interface{}{
//   								&TextTransformationProperty{
//   									Priority: jsii.Number(123),
//   									Type: jsii.String("type"),
//   								},
//   							},
//   						},
//   					},
//   				},
//   				EvaluationWindowSec: jsii.Number(123),
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
//   					HeaderOrder: &HeaderOrderProperty{
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
//   					Ja3Fingerprint: &JA3FingerprintProperty{
//   						FallbackBehavior: jsii.String("fallbackBehavior"),
//   					},
//   					Ja4Fingerprint: &JA4FingerprintProperty{
//   						FallbackBehavior: jsii.String("fallbackBehavior"),
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
//   					UriFragment: &UriFragmentProperty{
//   						FallbackBehavior: jsii.String("fallbackBehavior"),
//   					},
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
//   					HeaderOrder: &HeaderOrderProperty{
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
//   					Ja3Fingerprint: &JA3FingerprintProperty{
//   						FallbackBehavior: jsii.String("fallbackBehavior"),
//   					},
//   					Ja4Fingerprint: &JA4FingerprintProperty{
//   						FallbackBehavior: jsii.String("fallbackBehavior"),
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
//   					UriFragment: &UriFragmentProperty{
//   						FallbackBehavior: jsii.String("fallbackBehavior"),
//   					},
//   					UriPath: uriPath,
//   				},
//   				TextTransformations: []interface{}{
//   					&TextTransformationProperty{
//   						Priority: jsii.Number(123),
//   						Type: jsii.String("type"),
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
//   					HeaderOrder: &HeaderOrderProperty{
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
//   					Ja3Fingerprint: &JA3FingerprintProperty{
//   						FallbackBehavior: jsii.String("fallbackBehavior"),
//   					},
//   					Ja4Fingerprint: &JA4FingerprintProperty{
//   						FallbackBehavior: jsii.String("fallbackBehavior"),
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
//   					UriFragment: &UriFragmentProperty{
//   						FallbackBehavior: jsii.String("fallbackBehavior"),
//   					},
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
//   					HeaderOrder: &HeaderOrderProperty{
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
//   					Ja3Fingerprint: &JA3FingerprintProperty{
//   						FallbackBehavior: jsii.String("fallbackBehavior"),
//   					},
//   					Ja4Fingerprint: &JA4FingerprintProperty{
//   						FallbackBehavior: jsii.String("fallbackBehavior"),
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
//   					UriFragment: &UriFragmentProperty{
//   						FallbackBehavior: jsii.String("fallbackBehavior"),
//   					},
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
//   					HeaderOrder: &HeaderOrderProperty{
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
//   					Ja3Fingerprint: &JA3FingerprintProperty{
//   						FallbackBehavior: jsii.String("fallbackBehavior"),
//   					},
//   					Ja4Fingerprint: &JA4FingerprintProperty{
//   						FallbackBehavior: jsii.String("fallbackBehavior"),
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
//   					UriFragment: &UriFragmentProperty{
//   						FallbackBehavior: jsii.String("fallbackBehavior"),
//   					},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-andstatement.html
//
type CfnRuleGroup_AndStatementProperty struct {
	// The statements to combine with AND logic.
	//
	// You can use any statements that can be nested.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-andstatement.html#cfn-wafv2-rulegroup-andstatement-statements
	//
	Statements interface{} `field:"required" json:"statements" yaml:"statements"`
}

