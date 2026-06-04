package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   mcpServerTargetConfigurationProperty := &McpServerTargetConfigurationProperty{
//   	Endpoint: jsii.String("endpoint"),
//   	ListingMode: jsii.String("listingMode"),
//   	McpToolSchema: &McpToolSchemaConfigurationProperty{
//   		InlinePayload: jsii.String("inlinePayload"),
//   		S3: &S3ConfigurationProperty{
//   			BucketOwnerAccountId: jsii.String("bucketOwnerAccountId"),
//   			Uri: jsii.String("uri"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-mcpservertargetconfiguration.html
//
type CfnGatewayTargetPropsMixin_McpServerTargetConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-mcpservertargetconfiguration.html#cfn-bedrockagentcore-gatewaytarget-mcpservertargetconfiguration-endpoint
	//
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-mcpservertargetconfiguration.html#cfn-bedrockagentcore-gatewaytarget-mcpservertargetconfiguration-listingmode
	//
	ListingMode *string `field:"optional" json:"listingMode" yaml:"listingMode"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-mcpservertargetconfiguration.html#cfn-bedrockagentcore-gatewaytarget-mcpservertargetconfiguration-mcptoolschema
	//
	McpToolSchema interface{} `field:"optional" json:"mcpToolSchema" yaml:"mcpToolSchema"`
}

