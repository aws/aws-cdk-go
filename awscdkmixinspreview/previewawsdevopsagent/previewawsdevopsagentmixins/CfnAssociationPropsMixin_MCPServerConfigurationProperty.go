package previewawsdevopsagentmixins


// MCP server configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   mCPServerConfigurationProperty := &MCPServerConfigurationProperty{
//   	Description: jsii.String("description"),
//   	EnableWebhookUpdates: jsii.Boolean(false),
//   	Endpoint: jsii.String("endpoint"),
//   	Name: jsii.String("name"),
//   	Tools: []*string{
//   		jsii.String("tools"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-mcpserverconfiguration.html
//
type CfnAssociationPropsMixin_MCPServerConfigurationProperty struct {
	// The description of the MCP server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-mcpserverconfiguration.html#cfn-devopsagent-association-mcpserverconfiguration-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// When set to true, enables the Agent Space to create and update webhooks for receiving notifications and events from the service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-mcpserverconfiguration.html#cfn-devopsagent-association-mcpserverconfiguration-enablewebhookupdates
	//
	EnableWebhookUpdates interface{} `field:"optional" json:"enableWebhookUpdates" yaml:"enableWebhookUpdates"`
	// MCP server endpoint URL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-mcpserverconfiguration.html#cfn-devopsagent-association-mcpserverconfiguration-endpoint
	//
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// The name of the MCP server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-mcpserverconfiguration.html#cfn-devopsagent-association-mcpserverconfiguration-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// List of MCP tools that can be used with the association.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-mcpserverconfiguration.html#cfn-devopsagent-association-mcpserverconfiguration-tools
	//
	Tools *[]*string `field:"optional" json:"tools" yaml:"tools"`
}

