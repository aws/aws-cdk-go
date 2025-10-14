package awsbedrockagentcore


// The protocol configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gatewayProtocolConfigurationProperty := &GatewayProtocolConfigurationProperty{
//   	Mcp: &MCPGatewayConfigurationProperty{
//   		Instructions: jsii.String("instructions"),
//   		SearchType: jsii.String("searchType"),
//   		SupportedVersions: []*string{
//   			jsii.String("supportedVersions"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-gatewayprotocolconfiguration.html
//
type CfnGateway_GatewayProtocolConfigurationProperty struct {
	// The gateway protocol configuration for MCP.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-gatewayprotocolconfiguration.html#cfn-bedrockagentcore-gateway-gatewayprotocolconfiguration-mcp
	//
	Mcp interface{} `field:"required" json:"mcp" yaml:"mcp"`
}

