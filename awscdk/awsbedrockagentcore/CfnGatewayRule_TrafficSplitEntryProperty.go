package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   trafficSplitEntryProperty := &TrafficSplitEntryProperty{
//   	ConfigurationBundle: &ConfigurationBundleReferenceProperty{
//   		BundleArn: jsii.String("bundleArn"),
//   		BundleVersion: jsii.String("bundleVersion"),
//   	},
//   	Name: jsii.String("name"),
//   	Weight: jsii.Number(123),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Metadata: map[string]*string{
//   		"metadataKey": jsii.String("metadata"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-trafficsplitentry.html
//
type CfnGatewayRule_TrafficSplitEntryProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-trafficsplitentry.html#cfn-bedrockagentcore-gatewayrule-trafficsplitentry-configurationbundle
	//
	ConfigurationBundle interface{} `field:"required" json:"configurationBundle" yaml:"configurationBundle"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-trafficsplitentry.html#cfn-bedrockagentcore-gatewayrule-trafficsplitentry-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-trafficsplitentry.html#cfn-bedrockagentcore-gatewayrule-trafficsplitentry-weight
	//
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-trafficsplitentry.html#cfn-bedrockagentcore-gatewayrule-trafficsplitentry-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-trafficsplitentry.html#cfn-bedrockagentcore-gatewayrule-trafficsplitentry-metadata
	//
	Metadata interface{} `field:"optional" json:"metadata" yaml:"metadata"`
}

