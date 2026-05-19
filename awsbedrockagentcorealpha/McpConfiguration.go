package awsbedrockagentcorealpha


// MCP protocol configuration The configuration for the Model Context Protocol (MCP).
//
// This protocol enables communication between Amazon Bedrock Agent and external tools.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//
//   mcpConfiguration := &McpConfiguration{
//   	Instructions: jsii.String("instructions"),
//   	SearchType: bedrock_agentcore_alpha.McpGatewaySearchType_SEMANTIC,
//   	SupportedVersions: []MCPProtocolVersion{
//   		bedrock_agentcore_alpha.MCPProtocolVersion_MCP_2025_06_18,
//   	},
//   }
//
// Experimental.
type McpConfiguration struct {
	// The instructions for using the Model Context Protocol gateway.
	//
	// These instructions provide guidance on how to interact with the gateway.
	// Default: - No instructions provided.
	//
	// Experimental.
	Instructions *string `field:"optional" json:"instructions" yaml:"instructions"`
	// The search type for the Model Context Protocol gateway.
	//
	// This field specifies how the gateway handles search operations.
	// Default: - No search type specified.
	//
	// Experimental.
	SearchType McpGatewaySearchType `field:"optional" json:"searchType" yaml:"searchType"`
	// The supported versions of the Model Context Protocol.
	//
	// This field specifies which versions of the protocol the gateway can use.
	// Default: - No specific versions specified.
	//
	// Experimental.
	SupportedVersions *[]MCPProtocolVersion `field:"optional" json:"supportedVersions" yaml:"supportedVersions"`
}

