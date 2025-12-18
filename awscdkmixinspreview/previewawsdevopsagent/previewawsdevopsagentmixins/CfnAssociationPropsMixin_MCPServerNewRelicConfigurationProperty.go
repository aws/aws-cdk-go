package previewawsdevopsagentmixins


// Configuration for New Relic MCP server integration.
//
// Defines the New Relic account ID and MCP server endpoint URL required for the Agent Space to authenticate and query observability data from New Relic.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   mCPServerNewRelicConfigurationProperty := &MCPServerNewRelicConfigurationProperty{
//   	AccountId: jsii.String("accountId"),
//   	Endpoint: jsii.String("endpoint"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-mcpservernewrelicconfiguration.html
//
type CfnAssociationPropsMixin_MCPServerNewRelicConfigurationProperty struct {
	// New Relic Account ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-mcpservernewrelicconfiguration.html#cfn-devopsagent-association-mcpservernewrelicconfiguration-accountid
	//
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// MCP server endpoint URL (e.g., https://mcp.newrelic.com/mcp/).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-mcpservernewrelicconfiguration.html#cfn-devopsagent-association-mcpservernewrelicconfiguration-endpoint
	//
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
}

