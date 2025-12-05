package previewawsdevopsagentmixins


// Splunk MCP server configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   mCPServerSplunkConfigurationProperty := &MCPServerSplunkConfigurationProperty{
//   	Description: jsii.String("description"),
//   	EnableWebhookUpdates: jsii.Boolean(false),
//   	Endpoint: jsii.String("endpoint"),
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-mcpserversplunkconfiguration.html
//
type CfnAssociationPropsMixin_MCPServerSplunkConfigurationProperty struct {
	// The description of the MCP server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-mcpserversplunkconfiguration.html#cfn-devopsagent-association-mcpserversplunkconfiguration-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// When set to true, enables the Agent Space to create and update webhooks for receiving notifications and events from the service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-mcpserversplunkconfiguration.html#cfn-devopsagent-association-mcpserversplunkconfiguration-enablewebhookupdates
	//
	EnableWebhookUpdates interface{} `field:"optional" json:"enableWebhookUpdates" yaml:"enableWebhookUpdates"`
	// MCP server endpoint URL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-mcpserversplunkconfiguration.html#cfn-devopsagent-association-mcpserversplunkconfiguration-endpoint
	//
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// The name of the MCP server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-mcpserversplunkconfiguration.html#cfn-devopsagent-association-mcpserversplunkconfiguration-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

