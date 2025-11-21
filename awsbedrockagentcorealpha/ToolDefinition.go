package awsbedrockagentcorealpha


// Tool definition for inline payload.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//
//   var schemaDefinition_ SchemaDefinition
//
//   toolDefinition := &ToolDefinition{
//   	Description: jsii.String("description"),
//   	InputSchema: &SchemaDefinition{
//   		Type: bedrock_agentcore_alpha.SchemaDefinitionType_STRING,
//
//   		// the properties below are optional
//   		Description: jsii.String("description"),
//   		Items: schemaDefinition_,
//   		Properties: map[string]SchemaDefinition{
//   			"propertiesKey": schemaDefinition_,
//   		},
//   		Required: []*string{
//   			jsii.String("required"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	OutputSchema: &SchemaDefinition{
//   		Type: bedrock_agentcore_alpha.SchemaDefinitionType_STRING,
//
//   		// the properties below are optional
//   		Description: jsii.String("description"),
//   		Items: schemaDefinition_,
//   		Properties: map[string]SchemaDefinition{
//   			"propertiesKey": schemaDefinition_,
//   		},
//   		Required: []*string{
//   			jsii.String("required"),
//   		},
//   	},
//   }
//
// Experimental.
type ToolDefinition struct {
	// The description of the tool.
	//
	// This description provides information about the purpose and usage of the tool.
	// Experimental.
	Description *string `field:"required" json:"description" yaml:"description"`
	// The input schema for the tool.
	//
	// This schema defines the structure of the input that the tool accepts.
	// Experimental.
	InputSchema *SchemaDefinition `field:"required" json:"inputSchema" yaml:"inputSchema"`
	// The name of the tool.
	//
	// This name identifies the tool in the Model Context Protocol.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The output schema for the tool.
	//
	// This schema defines the structure of the output that the tool produces.
	// Default: - No output schema.
	//
	// Experimental.
	OutputSchema *SchemaDefinition `field:"optional" json:"outputSchema" yaml:"outputSchema"`
}

