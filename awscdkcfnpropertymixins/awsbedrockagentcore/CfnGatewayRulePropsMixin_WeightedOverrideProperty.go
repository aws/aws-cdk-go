package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   weightedOverrideProperty := &WeightedOverrideProperty{
//   	TrafficSplit: []interface{}{
//   		&TrafficSplitEntryProperty{
//   			ConfigurationBundle: &ConfigurationBundleReferenceProperty{
//   				BundleArn: jsii.String("bundleArn"),
//   				BundleVersion: jsii.String("bundleVersion"),
//   			},
//   			Description: jsii.String("description"),
//   			Metadata: map[string]*string{
//   				"metadataKey": jsii.String("metadata"),
//   			},
//   			Name: jsii.String("name"),
//   			Weight: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-weightedoverride.html
//
type CfnGatewayRulePropsMixin_WeightedOverrideProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-weightedoverride.html#cfn-bedrockagentcore-gatewayrule-weightedoverride-trafficsplit
	//
	TrafficSplit interface{} `field:"optional" json:"trafficSplit" yaml:"trafficSplit"`
}

