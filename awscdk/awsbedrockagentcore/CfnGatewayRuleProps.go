package awsbedrockagentcore


// Properties for defining a `CfnGatewayRule`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGatewayRuleProps := &CfnGatewayRuleProps{
//   	Actions: []interface{}{
//   		&ActionProperty{
//   			ConfigurationBundle: &ConfigurationBundleActionProperty{
//   				StaticOverride: &StaticOverrideProperty{
//   					BundleArn: jsii.String("bundleArn"),
//   					BundleVersion: jsii.String("bundleVersion"),
//   				},
//   				WeightedOverride: &WeightedOverrideProperty{
//   					TrafficSplit: []interface{}{
//   						&TrafficSplitEntryProperty{
//   							ConfigurationBundle: &ConfigurationBundleReferenceProperty{
//   								BundleArn: jsii.String("bundleArn"),
//   								BundleVersion: jsii.String("bundleVersion"),
//   							},
//   							Name: jsii.String("name"),
//   							Weight: jsii.Number(123),
//
//   							// the properties below are optional
//   							Description: jsii.String("description"),
//   							Metadata: map[string]*string{
//   								"metadataKey": jsii.String("metadata"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   			RouteToTarget: &RouteToTargetActionProperty{
//   				StaticRoute: &StaticRouteProperty{
//   					TargetName: jsii.String("targetName"),
//
//   					// the properties below are optional
//   					Fallback: &FallbackProperty{
//   						TargetFallback: []interface{}{
//   							&FallbackTargetProperty{
//   								TargetName: jsii.String("targetName"),
//
//   								// the properties below are optional
//   								FirstByteTimeoutMilliseconds: jsii.Number(123),
//   								Model: jsii.String("model"),
//   								TriggerOn: &FallbackTriggersProperty{
//   									Errors: []*string{
//   										jsii.String("errors"),
//   									},
//   									StatusCodes: []*string{
//   										jsii.String("statusCodes"),
//   									},
//   								},
//   							},
//   						},
//
//   						// the properties below are optional
//   						CircuitBreaker: &CircuitBreakerConfigurationProperty{
//   							Mode: jsii.String("mode"),
//   						},
//   						FirstByteTimeoutMilliseconds: jsii.Number(123),
//   						TriggerOn: &FallbackTriggersProperty{
//   							Errors: []*string{
//   								jsii.String("errors"),
//   							},
//   							StatusCodes: []*string{
//   								jsii.String("statusCodes"),
//   							},
//   						},
//   					},
//   					Model: jsii.String("model"),
//   				},
//   				WeightedRoute: &WeightedRouteProperty{
//   					TrafficSplit: []interface{}{
//   						&TargetTrafficSplitEntryProperty{
//   							Name: jsii.String("name"),
//   							TargetName: jsii.String("targetName"),
//   							Weight: jsii.Number(123),
//
//   							// the properties below are optional
//   							Description: jsii.String("description"),
//   							Fallback: &FallbackProperty{
//   								TargetFallback: []interface{}{
//   									&FallbackTargetProperty{
//   										TargetName: jsii.String("targetName"),
//
//   										// the properties below are optional
//   										FirstByteTimeoutMilliseconds: jsii.Number(123),
//   										Model: jsii.String("model"),
//   										TriggerOn: &FallbackTriggersProperty{
//   											Errors: []*string{
//   												jsii.String("errors"),
//   											},
//   											StatusCodes: []*string{
//   												jsii.String("statusCodes"),
//   											},
//   										},
//   									},
//   								},
//
//   								// the properties below are optional
//   								CircuitBreaker: &CircuitBreakerConfigurationProperty{
//   									Mode: jsii.String("mode"),
//   								},
//   								FirstByteTimeoutMilliseconds: jsii.Number(123),
//   								TriggerOn: &FallbackTriggersProperty{
//   									Errors: []*string{
//   										jsii.String("errors"),
//   									},
//   									StatusCodes: []*string{
//   										jsii.String("statusCodes"),
//   									},
//   								},
//   							},
//   							Metadata: map[string]*string{
//   								"metadataKey": jsii.String("metadata"),
//   							},
//   							Model: jsii.String("model"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	Priority: jsii.Number(123),
//
//   	// the properties below are optional
//   	Conditions: []interface{}{
//   		&ConditionProperty{
//   			MatchModels: &MatchModelsProperty{
//   				AnyOf: []interface{}{
//   					&MatchModelEntryProperty{
//   						Value: jsii.String("value"),
//
//   						// the properties below are optional
//   						Operator: jsii.String("operator"),
//   					},
//   				},
//   			},
//   			MatchPaths: &MatchPathsProperty{
//   				AnyOf: []*string{
//   					jsii.String("anyOf"),
//   				},
//   			},
//   			MatchPrincipals: &MatchPrincipalsProperty{
//   				AnyOf: []interface{}{
//   					&MatchPrincipalEntryProperty{
//   						IamPrincipal: &IamPrincipalProperty{
//   							Arn: jsii.String("arn"),
//
//   							// the properties below are optional
//   							Operator: jsii.String("operator"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	GatewayIdentifier: jsii.String("gatewayIdentifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-gatewayrule.html
//
type CfnGatewayRuleProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-gatewayrule.html#cfn-bedrockagentcore-gatewayrule-actions
	//
	Actions interface{} `field:"required" json:"actions" yaml:"actions"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-gatewayrule.html#cfn-bedrockagentcore-gatewayrule-priority
	//
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-gatewayrule.html#cfn-bedrockagentcore-gatewayrule-conditions
	//
	Conditions interface{} `field:"optional" json:"conditions" yaml:"conditions"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-gatewayrule.html#cfn-bedrockagentcore-gatewayrule-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-gatewayrule.html#cfn-bedrockagentcore-gatewayrule-gatewayidentifier
	//
	GatewayIdentifier *string `field:"optional" json:"gatewayIdentifier" yaml:"gatewayIdentifier"`
}

