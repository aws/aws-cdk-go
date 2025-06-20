package awswafv2


// The processing guidance for a rule, used by AWS WAF to determine whether a web request matches the rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var all interface{}
//   var allQueryArguments interface{}
//   var asn interface{}
//   var forwardedIp interface{}
//   var httpMethod interface{}
//   var ip interface{}
//   var method interface{}
//   var queryString interface{}
//   var singleHeader interface{}
//   var singleQueryArgument interface{}
//   var statementProperty_ statementProperty
//   var uriPath interface{}
//
//   statementProperty := &statementProperty{
//   	AndStatement: &AndStatementProperty{
//   		Statements: []interface{}{
//   			statementProperty_,
//   		},
//   	},
//   	AsnMatchStatement: &AsnMatchStatementProperty{
//   		AsnList: []interface{}{
//   			jsii.Number(123),
//   		},
//   		ForwardedIpConfig: &ForwardedIPConfigurationProperty{
//   			FallbackBehavior: jsii.String("fallbackBehavior"),
//   			HeaderName: jsii.String("headerName"),
//   		},
//   	},
//   	ByteMatchStatement: &ByteMatchStatementProperty{
//   		FieldToMatch: &FieldToMatchProperty{
//   			AllQueryArguments: allQueryArguments,
//   			Body: &BodyProperty{
//   				OversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			Cookies: &CookiesProperty{
//   				MatchPattern: &CookieMatchPatternProperty{
//   					All: all,
//   					ExcludedCookies: []*string{
//   						jsii.String("excludedCookies"),
//   					},
//   					IncludedCookies: []*string{
//   						jsii.String("includedCookies"),
//   					},
//   				},
//   				MatchScope: jsii.String("matchScope"),
//   				OversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			Headers: &HeadersProperty{
//   				MatchPattern: &HeaderMatchPatternProperty{
//   					All: all,
//   					ExcludedHeaders: []*string{
//   						jsii.String("excludedHeaders"),
//   					},
//   					IncludedHeaders: []*string{
//   						jsii.String("includedHeaders"),
//   					},
//   				},
//   				MatchScope: jsii.String("matchScope"),
//   				OversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			Ja3Fingerprint: &JA3FingerprintProperty{
//   				FallbackBehavior: jsii.String("fallbackBehavior"),
//   			},
//   			Ja4Fingerprint: &JA4FingerprintProperty{
//   				FallbackBehavior: jsii.String("fallbackBehavior"),
//   			},
//   			JsonBody: &JsonBodyProperty{
//   				MatchPattern: &JsonMatchPatternProperty{
//   					All: all,
//   					IncludedPaths: []*string{
//   						jsii.String("includedPaths"),
//   					},
//   				},
//   				MatchScope: jsii.String("matchScope"),
//
//   				// the properties below are optional
//   				InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   				OversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			Method: method,
//   			QueryString: queryString,
//   			SingleHeader: singleHeader,
//   			SingleQueryArgument: singleQueryArgument,
//   			UriFragment: &UriFragmentProperty{
//   				FallbackBehavior: jsii.String("fallbackBehavior"),
//   			},
//   			UriPath: uriPath,
//   		},
//   		PositionalConstraint: jsii.String("positionalConstraint"),
//   		TextTransformations: []interface{}{
//   			&TextTransformationProperty{
//   				Priority: jsii.Number(123),
//   				Type: jsii.String("type"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		SearchString: jsii.String("searchString"),
//   		SearchStringBase64: jsii.String("searchStringBase64"),
//   	},
//   	GeoMatchStatement: &GeoMatchStatementProperty{
//   		CountryCodes: []*string{
//   			jsii.String("countryCodes"),
//   		},
//   		ForwardedIpConfig: &ForwardedIPConfigurationProperty{
//   			FallbackBehavior: jsii.String("fallbackBehavior"),
//   			HeaderName: jsii.String("headerName"),
//   		},
//   	},
//   	IpSetReferenceStatement: map[string]interface{}{
//   		"arn": jsii.String("arn"),
//
//   		// the properties below are optional
//   		"ipSetForwardedIpConfig": map[string]*string{
//   			"fallbackBehavior": jsii.String("fallbackBehavior"),
//   			"headerName": jsii.String("headerName"),
//   			"position": jsii.String("position"),
//   		},
//   	},
//   	LabelMatchStatement: &LabelMatchStatementProperty{
//   		Key: jsii.String("key"),
//   		Scope: jsii.String("scope"),
//   	},
//   	NotStatement: &NotStatementProperty{
//   		Statement: statementProperty_,
//   	},
//   	OrStatement: &OrStatementProperty{
//   		Statements: []interface{}{
//   			statementProperty_,
//   		},
//   	},
//   	RateBasedStatement: &RateBasedStatementProperty{
//   		AggregateKeyType: jsii.String("aggregateKeyType"),
//   		Limit: jsii.Number(123),
//
//   		// the properties below are optional
//   		CustomKeys: []interface{}{
//   			&RateBasedStatementCustomKeyProperty{
//   				Asn: asn,
//   				Cookie: &RateLimitCookieProperty{
//   					Name: jsii.String("name"),
//   					TextTransformations: []interface{}{
//   						&TextTransformationProperty{
//   							Priority: jsii.Number(123),
//   							Type: jsii.String("type"),
//   						},
//   					},
//   				},
//   				ForwardedIp: forwardedIp,
//   				Header: &RateLimitHeaderProperty{
//   					Name: jsii.String("name"),
//   					TextTransformations: []interface{}{
//   						&TextTransformationProperty{
//   							Priority: jsii.Number(123),
//   							Type: jsii.String("type"),
//   						},
//   					},
//   				},
//   				HttpMethod: httpMethod,
//   				Ip: ip,
//   				Ja3Fingerprint: &RateLimitJA3FingerprintProperty{
//   					FallbackBehavior: jsii.String("fallbackBehavior"),
//   				},
//   				Ja4Fingerprint: &RateLimitJA4FingerprintProperty{
//   					FallbackBehavior: jsii.String("fallbackBehavior"),
//   				},
//   				LabelNamespace: &RateLimitLabelNamespaceProperty{
//   					Namespace: jsii.String("namespace"),
//   				},
//   				QueryArgument: &RateLimitQueryArgumentProperty{
//   					Name: jsii.String("name"),
//   					TextTransformations: []interface{}{
//   						&TextTransformationProperty{
//   							Priority: jsii.Number(123),
//   							Type: jsii.String("type"),
//   						},
//   					},
//   				},
//   				QueryString: &RateLimitQueryStringProperty{
//   					TextTransformations: []interface{}{
//   						&TextTransformationProperty{
//   							Priority: jsii.Number(123),
//   							Type: jsii.String("type"),
//   						},
//   					},
//   				},
//   				UriPath: &RateLimitUriPathProperty{
//   					TextTransformations: []interface{}{
//   						&TextTransformationProperty{
//   							Priority: jsii.Number(123),
//   							Type: jsii.String("type"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		EvaluationWindowSec: jsii.Number(123),
//   		ForwardedIpConfig: &ForwardedIPConfigurationProperty{
//   			FallbackBehavior: jsii.String("fallbackBehavior"),
//   			HeaderName: jsii.String("headerName"),
//   		},
//   		ScopeDownStatement: statementProperty_,
//   	},
//   	RegexMatchStatement: &RegexMatchStatementProperty{
//   		FieldToMatch: &FieldToMatchProperty{
//   			AllQueryArguments: allQueryArguments,
//   			Body: &BodyProperty{
//   				OversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			Cookies: &CookiesProperty{
//   				MatchPattern: &CookieMatchPatternProperty{
//   					All: all,
//   					ExcludedCookies: []*string{
//   						jsii.String("excludedCookies"),
//   					},
//   					IncludedCookies: []*string{
//   						jsii.String("includedCookies"),
//   					},
//   				},
//   				MatchScope: jsii.String("matchScope"),
//   				OversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			Headers: &HeadersProperty{
//   				MatchPattern: &HeaderMatchPatternProperty{
//   					All: all,
//   					ExcludedHeaders: []*string{
//   						jsii.String("excludedHeaders"),
//   					},
//   					IncludedHeaders: []*string{
//   						jsii.String("includedHeaders"),
//   					},
//   				},
//   				MatchScope: jsii.String("matchScope"),
//   				OversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			Ja3Fingerprint: &JA3FingerprintProperty{
//   				FallbackBehavior: jsii.String("fallbackBehavior"),
//   			},
//   			Ja4Fingerprint: &JA4FingerprintProperty{
//   				FallbackBehavior: jsii.String("fallbackBehavior"),
//   			},
//   			JsonBody: &JsonBodyProperty{
//   				MatchPattern: &JsonMatchPatternProperty{
//   					All: all,
//   					IncludedPaths: []*string{
//   						jsii.String("includedPaths"),
//   					},
//   				},
//   				MatchScope: jsii.String("matchScope"),
//
//   				// the properties below are optional
//   				InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   				OversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			Method: method,
//   			QueryString: queryString,
//   			SingleHeader: singleHeader,
//   			SingleQueryArgument: singleQueryArgument,
//   			UriFragment: &UriFragmentProperty{
//   				FallbackBehavior: jsii.String("fallbackBehavior"),
//   			},
//   			UriPath: uriPath,
//   		},
//   		RegexString: jsii.String("regexString"),
//   		TextTransformations: []interface{}{
//   			&TextTransformationProperty{
//   				Priority: jsii.Number(123),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   	},
//   	RegexPatternSetReferenceStatement: &RegexPatternSetReferenceStatementProperty{
//   		Arn: jsii.String("arn"),
//   		FieldToMatch: &FieldToMatchProperty{
//   			AllQueryArguments: allQueryArguments,
//   			Body: &BodyProperty{
//   				OversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			Cookies: &CookiesProperty{
//   				MatchPattern: &CookieMatchPatternProperty{
//   					All: all,
//   					ExcludedCookies: []*string{
//   						jsii.String("excludedCookies"),
//   					},
//   					IncludedCookies: []*string{
//   						jsii.String("includedCookies"),
//   					},
//   				},
//   				MatchScope: jsii.String("matchScope"),
//   				OversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			Headers: &HeadersProperty{
//   				MatchPattern: &HeaderMatchPatternProperty{
//   					All: all,
//   					ExcludedHeaders: []*string{
//   						jsii.String("excludedHeaders"),
//   					},
//   					IncludedHeaders: []*string{
//   						jsii.String("includedHeaders"),
//   					},
//   				},
//   				MatchScope: jsii.String("matchScope"),
//   				OversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			Ja3Fingerprint: &JA3FingerprintProperty{
//   				FallbackBehavior: jsii.String("fallbackBehavior"),
//   			},
//   			Ja4Fingerprint: &JA4FingerprintProperty{
//   				FallbackBehavior: jsii.String("fallbackBehavior"),
//   			},
//   			JsonBody: &JsonBodyProperty{
//   				MatchPattern: &JsonMatchPatternProperty{
//   					All: all,
//   					IncludedPaths: []*string{
//   						jsii.String("includedPaths"),
//   					},
//   				},
//   				MatchScope: jsii.String("matchScope"),
//
//   				// the properties below are optional
//   				InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   				OversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			Method: method,
//   			QueryString: queryString,
//   			SingleHeader: singleHeader,
//   			SingleQueryArgument: singleQueryArgument,
//   			UriFragment: &UriFragmentProperty{
//   				FallbackBehavior: jsii.String("fallbackBehavior"),
//   			},
//   			UriPath: uriPath,
//   		},
//   		TextTransformations: []interface{}{
//   			&TextTransformationProperty{
//   				Priority: jsii.Number(123),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   	},
//   	SizeConstraintStatement: &SizeConstraintStatementProperty{
//   		ComparisonOperator: jsii.String("comparisonOperator"),
//   		FieldToMatch: &FieldToMatchProperty{
//   			AllQueryArguments: allQueryArguments,
//   			Body: &BodyProperty{
//   				OversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			Cookies: &CookiesProperty{
//   				MatchPattern: &CookieMatchPatternProperty{
//   					All: all,
//   					ExcludedCookies: []*string{
//   						jsii.String("excludedCookies"),
//   					},
//   					IncludedCookies: []*string{
//   						jsii.String("includedCookies"),
//   					},
//   				},
//   				MatchScope: jsii.String("matchScope"),
//   				OversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			Headers: &HeadersProperty{
//   				MatchPattern: &HeaderMatchPatternProperty{
//   					All: all,
//   					ExcludedHeaders: []*string{
//   						jsii.String("excludedHeaders"),
//   					},
//   					IncludedHeaders: []*string{
//   						jsii.String("includedHeaders"),
//   					},
//   				},
//   				MatchScope: jsii.String("matchScope"),
//   				OversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			Ja3Fingerprint: &JA3FingerprintProperty{
//   				FallbackBehavior: jsii.String("fallbackBehavior"),
//   			},
//   			Ja4Fingerprint: &JA4FingerprintProperty{
//   				FallbackBehavior: jsii.String("fallbackBehavior"),
//   			},
//   			JsonBody: &JsonBodyProperty{
//   				MatchPattern: &JsonMatchPatternProperty{
//   					All: all,
//   					IncludedPaths: []*string{
//   						jsii.String("includedPaths"),
//   					},
//   				},
//   				MatchScope: jsii.String("matchScope"),
//
//   				// the properties below are optional
//   				InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   				OversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			Method: method,
//   			QueryString: queryString,
//   			SingleHeader: singleHeader,
//   			SingleQueryArgument: singleQueryArgument,
//   			UriFragment: &UriFragmentProperty{
//   				FallbackBehavior: jsii.String("fallbackBehavior"),
//   			},
//   			UriPath: uriPath,
//   		},
//   		Size: jsii.Number(123),
//   		TextTransformations: []interface{}{
//   			&TextTransformationProperty{
//   				Priority: jsii.Number(123),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   	},
//   	SqliMatchStatement: &SqliMatchStatementProperty{
//   		FieldToMatch: &FieldToMatchProperty{
//   			AllQueryArguments: allQueryArguments,
//   			Body: &BodyProperty{
//   				OversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			Cookies: &CookiesProperty{
//   				MatchPattern: &CookieMatchPatternProperty{
//   					All: all,
//   					ExcludedCookies: []*string{
//   						jsii.String("excludedCookies"),
//   					},
//   					IncludedCookies: []*string{
//   						jsii.String("includedCookies"),
//   					},
//   				},
//   				MatchScope: jsii.String("matchScope"),
//   				OversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			Headers: &HeadersProperty{
//   				MatchPattern: &HeaderMatchPatternProperty{
//   					All: all,
//   					ExcludedHeaders: []*string{
//   						jsii.String("excludedHeaders"),
//   					},
//   					IncludedHeaders: []*string{
//   						jsii.String("includedHeaders"),
//   					},
//   				},
//   				MatchScope: jsii.String("matchScope"),
//   				OversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			Ja3Fingerprint: &JA3FingerprintProperty{
//   				FallbackBehavior: jsii.String("fallbackBehavior"),
//   			},
//   			Ja4Fingerprint: &JA4FingerprintProperty{
//   				FallbackBehavior: jsii.String("fallbackBehavior"),
//   			},
//   			JsonBody: &JsonBodyProperty{
//   				MatchPattern: &JsonMatchPatternProperty{
//   					All: all,
//   					IncludedPaths: []*string{
//   						jsii.String("includedPaths"),
//   					},
//   				},
//   				MatchScope: jsii.String("matchScope"),
//
//   				// the properties below are optional
//   				InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   				OversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			Method: method,
//   			QueryString: queryString,
//   			SingleHeader: singleHeader,
//   			SingleQueryArgument: singleQueryArgument,
//   			UriFragment: &UriFragmentProperty{
//   				FallbackBehavior: jsii.String("fallbackBehavior"),
//   			},
//   			UriPath: uriPath,
//   		},
//   		TextTransformations: []interface{}{
//   			&TextTransformationProperty{
//   				Priority: jsii.Number(123),
//   				Type: jsii.String("type"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		SensitivityLevel: jsii.String("sensitivityLevel"),
//   	},
//   	XssMatchStatement: &XssMatchStatementProperty{
//   		FieldToMatch: &FieldToMatchProperty{
//   			AllQueryArguments: allQueryArguments,
//   			Body: &BodyProperty{
//   				OversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			Cookies: &CookiesProperty{
//   				MatchPattern: &CookieMatchPatternProperty{
//   					All: all,
//   					ExcludedCookies: []*string{
//   						jsii.String("excludedCookies"),
//   					},
//   					IncludedCookies: []*string{
//   						jsii.String("includedCookies"),
//   					},
//   				},
//   				MatchScope: jsii.String("matchScope"),
//   				OversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			Headers: &HeadersProperty{
//   				MatchPattern: &HeaderMatchPatternProperty{
//   					All: all,
//   					ExcludedHeaders: []*string{
//   						jsii.String("excludedHeaders"),
//   					},
//   					IncludedHeaders: []*string{
//   						jsii.String("includedHeaders"),
//   					},
//   				},
//   				MatchScope: jsii.String("matchScope"),
//   				OversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			Ja3Fingerprint: &JA3FingerprintProperty{
//   				FallbackBehavior: jsii.String("fallbackBehavior"),
//   			},
//   			Ja4Fingerprint: &JA4FingerprintProperty{
//   				FallbackBehavior: jsii.String("fallbackBehavior"),
//   			},
//   			JsonBody: &JsonBodyProperty{
//   				MatchPattern: &JsonMatchPatternProperty{
//   					All: all,
//   					IncludedPaths: []*string{
//   						jsii.String("includedPaths"),
//   					},
//   				},
//   				MatchScope: jsii.String("matchScope"),
//
//   				// the properties below are optional
//   				InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   				OversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			Method: method,
//   			QueryString: queryString,
//   			SingleHeader: singleHeader,
//   			SingleQueryArgument: singleQueryArgument,
//   			UriFragment: &UriFragmentProperty{
//   				FallbackBehavior: jsii.String("fallbackBehavior"),
//   			},
//   			UriPath: uriPath,
//   		},
//   		TextTransformations: []interface{}{
//   			&TextTransformationProperty{
//   				Priority: jsii.Number(123),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-statement.html
//
type CfnRuleGroup_StatementProperty struct {
	// A logical rule statement used to combine other rule statements with AND logic.
	//
	// You provide more than one `Statement` within the `AndStatement` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-statement.html#cfn-wafv2-rulegroup-statement-andstatement
	//
	AndStatement interface{} `field:"optional" json:"andStatement" yaml:"andStatement"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-statement.html#cfn-wafv2-rulegroup-statement-asnmatchstatement
	//
	AsnMatchStatement interface{} `field:"optional" json:"asnMatchStatement" yaml:"asnMatchStatement"`
	// A rule statement that defines a string match search for AWS WAF to apply to web requests.
	//
	// The byte match statement provides the bytes to search for, the location in requests that you want AWS WAF to search, and other settings. The bytes to search for are typically a string that corresponds with ASCII characters. In the AWS WAF console and the developer guide, this is called a string match statement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-statement.html#cfn-wafv2-rulegroup-statement-bytematchstatement
	//
	ByteMatchStatement interface{} `field:"optional" json:"byteMatchStatement" yaml:"byteMatchStatement"`
	// A rule statement that labels web requests by country and region and that matches against web requests based on country code.
	//
	// A geo match rule labels every request that it inspects regardless of whether it finds a match.
	//
	// - To manage requests only by country, you can use this statement by itself and specify the countries that you want to match against in the `CountryCodes` array.
	// - Otherwise, configure your geo match rule with Count action so that it only labels requests. Then, add one or more label match rules to run after the geo match rule and configure them to match against the geographic labels and handle the requests as needed.
	//
	// AWS WAF labels requests using the alpha-2 country and region codes from the International Organization for Standardization (ISO) 3166 standard. AWS WAF determines the codes using either the IP address in the web request origin or, if you specify it, the address in the geo match `ForwardedIPConfig` .
	//
	// If you use the web request origin, the label formats are `awswaf:clientip:geo:region:<ISO country code>-<ISO region code>` and `awswaf:clientip:geo:country:<ISO country code>` .
	//
	// If you use a forwarded IP address, the label formats are `awswaf:forwardedip:geo:region:<ISO country code>-<ISO region code>` and `awswaf:forwardedip:geo:country:<ISO country code>` .
	//
	// For additional details, see [Geographic match rule statement](https://docs.aws.amazon.com/waf/latest/developerguide/waf-rule-statement-type-geo-match.html) in the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-statement.html#cfn-wafv2-rulegroup-statement-geomatchstatement
	//
	GeoMatchStatement interface{} `field:"optional" json:"geoMatchStatement" yaml:"geoMatchStatement"`
	// A rule statement used to detect web requests coming from particular IP addresses or address ranges.
	//
	// To use this, create an `IPSet` that specifies the addresses you want to detect, then use the ARN of that set in this statement.
	//
	// Each IP set rule statement references an IP set. You create and maintain the set independent of your rules. This allows you to use the single set in multiple rules. When you update the referenced set, AWS WAF automatically updates all rules that reference it.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-statement.html#cfn-wafv2-rulegroup-statement-ipsetreferencestatement
	//
	IpSetReferenceStatement interface{} `field:"optional" json:"ipSetReferenceStatement" yaml:"ipSetReferenceStatement"`
	// A rule statement to match against labels that have been added to the web request by rules that have already run in the web ACL.
	//
	// The label match statement provides the label or namespace string to search for. The label string can represent a part or all of the fully qualified label name that had been added to the web request. Fully qualified labels have a prefix, optional namespaces, and label name. The prefix identifies the rule group or web ACL context of the rule that added the label. If you do not provide the fully qualified name in your label match string, AWS WAF performs the search for labels that were added in the same context as the label match statement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-statement.html#cfn-wafv2-rulegroup-statement-labelmatchstatement
	//
	LabelMatchStatement interface{} `field:"optional" json:"labelMatchStatement" yaml:"labelMatchStatement"`
	// A logical rule statement used to negate the results of another rule statement.
	//
	// You provide one `Statement` within the `NotStatement` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-statement.html#cfn-wafv2-rulegroup-statement-notstatement
	//
	NotStatement interface{} `field:"optional" json:"notStatement" yaml:"notStatement"`
	// A logical rule statement used to combine other rule statements with OR logic.
	//
	// You provide more than one `Statement` within the `OrStatement` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-statement.html#cfn-wafv2-rulegroup-statement-orstatement
	//
	OrStatement interface{} `field:"optional" json:"orStatement" yaml:"orStatement"`
	// A rate-based rule counts incoming requests and rate limits requests when they are coming at too fast a rate.
	//
	// The rule categorizes requests according to your aggregation criteria, collects them into aggregation instances, and counts and rate limits the requests for each instance.
	//
	// > If you change any of these settings in a rule that's currently in use, the change resets the rule's rate limiting counts. This can pause the rule's rate limiting activities for up to a minute.
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-statement.html#cfn-wafv2-rulegroup-statement-ratebasedstatement
	//
	RateBasedStatement interface{} `field:"optional" json:"rateBasedStatement" yaml:"rateBasedStatement"`
	// A rule statement used to search web request components for a match against a single regular expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-statement.html#cfn-wafv2-rulegroup-statement-regexmatchstatement
	//
	RegexMatchStatement interface{} `field:"optional" json:"regexMatchStatement" yaml:"regexMatchStatement"`
	// A rule statement used to search web request components for matches with regular expressions.
	//
	// To use this, create a `RegexPatternSet` that specifies the expressions that you want to detect, then use the ARN of that set in this statement. A web request matches the pattern set rule statement if the request component matches any of the patterns in the set.
	//
	// Each regex pattern set rule statement references a regex pattern set. You create and maintain the set independent of your rules. This allows you to use the single set in multiple rules. When you update the referenced set, AWS WAF automatically updates all rules that reference it.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-statement.html#cfn-wafv2-rulegroup-statement-regexpatternsetreferencestatement
	//
	RegexPatternSetReferenceStatement interface{} `field:"optional" json:"regexPatternSetReferenceStatement" yaml:"regexPatternSetReferenceStatement"`
	// A rule statement that compares a number of bytes against the size of a request component, using a comparison operator, such as greater than (>) or less than (<).
	//
	// For example, you can use a size constraint statement to look for query strings that are longer than 100 bytes.
	//
	// If you configure AWS WAF to inspect the request body, AWS WAF inspects only the number of bytes in the body up to the limit for the web ACL and protected resource type. If you know that the request body for your web requests should never exceed the inspection limit, you can use a size constraint statement to block requests that have a larger request body size. For more information about the inspection limits, see `Body` and `JsonBody` settings for the `FieldToMatch` data type.
	//
	// If you choose URI for the value of Part of the request to filter on, the slash (/) in the URI counts as one character. For example, the URI `/logo.jpg` is nine characters long.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-statement.html#cfn-wafv2-rulegroup-statement-sizeconstraintstatement
	//
	SizeConstraintStatement interface{} `field:"optional" json:"sizeConstraintStatement" yaml:"sizeConstraintStatement"`
	// A rule statement that inspects for malicious SQL code.
	//
	// Attackers insert malicious SQL code into web requests to do things like modify your database or extract data from it.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-statement.html#cfn-wafv2-rulegroup-statement-sqlimatchstatement
	//
	SqliMatchStatement interface{} `field:"optional" json:"sqliMatchStatement" yaml:"sqliMatchStatement"`
	// A rule statement that inspects for cross-site scripting (XSS) attacks.
	//
	// In XSS attacks, the attacker uses vulnerabilities in a benign website as a vehicle to inject malicious client-site scripts into other legitimate web browsers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-statement.html#cfn-wafv2-rulegroup-statement-xssmatchstatement
	//
	XssMatchStatement interface{} `field:"optional" json:"xssMatchStatement" yaml:"xssMatchStatement"`
}

