package awsdevopsagent


// Datadog MCP server configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mCPServerDatadogConfigurationProperty := &MCPServerDatadogConfigurationProperty{
//   	Endpoint: jsii.String("endpoint"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	EnableWebhookUpdates: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-mcpserverdatadogconfiguration.html
//
type CfnAssociation_MCPServerDatadogConfigurationProperty struct {
	// MCP server endpoint URL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-mcpserverdatadogconfiguration.html#cfn-devopsagent-association-mcpserverdatadogconfiguration-endpoint
	//
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// The name of the MCP server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-mcpserverdatadogconfiguration.html#cfn-devopsagent-association-mcpserverdatadogconfiguration-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The description of the MCP server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-mcpserverdatadogconfiguration.html#cfn-devopsagent-association-mcpserverdatadogconfiguration-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// When set to true, enables the Agent Space to create and update webhooks for receiving notifications and events from the service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-mcpserverdatadogconfiguration.html#cfn-devopsagent-association-mcpserverdatadogconfiguration-enablewebhookupdates
	//
	EnableWebhookUpdates interface{} `field:"optional" json:"enableWebhookUpdates" yaml:"enableWebhookUpdates"`
}

