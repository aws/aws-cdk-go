package awswafv2


// A rate-based rule tracks the rate of requests for each originating IP address, and triggers the rule action when the rate exceeds a limit that you specify on the number of requests in any 5-minute time span.
//
// You can use this to put a temporary block on requests from an IP address that is sending excessive requests.
//
// AWS WAF tracks and manages web requests separately for each instance of a rate-based rule that you use. For example, if you provide the same rate-based rule settings in two web ACLs, each of the two rule statements represents a separate instance of the rate-based rule and gets its own tracking and management by AWS WAF . If you define a rate-based rule inside a rule group, and then use that rule group in multiple places, each use creates a separate instance of the rate-based rule that gets its own tracking and management by AWS WAF .
//
// When the rule action triggers, AWS WAF blocks additional requests from the IP address until the request rate falls below the limit.
//
// You can optionally nest another statement inside the rate-based statement, to narrow the scope of the rule so that it only counts requests that match the nested statement. For example, based on recent requests that you have seen from an attacker, you might create a rate-based rule with a nested AND rule statement that contains the following nested statements:
//
// - An IP match statement with an IP set that specified the address 192.0.2.44.
// - A string match statement that searches in the User-Agent header for the string BadBot.
//
// In this rate-based rule, you also define a rate limit. For this example, the rate limit is 1,000. Requests that meet the criteria of both of the nested statements are counted. If the count exceeds 1,000 requests per five minutes, the rule action triggers. Requests that do not meet the criteria of both of the nested statements are not counted towards the rate limit and are not affected by this rule.
//
// You cannot nest a `RateBasedStatement` inside another statement, for example inside a `NotStatement` or `OrStatement` . You can define a `RateBasedStatement` inside a web ACL and inside a rule group.
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
//   rateBasedStatementProperty := &rateBasedStatementProperty{
//   	aggregateKeyType: jsii.String("aggregateKeyType"),
//   	limit: jsii.Number(123),
//
//   	// the properties below are optional
//   	forwardedIpConfig: &forwardedIPConfigurationProperty{
//   		fallbackBehavior: jsii.String("fallbackBehavior"),
//   		headerName: jsii.String("headerName"),
//   	},
//   	scopeDownStatement: &statementProperty{
//   		andStatement: &andStatementProperty{
//   			statements: []interface{}{
//   				statementProperty_,
//   			},
//   		},
//   		byteMatchStatement: &byteMatchStatementProperty{
//   			fieldToMatch: &fieldToMatchProperty{
//   				allQueryArguments: allQueryArguments,
//   				body: &bodyProperty{
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				cookies: &cookiesProperty{
//   					matchPattern: &cookieMatchPatternProperty{
//   						all: all,
//   						excludedCookies: []*string{
//   							jsii.String("excludedCookies"),
//   						},
//   						includedCookies: []*string{
//   							jsii.String("includedCookies"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				headers: &headersProperty{
//   					matchPattern: &headerMatchPatternProperty{
//   						all: all,
//   						excludedHeaders: []*string{
//   							jsii.String("excludedHeaders"),
//   						},
//   						includedHeaders: []*string{
//   							jsii.String("includedHeaders"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				jsonBody: &jsonBodyProperty{
//   					matchPattern: &jsonMatchPatternProperty{
//   						all: all,
//   						includedPaths: []*string{
//   							jsii.String("includedPaths"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//
//   					// the properties below are optional
//   					invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				method: method,
//   				queryString: queryString,
//   				singleHeader: singleHeader,
//   				singleQueryArgument: singleQueryArgument,
//   				uriPath: uriPath,
//   			},
//   			positionalConstraint: jsii.String("positionalConstraint"),
//   			textTransformations: []interface{}{
//   				&textTransformationProperty{
//   					priority: jsii.Number(123),
//   					type: jsii.String("type"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			searchString: jsii.String("searchString"),
//   			searchStringBase64: jsii.String("searchStringBase64"),
//   		},
//   		geoMatchStatement: &geoMatchStatementProperty{
//   			countryCodes: []*string{
//   				jsii.String("countryCodes"),
//   			},
//   			forwardedIpConfig: &forwardedIPConfigurationProperty{
//   				fallbackBehavior: jsii.String("fallbackBehavior"),
//   				headerName: jsii.String("headerName"),
//   			},
//   		},
//   		ipSetReferenceStatement: map[string]interface{}{
//   			"arn": jsii.String("arn"),
//
//   			// the properties below are optional
//   			"ipSetForwardedIpConfig": map[string]*string{
//   				"fallbackBehavior": jsii.String("fallbackBehavior"),
//   				"headerName": jsii.String("headerName"),
//   				"position": jsii.String("position"),
//   			},
//   		},
//   		labelMatchStatement: &labelMatchStatementProperty{
//   			key: jsii.String("key"),
//   			scope: jsii.String("scope"),
//   		},
//   		notStatement: &notStatementProperty{
//   			statement: statementProperty_,
//   		},
//   		orStatement: &orStatementProperty{
//   			statements: []interface{}{
//   				statementProperty_,
//   			},
//   		},
//   		rateBasedStatement: &rateBasedStatementProperty{
//   			aggregateKeyType: jsii.String("aggregateKeyType"),
//   			limit: jsii.Number(123),
//
//   			// the properties below are optional
//   			forwardedIpConfig: &forwardedIPConfigurationProperty{
//   				fallbackBehavior: jsii.String("fallbackBehavior"),
//   				headerName: jsii.String("headerName"),
//   			},
//   			scopeDownStatement: statementProperty_,
//   		},
//   		regexMatchStatement: &regexMatchStatementProperty{
//   			fieldToMatch: &fieldToMatchProperty{
//   				allQueryArguments: allQueryArguments,
//   				body: &bodyProperty{
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				cookies: &cookiesProperty{
//   					matchPattern: &cookieMatchPatternProperty{
//   						all: all,
//   						excludedCookies: []*string{
//   							jsii.String("excludedCookies"),
//   						},
//   						includedCookies: []*string{
//   							jsii.String("includedCookies"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				headers: &headersProperty{
//   					matchPattern: &headerMatchPatternProperty{
//   						all: all,
//   						excludedHeaders: []*string{
//   							jsii.String("excludedHeaders"),
//   						},
//   						includedHeaders: []*string{
//   							jsii.String("includedHeaders"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				jsonBody: &jsonBodyProperty{
//   					matchPattern: &jsonMatchPatternProperty{
//   						all: all,
//   						includedPaths: []*string{
//   							jsii.String("includedPaths"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//
//   					// the properties below are optional
//   					invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				method: method,
//   				queryString: queryString,
//   				singleHeader: singleHeader,
//   				singleQueryArgument: singleQueryArgument,
//   				uriPath: uriPath,
//   			},
//   			regexString: jsii.String("regexString"),
//   			textTransformations: []interface{}{
//   				&textTransformationProperty{
//   					priority: jsii.Number(123),
//   					type: jsii.String("type"),
//   				},
//   			},
//   		},
//   		regexPatternSetReferenceStatement: &regexPatternSetReferenceStatementProperty{
//   			arn: jsii.String("arn"),
//   			fieldToMatch: &fieldToMatchProperty{
//   				allQueryArguments: allQueryArguments,
//   				body: &bodyProperty{
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				cookies: &cookiesProperty{
//   					matchPattern: &cookieMatchPatternProperty{
//   						all: all,
//   						excludedCookies: []*string{
//   							jsii.String("excludedCookies"),
//   						},
//   						includedCookies: []*string{
//   							jsii.String("includedCookies"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				headers: &headersProperty{
//   					matchPattern: &headerMatchPatternProperty{
//   						all: all,
//   						excludedHeaders: []*string{
//   							jsii.String("excludedHeaders"),
//   						},
//   						includedHeaders: []*string{
//   							jsii.String("includedHeaders"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				jsonBody: &jsonBodyProperty{
//   					matchPattern: &jsonMatchPatternProperty{
//   						all: all,
//   						includedPaths: []*string{
//   							jsii.String("includedPaths"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//
//   					// the properties below are optional
//   					invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				method: method,
//   				queryString: queryString,
//   				singleHeader: singleHeader,
//   				singleQueryArgument: singleQueryArgument,
//   				uriPath: uriPath,
//   			},
//   			textTransformations: []interface{}{
//   				&textTransformationProperty{
//   					priority: jsii.Number(123),
//   					type: jsii.String("type"),
//   				},
//   			},
//   		},
//   		sizeConstraintStatement: &sizeConstraintStatementProperty{
//   			comparisonOperator: jsii.String("comparisonOperator"),
//   			fieldToMatch: &fieldToMatchProperty{
//   				allQueryArguments: allQueryArguments,
//   				body: &bodyProperty{
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				cookies: &cookiesProperty{
//   					matchPattern: &cookieMatchPatternProperty{
//   						all: all,
//   						excludedCookies: []*string{
//   							jsii.String("excludedCookies"),
//   						},
//   						includedCookies: []*string{
//   							jsii.String("includedCookies"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				headers: &headersProperty{
//   					matchPattern: &headerMatchPatternProperty{
//   						all: all,
//   						excludedHeaders: []*string{
//   							jsii.String("excludedHeaders"),
//   						},
//   						includedHeaders: []*string{
//   							jsii.String("includedHeaders"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				jsonBody: &jsonBodyProperty{
//   					matchPattern: &jsonMatchPatternProperty{
//   						all: all,
//   						includedPaths: []*string{
//   							jsii.String("includedPaths"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//
//   					// the properties below are optional
//   					invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				method: method,
//   				queryString: queryString,
//   				singleHeader: singleHeader,
//   				singleQueryArgument: singleQueryArgument,
//   				uriPath: uriPath,
//   			},
//   			size: jsii.Number(123),
//   			textTransformations: []interface{}{
//   				&textTransformationProperty{
//   					priority: jsii.Number(123),
//   					type: jsii.String("type"),
//   				},
//   			},
//   		},
//   		sqliMatchStatement: &sqliMatchStatementProperty{
//   			fieldToMatch: &fieldToMatchProperty{
//   				allQueryArguments: allQueryArguments,
//   				body: &bodyProperty{
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				cookies: &cookiesProperty{
//   					matchPattern: &cookieMatchPatternProperty{
//   						all: all,
//   						excludedCookies: []*string{
//   							jsii.String("excludedCookies"),
//   						},
//   						includedCookies: []*string{
//   							jsii.String("includedCookies"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				headers: &headersProperty{
//   					matchPattern: &headerMatchPatternProperty{
//   						all: all,
//   						excludedHeaders: []*string{
//   							jsii.String("excludedHeaders"),
//   						},
//   						includedHeaders: []*string{
//   							jsii.String("includedHeaders"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				jsonBody: &jsonBodyProperty{
//   					matchPattern: &jsonMatchPatternProperty{
//   						all: all,
//   						includedPaths: []*string{
//   							jsii.String("includedPaths"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//
//   					// the properties below are optional
//   					invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				method: method,
//   				queryString: queryString,
//   				singleHeader: singleHeader,
//   				singleQueryArgument: singleQueryArgument,
//   				uriPath: uriPath,
//   			},
//   			textTransformations: []interface{}{
//   				&textTransformationProperty{
//   					priority: jsii.Number(123),
//   					type: jsii.String("type"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			sensitivityLevel: jsii.String("sensitivityLevel"),
//   		},
//   		xssMatchStatement: &xssMatchStatementProperty{
//   			fieldToMatch: &fieldToMatchProperty{
//   				allQueryArguments: allQueryArguments,
//   				body: &bodyProperty{
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				cookies: &cookiesProperty{
//   					matchPattern: &cookieMatchPatternProperty{
//   						all: all,
//   						excludedCookies: []*string{
//   							jsii.String("excludedCookies"),
//   						},
//   						includedCookies: []*string{
//   							jsii.String("includedCookies"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				headers: &headersProperty{
//   					matchPattern: &headerMatchPatternProperty{
//   						all: all,
//   						excludedHeaders: []*string{
//   							jsii.String("excludedHeaders"),
//   						},
//   						includedHeaders: []*string{
//   							jsii.String("includedHeaders"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				jsonBody: &jsonBodyProperty{
//   					matchPattern: &jsonMatchPatternProperty{
//   						all: all,
//   						includedPaths: []*string{
//   							jsii.String("includedPaths"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//
//   					// the properties below are optional
//   					invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				method: method,
//   				queryString: queryString,
//   				singleHeader: singleHeader,
//   				singleQueryArgument: singleQueryArgument,
//   				uriPath: uriPath,
//   			},
//   			textTransformations: []interface{}{
//   				&textTransformationProperty{
//   					priority: jsii.Number(123),
//   					type: jsii.String("type"),
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnRuleGroup_RateBasedStatementProperty struct {
	// Setting that indicates how to aggregate the request counts. The options are the following:.
	//
	// - IP - Aggregate the request counts on the IP address from the web request origin.
	// - FORWARDED_IP - Aggregate the request counts on the first IP address in an HTTP header. If you use this, configure the `ForwardedIPConfig` , to specify the header to use.
	AggregateKeyType *string `field:"required" json:"aggregateKeyType" yaml:"aggregateKeyType"`
	// The limit on requests per 5-minute period for a single originating IP address.
	//
	// If the statement includes a `ScopeDownStatement` , this limit is applied only to the requests that match the statement.
	Limit *float64 `field:"required" json:"limit" yaml:"limit"`
	// The configuration for inspecting IP addresses in an HTTP header that you specify, instead of using the IP address that's reported by the web request origin.
	//
	// Commonly, this is the X-Forwarded-For (XFF) header, but you can specify any header name.
	//
	// > If the specified header isn't present in the request, AWS WAF doesn't apply the rule to the web request at all.
	//
	// This is required if `AggregateKeyType` is set to `FORWARDED_IP` .
	ForwardedIpConfig interface{} `field:"optional" json:"forwardedIpConfig" yaml:"forwardedIpConfig"`
	// An optional nested statement that narrows the scope of the web requests that are evaluated by the rate-based statement.
	//
	// Requests are only tracked by the rate-based statement if they match the scope-down statement. You can use any nestable `Statement` in the scope-down statement, and you can nest statements at any level, the same as you can for a rule statement.
	ScopeDownStatement interface{} `field:"optional" json:"scopeDownStatement" yaml:"scopeDownStatement"`
}

