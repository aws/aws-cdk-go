package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   actionProperty := &ActionProperty{
//   	ConfigurationBundle: &ConfigurationBundleActionProperty{
//   		StaticOverride: &StaticOverrideProperty{
//   			BundleArn: jsii.String("bundleArn"),
//   			BundleVersion: jsii.String("bundleVersion"),
//   		},
//   		WeightedOverride: &WeightedOverrideProperty{
//   			TrafficSplit: []interface{}{
//   				&TrafficSplitEntryProperty{
//   					ConfigurationBundle: &ConfigurationBundleReferenceProperty{
//   						BundleArn: jsii.String("bundleArn"),
//   						BundleVersion: jsii.String("bundleVersion"),
//   					},
//   					Description: jsii.String("description"),
//   					Metadata: map[string]*string{
//   						"metadataKey": jsii.String("metadata"),
//   					},
//   					Name: jsii.String("name"),
//   					Weight: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   	RouteToTarget: &RouteToTargetActionProperty{
//   		StaticRoute: &StaticRouteProperty{
//   			TargetName: jsii.String("targetName"),
//   		},
//   		WeightedRoute: &WeightedRouteProperty{
//   			TrafficSplit: []interface{}{
//   				&TargetTrafficSplitEntryProperty{
//   					Description: jsii.String("description"),
//   					Metadata: map[string]*string{
//   						"metadataKey": jsii.String("metadata"),
//   					},
//   					Name: jsii.String("name"),
//   					TargetName: jsii.String("targetName"),
//   					Weight: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-action.html
//
type CfnGatewayRulePropsMixin_ActionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-action.html#cfn-bedrockagentcore-gatewayrule-action-configurationbundle
	//
	ConfigurationBundle interface{} `field:"optional" json:"configurationBundle" yaml:"configurationBundle"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-action.html#cfn-bedrockagentcore-gatewayrule-action-routetotarget
	//
	RouteToTarget interface{} `field:"optional" json:"routeToTarget" yaml:"routeToTarget"`
}

