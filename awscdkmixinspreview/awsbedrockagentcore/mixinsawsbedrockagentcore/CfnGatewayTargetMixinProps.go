package mixinsawsbedrockagentcore


// Properties for CfnGatewayTargetPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var schemaDefinitionProperty_ SchemaDefinitionProperty
//
//   cfnGatewayTargetMixinProps := &CfnGatewayTargetMixinProps{
//   	CredentialProviderConfigurations: []interface{}{
//   		&CredentialProviderConfigurationProperty{
//   			CredentialProvider: &CredentialProviderProperty{
//   				ApiKeyCredentialProvider: &ApiKeyCredentialProviderProperty{
//   					CredentialLocation: jsii.String("credentialLocation"),
//   					CredentialParameterName: jsii.String("credentialParameterName"),
//   					CredentialPrefix: jsii.String("credentialPrefix"),
//   					ProviderArn: jsii.String("providerArn"),
//   				},
//   				OauthCredentialProvider: &OAuthCredentialProviderProperty{
//   					CustomParameters: map[string]*string{
//   						"customParametersKey": jsii.String("customParameters"),
//   					},
//   					ProviderArn: jsii.String("providerArn"),
//   					Scopes: []*string{
//   						jsii.String("scopes"),
//   					},
//   				},
//   			},
//   			CredentialProviderType: jsii.String("credentialProviderType"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	GatewayIdentifier: jsii.String("gatewayIdentifier"),
//   	Name: jsii.String("name"),
//   	TargetConfiguration: &TargetConfigurationProperty{
//   		Mcp: &McpTargetConfigurationProperty{
//   			Lambda: &McpLambdaTargetConfigurationProperty{
//   				LambdaArn: jsii.String("lambdaArn"),
//   				ToolSchema: &ToolSchemaProperty{
//   					InlinePayload: []interface{}{
//   						&ToolDefinitionProperty{
//   							Description: jsii.String("description"),
//   							InputSchema: &SchemaDefinitionProperty{
//   								Description: jsii.String("description"),
//   								Items: schemaDefinitionProperty_,
//   								Properties: map[string]interface{}{
//   									"propertiesKey": schemaDefinitionProperty_,
//   								},
//   								Required: []*string{
//   									jsii.String("required"),
//   								},
//   								Type: jsii.String("type"),
//   							},
//   							Name: jsii.String("name"),
//   							OutputSchema: &SchemaDefinitionProperty{
//   								Description: jsii.String("description"),
//   								Items: schemaDefinitionProperty_,
//   								Properties: map[string]interface{}{
//   									"propertiesKey": schemaDefinitionProperty_,
//   								},
//   								Required: []*string{
//   									jsii.String("required"),
//   								},
//   								Type: jsii.String("type"),
//   							},
//   						},
//   					},
//   					S3: &S3ConfigurationProperty{
//   						BucketOwnerAccountId: jsii.String("bucketOwnerAccountId"),
//   						Uri: jsii.String("uri"),
//   					},
//   				},
//   			},
//   			McpServer: &McpServerTargetConfigurationProperty{
//   				Endpoint: jsii.String("endpoint"),
//   			},
//   			OpenApiSchema: &ApiSchemaConfigurationProperty{
//   				InlinePayload: jsii.String("inlinePayload"),
//   				S3: &S3ConfigurationProperty{
//   					BucketOwnerAccountId: jsii.String("bucketOwnerAccountId"),
//   					Uri: jsii.String("uri"),
//   				},
//   			},
//   			SmithyModel: &ApiSchemaConfigurationProperty{
//   				InlinePayload: jsii.String("inlinePayload"),
//   				S3: &S3ConfigurationProperty{
//   					BucketOwnerAccountId: jsii.String("bucketOwnerAccountId"),
//   					Uri: jsii.String("uri"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-gatewaytarget.html
//
type CfnGatewayTargetMixinProps struct {
	// The OAuth credential provider configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-gatewaytarget.html#cfn-bedrockagentcore-gatewaytarget-credentialproviderconfigurations
	//
	CredentialProviderConfigurations interface{} `field:"optional" json:"credentialProviderConfigurations" yaml:"credentialProviderConfigurations"`
	// The description for the gateway target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-gatewaytarget.html#cfn-bedrockagentcore-gatewaytarget-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The gateway ID for the gateway target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-gatewaytarget.html#cfn-bedrockagentcore-gatewaytarget-gatewayidentifier
	//
	GatewayIdentifier *string `field:"optional" json:"gatewayIdentifier" yaml:"gatewayIdentifier"`
	// The name for the gateway target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-gatewaytarget.html#cfn-bedrockagentcore-gatewaytarget-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The target configuration for the Smithy model target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-gatewaytarget.html#cfn-bedrockagentcore-gatewaytarget-targetconfiguration
	//
	TargetConfiguration interface{} `field:"optional" json:"targetConfiguration" yaml:"targetConfiguration"`
}

