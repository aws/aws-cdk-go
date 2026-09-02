package awsagentregistry


// The MCP tools descriptor.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mcpToolsDescriptorProperty := &McpToolsDescriptorProperty{
//   	Data: jsii.String("data"),
//   	DataSchemaVersion: jsii.String("dataSchemaVersion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registryrecord-mcptoolsdescriptor.html
//
type CfnRegistryRecord_McpToolsDescriptorProperty struct {
	// Descriptor payload data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registryrecord-mcptoolsdescriptor.html#cfn-agentregistry-registryrecord-mcptoolsdescriptor-data
	//
	Data *string `field:"optional" json:"data" yaml:"data"`
	// Version of the tools descriptor schema.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registryrecord-mcptoolsdescriptor.html#cfn-agentregistry-registryrecord-mcptoolsdescriptor-dataschemaversion
	//
	DataSchemaVersion *string `field:"optional" json:"dataSchemaVersion" yaml:"dataSchemaVersion"`
}

