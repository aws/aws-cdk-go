package awsdevopsagent


// Grafana MCP server configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   mCPServerGrafanaConfigurationProperty := &MCPServerGrafanaConfigurationProperty{
//   	EnabledElevatedTools: []*string{
//   		jsii.String("enabledElevatedTools"),
//   	},
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
	// The mutative MCP tools enabled for this integration, by tool name.
	//
	// Only tools the service classifies as MUTATIVE may be listed; other names are rejected. Tool names can be found in the console or from the association's current configuration. The list fully replaces the stored value on update: removing a name disables that tool, and removing the property disables all elevated tools. Independent of Tools, which selects the toolsets available to the agent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-mcpservergrafanaconfiguration.html#cfn-devopsagent-association-mcpservergrafanaconfiguration-enabledelevatedtools
	//
	EnabledElevatedTools *[]*string `field:"optional" json:"enabledElevatedTools" yaml:"enabledElevatedTools"`
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

