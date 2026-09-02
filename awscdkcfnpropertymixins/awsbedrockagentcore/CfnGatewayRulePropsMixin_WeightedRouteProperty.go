package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   weightedRouteProperty := &WeightedRouteProperty{
//   	TrafficSplit: []interface{}{
//   		&TargetTrafficSplitEntryProperty{
//   			Description: jsii.String("description"),
//   			Metadata: map[string]*string{
//   				"metadataKey": jsii.String("metadata"),
//   			},
//   			Name: jsii.String("name"),
//   			TargetName: jsii.String("targetName"),
//   			Weight: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-weightedroute.html
//
type CfnGatewayRulePropsMixin_WeightedRouteProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-weightedroute.html#cfn-bedrockagentcore-gatewayrule-weightedroute-trafficsplit
	//
	TrafficSplit interface{} `field:"optional" json:"trafficSplit" yaml:"trafficSplit"`
}

