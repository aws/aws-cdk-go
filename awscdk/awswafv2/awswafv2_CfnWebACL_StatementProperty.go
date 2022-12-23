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
//   var andStatementProperty_ andStatementProperty
//   var managedRuleGroupStatementProperty_ managedRuleGroupStatementProperty
//   var method interface{}
//   var notStatementProperty_ notStatementProperty
//   var orStatementProperty_ orStatementProperty
//   var queryString interface{}
//   var rateBasedStatementProperty_ rateBasedStatementProperty
//   var singleHeader interface{}
//   var singleQueryArgument interface{}
//   var statementProperty_ statementProperty
//   var uriPath interface{}
//
//   statementProperty := &statementProperty{
//   	andStatement: &andStatementProperty{
//   		statements: []interface{}{
//   			&statementProperty{
//   				andStatement: andStatementProperty_,
//   				byteMatchStatement: &byteMatchStatementProperty{
//   					fieldToMatch: &fieldToMatchProperty{
//   						allQueryArguments: allQueryArguments,
//   						body: &bodyProperty{
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						cookies: &cookiesProperty{
//   							matchPattern: &cookieMatchPatternProperty{
//   								all: all,
//   								excludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								includedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						headers: &headersProperty{
//   							matchPattern: &headerMatchPatternProperty{
//   								all: all,
//   								excludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								includedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						jsonBody: &jsonBodyProperty{
//   							matchPattern: &jsonMatchPatternProperty{
//   								all: all,
//   								includedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						method: method,
//   						queryString: queryString,
//   						singleHeader: singleHeader,
//   						singleQueryArgument: singleQueryArgument,
//   						uriPath: uriPath,
//   					},
//   					positionalConstraint: jsii.String("positionalConstraint"),
//   					textTransformations: []interface{}{
//   						&textTransformationProperty{
//   							priority: jsii.Number(123),
//   							type: jsii.String("type"),
//   						},
//   					},
//
//   					// the properties below are optional
//   					searchString: jsii.String("searchString"),
//   					searchStringBase64: jsii.String("searchStringBase64"),
//   				},
//   				geoMatchStatement: &geoMatchStatementProperty{
//   					countryCodes: []*string{
//   						jsii.String("countryCodes"),
//   					},
//   					forwardedIpConfig: &forwardedIPConfigurationProperty{
//   						fallbackBehavior: jsii.String("fallbackBehavior"),
//   						headerName: jsii.String("headerName"),
//   					},
//   				},
//   				ipSetReferenceStatement: map[string]interface{}{
//   					"arn": jsii.String("arn"),
//
//   					// the properties below are optional
//   					"ipSetForwardedIpConfig": map[string]*string{
//   						"fallbackBehavior": jsii.String("fallbackBehavior"),
//   						"headerName": jsii.String("headerName"),
//   						"position": jsii.String("position"),
//   					},
//   				},
//   				labelMatchStatement: &labelMatchStatementProperty{
//   					key: jsii.String("key"),
//   					scope: jsii.String("scope"),
//   				},
//   				managedRuleGroupStatement: &managedRuleGroupStatementProperty{
//   					name: jsii.String("name"),
//   					vendorName: jsii.String("vendorName"),
//
//   					// the properties below are optional
//   					excludedRules: []interface{}{
//   						&excludedRuleProperty{
//   							name: jsii.String("name"),
//   						},
//   					},
//   					managedRuleGroupConfigs: []interface{}{
//   						&managedRuleGroupConfigProperty{
//   							awsManagedRulesBotControlRuleSet: &aWSManagedRulesBotControlRuleSetProperty{
//   								inspectionLevel: jsii.String("inspectionLevel"),
//   							},
//   							loginPath: jsii.String("loginPath"),
//   							passwordField: &fieldIdentifierProperty{
//   								identifier: jsii.String("identifier"),
//   							},
//   							payloadType: jsii.String("payloadType"),
//   							usernameField: &fieldIdentifierProperty{
//   								identifier: jsii.String("identifier"),
//   							},
//   						},
//   					},
//   					ruleActionOverrides: []interface{}{
//   						&ruleActionOverrideProperty{
//   							actionToUse: &ruleActionProperty{
//   								allow: &allowActionProperty{
//   									customRequestHandling: &customRequestHandlingProperty{
//   										insertHeaders: []interface{}{
//   											&customHTTPHeaderProperty{
//   												name: jsii.String("name"),
//   												value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   								block: &blockActionProperty{
//   									customResponse: &customResponseProperty{
//   										responseCode: jsii.Number(123),
//
//   										// the properties below are optional
//   										customResponseBodyKey: jsii.String("customResponseBodyKey"),
//   										responseHeaders: []interface{}{
//   											&customHTTPHeaderProperty{
//   												name: jsii.String("name"),
//   												value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   								captcha: &captchaActionProperty{
//   									customRequestHandling: &customRequestHandlingProperty{
//   										insertHeaders: []interface{}{
//   											&customHTTPHeaderProperty{
//   												name: jsii.String("name"),
//   												value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   								challenge: &challengeActionProperty{
//   									customRequestHandling: &customRequestHandlingProperty{
//   										insertHeaders: []interface{}{
//   											&customHTTPHeaderProperty{
//   												name: jsii.String("name"),
//   												value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   								count: &countActionProperty{
//   									customRequestHandling: &customRequestHandlingProperty{
//   										insertHeaders: []interface{}{
//   											&customHTTPHeaderProperty{
//   												name: jsii.String("name"),
//   												value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   							},
//   							name: jsii.String("name"),
//   						},
//   					},
//   					scopeDownStatement: statementProperty_,
//   					version: jsii.String("version"),
//   				},
//   				notStatement: &notStatementProperty{
//   					statement: statementProperty_,
//   				},
//   				orStatement: &orStatementProperty{
//   					statements: []interface{}{
//   						statementProperty_,
//   					},
//   				},
//   				rateBasedStatement: &rateBasedStatementProperty{
//   					aggregateKeyType: jsii.String("aggregateKeyType"),
//   					limit: jsii.Number(123),
//
//   					// the properties below are optional
//   					forwardedIpConfig: &forwardedIPConfigurationProperty{
//   						fallbackBehavior: jsii.String("fallbackBehavior"),
//   						headerName: jsii.String("headerName"),
//   					},
//   					scopeDownStatement: statementProperty_,
//   				},
//   				regexMatchStatement: &regexMatchStatementProperty{
//   					fieldToMatch: &fieldToMatchProperty{
//   						allQueryArguments: allQueryArguments,
//   						body: &bodyProperty{
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						cookies: &cookiesProperty{
//   							matchPattern: &cookieMatchPatternProperty{
//   								all: all,
//   								excludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								includedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						headers: &headersProperty{
//   							matchPattern: &headerMatchPatternProperty{
//   								all: all,
//   								excludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								includedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						jsonBody: &jsonBodyProperty{
//   							matchPattern: &jsonMatchPatternProperty{
//   								all: all,
//   								includedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						method: method,
//   						queryString: queryString,
//   						singleHeader: singleHeader,
//   						singleQueryArgument: singleQueryArgument,
//   						uriPath: uriPath,
//   					},
//   					regexString: jsii.String("regexString"),
//   					textTransformations: []interface{}{
//   						&textTransformationProperty{
//   							priority: jsii.Number(123),
//   							type: jsii.String("type"),
//   						},
//   					},
//   				},
//   				regexPatternSetReferenceStatement: &regexPatternSetReferenceStatementProperty{
//   					arn: jsii.String("arn"),
//   					fieldToMatch: &fieldToMatchProperty{
//   						allQueryArguments: allQueryArguments,
//   						body: &bodyProperty{
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						cookies: &cookiesProperty{
//   							matchPattern: &cookieMatchPatternProperty{
//   								all: all,
//   								excludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								includedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						headers: &headersProperty{
//   							matchPattern: &headerMatchPatternProperty{
//   								all: all,
//   								excludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								includedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						jsonBody: &jsonBodyProperty{
//   							matchPattern: &jsonMatchPatternProperty{
//   								all: all,
//   								includedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						method: method,
//   						queryString: queryString,
//   						singleHeader: singleHeader,
//   						singleQueryArgument: singleQueryArgument,
//   						uriPath: uriPath,
//   					},
//   					textTransformations: []interface{}{
//   						&textTransformationProperty{
//   							priority: jsii.Number(123),
//   							type: jsii.String("type"),
//   						},
//   					},
//   				},
//   				ruleGroupReferenceStatement: &ruleGroupReferenceStatementProperty{
//   					arn: jsii.String("arn"),
//
//   					// the properties below are optional
//   					excludedRules: []interface{}{
//   						&excludedRuleProperty{
//   							name: jsii.String("name"),
//   						},
//   					},
//   					ruleActionOverrides: []interface{}{
//   						&ruleActionOverrideProperty{
//   							actionToUse: &ruleActionProperty{
//   								allow: &allowActionProperty{
//   									customRequestHandling: &customRequestHandlingProperty{
//   										insertHeaders: []interface{}{
//   											&customHTTPHeaderProperty{
//   												name: jsii.String("name"),
//   												value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   								block: &blockActionProperty{
//   									customResponse: &customResponseProperty{
//   										responseCode: jsii.Number(123),
//
//   										// the properties below are optional
//   										customResponseBodyKey: jsii.String("customResponseBodyKey"),
//   										responseHeaders: []interface{}{
//   											&customHTTPHeaderProperty{
//   												name: jsii.String("name"),
//   												value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   								captcha: &captchaActionProperty{
//   									customRequestHandling: &customRequestHandlingProperty{
//   										insertHeaders: []interface{}{
//   											&customHTTPHeaderProperty{
//   												name: jsii.String("name"),
//   												value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   								challenge: &challengeActionProperty{
//   									customRequestHandling: &customRequestHandlingProperty{
//   										insertHeaders: []interface{}{
//   											&customHTTPHeaderProperty{
//   												name: jsii.String("name"),
//   												value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   								count: &countActionProperty{
//   									customRequestHandling: &customRequestHandlingProperty{
//   										insertHeaders: []interface{}{
//   											&customHTTPHeaderProperty{
//   												name: jsii.String("name"),
//   												value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   							},
//   							name: jsii.String("name"),
//   						},
//   					},
//   				},
//   				sizeConstraintStatement: &sizeConstraintStatementProperty{
//   					comparisonOperator: jsii.String("comparisonOperator"),
//   					fieldToMatch: &fieldToMatchProperty{
//   						allQueryArguments: allQueryArguments,
//   						body: &bodyProperty{
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						cookies: &cookiesProperty{
//   							matchPattern: &cookieMatchPatternProperty{
//   								all: all,
//   								excludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								includedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						headers: &headersProperty{
//   							matchPattern: &headerMatchPatternProperty{
//   								all: all,
//   								excludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								includedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						jsonBody: &jsonBodyProperty{
//   							matchPattern: &jsonMatchPatternProperty{
//   								all: all,
//   								includedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						method: method,
//   						queryString: queryString,
//   						singleHeader: singleHeader,
//   						singleQueryArgument: singleQueryArgument,
//   						uriPath: uriPath,
//   					},
//   					size: jsii.Number(123),
//   					textTransformations: []interface{}{
//   						&textTransformationProperty{
//   							priority: jsii.Number(123),
//   							type: jsii.String("type"),
//   						},
//   					},
//   				},
//   				sqliMatchStatement: &sqliMatchStatementProperty{
//   					fieldToMatch: &fieldToMatchProperty{
//   						allQueryArguments: allQueryArguments,
//   						body: &bodyProperty{
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						cookies: &cookiesProperty{
//   							matchPattern: &cookieMatchPatternProperty{
//   								all: all,
//   								excludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								includedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						headers: &headersProperty{
//   							matchPattern: &headerMatchPatternProperty{
//   								all: all,
//   								excludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								includedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						jsonBody: &jsonBodyProperty{
//   							matchPattern: &jsonMatchPatternProperty{
//   								all: all,
//   								includedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						method: method,
//   						queryString: queryString,
//   						singleHeader: singleHeader,
//   						singleQueryArgument: singleQueryArgument,
//   						uriPath: uriPath,
//   					},
//   					textTransformations: []interface{}{
//   						&textTransformationProperty{
//   							priority: jsii.Number(123),
//   							type: jsii.String("type"),
//   						},
//   					},
//
//   					// the properties below are optional
//   					sensitivityLevel: jsii.String("sensitivityLevel"),
//   				},
//   				xssMatchStatement: &xssMatchStatementProperty{
//   					fieldToMatch: &fieldToMatchProperty{
//   						allQueryArguments: allQueryArguments,
//   						body: &bodyProperty{
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						cookies: &cookiesProperty{
//   							matchPattern: &cookieMatchPatternProperty{
//   								all: all,
//   								excludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								includedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						headers: &headersProperty{
//   							matchPattern: &headerMatchPatternProperty{
//   								all: all,
//   								excludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								includedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						jsonBody: &jsonBodyProperty{
//   							matchPattern: &jsonMatchPatternProperty{
//   								all: all,
//   								includedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						method: method,
//   						queryString: queryString,
//   						singleHeader: singleHeader,
//   						singleQueryArgument: singleQueryArgument,
//   						uriPath: uriPath,
//   					},
//   					textTransformations: []interface{}{
//   						&textTransformationProperty{
//   							priority: jsii.Number(123),
//   							type: jsii.String("type"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	byteMatchStatement: &byteMatchStatementProperty{
//   		fieldToMatch: &fieldToMatchProperty{
//   			allQueryArguments: allQueryArguments,
//   			body: &bodyProperty{
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			cookies: &cookiesProperty{
//   				matchPattern: &cookieMatchPatternProperty{
//   					all: all,
//   					excludedCookies: []*string{
//   						jsii.String("excludedCookies"),
//   					},
//   					includedCookies: []*string{
//   						jsii.String("includedCookies"),
//   					},
//   				},
//   				matchScope: jsii.String("matchScope"),
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			headers: &headersProperty{
//   				matchPattern: &headerMatchPatternProperty{
//   					all: all,
//   					excludedHeaders: []*string{
//   						jsii.String("excludedHeaders"),
//   					},
//   					includedHeaders: []*string{
//   						jsii.String("includedHeaders"),
//   					},
//   				},
//   				matchScope: jsii.String("matchScope"),
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			jsonBody: &jsonBodyProperty{
//   				matchPattern: &jsonMatchPatternProperty{
//   					all: all,
//   					includedPaths: []*string{
//   						jsii.String("includedPaths"),
//   					},
//   				},
//   				matchScope: jsii.String("matchScope"),
//
//   				// the properties below are optional
//   				invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			method: method,
//   			queryString: queryString,
//   			singleHeader: singleHeader,
//   			singleQueryArgument: singleQueryArgument,
//   			uriPath: uriPath,
//   		},
//   		positionalConstraint: jsii.String("positionalConstraint"),
//   		textTransformations: []interface{}{
//   			&textTransformationProperty{
//   				priority: jsii.Number(123),
//   				type: jsii.String("type"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		searchString: jsii.String("searchString"),
//   		searchStringBase64: jsii.String("searchStringBase64"),
//   	},
//   	geoMatchStatement: &geoMatchStatementProperty{
//   		countryCodes: []*string{
//   			jsii.String("countryCodes"),
//   		},
//   		forwardedIpConfig: &forwardedIPConfigurationProperty{
//   			fallbackBehavior: jsii.String("fallbackBehavior"),
//   			headerName: jsii.String("headerName"),
//   		},
//   	},
//   	ipSetReferenceStatement: map[string]interface{}{
//   		"arn": jsii.String("arn"),
//
//   		// the properties below are optional
//   		"ipSetForwardedIpConfig": map[string]*string{
//   			"fallbackBehavior": jsii.String("fallbackBehavior"),
//   			"headerName": jsii.String("headerName"),
//   			"position": jsii.String("position"),
//   		},
//   	},
//   	labelMatchStatement: &labelMatchStatementProperty{
//   		key: jsii.String("key"),
//   		scope: jsii.String("scope"),
//   	},
//   	managedRuleGroupStatement: &managedRuleGroupStatementProperty{
//   		name: jsii.String("name"),
//   		vendorName: jsii.String("vendorName"),
//
//   		// the properties below are optional
//   		excludedRules: []interface{}{
//   			&excludedRuleProperty{
//   				name: jsii.String("name"),
//   			},
//   		},
//   		managedRuleGroupConfigs: []interface{}{
//   			&managedRuleGroupConfigProperty{
//   				awsManagedRulesBotControlRuleSet: &aWSManagedRulesBotControlRuleSetProperty{
//   					inspectionLevel: jsii.String("inspectionLevel"),
//   				},
//   				loginPath: jsii.String("loginPath"),
//   				passwordField: &fieldIdentifierProperty{
//   					identifier: jsii.String("identifier"),
//   				},
//   				payloadType: jsii.String("payloadType"),
//   				usernameField: &fieldIdentifierProperty{
//   					identifier: jsii.String("identifier"),
//   				},
//   			},
//   		},
//   		ruleActionOverrides: []interface{}{
//   			&ruleActionOverrideProperty{
//   				actionToUse: &ruleActionProperty{
//   					allow: &allowActionProperty{
//   						customRequestHandling: &customRequestHandlingProperty{
//   							insertHeaders: []interface{}{
//   								&customHTTPHeaderProperty{
//   									name: jsii.String("name"),
//   									value: jsii.String("value"),
//   								},
//   							},
//   						},
//   					},
//   					block: &blockActionProperty{
//   						customResponse: &customResponseProperty{
//   							responseCode: jsii.Number(123),
//
//   							// the properties below are optional
//   							customResponseBodyKey: jsii.String("customResponseBodyKey"),
//   							responseHeaders: []interface{}{
//   								&customHTTPHeaderProperty{
//   									name: jsii.String("name"),
//   									value: jsii.String("value"),
//   								},
//   							},
//   						},
//   					},
//   					captcha: &captchaActionProperty{
//   						customRequestHandling: &customRequestHandlingProperty{
//   							insertHeaders: []interface{}{
//   								&customHTTPHeaderProperty{
//   									name: jsii.String("name"),
//   									value: jsii.String("value"),
//   								},
//   							},
//   						},
//   					},
//   					challenge: &challengeActionProperty{
//   						customRequestHandling: &customRequestHandlingProperty{
//   							insertHeaders: []interface{}{
//   								&customHTTPHeaderProperty{
//   									name: jsii.String("name"),
//   									value: jsii.String("value"),
//   								},
//   							},
//   						},
//   					},
//   					count: &countActionProperty{
//   						customRequestHandling: &customRequestHandlingProperty{
//   							insertHeaders: []interface{}{
//   								&customHTTPHeaderProperty{
//   									name: jsii.String("name"),
//   									value: jsii.String("value"),
//   								},
//   							},
//   						},
//   					},
//   				},
//   				name: jsii.String("name"),
//   			},
//   		},
//   		scopeDownStatement: &statementProperty{
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
//   			managedRuleGroupStatement: managedRuleGroupStatementProperty_,
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
//   		version: jsii.String("version"),
//   	},
//   	notStatement: &notStatementProperty{
//   		statement: &statementProperty{
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
//   			notStatement: notStatementProperty_,
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
//   	orStatement: &orStatementProperty{
//   		statements: []interface{}{
//   			&statementProperty{
//   				andStatement: &andStatementProperty{
//   					statements: []interface{}{
//   						statementProperty_,
//   					},
//   				},
//   				byteMatchStatement: &byteMatchStatementProperty{
//   					fieldToMatch: &fieldToMatchProperty{
//   						allQueryArguments: allQueryArguments,
//   						body: &bodyProperty{
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						cookies: &cookiesProperty{
//   							matchPattern: &cookieMatchPatternProperty{
//   								all: all,
//   								excludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								includedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						headers: &headersProperty{
//   							matchPattern: &headerMatchPatternProperty{
//   								all: all,
//   								excludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								includedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						jsonBody: &jsonBodyProperty{
//   							matchPattern: &jsonMatchPatternProperty{
//   								all: all,
//   								includedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						method: method,
//   						queryString: queryString,
//   						singleHeader: singleHeader,
//   						singleQueryArgument: singleQueryArgument,
//   						uriPath: uriPath,
//   					},
//   					positionalConstraint: jsii.String("positionalConstraint"),
//   					textTransformations: []interface{}{
//   						&textTransformationProperty{
//   							priority: jsii.Number(123),
//   							type: jsii.String("type"),
//   						},
//   					},
//
//   					// the properties below are optional
//   					searchString: jsii.String("searchString"),
//   					searchStringBase64: jsii.String("searchStringBase64"),
//   				},
//   				geoMatchStatement: &geoMatchStatementProperty{
//   					countryCodes: []*string{
//   						jsii.String("countryCodes"),
//   					},
//   					forwardedIpConfig: &forwardedIPConfigurationProperty{
//   						fallbackBehavior: jsii.String("fallbackBehavior"),
//   						headerName: jsii.String("headerName"),
//   					},
//   				},
//   				ipSetReferenceStatement: map[string]interface{}{
//   					"arn": jsii.String("arn"),
//
//   					// the properties below are optional
//   					"ipSetForwardedIpConfig": map[string]*string{
//   						"fallbackBehavior": jsii.String("fallbackBehavior"),
//   						"headerName": jsii.String("headerName"),
//   						"position": jsii.String("position"),
//   					},
//   				},
//   				labelMatchStatement: &labelMatchStatementProperty{
//   					key: jsii.String("key"),
//   					scope: jsii.String("scope"),
//   				},
//   				managedRuleGroupStatement: &managedRuleGroupStatementProperty{
//   					name: jsii.String("name"),
//   					vendorName: jsii.String("vendorName"),
//
//   					// the properties below are optional
//   					excludedRules: []interface{}{
//   						&excludedRuleProperty{
//   							name: jsii.String("name"),
//   						},
//   					},
//   					managedRuleGroupConfigs: []interface{}{
//   						&managedRuleGroupConfigProperty{
//   							awsManagedRulesBotControlRuleSet: &aWSManagedRulesBotControlRuleSetProperty{
//   								inspectionLevel: jsii.String("inspectionLevel"),
//   							},
//   							loginPath: jsii.String("loginPath"),
//   							passwordField: &fieldIdentifierProperty{
//   								identifier: jsii.String("identifier"),
//   							},
//   							payloadType: jsii.String("payloadType"),
//   							usernameField: &fieldIdentifierProperty{
//   								identifier: jsii.String("identifier"),
//   							},
//   						},
//   					},
//   					ruleActionOverrides: []interface{}{
//   						&ruleActionOverrideProperty{
//   							actionToUse: &ruleActionProperty{
//   								allow: &allowActionProperty{
//   									customRequestHandling: &customRequestHandlingProperty{
//   										insertHeaders: []interface{}{
//   											&customHTTPHeaderProperty{
//   												name: jsii.String("name"),
//   												value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   								block: &blockActionProperty{
//   									customResponse: &customResponseProperty{
//   										responseCode: jsii.Number(123),
//
//   										// the properties below are optional
//   										customResponseBodyKey: jsii.String("customResponseBodyKey"),
//   										responseHeaders: []interface{}{
//   											&customHTTPHeaderProperty{
//   												name: jsii.String("name"),
//   												value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   								captcha: &captchaActionProperty{
//   									customRequestHandling: &customRequestHandlingProperty{
//   										insertHeaders: []interface{}{
//   											&customHTTPHeaderProperty{
//   												name: jsii.String("name"),
//   												value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   								challenge: &challengeActionProperty{
//   									customRequestHandling: &customRequestHandlingProperty{
//   										insertHeaders: []interface{}{
//   											&customHTTPHeaderProperty{
//   												name: jsii.String("name"),
//   												value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   								count: &countActionProperty{
//   									customRequestHandling: &customRequestHandlingProperty{
//   										insertHeaders: []interface{}{
//   											&customHTTPHeaderProperty{
//   												name: jsii.String("name"),
//   												value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   							},
//   							name: jsii.String("name"),
//   						},
//   					},
//   					scopeDownStatement: statementProperty_,
//   					version: jsii.String("version"),
//   				},
//   				notStatement: &notStatementProperty{
//   					statement: statementProperty_,
//   				},
//   				orStatement: orStatementProperty_,
//   				rateBasedStatement: &rateBasedStatementProperty{
//   					aggregateKeyType: jsii.String("aggregateKeyType"),
//   					limit: jsii.Number(123),
//
//   					// the properties below are optional
//   					forwardedIpConfig: &forwardedIPConfigurationProperty{
//   						fallbackBehavior: jsii.String("fallbackBehavior"),
//   						headerName: jsii.String("headerName"),
//   					},
//   					scopeDownStatement: statementProperty_,
//   				},
//   				regexMatchStatement: &regexMatchStatementProperty{
//   					fieldToMatch: &fieldToMatchProperty{
//   						allQueryArguments: allQueryArguments,
//   						body: &bodyProperty{
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						cookies: &cookiesProperty{
//   							matchPattern: &cookieMatchPatternProperty{
//   								all: all,
//   								excludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								includedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						headers: &headersProperty{
//   							matchPattern: &headerMatchPatternProperty{
//   								all: all,
//   								excludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								includedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						jsonBody: &jsonBodyProperty{
//   							matchPattern: &jsonMatchPatternProperty{
//   								all: all,
//   								includedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						method: method,
//   						queryString: queryString,
//   						singleHeader: singleHeader,
//   						singleQueryArgument: singleQueryArgument,
//   						uriPath: uriPath,
//   					},
//   					regexString: jsii.String("regexString"),
//   					textTransformations: []interface{}{
//   						&textTransformationProperty{
//   							priority: jsii.Number(123),
//   							type: jsii.String("type"),
//   						},
//   					},
//   				},
//   				regexPatternSetReferenceStatement: &regexPatternSetReferenceStatementProperty{
//   					arn: jsii.String("arn"),
//   					fieldToMatch: &fieldToMatchProperty{
//   						allQueryArguments: allQueryArguments,
//   						body: &bodyProperty{
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						cookies: &cookiesProperty{
//   							matchPattern: &cookieMatchPatternProperty{
//   								all: all,
//   								excludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								includedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						headers: &headersProperty{
//   							matchPattern: &headerMatchPatternProperty{
//   								all: all,
//   								excludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								includedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						jsonBody: &jsonBodyProperty{
//   							matchPattern: &jsonMatchPatternProperty{
//   								all: all,
//   								includedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						method: method,
//   						queryString: queryString,
//   						singleHeader: singleHeader,
//   						singleQueryArgument: singleQueryArgument,
//   						uriPath: uriPath,
//   					},
//   					textTransformations: []interface{}{
//   						&textTransformationProperty{
//   							priority: jsii.Number(123),
//   							type: jsii.String("type"),
//   						},
//   					},
//   				},
//   				ruleGroupReferenceStatement: &ruleGroupReferenceStatementProperty{
//   					arn: jsii.String("arn"),
//
//   					// the properties below are optional
//   					excludedRules: []interface{}{
//   						&excludedRuleProperty{
//   							name: jsii.String("name"),
//   						},
//   					},
//   					ruleActionOverrides: []interface{}{
//   						&ruleActionOverrideProperty{
//   							actionToUse: &ruleActionProperty{
//   								allow: &allowActionProperty{
//   									customRequestHandling: &customRequestHandlingProperty{
//   										insertHeaders: []interface{}{
//   											&customHTTPHeaderProperty{
//   												name: jsii.String("name"),
//   												value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   								block: &blockActionProperty{
//   									customResponse: &customResponseProperty{
//   										responseCode: jsii.Number(123),
//
//   										// the properties below are optional
//   										customResponseBodyKey: jsii.String("customResponseBodyKey"),
//   										responseHeaders: []interface{}{
//   											&customHTTPHeaderProperty{
//   												name: jsii.String("name"),
//   												value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   								captcha: &captchaActionProperty{
//   									customRequestHandling: &customRequestHandlingProperty{
//   										insertHeaders: []interface{}{
//   											&customHTTPHeaderProperty{
//   												name: jsii.String("name"),
//   												value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   								challenge: &challengeActionProperty{
//   									customRequestHandling: &customRequestHandlingProperty{
//   										insertHeaders: []interface{}{
//   											&customHTTPHeaderProperty{
//   												name: jsii.String("name"),
//   												value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   								count: &countActionProperty{
//   									customRequestHandling: &customRequestHandlingProperty{
//   										insertHeaders: []interface{}{
//   											&customHTTPHeaderProperty{
//   												name: jsii.String("name"),
//   												value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   							},
//   							name: jsii.String("name"),
//   						},
//   					},
//   				},
//   				sizeConstraintStatement: &sizeConstraintStatementProperty{
//   					comparisonOperator: jsii.String("comparisonOperator"),
//   					fieldToMatch: &fieldToMatchProperty{
//   						allQueryArguments: allQueryArguments,
//   						body: &bodyProperty{
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						cookies: &cookiesProperty{
//   							matchPattern: &cookieMatchPatternProperty{
//   								all: all,
//   								excludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								includedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						headers: &headersProperty{
//   							matchPattern: &headerMatchPatternProperty{
//   								all: all,
//   								excludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								includedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						jsonBody: &jsonBodyProperty{
//   							matchPattern: &jsonMatchPatternProperty{
//   								all: all,
//   								includedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						method: method,
//   						queryString: queryString,
//   						singleHeader: singleHeader,
//   						singleQueryArgument: singleQueryArgument,
//   						uriPath: uriPath,
//   					},
//   					size: jsii.Number(123),
//   					textTransformations: []interface{}{
//   						&textTransformationProperty{
//   							priority: jsii.Number(123),
//   							type: jsii.String("type"),
//   						},
//   					},
//   				},
//   				sqliMatchStatement: &sqliMatchStatementProperty{
//   					fieldToMatch: &fieldToMatchProperty{
//   						allQueryArguments: allQueryArguments,
//   						body: &bodyProperty{
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						cookies: &cookiesProperty{
//   							matchPattern: &cookieMatchPatternProperty{
//   								all: all,
//   								excludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								includedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						headers: &headersProperty{
//   							matchPattern: &headerMatchPatternProperty{
//   								all: all,
//   								excludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								includedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						jsonBody: &jsonBodyProperty{
//   							matchPattern: &jsonMatchPatternProperty{
//   								all: all,
//   								includedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						method: method,
//   						queryString: queryString,
//   						singleHeader: singleHeader,
//   						singleQueryArgument: singleQueryArgument,
//   						uriPath: uriPath,
//   					},
//   					textTransformations: []interface{}{
//   						&textTransformationProperty{
//   							priority: jsii.Number(123),
//   							type: jsii.String("type"),
//   						},
//   					},
//
//   					// the properties below are optional
//   					sensitivityLevel: jsii.String("sensitivityLevel"),
//   				},
//   				xssMatchStatement: &xssMatchStatementProperty{
//   					fieldToMatch: &fieldToMatchProperty{
//   						allQueryArguments: allQueryArguments,
//   						body: &bodyProperty{
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						cookies: &cookiesProperty{
//   							matchPattern: &cookieMatchPatternProperty{
//   								all: all,
//   								excludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								includedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						headers: &headersProperty{
//   							matchPattern: &headerMatchPatternProperty{
//   								all: all,
//   								excludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								includedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						jsonBody: &jsonBodyProperty{
//   							matchPattern: &jsonMatchPatternProperty{
//   								all: all,
//   								includedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						method: method,
//   						queryString: queryString,
//   						singleHeader: singleHeader,
//   						singleQueryArgument: singleQueryArgument,
//   						uriPath: uriPath,
//   					},
//   					textTransformations: []interface{}{
//   						&textTransformationProperty{
//   							priority: jsii.Number(123),
//   							type: jsii.String("type"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	rateBasedStatement: &rateBasedStatementProperty{
//   		aggregateKeyType: jsii.String("aggregateKeyType"),
//   		limit: jsii.Number(123),
//
//   		// the properties below are optional
//   		forwardedIpConfig: &forwardedIPConfigurationProperty{
//   			fallbackBehavior: jsii.String("fallbackBehavior"),
//   			headerName: jsii.String("headerName"),
//   		},
//   		scopeDownStatement: &statementProperty{
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
//   			rateBasedStatement: rateBasedStatementProperty_,
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
//   	regexMatchStatement: &regexMatchStatementProperty{
//   		fieldToMatch: &fieldToMatchProperty{
//   			allQueryArguments: allQueryArguments,
//   			body: &bodyProperty{
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			cookies: &cookiesProperty{
//   				matchPattern: &cookieMatchPatternProperty{
//   					all: all,
//   					excludedCookies: []*string{
//   						jsii.String("excludedCookies"),
//   					},
//   					includedCookies: []*string{
//   						jsii.String("includedCookies"),
//   					},
//   				},
//   				matchScope: jsii.String("matchScope"),
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			headers: &headersProperty{
//   				matchPattern: &headerMatchPatternProperty{
//   					all: all,
//   					excludedHeaders: []*string{
//   						jsii.String("excludedHeaders"),
//   					},
//   					includedHeaders: []*string{
//   						jsii.String("includedHeaders"),
//   					},
//   				},
//   				matchScope: jsii.String("matchScope"),
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			jsonBody: &jsonBodyProperty{
//   				matchPattern: &jsonMatchPatternProperty{
//   					all: all,
//   					includedPaths: []*string{
//   						jsii.String("includedPaths"),
//   					},
//   				},
//   				matchScope: jsii.String("matchScope"),
//
//   				// the properties below are optional
//   				invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			method: method,
//   			queryString: queryString,
//   			singleHeader: singleHeader,
//   			singleQueryArgument: singleQueryArgument,
//   			uriPath: uriPath,
//   		},
//   		regexString: jsii.String("regexString"),
//   		textTransformations: []interface{}{
//   			&textTransformationProperty{
//   				priority: jsii.Number(123),
//   				type: jsii.String("type"),
//   			},
//   		},
//   	},
//   	regexPatternSetReferenceStatement: &regexPatternSetReferenceStatementProperty{
//   		arn: jsii.String("arn"),
//   		fieldToMatch: &fieldToMatchProperty{
//   			allQueryArguments: allQueryArguments,
//   			body: &bodyProperty{
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			cookies: &cookiesProperty{
//   				matchPattern: &cookieMatchPatternProperty{
//   					all: all,
//   					excludedCookies: []*string{
//   						jsii.String("excludedCookies"),
//   					},
//   					includedCookies: []*string{
//   						jsii.String("includedCookies"),
//   					},
//   				},
//   				matchScope: jsii.String("matchScope"),
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			headers: &headersProperty{
//   				matchPattern: &headerMatchPatternProperty{
//   					all: all,
//   					excludedHeaders: []*string{
//   						jsii.String("excludedHeaders"),
//   					},
//   					includedHeaders: []*string{
//   						jsii.String("includedHeaders"),
//   					},
//   				},
//   				matchScope: jsii.String("matchScope"),
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			jsonBody: &jsonBodyProperty{
//   				matchPattern: &jsonMatchPatternProperty{
//   					all: all,
//   					includedPaths: []*string{
//   						jsii.String("includedPaths"),
//   					},
//   				},
//   				matchScope: jsii.String("matchScope"),
//
//   				// the properties below are optional
//   				invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			method: method,
//   			queryString: queryString,
//   			singleHeader: singleHeader,
//   			singleQueryArgument: singleQueryArgument,
//   			uriPath: uriPath,
//   		},
//   		textTransformations: []interface{}{
//   			&textTransformationProperty{
//   				priority: jsii.Number(123),
//   				type: jsii.String("type"),
//   			},
//   		},
//   	},
//   	ruleGroupReferenceStatement: &ruleGroupReferenceStatementProperty{
//   		arn: jsii.String("arn"),
//
//   		// the properties below are optional
//   		excludedRules: []interface{}{
//   			&excludedRuleProperty{
//   				name: jsii.String("name"),
//   			},
//   		},
//   		ruleActionOverrides: []interface{}{
//   			&ruleActionOverrideProperty{
//   				actionToUse: &ruleActionProperty{
//   					allow: &allowActionProperty{
//   						customRequestHandling: &customRequestHandlingProperty{
//   							insertHeaders: []interface{}{
//   								&customHTTPHeaderProperty{
//   									name: jsii.String("name"),
//   									value: jsii.String("value"),
//   								},
//   							},
//   						},
//   					},
//   					block: &blockActionProperty{
//   						customResponse: &customResponseProperty{
//   							responseCode: jsii.Number(123),
//
//   							// the properties below are optional
//   							customResponseBodyKey: jsii.String("customResponseBodyKey"),
//   							responseHeaders: []interface{}{
//   								&customHTTPHeaderProperty{
//   									name: jsii.String("name"),
//   									value: jsii.String("value"),
//   								},
//   							},
//   						},
//   					},
//   					captcha: &captchaActionProperty{
//   						customRequestHandling: &customRequestHandlingProperty{
//   							insertHeaders: []interface{}{
//   								&customHTTPHeaderProperty{
//   									name: jsii.String("name"),
//   									value: jsii.String("value"),
//   								},
//   							},
//   						},
//   					},
//   					challenge: &challengeActionProperty{
//   						customRequestHandling: &customRequestHandlingProperty{
//   							insertHeaders: []interface{}{
//   								&customHTTPHeaderProperty{
//   									name: jsii.String("name"),
//   									value: jsii.String("value"),
//   								},
//   							},
//   						},
//   					},
//   					count: &countActionProperty{
//   						customRequestHandling: &customRequestHandlingProperty{
//   							insertHeaders: []interface{}{
//   								&customHTTPHeaderProperty{
//   									name: jsii.String("name"),
//   									value: jsii.String("value"),
//   								},
//   							},
//   						},
//   					},
//   				},
//   				name: jsii.String("name"),
//   			},
//   		},
//   	},
//   	sizeConstraintStatement: &sizeConstraintStatementProperty{
//   		comparisonOperator: jsii.String("comparisonOperator"),
//   		fieldToMatch: &fieldToMatchProperty{
//   			allQueryArguments: allQueryArguments,
//   			body: &bodyProperty{
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			cookies: &cookiesProperty{
//   				matchPattern: &cookieMatchPatternProperty{
//   					all: all,
//   					excludedCookies: []*string{
//   						jsii.String("excludedCookies"),
//   					},
//   					includedCookies: []*string{
//   						jsii.String("includedCookies"),
//   					},
//   				},
//   				matchScope: jsii.String("matchScope"),
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			headers: &headersProperty{
//   				matchPattern: &headerMatchPatternProperty{
//   					all: all,
//   					excludedHeaders: []*string{
//   						jsii.String("excludedHeaders"),
//   					},
//   					includedHeaders: []*string{
//   						jsii.String("includedHeaders"),
//   					},
//   				},
//   				matchScope: jsii.String("matchScope"),
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			jsonBody: &jsonBodyProperty{
//   				matchPattern: &jsonMatchPatternProperty{
//   					all: all,
//   					includedPaths: []*string{
//   						jsii.String("includedPaths"),
//   					},
//   				},
//   				matchScope: jsii.String("matchScope"),
//
//   				// the properties below are optional
//   				invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			method: method,
//   			queryString: queryString,
//   			singleHeader: singleHeader,
//   			singleQueryArgument: singleQueryArgument,
//   			uriPath: uriPath,
//   		},
//   		size: jsii.Number(123),
//   		textTransformations: []interface{}{
//   			&textTransformationProperty{
//   				priority: jsii.Number(123),
//   				type: jsii.String("type"),
//   			},
//   		},
//   	},
//   	sqliMatchStatement: &sqliMatchStatementProperty{
//   		fieldToMatch: &fieldToMatchProperty{
//   			allQueryArguments: allQueryArguments,
//   			body: &bodyProperty{
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			cookies: &cookiesProperty{
//   				matchPattern: &cookieMatchPatternProperty{
//   					all: all,
//   					excludedCookies: []*string{
//   						jsii.String("excludedCookies"),
//   					},
//   					includedCookies: []*string{
//   						jsii.String("includedCookies"),
//   					},
//   				},
//   				matchScope: jsii.String("matchScope"),
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			headers: &headersProperty{
//   				matchPattern: &headerMatchPatternProperty{
//   					all: all,
//   					excludedHeaders: []*string{
//   						jsii.String("excludedHeaders"),
//   					},
//   					includedHeaders: []*string{
//   						jsii.String("includedHeaders"),
//   					},
//   				},
//   				matchScope: jsii.String("matchScope"),
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			jsonBody: &jsonBodyProperty{
//   				matchPattern: &jsonMatchPatternProperty{
//   					all: all,
//   					includedPaths: []*string{
//   						jsii.String("includedPaths"),
//   					},
//   				},
//   				matchScope: jsii.String("matchScope"),
//
//   				// the properties below are optional
//   				invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			method: method,
//   			queryString: queryString,
//   			singleHeader: singleHeader,
//   			singleQueryArgument: singleQueryArgument,
//   			uriPath: uriPath,
//   		},
//   		textTransformations: []interface{}{
//   			&textTransformationProperty{
//   				priority: jsii.Number(123),
//   				type: jsii.String("type"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		sensitivityLevel: jsii.String("sensitivityLevel"),
//   	},
//   	xssMatchStatement: &xssMatchStatementProperty{
//   		fieldToMatch: &fieldToMatchProperty{
//   			allQueryArguments: allQueryArguments,
//   			body: &bodyProperty{
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			cookies: &cookiesProperty{
//   				matchPattern: &cookieMatchPatternProperty{
//   					all: all,
//   					excludedCookies: []*string{
//   						jsii.String("excludedCookies"),
//   					},
//   					includedCookies: []*string{
//   						jsii.String("includedCookies"),
//   					},
//   				},
//   				matchScope: jsii.String("matchScope"),
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			headers: &headersProperty{
//   				matchPattern: &headerMatchPatternProperty{
//   					all: all,
//   					excludedHeaders: []*string{
//   						jsii.String("excludedHeaders"),
//   					},
//   					includedHeaders: []*string{
//   						jsii.String("includedHeaders"),
//   					},
//   				},
//   				matchScope: jsii.String("matchScope"),
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			jsonBody: &jsonBodyProperty{
//   				matchPattern: &jsonMatchPatternProperty{
//   					all: all,
//   					includedPaths: []*string{
//   						jsii.String("includedPaths"),
//   					},
//   				},
//   				matchScope: jsii.String("matchScope"),
//
//   				// the properties below are optional
//   				invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			method: method,
//   			queryString: queryString,
//   			singleHeader: singleHeader,
//   			singleQueryArgument: singleQueryArgument,
//   			uriPath: uriPath,
//   		},
//   		textTransformations: []interface{}{
//   			&textTransformationProperty{
//   				priority: jsii.Number(123),
//   				type: jsii.String("type"),
//   			},
//   		},
//   	},
//   }
//
type CfnWebACL_StatementProperty struct {
	// A logical rule statement used to combine other rule statements with AND logic.
	//
	// You provide more than one `Statement` within the `AndStatement` .
	AndStatement interface{} `field:"optional" json:"andStatement" yaml:"andStatement"`
	// A rule statement that defines a string match search for AWS WAF to apply to web requests.
	//
	// The byte match statement provides the bytes to search for, the location in requests that you want AWS WAF to search, and other settings. The bytes to search for are typically a string that corresponds with ASCII characters. In the AWS WAF console and the developer guide, this is called a string match statement.
	ByteMatchStatement interface{} `field:"optional" json:"byteMatchStatement" yaml:"byteMatchStatement"`
	// A rule statement used to identify web requests based on country of origin.
	GeoMatchStatement interface{} `field:"optional" json:"geoMatchStatement" yaml:"geoMatchStatement"`
	// A rule statement used to detect web requests coming from particular IP addresses or address ranges.
	//
	// To use this, create an `IPSet` that specifies the addresses you want to detect, then use the ARN of that set in this statement.
	//
	// Each IP set rule statement references an IP set. You create and maintain the set independent of your rules. This allows you to use the single set in multiple rules. When you update the referenced set, AWS WAF automatically updates all rules that reference it.
	IpSetReferenceStatement interface{} `field:"optional" json:"ipSetReferenceStatement" yaml:"ipSetReferenceStatement"`
	// A rule statement that defines a string match search against labels that have been added to the web request by rules that have already run in the web ACL.
	//
	// The label match statement provides the label or namespace string to search for. The label string can represent a part or all of the fully qualified label name that had been added to the web request. Fully qualified labels have a prefix, optional namespaces, and label name. The prefix identifies the rule group or web ACL context of the rule that added the label. If you do not provide the fully qualified name in your label match string, AWS WAF performs the search for labels that were added in the same context as the label match statement.
	LabelMatchStatement interface{} `field:"optional" json:"labelMatchStatement" yaml:"labelMatchStatement"`
	// A rule statement used to run the rules that are defined in a managed rule group.
	//
	// To use this, provide the vendor name and the name of the rule group in this statement.
	//
	// You cannot nest a `ManagedRuleGroupStatement` , for example for use inside a `NotStatement` or `OrStatement` . It can only be referenced as a top-level statement within a rule.
	ManagedRuleGroupStatement interface{} `field:"optional" json:"managedRuleGroupStatement" yaml:"managedRuleGroupStatement"`
	// A logical rule statement used to negate the results of another rule statement.
	//
	// You provide one `Statement` within the `NotStatement` .
	NotStatement interface{} `field:"optional" json:"notStatement" yaml:"notStatement"`
	// A logical rule statement used to combine other rule statements with OR logic.
	//
	// You provide more than one `Statement` within the `OrStatement` .
	OrStatement interface{} `field:"optional" json:"orStatement" yaml:"orStatement"`
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
	RateBasedStatement interface{} `field:"optional" json:"rateBasedStatement" yaml:"rateBasedStatement"`
	// A rule statement used to search web request components for a match against a single regular expression.
	RegexMatchStatement interface{} `field:"optional" json:"regexMatchStatement" yaml:"regexMatchStatement"`
	// A rule statement used to search web request components for matches with regular expressions.
	//
	// To use this, create a `RegexPatternSet` that specifies the expressions that you want to detect, then use the ARN of that set in this statement. A web request matches the pattern set rule statement if the request component matches any of the patterns in the set.
	//
	// Each regex pattern set rule statement references a regex pattern set. You create and maintain the set independent of your rules. This allows you to use the single set in multiple rules. When you update the referenced set, AWS WAF automatically updates all rules that reference it.
	RegexPatternSetReferenceStatement interface{} `field:"optional" json:"regexPatternSetReferenceStatement" yaml:"regexPatternSetReferenceStatement"`
	// A rule statement used to run the rules that are defined in a `RuleGroup` .
	//
	// To use this, create a rule group with your rules, then provide the ARN of the rule group in this statement.
	//
	// You cannot nest a `RuleGroupReferenceStatement` , for example for use inside a `NotStatement` or `OrStatement` . You can only use a rule group reference statement at the top level inside a web ACL.
	RuleGroupReferenceStatement interface{} `field:"optional" json:"ruleGroupReferenceStatement" yaml:"ruleGroupReferenceStatement"`
	// A rule statement that compares a number of bytes against the size of a request component, using a comparison operator, such as greater than (>) or less than (<).
	//
	// For example, you can use a size constraint statement to look for query strings that are longer than 100 bytes.
	//
	// If you configure AWS WAF to inspect the request body, AWS WAF inspects only the first 8192 bytes (8 KB). If the request body for your web requests never exceeds 8192 bytes, you could use a size constraint statement to block requests that have a request body greater than 8192 bytes.
	//
	// If you choose URI for the value of Part of the request to filter on, the slash (/) in the URI counts as one character. For example, the URI `/logo.jpg` is nine characters long.
	SizeConstraintStatement interface{} `field:"optional" json:"sizeConstraintStatement" yaml:"sizeConstraintStatement"`
	// A rule statement that inspects for malicious SQL code.
	//
	// Attackers insert malicious SQL code into web requests to do things like modify your database or extract data from it.
	SqliMatchStatement interface{} `field:"optional" json:"sqliMatchStatement" yaml:"sqliMatchStatement"`
	// A rule statement that inspects for cross-site scripting (XSS) attacks.
	//
	// In XSS attacks, the attacker uses vulnerabilities in a benign website as a vehicle to inject malicious client-site scripts into other legitimate web browsers.
	XssMatchStatement interface{} `field:"optional" json:"xssMatchStatement" yaml:"xssMatchStatement"`
}

