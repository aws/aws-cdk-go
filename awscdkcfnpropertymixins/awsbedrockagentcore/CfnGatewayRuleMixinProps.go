package awsbedrockagentcore


// Properties for CfnGatewayRulePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnGatewayRuleMixinProps := &CfnGatewayRuleMixinProps{
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
//   							Description: jsii.String("description"),
//   							Metadata: map[string]*string{
//   								"metadataKey": jsii.String("metadata"),
//   							},
//   							Name: jsii.String("name"),
//   							Weight: jsii.Number(123),
//   						},
//   					},
//   				},
//   			},
//   			RouteToTarget: &RouteToTargetActionProperty{
//   				StaticRoute: &StaticRouteProperty{
//   					TargetName: jsii.String("targetName"),
//   				},
//   				WeightedRoute: &WeightedRouteProperty{
//   					TrafficSplit: []interface{}{
//   						&TargetTrafficSplitEntryProperty{
//   							Description: jsii.String("description"),
//   							Metadata: map[string]*string{
//   								"metadataKey": jsii.String("metadata"),
//   							},
//   							Name: jsii.String("name"),
//   							TargetName: jsii.String("targetName"),
//   							Weight: jsii.Number(123),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	Conditions: []interface{}{
//   		&ConditionProperty{
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
//   							Operator: jsii.String("operator"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	GatewayIdentifier: jsii.String("gatewayIdentifier"),
//   	Priority: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-gatewayrule.html
//
type CfnGatewayRuleMixinProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-gatewayrule.html#cfn-bedrockagentcore-gatewayrule-actions
	//
	Actions interface{} `field:"optional" json:"actions" yaml:"actions"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-gatewayrule.html#cfn-bedrockagentcore-gatewayrule-conditions
	//
	Conditions interface{} `field:"optional" json:"conditions" yaml:"conditions"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-gatewayrule.html#cfn-bedrockagentcore-gatewayrule-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-gatewayrule.html#cfn-bedrockagentcore-gatewayrule-gatewayidentifier
	//
	GatewayIdentifier *string `field:"optional" json:"gatewayIdentifier" yaml:"gatewayIdentifier"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-gatewayrule.html#cfn-bedrockagentcore-gatewayrule-priority
	//
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
}

