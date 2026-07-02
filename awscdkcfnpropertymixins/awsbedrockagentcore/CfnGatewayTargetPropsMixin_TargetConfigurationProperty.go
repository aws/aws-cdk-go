package awsbedrockagentcore


// The target configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   var parameterValues interface{}
//   var schemaDefinitionProperty_ SchemaDefinitionProperty
//
//   targetConfigurationProperty := &TargetConfigurationProperty{
//   	Http: &HttpTargetConfigurationProperty{
//   		AgentcoreRuntime: &RuntimeTargetConfigurationProperty{
//   			Arn: jsii.String("arn"),
//   			Qualifier: jsii.String("qualifier"),
//   			Schema: &HttpApiSchemaConfigurationProperty{
//   				Source: &ApiSchemaConfigurationProperty{
//   					InlinePayload: jsii.String("inlinePayload"),
//   					S3: &S3ConfigurationProperty{
//   						BucketOwnerAccountId: jsii.String("bucketOwnerAccountId"),
//   						Uri: jsii.String("uri"),
//   					},
//   				},
//   			},
//   		},
//   		Passthrough: &PassthroughTargetConfigurationProperty{
//   			Endpoint: jsii.String("endpoint"),
//   			ProtocolType: jsii.String("protocolType"),
//   			Schema: &HttpApiSchemaConfigurationProperty{
//   				Source: &ApiSchemaConfigurationProperty{
//   					InlinePayload: jsii.String("inlinePayload"),
//   					S3: &S3ConfigurationProperty{
//   						BucketOwnerAccountId: jsii.String("bucketOwnerAccountId"),
//   						Uri: jsii.String("uri"),
//   					},
//   				},
//   			},
//   			StickinessConfiguration: &StickinessConfigurationProperty{
//   				Identifier: jsii.String("identifier"),
//   				Timeout: jsii.Number(123),
//   			},
//   		},
//   	},
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
//   				ToolOverrides: []interface{}{
//   					&ApiGatewayToolOverrideProperty{
//   						Description: jsii.String("description"),
//   						Method: jsii.String("method"),
//   						Name: jsii.String("name"),
//   						Path: jsii.String("path"),
//   					},
//   				},
//   			},
//   			RestApiId: jsii.String("restApiId"),
//   			Stage: jsii.String("stage"),
//   		},
//   		Connector: &ConnectorTargetConfigurationProperty{
//   			Configurations: []interface{}{
//   				&ConnectorConfigurationProperty{
//   					Description: jsii.String("description"),
//   					Name: jsii.String("name"),
//   					ParameterOverrides: []interface{}{
//   						&ConnectorParameterOverrideProperty{
//   							Description: jsii.String("description"),
//   							Path: jsii.String("path"),
//   							Visible: jsii.Boolean(false),
//   						},
//   					},
//   					ParameterValues: parameterValues,
//   				},
//   			},
//   			Enabled: []*string{
//   				jsii.String("enabled"),
//   			},
//   			Source: &ConnectorSourceProperty{
//   				ConnectorId: jsii.String("connectorId"),
//   			},
//   		},
//   		Lambda: &McpLambdaTargetConfigurationProperty{
//   			LambdaArn: jsii.String("lambdaArn"),
//   			ToolSchema: &ToolSchemaProperty{
//   				InlinePayload: []interface{}{
//   					&ToolDefinitionProperty{
//   						Description: jsii.String("description"),
//   						InputSchema: &SchemaDefinitionProperty{
//   							Description: jsii.String("description"),
//   							Items: schemaDefinitionProperty_,
//   							Properties: map[string]interface{}{
//   								"propertiesKey": schemaDefinitionProperty_,
//   							},
//   							Required: []*string{
//   								jsii.String("required"),
//   							},
//   							Type: jsii.String("type"),
//   						},
//   						Name: jsii.String("name"),
//   						OutputSchema: &SchemaDefinitionProperty{
//   							Description: jsii.String("description"),
//   							Items: schemaDefinitionProperty_,
//   							Properties: map[string]interface{}{
//   								"propertiesKey": schemaDefinitionProperty_,
//   							},
//   							Required: []*string{
//   								jsii.String("required"),
//   							},
//   							Type: jsii.String("type"),
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
//   			ListingMode: jsii.String("listingMode"),
//   			McpToolSchema: &McpToolSchemaConfigurationProperty{
//   				InlinePayload: jsii.String("inlinePayload"),
//   				S3: &S3ConfigurationProperty{
//   					BucketOwnerAccountId: jsii.String("bucketOwnerAccountId"),
//   					Uri: jsii.String("uri"),
//   				},
//   			},
//   			ResourcePriority: jsii.Number(123),
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
type CfnGatewayTargetPropsMixin_TargetConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-targetconfiguration.html#cfn-bedrockagentcore-gatewaytarget-targetconfiguration-http
	//
	Http interface{} `field:"optional" json:"http" yaml:"http"`
	// The target configuration definition for MCP.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-targetconfiguration.html#cfn-bedrockagentcore-gatewaytarget-targetconfiguration-mcp
	//
	Mcp interface{} `field:"optional" json:"mcp" yaml:"mcp"`
}

