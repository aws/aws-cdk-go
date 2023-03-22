package awswafv2

import (
	"github.com/aws/aws-cdk-go/awscdk"
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
//   cfnWebACLProps := &cfnWebACLProps{
//   	defaultAction: &defaultActionProperty{
//   		allow: &allowActionProperty{
//   			customRequestHandling: &customRequestHandlingProperty{
//   				insertHeaders: []interface{}{
//   					&customHTTPHeaderProperty{
//   						name: jsii.String("name"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   		block: &blockActionProperty{
//   			customResponse: &customResponseProperty{
//   				responseCode: jsii.Number(123),
//
//   				// the properties below are optional
//   				customResponseBodyKey: jsii.String("customResponseBodyKey"),
//   				responseHeaders: []interface{}{
//   					&customHTTPHeaderProperty{
//   						name: jsii.String("name"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	scope: jsii.String("scope"),
//   	visibilityConfig: &visibilityConfigProperty{
//   		cloudWatchMetricsEnabled: jsii.Boolean(false),
//   		metricName: jsii.String("metricName"),
//   		sampledRequestsEnabled: jsii.Boolean(false),
//   	},
//
//   	// the properties below are optional
//   	captchaConfig: &captchaConfigProperty{
//   		immunityTimeProperty: &immunityTimePropertyProperty{
//   			immunityTime: jsii.Number(123),
//   		},
//   	},
//   	challengeConfig: &challengeConfigProperty{
//   		immunityTimeProperty: &immunityTimePropertyProperty{
//   			immunityTime: jsii.Number(123),
//   		},
//   	},
//   	customResponseBodies: map[string]interface{}{
//   		"customResponseBodiesKey": &CustomResponseBodyProperty{
//   			"content": jsii.String("content"),
//   			"contentType": jsii.String("contentType"),
//   		},
//   	},
//   	description: jsii.String("description"),
//   	name: jsii.String("name"),
//   	rules: []interface{}{
//   		&ruleProperty{
//   			name: jsii.String("name"),
//   			priority: jsii.Number(123),
//   			statement: &statementProperty{
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
//   							awsManagedRulesAtpRuleSet: &aWSManagedRulesATPRuleSetProperty{
//   								loginPath: jsii.String("loginPath"),
//
//   								// the properties below are optional
//   								requestInspection: &requestInspectionProperty{
//   									passwordField: &fieldIdentifierProperty{
//   										identifier: jsii.String("identifier"),
//   									},
//   									payloadType: jsii.String("payloadType"),
//   									usernameField: &fieldIdentifierProperty{
//   										identifier: jsii.String("identifier"),
//   									},
//   								},
//   								responseInspection: &responseInspectionProperty{
//   									bodyContains: &responseInspectionBodyContainsProperty{
//   										failureStrings: []*string{
//   											jsii.String("failureStrings"),
//   										},
//   										successStrings: []*string{
//   											jsii.String("successStrings"),
//   										},
//   									},
//   									header: &responseInspectionHeaderProperty{
//   										failureValues: []*string{
//   											jsii.String("failureValues"),
//   										},
//   										name: jsii.String("name"),
//   										successValues: []*string{
//   											jsii.String("successValues"),
//   										},
//   									},
//   									json: &responseInspectionJsonProperty{
//   										failureValues: []*string{
//   											jsii.String("failureValues"),
//   										},
//   										identifier: jsii.String("identifier"),
//   										successValues: []*string{
//   											jsii.String("successValues"),
//   										},
//   									},
//   									statusCode: &responseInspectionStatusCodeProperty{
//   										failureCodes: []interface{}{
//   											jsii.Number(123),
//   										},
//   										successCodes: []interface{}{
//   											jsii.Number(123),
//   										},
//   									},
//   								},
//   							},
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
//   			visibilityConfig: &visibilityConfigProperty{
//   				cloudWatchMetricsEnabled: jsii.Boolean(false),
//   				metricName: jsii.String("metricName"),
//   				sampledRequestsEnabled: jsii.Boolean(false),
//   			},
//
//   			// the properties below are optional
//   			action: &ruleActionProperty{
//   				allow: &allowActionProperty{
//   					customRequestHandling: &customRequestHandlingProperty{
//   						insertHeaders: []interface{}{
//   							&customHTTPHeaderProperty{
//   								name: jsii.String("name"),
//   								value: jsii.String("value"),
//   							},
//   						},
//   					},
//   				},
//   				block: &blockActionProperty{
//   					customResponse: &customResponseProperty{
//   						responseCode: jsii.Number(123),
//
//   						// the properties below are optional
//   						customResponseBodyKey: jsii.String("customResponseBodyKey"),
//   						responseHeaders: []interface{}{
//   							&customHTTPHeaderProperty{
//   								name: jsii.String("name"),
//   								value: jsii.String("value"),
//   							},
//   						},
//   					},
//   				},
//   				captcha: &captchaActionProperty{
//   					customRequestHandling: &customRequestHandlingProperty{
//   						insertHeaders: []interface{}{
//   							&customHTTPHeaderProperty{
//   								name: jsii.String("name"),
//   								value: jsii.String("value"),
//   							},
//   						},
//   					},
//   				},
//   				challenge: &challengeActionProperty{
//   					customRequestHandling: &customRequestHandlingProperty{
//   						insertHeaders: []interface{}{
//   							&customHTTPHeaderProperty{
//   								name: jsii.String("name"),
//   								value: jsii.String("value"),
//   							},
//   						},
//   					},
//   				},
//   				count: &countActionProperty{
//   					customRequestHandling: &customRequestHandlingProperty{
//   						insertHeaders: []interface{}{
//   							&customHTTPHeaderProperty{
//   								name: jsii.String("name"),
//   								value: jsii.String("value"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   			captchaConfig: &captchaConfigProperty{
//   				immunityTimeProperty: &immunityTimePropertyProperty{
//   					immunityTime: jsii.Number(123),
//   				},
//   			},
//   			challengeConfig: &challengeConfigProperty{
//   				immunityTimeProperty: &immunityTimePropertyProperty{
//   					immunityTime: jsii.Number(123),
//   				},
//   			},
//   			overrideAction: &overrideActionProperty{
//   				count: count,
//   				none: none,
//   			},
//   			ruleLabels: []interface{}{
//   				&labelProperty{
//   					name: jsii.String("name"),
//   				},
//   			},
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	tokenDomains: []*string{
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

