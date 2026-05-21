package awsdevopsagent


// Grafana MCP server configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mCPServerGrafanaConfigurationProperty := &MCPServerGrafanaConfigurationProperty{
//   	Endpoint: jsii.String("endpoint"),
//
//   	// the properties below are optional
//   	EnableWebhookUpdates: jsii.Boolean(false),
//   	Tools: []*string{
//   		jsii.String("tools"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-mcpservergrafanaconfiguration.html
//
type CfnAssociation_MCPServerGrafanaConfigurationProperty struct {
	// MCP server endpoint URL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-mcpservergrafanaconfiguration.html#cfn-devopsagent-association-mcpservergrafanaconfiguration-endpoint
	//
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// When set to true, enables the Agent Space to create and update webhooks for receiving notifications and events from the service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-mcpservergrafanaconfiguration.html#cfn-devopsagent-association-mcpservergrafanaconfiguration-enablewebhookupdates
	//
	EnableWebhookUpdates interface{} `field:"optional" json:"enableWebhookUpdates" yaml:"enableWebhookUpdates"`
	// List of tool categories to enable for the Grafana MCP server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-mcpservergrafanaconfiguration.html#cfn-devopsagent-association-mcpservergrafanaconfiguration-tools
	//
	Tools *[]*string `field:"optional" json:"tools" yaml:"tools"`
}

