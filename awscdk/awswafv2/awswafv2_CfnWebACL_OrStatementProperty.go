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
//   orStatementProperty := &orStatementProperty{
//   	statements: []interface{}{
//   		&statementProperty{
//   			andStatement: &andStatementProperty{
//   				statements: []interface{}{
//   					statementProperty_,
//   				},
//   			},
//   			byteMatchStatement: &byteMatchStatementProperty{
//   				fieldToMatch: &fieldToMatchProperty{
//   					allQueryArguments: allQueryArguments,
//   					body: &bodyProperty{
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					cookies: &cookiesProperty{
//   						matchPattern: &cookieMatchPatternProperty{
//   							all: all,
//   							excludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							includedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					headers: &headersProperty{
//   						matchPattern: &headerMatchPatternProperty{
//   							all: all,
//   							excludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							includedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					jsonBody: &jsonBodyProperty{
//   						matchPattern: &jsonMatchPatternProperty{
//   							all: all,
//   							includedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					method: method,
//   					queryString: queryString,
//   					singleHeader: singleHeader,
//   					singleQueryArgument: singleQueryArgument,
//   					uriPath: uriPath,
//   				},
//   				positionalConstraint: jsii.String("positionalConstraint"),
//   				textTransformations: []interface{}{
//   					&textTransformationProperty{
//   						priority: jsii.Number(123),
//   						type: jsii.String("type"),
//   					},
//   				},
//
//   				// the properties below are optional
//   				searchString: jsii.String("searchString"),
//   				searchStringBase64: jsii.String("searchStringBase64"),
//   			},
//   			geoMatchStatement: &geoMatchStatementProperty{
//   				countryCodes: []*string{
//   					jsii.String("countryCodes"),
//   				},
//   				forwardedIpConfig: &forwardedIPConfigurationProperty{
//   					fallbackBehavior: jsii.String("fallbackBehavior"),
//   					headerName: jsii.String("headerName"),
//   				},
//   			},
//   			ipSetReferenceStatement: map[string]interface{}{
//   				"arn": jsii.String("arn"),
//
//   				// the properties below are optional
//   				"ipSetForwardedIpConfig": map[string]*string{
//   					"fallbackBehavior": jsii.String("fallbackBehavior"),
//   					"headerName": jsii.String("headerName"),
//   					"position": jsii.String("position"),
//   				},
//   			},
//   			labelMatchStatement: &labelMatchStatementProperty{
//   				key: jsii.String("key"),
//   				scope: jsii.String("scope"),
//   			},
//   			managedRuleGroupStatement: &managedRuleGroupStatementProperty{
//   				name: jsii.String("name"),
//   				vendorName: jsii.String("vendorName"),
//
//   				// the properties below are optional
//   				excludedRules: []interface{}{
//   					&excludedRuleProperty{
//   						name: jsii.String("name"),
//   					},
//   				},
//   				managedRuleGroupConfigs: []interface{}{
//   					&managedRuleGroupConfigProperty{
//   						awsManagedRulesBotControlRuleSet: &aWSManagedRulesBotControlRuleSetProperty{
//   							inspectionLevel: jsii.String("inspectionLevel"),
//   						},
//   						loginPath: jsii.String("loginPath"),
//   						passwordField: &fieldIdentifierProperty{
//   							identifier: jsii.String("identifier"),
//   						},
//   						payloadType: jsii.String("payloadType"),
//   						usernameField: &fieldIdentifierProperty{
//   							identifier: jsii.String("identifier"),
//   						},
//   					},
//   				},
//   				ruleActionOverrides: []interface{}{
//   					&ruleActionOverrideProperty{
//   						actionToUse: &ruleActionProperty{
//   							allow: &allowActionProperty{
//   								customRequestHandling: &customRequestHandlingProperty{
//   									insertHeaders: []interface{}{
//   										&customHTTPHeaderProperty{
//   											name: jsii.String("name"),
//   											value: jsii.String("value"),
//   										},
//   									},
//   								},
//   							},
//   							block: &blockActionProperty{
//   								customResponse: &customResponseProperty{
//   									responseCode: jsii.Number(123),
//
//   									// the properties below are optional
//   									customResponseBodyKey: jsii.String("customResponseBodyKey"),
//   									responseHeaders: []interface{}{
//   										&customHTTPHeaderProperty{
//   											name: jsii.String("name"),
//   											value: jsii.String("value"),
//   										},
//   									},
//   								},
//   							},
//   							captcha: &captchaActionProperty{
//   								customRequestHandling: &customRequestHandlingProperty{
//   									insertHeaders: []interface{}{
//   										&customHTTPHeaderProperty{
//   											name: jsii.String("name"),
//   											value: jsii.String("value"),
//   										},
//   									},
//   								},
//   							},
//   							challenge: &challengeActionProperty{
//   								customRequestHandling: &customRequestHandlingProperty{
//   									insertHeaders: []interface{}{
//   										&customHTTPHeaderProperty{
//   											name: jsii.String("name"),
//   											value: jsii.String("value"),
//   										},
//   									},
//   								},
//   							},
//   							count: &countActionProperty{
//   								customRequestHandling: &customRequestHandlingProperty{
//   									insertHeaders: []interface{}{
//   										&customHTTPHeaderProperty{
//   											name: jsii.String("name"),
//   											value: jsii.String("value"),
//   										},
//   									},
//   								},
//   							},
//   						},
//   						name: jsii.String("name"),
//   					},
//   				},
//   				scopeDownStatement: statementProperty_,
//   				version: jsii.String("version"),
//   			},
//   			notStatement: &notStatementProperty{
//   				statement: statementProperty_,
//   			},
//   			orStatement: &orStatementProperty{
//   				statements: []interface{}{
//   					statementProperty_,
//   				},
//   			},
//   			rateBasedStatement: &rateBasedStatementProperty{
//   				aggregateKeyType: jsii.String("aggregateKeyType"),
//   				limit: jsii.Number(123),
//
//   				// the properties below are optional
//   				forwardedIpConfig: &forwardedIPConfigurationProperty{
//   					fallbackBehavior: jsii.String("fallbackBehavior"),
//   					headerName: jsii.String("headerName"),
//   				},
//   				scopeDownStatement: statementProperty_,
//   			},
//   			regexMatchStatement: &regexMatchStatementProperty{
//   				fieldToMatch: &fieldToMatchProperty{
//   					allQueryArguments: allQueryArguments,
//   					body: &bodyProperty{
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					cookies: &cookiesProperty{
//   						matchPattern: &cookieMatchPatternProperty{
//   							all: all,
//   							excludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							includedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					headers: &headersProperty{
//   						matchPattern: &headerMatchPatternProperty{
//   							all: all,
//   							excludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							includedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					jsonBody: &jsonBodyProperty{
//   						matchPattern: &jsonMatchPatternProperty{
//   							all: all,
//   							includedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					method: method,
//   					queryString: queryString,
//   					singleHeader: singleHeader,
//   					singleQueryArgument: singleQueryArgument,
//   					uriPath: uriPath,
//   				},
//   				regexString: jsii.String("regexString"),
//   				textTransformations: []interface{}{
//   					&textTransformationProperty{
//   						priority: jsii.Number(123),
//   						type: jsii.String("type"),
//   					},
//   				},
//   			},
//   			regexPatternSetReferenceStatement: &regexPatternSetReferenceStatementProperty{
//   				arn: jsii.String("arn"),
//   				fieldToMatch: &fieldToMatchProperty{
//   					allQueryArguments: allQueryArguments,
//   					body: &bodyProperty{
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					cookies: &cookiesProperty{
//   						matchPattern: &cookieMatchPatternProperty{
//   							all: all,
//   							excludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							includedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					headers: &headersProperty{
//   						matchPattern: &headerMatchPatternProperty{
//   							all: all,
//   							excludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							includedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					jsonBody: &jsonBodyProperty{
//   						matchPattern: &jsonMatchPatternProperty{
//   							all: all,
//   							includedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					method: method,
//   					queryString: queryString,
//   					singleHeader: singleHeader,
//   					singleQueryArgument: singleQueryArgument,
//   					uriPath: uriPath,
//   				},
//   				textTransformations: []interface{}{
//   					&textTransformationProperty{
//   						priority: jsii.Number(123),
//   						type: jsii.String("type"),
//   					},
//   				},
//   			},
//   			ruleGroupReferenceStatement: &ruleGroupReferenceStatementProperty{
//   				arn: jsii.String("arn"),
//
//   				// the properties below are optional
//   				excludedRules: []interface{}{
//   					&excludedRuleProperty{
//   						name: jsii.String("name"),
//   					},
//   				},
//   				ruleActionOverrides: []interface{}{
//   					&ruleActionOverrideProperty{
//   						actionToUse: &ruleActionProperty{
//   							allow: &allowActionProperty{
//   								customRequestHandling: &customRequestHandlingProperty{
//   									insertHeaders: []interface{}{
//   										&customHTTPHeaderProperty{
//   											name: jsii.String("name"),
//   											value: jsii.String("value"),
//   										},
//   									},
//   								},
//   							},
//   							block: &blockActionProperty{
//   								customResponse: &customResponseProperty{
//   									responseCode: jsii.Number(123),
//
//   									// the properties below are optional
//   									customResponseBodyKey: jsii.String("customResponseBodyKey"),
//   									responseHeaders: []interface{}{
//   										&customHTTPHeaderProperty{
//   											name: jsii.String("name"),
//   											value: jsii.String("value"),
//   										},
//   									},
//   								},
//   							},
//   							captcha: &captchaActionProperty{
//   								customRequestHandling: &customRequestHandlingProperty{
//   									insertHeaders: []interface{}{
//   										&customHTTPHeaderProperty{
//   											name: jsii.String("name"),
//   											value: jsii.String("value"),
//   										},
//   									},
//   								},
//   							},
//   							challenge: &challengeActionProperty{
//   								customRequestHandling: &customRequestHandlingProperty{
//   									insertHeaders: []interface{}{
//   										&customHTTPHeaderProperty{
//   											name: jsii.String("name"),
//   											value: jsii.String("value"),
//   										},
//   									},
//   								},
//   							},
//   							count: &countActionProperty{
//   								customRequestHandling: &customRequestHandlingProperty{
//   									insertHeaders: []interface{}{
//   										&customHTTPHeaderProperty{
//   											name: jsii.String("name"),
//   											value: jsii.String("value"),
//   										},
//   									},
//   								},
//   							},
//   						},
//   						name: jsii.String("name"),
//   					},
//   				},
//   			},
//   			sizeConstraintStatement: &sizeConstraintStatementProperty{
//   				comparisonOperator: jsii.String("comparisonOperator"),
//   				fieldToMatch: &fieldToMatchProperty{
//   					allQueryArguments: allQueryArguments,
//   					body: &bodyProperty{
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					cookies: &cookiesProperty{
//   						matchPattern: &cookieMatchPatternProperty{
//   							all: all,
//   							excludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							includedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					headers: &headersProperty{
//   						matchPattern: &headerMatchPatternProperty{
//   							all: all,
//   							excludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							includedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					jsonBody: &jsonBodyProperty{
//   						matchPattern: &jsonMatchPatternProperty{
//   							all: all,
//   							includedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					method: method,
//   					queryString: queryString,
//   					singleHeader: singleHeader,
//   					singleQueryArgument: singleQueryArgument,
//   					uriPath: uriPath,
//   				},
//   				size: jsii.Number(123),
//   				textTransformations: []interface{}{
//   					&textTransformationProperty{
//   						priority: jsii.Number(123),
//   						type: jsii.String("type"),
//   					},
//   				},
//   			},
//   			sqliMatchStatement: &sqliMatchStatementProperty{
//   				fieldToMatch: &fieldToMatchProperty{
//   					allQueryArguments: allQueryArguments,
//   					body: &bodyProperty{
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					cookies: &cookiesProperty{
//   						matchPattern: &cookieMatchPatternProperty{
//   							all: all,
//   							excludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							includedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					headers: &headersProperty{
//   						matchPattern: &headerMatchPatternProperty{
//   							all: all,
//   							excludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							includedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					jsonBody: &jsonBodyProperty{
//   						matchPattern: &jsonMatchPatternProperty{
//   							all: all,
//   							includedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					method: method,
//   					queryString: queryString,
//   					singleHeader: singleHeader,
//   					singleQueryArgument: singleQueryArgument,
//   					uriPath: uriPath,
//   				},
//   				textTransformations: []interface{}{
//   					&textTransformationProperty{
//   						priority: jsii.Number(123),
//   						type: jsii.String("type"),
//   					},
//   				},
//
//   				// the properties below are optional
//   				sensitivityLevel: jsii.String("sensitivityLevel"),
//   			},
//   			xssMatchStatement: &xssMatchStatementProperty{
//   				fieldToMatch: &fieldToMatchProperty{
//   					allQueryArguments: allQueryArguments,
//   					body: &bodyProperty{
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					cookies: &cookiesProperty{
//   						matchPattern: &cookieMatchPatternProperty{
//   							all: all,
//   							excludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							includedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					headers: &headersProperty{
//   						matchPattern: &headerMatchPatternProperty{
//   							all: all,
//   							excludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							includedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					jsonBody: &jsonBodyProperty{
//   						matchPattern: &jsonMatchPatternProperty{
//   							all: all,
//   							includedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					method: method,
//   					queryString: queryString,
//   					singleHeader: singleHeader,
//   					singleQueryArgument: singleQueryArgument,
//   					uriPath: uriPath,
//   				},
//   				textTransformations: []interface{}{
//   					&textTransformationProperty{
//   						priority: jsii.Number(123),
//   						type: jsii.String("type"),
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

