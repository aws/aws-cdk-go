package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   configurationBundleActionProperty := &ConfigurationBundleActionProperty{
//   	StaticOverride: &StaticOverrideProperty{
//   		BundleArn: jsii.String("bundleArn"),
//   		BundleVersion: jsii.String("bundleVersion"),
//   	},
//   	WeightedOverride: &WeightedOverrideProperty{
//   		TrafficSplit: []interface{}{
//   			&TrafficSplitEntryProperty{
//   				ConfigurationBundle: &ConfigurationBundleReferenceProperty{
//   					BundleArn: jsii.String("bundleArn"),
//   					BundleVersion: jsii.String("bundleVersion"),
//   				},
//   				Name: jsii.String("name"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-configurationbundleaction.html
//
type CfnGatewayRule_ConfigurationBundleActionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-configurationbundleaction.html#cfn-bedrockagentcore-gatewayrule-configurationbundleaction-staticoverride
	//
	StaticOverride interface{} `field:"optional" json:"staticOverride" yaml:"staticOverride"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-configurationbundleaction.html#cfn-bedrockagentcore-gatewayrule-configurationbundleaction-weightedoverride
	//
	WeightedOverride interface{} `field:"optional" json:"weightedOverride" yaml:"weightedOverride"`
}

