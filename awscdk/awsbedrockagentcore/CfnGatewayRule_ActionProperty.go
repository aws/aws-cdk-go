package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
//   					Name: jsii.String("name"),
//   					Weight: jsii.Number(123),
//
//   					// the properties below are optional
//   					Description: jsii.String("description"),
//   					Metadata: map[string]*string{
//   						"metadataKey": jsii.String("metadata"),
//   					},
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
//   					Name: jsii.String("name"),
//   					TargetName: jsii.String("targetName"),
//   					Weight: jsii.Number(123),
//
//   					// the properties below are optional
//   					Description: jsii.String("description"),
//   					Metadata: map[string]*string{
//   						"metadataKey": jsii.String("metadata"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-action.html
//
type CfnGatewayRule_ActionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-action.html#cfn-bedrockagentcore-gatewayrule-action-configurationbundle
	//
	ConfigurationBundle interface{} `field:"optional" json:"configurationBundle" yaml:"configurationBundle"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-action.html#cfn-bedrockagentcore-gatewayrule-action-routetotarget
	//
	RouteToTarget interface{} `field:"optional" json:"routeToTarget" yaml:"routeToTarget"`
}

