package awsbedrockagentcore


// The target configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var schemaDefinitionProperty_ SchemaDefinitionProperty
//
//   targetConfigurationProperty := &TargetConfigurationProperty{
//   	Mcp: &McpTargetConfigurationProperty{
//   		ApiGateway: &ApiGatewayTargetConfigurationProperty{
//   			ApiGatewayToolConfiguration: &ApiGatewayToolConfigurationProperty{
//   				ToolFilters: []interface{}{
//   					&ApiGatewayToolFilterProperty{
//   						FilterPath: jsii.String("filterPath"),
//   						Methods: []*string{
//   							jsii.String("methods"),
//   						},
//   					},
//   				},
//
//   				// the properties below are optional
//   				ToolOverrides: []interface{}{
//   					&ApiGatewayToolOverrideProperty{
//   						Method: jsii.String("method"),
//   						Name: jsii.String("name"),
//   						Path: jsii.String("path"),
//
//   						// the properties below are optional
//   						Description: jsii.String("description"),
//   					},
//   				},
//   			},
//   			RestApiId: jsii.String("restApiId"),
//   			Stage: jsii.String("stage"),
//   		},
//   		Lambda: &McpLambdaTargetConfigurationProperty{
//   			LambdaArn: jsii.String("lambdaArn"),
//   			ToolSchema: &ToolSchemaProperty{
//   				InlinePayload: []interface{}{
//   					&ToolDefinitionProperty{
//   						Description: jsii.String("description"),
//   						InputSchema: &SchemaDefinitionProperty{
//   							Type: jsii.String("type"),
//
//   							// the properties below are optional
//   							Description: jsii.String("description"),
//   							Items: schemaDefinitionProperty_,
//   							Properties: map[string]interface{}{
//   								"propertiesKey": schemaDefinitionProperty_,
//   							},
//   							Required: []*string{
//   								jsii.String("required"),
//   							},
//   						},
//   						Name: jsii.String("name"),
//
//   						// the properties below are optional
//   						OutputSchema: &SchemaDefinitionProperty{
//   							Type: jsii.String("type"),
//
//   							// the properties below are optional
//   							Description: jsii.String("description"),
//   							Items: schemaDefinitionProperty_,
//   							Properties: map[string]interface{}{
//   								"propertiesKey": schemaDefinitionProperty_,
//   							},
//   							Required: []*string{
//   								jsii.String("required"),
//   							},
//   						},
//   					},
//   				},
//   				S3: &S3ConfigurationProperty{
//   					BucketOwnerAccountId: jsii.String("bucketOwnerAccountId"),
//   					Uri: jsii.String("uri"),
//   				},
//   			},
//   		},
//   		McpServer: &McpServerTargetConfigurationProperty{
//   			Endpoint: jsii.String("endpoint"),
//   		},
//   		OpenApiSchema: &ApiSchemaConfigurationProperty{
//   			InlinePayload: jsii.String("inlinePayload"),
//   			S3: &S3ConfigurationProperty{
//   				BucketOwnerAccountId: jsii.String("bucketOwnerAccountId"),
//   				Uri: jsii.String("uri"),
//   			},
//   		},
//   		SmithyModel: &ApiSchemaConfigurationProperty{
//   			InlinePayload: jsii.String("inlinePayload"),
//   			S3: &S3ConfigurationProperty{
//   				BucketOwnerAccountId: jsii.String("bucketOwnerAccountId"),
//   				Uri: jsii.String("uri"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-targetconfiguration.html
//
type CfnGatewayTarget_TargetConfigurationProperty struct {
	// The target configuration definition for MCP.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-targetconfiguration.html#cfn-bedrockagentcore-gatewaytarget-targetconfiguration-mcp
	//
	Mcp interface{} `field:"required" json:"mcp" yaml:"mcp"`
}

