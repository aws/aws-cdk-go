package previewawsbedrockagentcoremixins


// The gateway configuration for MCP.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   mCPGatewayConfigurationProperty := &MCPGatewayConfigurationProperty{
//   	Instructions: jsii.String("instructions"),
//   	SearchType: jsii.String("searchType"),
//   	SupportedVersions: []*string{
//   		jsii.String("supportedVersions"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-mcpgatewayconfiguration.html
//
type CfnGatewayPropsMixin_MCPGatewayConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-mcpgatewayconfiguration.html#cfn-bedrockagentcore-gateway-mcpgatewayconfiguration-instructions
	//
	Instructions *string `field:"optional" json:"instructions" yaml:"instructions"`
	// The MCP gateway configuration search type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-mcpgatewayconfiguration.html#cfn-bedrockagentcore-gateway-mcpgatewayconfiguration-searchtype
	//
	SearchType *string `field:"optional" json:"searchType" yaml:"searchType"`
	// The supported versions for the MCP configuration for the gateway target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-mcpgatewayconfiguration.html#cfn-bedrockagentcore-gateway-mcpgatewayconfiguration-supportedversions
	//
	SupportedVersions *[]*string `field:"optional" json:"supportedVersions" yaml:"supportedVersions"`
}

