package awsbedrockagentcorealpha


// Schema definition for tool input/output.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//
//   var schemaDefinition_ SchemaDefinition
//
//   schemaDefinition := &SchemaDefinition{
//   	Type: bedrock_agentcore_alpha.SchemaDefinitionType_STRING,
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Items: schemaDefinition_,
//   	Properties: map[string]SchemaDefinition{
//   		"propertiesKey": schemaDefinition_,
//   	},
//   	Required: []*string{
//   		jsii.String("required"),
//   	},
//   }
//
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
type SchemaDefinition struct {
	// The type of the schema definition.
	//
	// This field specifies the data type of the schema.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	Type SchemaDefinitionType `field:"required" json:"type" yaml:"type"`
	// The description of the schema definition.
	//
	// This description provides information about the purpose and usage of the schema.
	// Default: - No description.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The items in the schema definition.
	//
	// This field is used for array types to define the structure of the array elements.
	// Default: - No items definition.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	Items **SchemaDefinition `field:"optional" json:"items" yaml:"items"`
	// The properties of the schema definition.
	//
	// These properties define the fields in the schema.
	// Default: - No properties.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	Properties *map[string]*SchemaDefinition `field:"optional" json:"properties" yaml:"properties"`
	// The required fields in the schema definition.
	//
	// These fields must be provided when using the schema.
	// Default: - No required fields.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	Required *[]*string `field:"optional" json:"required" yaml:"required"`
}

