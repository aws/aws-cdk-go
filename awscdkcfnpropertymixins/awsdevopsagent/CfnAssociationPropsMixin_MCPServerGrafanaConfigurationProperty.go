package awsdevopsagent


// Grafana MCP server configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   mCPServerGrafanaConfigurationProperty := &MCPServerGrafanaConfigurationProperty{
//   	EnableWebhookUpdates: jsii.Boolean(false),
//   	Endpoint: jsii.String("endpoint"),
//   	Tools: []*string{
//   		jsii.String("tools"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-mcpservergrafanaconfiguration.html
//
type CfnAssociationPropsMixin_MCPServerGrafanaConfigurationProperty struct {
	// When set to true, enables the Agent Space to create and update webhooks for receiving notifications and events from the service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-mcpservergrafanaconfiguration.html#cfn-devopsagent-association-mcpservergrafanaconfiguration-enablewebhookupdates
	//
	EnableWebhookUpdates interface{} `field:"optional" json:"enableWebhookUpdates" yaml:"enableWebhookUpdates"`
	// MCP server endpoint URL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-mcpservergrafanaconfiguration.html#cfn-devopsagent-association-mcpservergrafanaconfiguration-endpoint
	//
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// List of tool categories to enable for the Grafana MCP server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-mcpservergrafanaconfiguration.html#cfn-devopsagent-association-mcpservergrafanaconfiguration-tools
	//
	Tools *[]*string `field:"optional" json:"tools" yaml:"tools"`
}

