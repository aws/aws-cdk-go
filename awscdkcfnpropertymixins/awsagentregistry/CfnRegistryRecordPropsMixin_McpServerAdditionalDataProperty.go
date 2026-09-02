package awsagentregistry


// Additional data associated with an MCP server descriptor.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   mcpServerAdditionalDataProperty := &McpServerAdditionalDataProperty{
//   	Tools: &McpToolsDescriptorProperty{
//   		Data: jsii.String("data"),
//   		DataSchemaVersion: jsii.String("dataSchemaVersion"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registryrecord-mcpserveradditionaldata.html
//
type CfnRegistryRecordPropsMixin_McpServerAdditionalDataProperty struct {
	// The MCP tools descriptor.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registryrecord-mcpserveradditionaldata.html#cfn-agentregistry-registryrecord-mcpserveradditionaldata-tools
	//
	Tools interface{} `field:"optional" json:"tools" yaml:"tools"`
}

