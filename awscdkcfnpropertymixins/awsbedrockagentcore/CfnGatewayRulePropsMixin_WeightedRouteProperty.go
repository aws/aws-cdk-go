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
//   			Fallback: &FallbackProperty{
//   				CircuitBreaker: &CircuitBreakerConfigurationProperty{
//   					Mode: jsii.String("mode"),
//   				},
//   				FirstByteTimeoutMilliseconds: jsii.Number(123),
//   				TargetFallback: []interface{}{
//   					&FallbackTargetProperty{
//   						FirstByteTimeoutMilliseconds: jsii.Number(123),
//   						Model: jsii.String("model"),
//   						TargetName: jsii.String("targetName"),
//   						TriggerOn: &FallbackTriggersProperty{
//   							Errors: []*string{
//   								jsii.String("errors"),
//   							},
//   							StatusCodes: []*string{
//   								jsii.String("statusCodes"),
//   							},
//   						},
//   					},
//   				},
//   				TriggerOn: &FallbackTriggersProperty{
//   					Errors: []*string{
//   						jsii.String("errors"),
//   					},
//   					StatusCodes: []*string{
//   						jsii.String("statusCodes"),
//   					},
//   				},
//   			},
//   			Metadata: map[string]*string{
//   				"metadataKey": jsii.String("metadata"),
//   			},
//   			Model: jsii.String("model"),
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

