package previewawsbedrockagentcoremixins


// The Lambda target configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var schemaDefinitionProperty_ SchemaDefinitionProperty
//
//   mcpLambdaTargetConfigurationProperty := &McpLambdaTargetConfigurationProperty{
//   	LambdaArn: jsii.String("lambdaArn"),
//   	ToolSchema: &ToolSchemaProperty{
//   		InlinePayload: []interface{}{
//   			&ToolDefinitionProperty{
//   				Description: jsii.String("description"),
//   				InputSchema: &SchemaDefinitionProperty{
//   					Description: jsii.String("description"),
//   					Items: schemaDefinitionProperty_,
//   					Properties: map[string]interface{}{
//   						"propertiesKey": schemaDefinitionProperty_,
//   					},
//   					Required: []*string{
//   						jsii.String("required"),
//   					},
//   					Type: jsii.String("type"),
//   				},
//   				Name: jsii.String("name"),
//   				OutputSchema: &SchemaDefinitionProperty{
//   					Description: jsii.String("description"),
//   					Items: schemaDefinitionProperty_,
//   					Properties: map[string]interface{}{
//   						"propertiesKey": schemaDefinitionProperty_,
//   					},
//   					Required: []*string{
//   						jsii.String("required"),
//   					},
//   					Type: jsii.String("type"),
//   				},
//   			},
//   		},
//   		S3: &S3ConfigurationProperty{
//   			BucketOwnerAccountId: jsii.String("bucketOwnerAccountId"),
//   			Uri: jsii.String("uri"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-mcplambdatargetconfiguration.html
//
type CfnGatewayTargetPropsMixin_McpLambdaTargetConfigurationProperty struct {
	// The ARN of the Lambda target configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-mcplambdatargetconfiguration.html#cfn-bedrockagentcore-gatewaytarget-mcplambdatargetconfiguration-lambdaarn
	//
	LambdaArn *string `field:"optional" json:"lambdaArn" yaml:"lambdaArn"`
	// The tool schema configuration for the gateway target MCP configuration for Lambda.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-mcplambdatargetconfiguration.html#cfn-bedrockagentcore-gatewaytarget-mcplambdatargetconfiguration-toolschema
	//
	ToolSchema interface{} `field:"optional" json:"toolSchema" yaml:"toolSchema"`
}

