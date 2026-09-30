package awsglue

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnConnectionType`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnConnectionTypeProps := &CfnConnectionTypeProps{
//   	ConnectionType: jsii.String("connectionType"),
//   	RestConfiguration: &RestConfigurationProperty{
//   		EntityConfigurations: map[string]interface{}{
//   			"entityConfigurationsKey": &EntityConfigurationProperty{
//   				"schema": map[string]interface{}{
//   					"schemaKey": &FieldDefinitionProperty{
//   						"fieldDataType": jsii.String("fieldDataType"),
//   						"name": jsii.String("name"),
//
//   						// the properties below are optional
//   						"filterOverrides": &FilterOverridesProperty{
//   							"betweenConfiguration": &BetweenConfigurationProperty{
//   								"highBoundKey": jsii.String("highBoundKey"),
//   								"lowBoundKey": jsii.String("lowBoundKey"),
//   								"template": jsii.String("template"),
//   							},
//   							"dateTimeFormat": jsii.String("dateTimeFormat"),
//   							"fieldName": jsii.String("fieldName"),
//   							"operatorMappings": map[string]*string{
//   								"operatorMappingsKey": jsii.String("operatorMappings"),
//   							},
//   						},
//   						"isNullable": jsii.Boolean(false),
//   						"isOrderable": jsii.Boolean(false),
//   						"isPartitionable": jsii.Boolean(false),
//   						"isQueryable": jsii.Boolean(false),
//   						"responseDateFormat": jsii.String("responseDateFormat"),
//   					},
//   				},
//   				"sourceConfiguration": &SourceConfigurationProperty{
//   					"filterConfiguration": &FilterConfigurationProperty{
//   						"filterMode": jsii.String("filterMode"),
//
//   						// the properties below are optional
//   						"betweenConfiguration": &BetweenConfigurationProperty{
//   							"highBoundKey": jsii.String("highBoundKey"),
//   							"lowBoundKey": jsii.String("lowBoundKey"),
//   							"template": jsii.String("template"),
//   						},
//   						"dateTimeFormat": jsii.String("dateTimeFormat"),
//   						"filterStringConfiguration": &FilterStringConfigurationProperty{
//   							"queryParameterName": jsii.String("queryParameterName"),
//
//   							// the properties below are optional
//   							"quoteCharacter": jsii.String("quoteCharacter"),
//   							"quoteStringValues": jsii.Boolean(false),
//   						},
//   						"operatorMappings": map[string]*string{
//   							"operatorMappingsKey": jsii.String("operatorMappings"),
//   						},
//   						"stripQuotes": jsii.Boolean(false),
//   					},
//   					"paginationConfiguration": &PaginationConfigurationProperty{
//   						"cursorConfiguration": &CursorConfigurationProperty{
//   							"nextPage": &ExtractedParameterProperty{
//   								"defaultValue": jsii.String("defaultValue"),
//   								"key": jsii.String("key"),
//   								"propertyLocation": jsii.String("propertyLocation"),
//   								"value": &ResponseExtractionMappingProperty{
//   									"contentPath": jsii.String("contentPath"),
//   									"headerKey": jsii.String("headerKey"),
//   								},
//   							},
//
//   							// the properties below are optional
//   							"limitParameter": &ExtractedParameterProperty{
//   								"defaultValue": jsii.String("defaultValue"),
//   								"key": jsii.String("key"),
//   								"propertyLocation": jsii.String("propertyLocation"),
//   								"value": &ResponseExtractionMappingProperty{
//   									"contentPath": jsii.String("contentPath"),
//   									"headerKey": jsii.String("headerKey"),
//   								},
//   							},
//   						},
//   						"offsetConfiguration": &OffsetConfigurationProperty{
//   							"limitParameter": &ExtractedParameterProperty{
//   								"defaultValue": jsii.String("defaultValue"),
//   								"key": jsii.String("key"),
//   								"propertyLocation": jsii.String("propertyLocation"),
//   								"value": &ResponseExtractionMappingProperty{
//   									"contentPath": jsii.String("contentPath"),
//   									"headerKey": jsii.String("headerKey"),
//   								},
//   							},
//   							"offsetParameter": &ExtractedParameterProperty{
//   								"defaultValue": jsii.String("defaultValue"),
//   								"key": jsii.String("key"),
//   								"propertyLocation": jsii.String("propertyLocation"),
//   								"value": &ResponseExtractionMappingProperty{
//   									"contentPath": jsii.String("contentPath"),
//   									"headerKey": jsii.String("headerKey"),
//   								},
//   							},
//   						},
//   					},
//   					"requestMethod": jsii.String("requestMethod"),
//   					"requestParameters": []interface{}{
//   						&ConnectorPropertyProperty{
//   							"name": jsii.String("name"),
//   							"propertyType": jsii.String("propertyType"),
//   							"required": jsii.Boolean(false),
//
//   							// the properties below are optional
//   							"allowedValues": []*string{
//   								jsii.String("allowedValues"),
//   							},
//   							"defaultValue": jsii.String("defaultValue"),
//   							"keyOverride": jsii.String("keyOverride"),
//   							"propertyLocation": jsii.String("propertyLocation"),
//   						},
//   					},
//   					"requestPath": jsii.String("requestPath"),
//   					"responseConfiguration": &ResponseConfigurationProperty{
//   						"resultPath": jsii.String("resultPath"),
//
//   						// the properties below are optional
//   						"errorPath": jsii.String("errorPath"),
//   					},
//   				},
//   			},
//   		},
//   		GlobalSourceConfiguration: &SourceConfigurationProperty{
//   			FilterConfiguration: &FilterConfigurationProperty{
//   				FilterMode: jsii.String("filterMode"),
//
//   				// the properties below are optional
//   				BetweenConfiguration: &BetweenConfigurationProperty{
//   					HighBoundKey: jsii.String("highBoundKey"),
//   					LowBoundKey: jsii.String("lowBoundKey"),
//   					Template: jsii.String("template"),
//   				},
//   				DateTimeFormat: jsii.String("dateTimeFormat"),
//   				FilterStringConfiguration: &FilterStringConfigurationProperty{
//   					QueryParameterName: jsii.String("queryParameterName"),
//
//   					// the properties below are optional
//   					QuoteCharacter: jsii.String("quoteCharacter"),
//   					QuoteStringValues: jsii.Boolean(false),
//   				},
//   				OperatorMappings: map[string]*string{
//   					"operatorMappingsKey": jsii.String("operatorMappings"),
//   				},
//   				StripQuotes: jsii.Boolean(false),
//   			},
//   			PaginationConfiguration: &PaginationConfigurationProperty{
//   				CursorConfiguration: &CursorConfigurationProperty{
//   					NextPage: &ExtractedParameterProperty{
//   						DefaultValue: jsii.String("defaultValue"),
//   						Key: jsii.String("key"),
//   						PropertyLocation: jsii.String("propertyLocation"),
//   						Value: &ResponseExtractionMappingProperty{
//   							ContentPath: jsii.String("contentPath"),
//   							HeaderKey: jsii.String("headerKey"),
//   						},
//   					},
//
//   					// the properties below are optional
//   					LimitParameter: &ExtractedParameterProperty{
//   						DefaultValue: jsii.String("defaultValue"),
//   						Key: jsii.String("key"),
//   						PropertyLocation: jsii.String("propertyLocation"),
//   						Value: &ResponseExtractionMappingProperty{
//   							ContentPath: jsii.String("contentPath"),
//   							HeaderKey: jsii.String("headerKey"),
//   						},
//   					},
//   				},
//   				OffsetConfiguration: &OffsetConfigurationProperty{
//   					LimitParameter: &ExtractedParameterProperty{
//   						DefaultValue: jsii.String("defaultValue"),
//   						Key: jsii.String("key"),
//   						PropertyLocation: jsii.String("propertyLocation"),
//   						Value: &ResponseExtractionMappingProperty{
//   							ContentPath: jsii.String("contentPath"),
//   							HeaderKey: jsii.String("headerKey"),
//   						},
//   					},
//   					OffsetParameter: &ExtractedParameterProperty{
//   						DefaultValue: jsii.String("defaultValue"),
//   						Key: jsii.String("key"),
//   						PropertyLocation: jsii.String("propertyLocation"),
//   						Value: &ResponseExtractionMappingProperty{
//   							ContentPath: jsii.String("contentPath"),
//   							HeaderKey: jsii.String("headerKey"),
//   						},
//   					},
//   				},
//   			},
//   			RequestMethod: jsii.String("requestMethod"),
//   			RequestParameters: []interface{}{
//   				&ConnectorPropertyProperty{
//   					Name: jsii.String("name"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//
//   					// the properties below are optional
//   					AllowedValues: []*string{
//   						jsii.String("allowedValues"),
//   					},
//   					DefaultValue: jsii.String("defaultValue"),
//   					KeyOverride: jsii.String("keyOverride"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   				},
//   			},
//   			RequestPath: jsii.String("requestPath"),
//   			ResponseConfiguration: &ResponseConfigurationProperty{
//   				ResultPath: jsii.String("resultPath"),
//
//   				// the properties below are optional
//   				ErrorPath: jsii.String("errorPath"),
//   			},
//   		},
//   		ValidationEndpointConfiguration: &ValidationEndpointConfigurationProperty{
//   			RequestMethod: jsii.String("requestMethod"),
//   			RequestPath: jsii.String("requestPath"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	ConnectionProperties: &ConnectionPropertiesConfigurationProperty{
//   		AdditionalRequestParameters: []interface{}{
//   			&ConnectorPropertyProperty{
//   				Name: jsii.String("name"),
//   				PropertyType: jsii.String("propertyType"),
//   				Required: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				AllowedValues: []*string{
//   					jsii.String("allowedValues"),
//   				},
//   				DefaultValue: jsii.String("defaultValue"),
//   				KeyOverride: jsii.String("keyOverride"),
//   				PropertyLocation: jsii.String("propertyLocation"),
//   			},
//   		},
//   		Url: &ConnectorPropertyProperty{
//   			Name: jsii.String("name"),
//   			PropertyType: jsii.String("propertyType"),
//   			Required: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			AllowedValues: []*string{
//   				jsii.String("allowedValues"),
//   			},
//   			DefaultValue: jsii.String("defaultValue"),
//   			KeyOverride: jsii.String("keyOverride"),
//   			PropertyLocation: jsii.String("propertyLocation"),
//   		},
//   	},
//   	ConnectorAuthenticationConfiguration: &ConnectorAuthenticationConfigurationProperty{
//   		AuthenticationTypes: []*string{
//   			jsii.String("authenticationTypes"),
//   		},
//
//   		// the properties below are optional
//   		BasicAuthenticationProperties: &BasicAuthenticationPropertiesProperty{
//   			Password: &SecretConnectorPropertyProperty{
//   				Name: jsii.String("name"),
//   				PropertyType: jsii.String("propertyType"),
//   				Required: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				KeyOverride: jsii.String("keyOverride"),
//   				PropertyLocation: jsii.String("propertyLocation"),
//   			},
//   			Username: &SecretConnectorPropertyProperty{
//   				Name: jsii.String("name"),
//   				PropertyType: jsii.String("propertyType"),
//   				Required: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				KeyOverride: jsii.String("keyOverride"),
//   				PropertyLocation: jsii.String("propertyLocation"),
//   			},
//   		},
//   		CustomAuthenticationProperties: &CustomAuthenticationPropertiesProperty{
//   			AuthenticationParameters: []interface{}{
//   				&SecretConnectorPropertyProperty{
//   					Name: jsii.String("name"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//
//   					// the properties below are optional
//   					KeyOverride: jsii.String("keyOverride"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   				},
//   			},
//   		},
//   		OAuth2Properties: &ConnectorOAuth2PropertiesProperty{
//   			OAuth2GrantType: jsii.String("oAuth2GrantType"),
//
//   			// the properties below are optional
//   			AuthorizationCodeProperties: &ConnectorAuthorizationCodePropertiesProperty{
//   				AuthorizationCode: &ConnectorPropertyProperty{
//   					Name: jsii.String("name"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//
//   					// the properties below are optional
//   					AllowedValues: []*string{
//   						jsii.String("allowedValues"),
//   					},
//   					DefaultValue: jsii.String("defaultValue"),
//   					KeyOverride: jsii.String("keyOverride"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   				},
//   				AuthorizationCodeUrl: &ConnectorPropertyProperty{
//   					Name: jsii.String("name"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//
//   					// the properties below are optional
//   					AllowedValues: []*string{
//   						jsii.String("allowedValues"),
//   					},
//   					DefaultValue: jsii.String("defaultValue"),
//   					KeyOverride: jsii.String("keyOverride"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   				},
//   				ClientId: &SecretConnectorPropertyProperty{
//   					Name: jsii.String("name"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//
//   					// the properties below are optional
//   					KeyOverride: jsii.String("keyOverride"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   				},
//   				ClientSecret: &SecretConnectorPropertyProperty{
//   					Name: jsii.String("name"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//
//   					// the properties below are optional
//   					KeyOverride: jsii.String("keyOverride"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   				},
//   				ContentType: jsii.String("contentType"),
//   				Prompt: &ConnectorPropertyProperty{
//   					Name: jsii.String("name"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//
//   					// the properties below are optional
//   					AllowedValues: []*string{
//   						jsii.String("allowedValues"),
//   					},
//   					DefaultValue: jsii.String("defaultValue"),
//   					KeyOverride: jsii.String("keyOverride"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   				},
//   				RedirectUri: &ConnectorPropertyProperty{
//   					Name: jsii.String("name"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//
//   					// the properties below are optional
//   					AllowedValues: []*string{
//   						jsii.String("allowedValues"),
//   					},
//   					DefaultValue: jsii.String("defaultValue"),
//   					KeyOverride: jsii.String("keyOverride"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   				},
//   				RequestMethod: jsii.String("requestMethod"),
//   				Scope: &ConnectorPropertyProperty{
//   					Name: jsii.String("name"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//
//   					// the properties below are optional
//   					AllowedValues: []*string{
//   						jsii.String("allowedValues"),
//   					},
//   					DefaultValue: jsii.String("defaultValue"),
//   					KeyOverride: jsii.String("keyOverride"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   				},
//   				TokenUrl: &ConnectorPropertyProperty{
//   					Name: jsii.String("name"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//
//   					// the properties below are optional
//   					AllowedValues: []*string{
//   						jsii.String("allowedValues"),
//   					},
//   					DefaultValue: jsii.String("defaultValue"),
//   					KeyOverride: jsii.String("keyOverride"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   				},
//   				TokenUrlParameters: []interface{}{
//   					&ConnectorPropertyProperty{
//   						Name: jsii.String("name"),
//   						PropertyType: jsii.String("propertyType"),
//   						Required: jsii.Boolean(false),
//
//   						// the properties below are optional
//   						AllowedValues: []*string{
//   							jsii.String("allowedValues"),
//   						},
//   						DefaultValue: jsii.String("defaultValue"),
//   						KeyOverride: jsii.String("keyOverride"),
//   						PropertyLocation: jsii.String("propertyLocation"),
//   					},
//   				},
//   			},
//   			ClientCredentialsProperties: &ClientCredentialsPropertiesProperty{
//   				ClientId: &SecretConnectorPropertyProperty{
//   					Name: jsii.String("name"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//
//   					// the properties below are optional
//   					KeyOverride: jsii.String("keyOverride"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   				},
//   				ClientSecret: &SecretConnectorPropertyProperty{
//   					Name: jsii.String("name"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//
//   					// the properties below are optional
//   					KeyOverride: jsii.String("keyOverride"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   				},
//   				ContentType: jsii.String("contentType"),
//   				RequestMethod: jsii.String("requestMethod"),
//   				Scope: &ConnectorPropertyProperty{
//   					Name: jsii.String("name"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//
//   					// the properties below are optional
//   					AllowedValues: []*string{
//   						jsii.String("allowedValues"),
//   					},
//   					DefaultValue: jsii.String("defaultValue"),
//   					KeyOverride: jsii.String("keyOverride"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   				},
//   				TokenUrl: &ConnectorPropertyProperty{
//   					Name: jsii.String("name"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//
//   					// the properties below are optional
//   					AllowedValues: []*string{
//   						jsii.String("allowedValues"),
//   					},
//   					DefaultValue: jsii.String("defaultValue"),
//   					KeyOverride: jsii.String("keyOverride"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   				},
//   				TokenUrlParameters: []interface{}{
//   					&ConnectorPropertyProperty{
//   						Name: jsii.String("name"),
//   						PropertyType: jsii.String("propertyType"),
//   						Required: jsii.Boolean(false),
//
//   						// the properties below are optional
//   						AllowedValues: []*string{
//   							jsii.String("allowedValues"),
//   						},
//   						DefaultValue: jsii.String("defaultValue"),
//   						KeyOverride: jsii.String("keyOverride"),
//   						PropertyLocation: jsii.String("propertyLocation"),
//   					},
//   				},
//   			},
//   			JwtBearerProperties: &JWTBearerPropertiesProperty{
//   				ContentType: jsii.String("contentType"),
//   				JwtToken: &SecretConnectorPropertyProperty{
//   					Name: jsii.String("name"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//
//   					// the properties below are optional
//   					KeyOverride: jsii.String("keyOverride"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   				},
//   				RequestMethod: jsii.String("requestMethod"),
//   				TokenUrl: &ConnectorPropertyProperty{
//   					Name: jsii.String("name"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//
//   					// the properties below are optional
//   					AllowedValues: []*string{
//   						jsii.String("allowedValues"),
//   					},
//   					DefaultValue: jsii.String("defaultValue"),
//   					KeyOverride: jsii.String("keyOverride"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   				},
//   				TokenUrlParameters: []interface{}{
//   					&ConnectorPropertyProperty{
//   						Name: jsii.String("name"),
//   						PropertyType: jsii.String("propertyType"),
//   						Required: jsii.Boolean(false),
//
//   						// the properties below are optional
//   						AllowedValues: []*string{
//   							jsii.String("allowedValues"),
//   						},
//   						DefaultValue: jsii.String("defaultValue"),
//   						KeyOverride: jsii.String("keyOverride"),
//   						PropertyLocation: jsii.String("propertyLocation"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	IntegrationType: jsii.String("integrationType"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-connectiontype.html
//
type CfnConnectionTypeProps struct {
	// The name of the connection type.
	//
	// Must be prefixed with REST-.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-connectiontype.html#cfn-glue-connectiontype-connectiontype
	//
	ConnectionType *string `field:"required" json:"connectionType" yaml:"connectionType"`
	// Configuration for HTTP request and response handling.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-connectiontype.html#cfn-glue-connectiontype-restconfiguration
	//
	RestConfiguration interface{} `field:"required" json:"restConfiguration" yaml:"restConfiguration"`
	// Configuration that defines the base URL and additional request parameters needed during connection creation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-connectiontype.html#cfn-glue-connectiontype-connectionproperties
	//
	ConnectionProperties interface{} `field:"optional" json:"connectionProperties" yaml:"connectionProperties"`
	// Configuration that defines supported authentication types and required properties.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-connectiontype.html#cfn-glue-connectiontype-connectorauthenticationconfiguration
	//
	ConnectorAuthenticationConfiguration interface{} `field:"optional" json:"connectorAuthenticationConfiguration" yaml:"connectorAuthenticationConfiguration"`
	// A description of the connection type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-connectiontype.html#cfn-glue-connectiontype-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The integration type for the connection.
	//
	// Currently only REST is supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-connectiontype.html#cfn-glue-connectiontype-integrationtype
	//
	IntegrationType *string `field:"optional" json:"integrationType" yaml:"integrationType"`
	// Tags to assign to the connection type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-connectiontype.html#cfn-glue-connectiontype-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

