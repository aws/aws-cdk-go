package mixinsawsbedrockagentcore


// The tool definition for the gateway.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var schemaDefinitionProperty_ SchemaDefinitionProperty
//
//   toolDefinitionProperty := &ToolDefinitionProperty{
//   	Description: jsii.String("description"),
//   	InputSchema: &SchemaDefinitionProperty{
//   		Description: jsii.String("description"),
//   		Items: schemaDefinitionProperty_,
//   		Properties: map[string]interface{}{
//   			"propertiesKey": schemaDefinitionProperty_,
//   		},
//   		Required: []*string{
//   			jsii.String("required"),
//   		},
//   		Type: jsii.String("type"),
//   	},
//   	Name: jsii.String("name"),
//   	OutputSchema: &SchemaDefinitionProperty{
//   		Description: jsii.String("description"),
//   		Items: schemaDefinitionProperty_,
//   		Properties: map[string]interface{}{
//   			"propertiesKey": schemaDefinitionProperty_,
//   		},
//   		Required: []*string{
//   			jsii.String("required"),
//   		},
//   		Type: jsii.String("type"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-tooldefinition.html
//
type CfnGatewayTargetPropsMixin_ToolDefinitionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-tooldefinition.html#cfn-bedrockagentcore-gatewaytarget-tooldefinition-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The input schema for the gateway target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-tooldefinition.html#cfn-bedrockagentcore-gatewaytarget-tooldefinition-inputschema
	//
	InputSchema interface{} `field:"optional" json:"inputSchema" yaml:"inputSchema"`
	// The tool name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-tooldefinition.html#cfn-bedrockagentcore-gatewaytarget-tooldefinition-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The tool definition output schema for the gateway target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-tooldefinition.html#cfn-bedrockagentcore-gatewaytarget-tooldefinition-outputschema
	//
	OutputSchema interface{} `field:"optional" json:"outputSchema" yaml:"outputSchema"`
}

