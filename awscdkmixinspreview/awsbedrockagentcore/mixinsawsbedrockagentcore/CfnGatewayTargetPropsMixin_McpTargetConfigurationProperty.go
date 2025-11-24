package mixinsawsbedrockagentcore


// The MCP target configuration for the gateway target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var schemaDefinitionProperty_ SchemaDefinitionProperty
//
//   mcpTargetConfigurationProperty := &McpTargetConfigurationProperty{
//   	Lambda: &McpLambdaTargetConfigurationProperty{
//   		LambdaArn: jsii.String("lambdaArn"),
//   		ToolSchema: &ToolSchemaProperty{
//   			InlinePayload: []interface{}{
//   				&ToolDefinitionProperty{
//   					Description: jsii.String("description"),
//   					InputSchema: &SchemaDefinitionProperty{
//   						Description: jsii.String("description"),
//   						Items: schemaDefinitionProperty_,
//   						Properties: map[string]interface{}{
//   							"propertiesKey": schemaDefinitionProperty_,
//   						},
//   						Required: []*string{
//   							jsii.String("required"),
//   						},
//   						Type: jsii.String("type"),
//   					},
//   					Name: jsii.String("name"),
//   					OutputSchema: &SchemaDefinitionProperty{
//   						Description: jsii.String("description"),
//   						Items: schemaDefinitionProperty_,
//   						Properties: map[string]interface{}{
//   							"propertiesKey": schemaDefinitionProperty_,
//   						},
//   						Required: []*string{
//   							jsii.String("required"),
//   						},
//   						Type: jsii.String("type"),
//   					},
//   				},
//   			},
//   			S3: &S3ConfigurationProperty{
//   				BucketOwnerAccountId: jsii.String("bucketOwnerAccountId"),
//   				Uri: jsii.String("uri"),
//   			},
//   		},
//   	},
//   	McpServer: &McpServerTargetConfigurationProperty{
//   		Endpoint: jsii.String("endpoint"),
//   	},
//   	OpenApiSchema: &ApiSchemaConfigurationProperty{
//   		InlinePayload: jsii.String("inlinePayload"),
//   		S3: &S3ConfigurationProperty{
//   			BucketOwnerAccountId: jsii.String("bucketOwnerAccountId"),
//   			Uri: jsii.String("uri"),
//   		},
//   	},
//   	SmithyModel: &ApiSchemaConfigurationProperty{
//   		InlinePayload: jsii.String("inlinePayload"),
//   		S3: &S3ConfigurationProperty{
//   			BucketOwnerAccountId: jsii.String("bucketOwnerAccountId"),
//   			Uri: jsii.String("uri"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-mcptargetconfiguration.html
//
type CfnGatewayTargetPropsMixin_McpTargetConfigurationProperty struct {
	// The Lambda MCP configuration for the gateway target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-mcptargetconfiguration.html#cfn-bedrockagentcore-gatewaytarget-mcptargetconfiguration-lambda
	//
	Lambda interface{} `field:"optional" json:"lambda" yaml:"lambda"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-mcptargetconfiguration.html#cfn-bedrockagentcore-gatewaytarget-mcptargetconfiguration-mcpserver
	//
	McpServer interface{} `field:"optional" json:"mcpServer" yaml:"mcpServer"`
	// The OpenApi schema for the gateway target MCP configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-mcptargetconfiguration.html#cfn-bedrockagentcore-gatewaytarget-mcptargetconfiguration-openapischema
	//
	OpenApiSchema interface{} `field:"optional" json:"openApiSchema" yaml:"openApiSchema"`
	// The target configuration for the Smithy model target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-mcptargetconfiguration.html#cfn-bedrockagentcore-gatewaytarget-mcptargetconfiguration-smithymodel
	//
	SmithyModel interface{} `field:"optional" json:"smithyModel" yaml:"smithyModel"`
}

