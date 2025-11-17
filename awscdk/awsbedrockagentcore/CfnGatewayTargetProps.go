package awsbedrockagentcore


// Properties for defining a `CfnGatewayTarget`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var schemaDefinitionProperty_ SchemaDefinitionProperty
//
//   cfnGatewayTargetProps := &CfnGatewayTargetProps{
//   	CredentialProviderConfigurations: []interface{}{
//   		&CredentialProviderConfigurationProperty{
//   			CredentialProviderType: jsii.String("credentialProviderType"),
//
//   			// the properties below are optional
//   			CredentialProvider: &CredentialProviderProperty{
//   				ApiKeyCredentialProvider: &ApiKeyCredentialProviderProperty{
//   					ProviderArn: jsii.String("providerArn"),
//
//   					// the properties below are optional
//   					CredentialLocation: jsii.String("credentialLocation"),
//   					CredentialParameterName: jsii.String("credentialParameterName"),
//   					CredentialPrefix: jsii.String("credentialPrefix"),
//   				},
//   				OauthCredentialProvider: &OAuthCredentialProviderProperty{
//   					ProviderArn: jsii.String("providerArn"),
//   					Scopes: []*string{
//   						jsii.String("scopes"),
//   					},
//
//   					// the properties below are optional
//   					CustomParameters: map[string]*string{
//   						"customParametersKey": jsii.String("customParameters"),
//   					},
//   				},
//   			},
//   		},
//   	},
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
//   								Type: jsii.String("type"),
//
//   								// the properties below are optional
//   								Description: jsii.String("description"),
//   								Items: schemaDefinitionProperty_,
//   								Properties: map[string]interface{}{
//   									"propertiesKey": schemaDefinitionProperty_,
//   								},
//   								Required: []*string{
//   									jsii.String("required"),
//   								},
//   							},
//   							Name: jsii.String("name"),
//
//   							// the properties below are optional
//   							OutputSchema: &SchemaDefinitionProperty{
//   								Type: jsii.String("type"),
//
//   								// the properties below are optional
//   								Description: jsii.String("description"),
//   								Items: schemaDefinitionProperty_,
//   								Properties: map[string]interface{}{
//   									"propertiesKey": schemaDefinitionProperty_,
//   								},
//   								Required: []*string{
//   									jsii.String("required"),
//   								},
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
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	GatewayIdentifier: jsii.String("gatewayIdentifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-gatewaytarget.html
//
type CfnGatewayTargetProps struct {
	// The OAuth credential provider configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-gatewaytarget.html#cfn-bedrockagentcore-gatewaytarget-credentialproviderconfigurations
	//
	CredentialProviderConfigurations interface{} `field:"required" json:"credentialProviderConfigurations" yaml:"credentialProviderConfigurations"`
	// The name for the gateway target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-gatewaytarget.html#cfn-bedrockagentcore-gatewaytarget-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The target configuration for the Smithy model target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-gatewaytarget.html#cfn-bedrockagentcore-gatewaytarget-targetconfiguration
	//
	TargetConfiguration interface{} `field:"required" json:"targetConfiguration" yaml:"targetConfiguration"`
	// The description for the gateway target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-gatewaytarget.html#cfn-bedrockagentcore-gatewaytarget-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The gateway ID for the gateway target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-gatewaytarget.html#cfn-bedrockagentcore-gatewaytarget-gatewayidentifier
	//
	GatewayIdentifier *string `field:"optional" json:"gatewayIdentifier" yaml:"gatewayIdentifier"`
}

