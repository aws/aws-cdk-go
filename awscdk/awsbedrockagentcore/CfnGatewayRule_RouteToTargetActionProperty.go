package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   routeToTargetActionProperty := &RouteToTargetActionProperty{
//   	StaticRoute: &StaticRouteProperty{
//   		TargetName: jsii.String("targetName"),
//   	},
//   	WeightedRoute: &WeightedRouteProperty{
//   		TrafficSplit: []interface{}{
//   			&TargetTrafficSplitEntryProperty{
//   				Name: jsii.String("name"),
//   				TargetName: jsii.String("targetName"),
//   				Weight: jsii.Number(123),
//
//   				// the properties below are optional
//   				Description: jsii.String("description"),
//   				Metadata: map[string]*string{
//   					"metadataKey": jsii.String("metadata"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-routetotargetaction.html
//
type CfnGatewayRule_RouteToTargetActionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-routetotargetaction.html#cfn-bedrockagentcore-gatewayrule-routetotargetaction-staticroute
	//
	StaticRoute interface{} `field:"optional" json:"staticRoute" yaml:"staticRoute"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-routetotargetaction.html#cfn-bedrockagentcore-gatewayrule-routetotargetaction-weightedroute
	//
	WeightedRoute interface{} `field:"optional" json:"weightedRoute" yaml:"weightedRoute"`
}

