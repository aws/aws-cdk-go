package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   weightedOverrideProperty := &WeightedOverrideProperty{
//   	TrafficSplit: []interface{}{
//   		&TrafficSplitEntryProperty{
//   			ConfigurationBundle: &ConfigurationBundleReferenceProperty{
//   				BundleArn: jsii.String("bundleArn"),
//   				BundleVersion: jsii.String("bundleVersion"),
//   			},
//   			Name: jsii.String("name"),
//   			Weight: jsii.Number(123),
//
//   			// the properties below are optional
//   			Description: jsii.String("description"),
//   			Metadata: map[string]*string{
//   				"metadataKey": jsii.String("metadata"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-weightedoverride.html
//
type CfnGatewayRule_WeightedOverrideProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-weightedoverride.html#cfn-bedrockagentcore-gatewayrule-weightedoverride-trafficsplit
	//
	TrafficSplit interface{} `field:"required" json:"trafficSplit" yaml:"trafficSplit"`
}

