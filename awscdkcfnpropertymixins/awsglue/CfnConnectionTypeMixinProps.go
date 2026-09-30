package awsglue

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnConnectionTypePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnConnectionTypeMixinProps := &CfnConnectionTypeMixinProps{
//   	ConnectionProperties: &ConnectionPropertiesConfigurationProperty{
//   		AdditionalRequestParameters: []interface{}{
//   			&ConnectorPropertyProperty{
//   				AllowedValues: []*string{
//   					jsii.String("allowedValues"),
//   				},
//   				DefaultValue: jsii.String("defaultValue"),
//   				KeyOverride: jsii.String("keyOverride"),
//   				Name: jsii.String("name"),
//   				PropertyLocation: jsii.String("propertyLocation"),
//   				PropertyType: jsii.String("propertyType"),
//   				Required: jsii.Boolean(false),
//   			},
//   		},
//   		Url: &ConnectorPropertyProperty{
//   			AllowedValues: []*string{
//   				jsii.String("allowedValues"),
//   			},
//   			DefaultValue: jsii.String("defaultValue"),
//   			KeyOverride: jsii.String("keyOverride"),
//   			Name: jsii.String("name"),
//   			PropertyLocation: jsii.String("propertyLocation"),
//   			PropertyType: jsii.String("propertyType"),
//   			Required: jsii.Boolean(false),
//   		},
//   	},
//   	ConnectionType: jsii.String("connectionType"),
//   	ConnectorAuthenticationConfiguration: &ConnectorAuthenticationConfigurationProperty{
//   		AuthenticationTypes: []*string{
//   			jsii.String("authenticationTypes"),
//   		},
//   		BasicAuthenticationProperties: &BasicAuthenticationPropertiesProperty{
//   			Password: &SecretConnectorPropertyProperty{
//   				KeyOverride: jsii.String("keyOverride"),
//   				Name: jsii.String("name"),
//   				PropertyLocation: jsii.String("propertyLocation"),
//   				PropertyType: jsii.String("propertyType"),
//   				Required: jsii.Boolean(false),
//   			},
//   			Username: &SecretConnectorPropertyProperty{
//   				KeyOverride: jsii.String("keyOverride"),
//   				Name: jsii.String("name"),
//   				PropertyLocation: jsii.String("propertyLocation"),
//   				PropertyType: jsii.String("propertyType"),
//   				Required: jsii.Boolean(false),
//   			},
//   		},
//   		CustomAuthenticationProperties: &CustomAuthenticationPropertiesProperty{
//   			AuthenticationParameters: []interface{}{
//   				&SecretConnectorPropertyProperty{
//   					KeyOverride: jsii.String("keyOverride"),
//   					Name: jsii.String("name"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//   				},
//   			},
//   		},
//   		OAuth2Properties: &ConnectorOAuth2PropertiesProperty{
//   			AuthorizationCodeProperties: &ConnectorAuthorizationCodePropertiesProperty{
//   				AuthorizationCode: &ConnectorPropertyProperty{
//   					AllowedValues: []*string{
//   						jsii.String("allowedValues"),
//   					},
//   					DefaultValue: jsii.String("defaultValue"),
//   					KeyOverride: jsii.String("keyOverride"),
//   					Name: jsii.String("name"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//   				},
//   				AuthorizationCodeUrl: &ConnectorPropertyProperty{
//   					AllowedValues: []*string{
//   						jsii.String("allowedValues"),
//   					},
//   					DefaultValue: jsii.String("defaultValue"),
//   					KeyOverride: jsii.String("keyOverride"),
//   					Name: jsii.String("name"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//   				},
//   				ClientId: &SecretConnectorPropertyProperty{
//   					KeyOverride: jsii.String("keyOverride"),
//   					Name: jsii.String("name"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//   				},
//   				ClientSecret: &SecretConnectorPropertyProperty{
//   					KeyOverride: jsii.String("keyOverride"),
//   					Name: jsii.String("name"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//   				},
//   				ContentType: jsii.String("contentType"),
//   				Prompt: &ConnectorPropertyProperty{
//   					AllowedValues: []*string{
//   						jsii.String("allowedValues"),
//   					},
//   					DefaultValue: jsii.String("defaultValue"),
//   					KeyOverride: jsii.String("keyOverride"),
//   					Name: jsii.String("name"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//   				},
//   				RedirectUri: &ConnectorPropertyProperty{
//   					AllowedValues: []*string{
//   						jsii.String("allowedValues"),
//   					},
//   					DefaultValue: jsii.String("defaultValue"),
//   					KeyOverride: jsii.String("keyOverride"),
//   					Name: jsii.String("name"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//   				},
//   				RequestMethod: jsii.String("requestMethod"),
//   				Scope: &ConnectorPropertyProperty{
//   					AllowedValues: []*string{
//   						jsii.String("allowedValues"),
//   					},
//   					DefaultValue: jsii.String("defaultValue"),
//   					KeyOverride: jsii.String("keyOverride"),
//   					Name: jsii.String("name"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//   				},
//   				TokenUrl: &ConnectorPropertyProperty{
//   					AllowedValues: []*string{
//   						jsii.String("allowedValues"),
//   					},
//   					DefaultValue: jsii.String("defaultValue"),
//   					KeyOverride: jsii.String("keyOverride"),
//   					Name: jsii.String("name"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//   				},
//   				TokenUrlParameters: []interface{}{
//   					&ConnectorPropertyProperty{
//   						AllowedValues: []*string{
//   							jsii.String("allowedValues"),
//   						},
//   						DefaultValue: jsii.String("defaultValue"),
//   						KeyOverride: jsii.String("keyOverride"),
//   						Name: jsii.String("name"),
//   						PropertyLocation: jsii.String("propertyLocation"),
//   						PropertyType: jsii.String("propertyType"),
//   						Required: jsii.Boolean(false),
//   					},
//   				},
//   			},
//   			ClientCredentialsProperties: &ClientCredentialsPropertiesProperty{
//   				ClientId: &SecretConnectorPropertyProperty{
//   					KeyOverride: jsii.String("keyOverride"),
//   					Name: jsii.String("name"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//   				},
//   				ClientSecret: &SecretConnectorPropertyProperty{
//   					KeyOverride: jsii.String("keyOverride"),
//   					Name: jsii.String("name"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//   				},
//   				ContentType: jsii.String("contentType"),
//   				RequestMethod: jsii.String("requestMethod"),
//   				Scope: &ConnectorPropertyProperty{
//   					AllowedValues: []*string{
//   						jsii.String("allowedValues"),
//   					},
//   					DefaultValue: jsii.String("defaultValue"),
//   					KeyOverride: jsii.String("keyOverride"),
//   					Name: jsii.String("name"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//   				},
//   				TokenUrl: &ConnectorPropertyProperty{
//   					AllowedValues: []*string{
//   						jsii.String("allowedValues"),
//   					},
//   					DefaultValue: jsii.String("defaultValue"),
//   					KeyOverride: jsii.String("keyOverride"),
//   					Name: jsii.String("name"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//   				},
//   				TokenUrlParameters: []interface{}{
//   					&ConnectorPropertyProperty{
//   						AllowedValues: []*string{
//   							jsii.String("allowedValues"),
//   						},
//   						DefaultValue: jsii.String("defaultValue"),
//   						KeyOverride: jsii.String("keyOverride"),
//   						Name: jsii.String("name"),
//   						PropertyLocation: jsii.String("propertyLocation"),
//   						PropertyType: jsii.String("propertyType"),
//   						Required: jsii.Boolean(false),
//   					},
//   				},
//   			},
//   			JwtBearerProperties: &JWTBearerPropertiesProperty{
//   				ContentType: jsii.String("contentType"),
//   				JwtToken: &SecretConnectorPropertyProperty{
//   					KeyOverride: jsii.String("keyOverride"),
//   					Name: jsii.String("name"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//   				},
//   				RequestMethod: jsii.String("requestMethod"),
//   				TokenUrl: &ConnectorPropertyProperty{
//   					AllowedValues: []*string{
//   						jsii.String("allowedValues"),
//   					},
//   					DefaultValue: jsii.String("defaultValue"),
//   					KeyOverride: jsii.String("keyOverride"),
//   					Name: jsii.String("name"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//   				},
//   				TokenUrlParameters: []interface{}{
//   					&ConnectorPropertyProperty{
//   						AllowedValues: []*string{
//   							jsii.String("allowedValues"),
//   						},
//   						DefaultValue: jsii.String("defaultValue"),
//   						KeyOverride: jsii.String("keyOverride"),
//   						Name: jsii.String("name"),
//   						PropertyLocation: jsii.String("propertyLocation"),
//   						PropertyType: jsii.String("propertyType"),
//   						Required: jsii.Boolean(false),
//   					},
//   				},
//   			},
//   			OAuth2GrantType: jsii.String("oAuth2GrantType"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	IntegrationType: jsii.String("integrationType"),
//   	RestConfiguration: &RestConfigurationProperty{
//   		EntityConfigurations: map[string]interface{}{
//   			"entityConfigurationsKey": &EntityConfigurationProperty{
//   				"schema": map[string]interface{}{
//   					"schemaKey": &FieldDefinitionProperty{
//   						"fieldDataType": jsii.String("fieldDataType"),
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
//   						"name": jsii.String("name"),
//   						"responseDateFormat": jsii.String("responseDateFormat"),
//   					},
//   				},
//   				"sourceConfiguration": &SourceConfigurationProperty{
//   					"filterConfiguration": &FilterConfigurationProperty{
//   						"betweenConfiguration": &BetweenConfigurationProperty{
//   							"highBoundKey": jsii.String("highBoundKey"),
//   							"lowBoundKey": jsii.String("lowBoundKey"),
//   							"template": jsii.String("template"),
//   						},
//   						"dateTimeFormat": jsii.String("dateTimeFormat"),
//   						"filterMode": jsii.String("filterMode"),
//   						"filterStringConfiguration": &FilterStringConfigurationProperty{
//   							"queryParameterName": jsii.String("queryParameterName"),
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
//   							"limitParameter": &ExtractedParameterProperty{
//   								"defaultValue": jsii.String("defaultValue"),
//   								"key": jsii.String("key"),
//   								"propertyLocation": jsii.String("propertyLocation"),
//   								"value": &ResponseExtractionMappingProperty{
//   									"contentPath": jsii.String("contentPath"),
//   									"headerKey": jsii.String("headerKey"),
//   								},
//   							},
//   							"nextPage": &ExtractedParameterProperty{
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
//   							"allowedValues": []*string{
//   								jsii.String("allowedValues"),
//   							},
//   							"defaultValue": jsii.String("defaultValue"),
//   							"keyOverride": jsii.String("keyOverride"),
//   							"name": jsii.String("name"),
//   							"propertyLocation": jsii.String("propertyLocation"),
//   							"propertyType": jsii.String("propertyType"),
//   							"required": jsii.Boolean(false),
//   						},
//   					},
//   					"requestPath": jsii.String("requestPath"),
//   					"responseConfiguration": &ResponseConfigurationProperty{
//   						"errorPath": jsii.String("errorPath"),
//   						"resultPath": jsii.String("resultPath"),
//   					},
//   				},
//   			},
//   		},
//   		GlobalSourceConfiguration: &SourceConfigurationProperty{
//   			FilterConfiguration: &FilterConfigurationProperty{
//   				BetweenConfiguration: &BetweenConfigurationProperty{
//   					HighBoundKey: jsii.String("highBoundKey"),
//   					LowBoundKey: jsii.String("lowBoundKey"),
//   					Template: jsii.String("template"),
//   				},
//   				DateTimeFormat: jsii.String("dateTimeFormat"),
//   				FilterMode: jsii.String("filterMode"),
//   				FilterStringConfiguration: &FilterStringConfigurationProperty{
//   					QueryParameterName: jsii.String("queryParameterName"),
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
//   					LimitParameter: &ExtractedParameterProperty{
//   						DefaultValue: jsii.String("defaultValue"),
//   						Key: jsii.String("key"),
//   						PropertyLocation: jsii.String("propertyLocation"),
//   						Value: &ResponseExtractionMappingProperty{
//   							ContentPath: jsii.String("contentPath"),
//   							HeaderKey: jsii.String("headerKey"),
//   						},
//   					},
//   					NextPage: &ExtractedParameterProperty{
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
//   					AllowedValues: []*string{
//   						jsii.String("allowedValues"),
//   					},
//   					DefaultValue: jsii.String("defaultValue"),
//   					KeyOverride: jsii.String("keyOverride"),
//   					Name: jsii.String("name"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//   				},
//   			},
//   			RequestPath: jsii.String("requestPath"),
//   			ResponseConfiguration: &ResponseConfigurationProperty{
//   				ErrorPath: jsii.String("errorPath"),
//   				ResultPath: jsii.String("resultPath"),
//   			},
//   		},
//   		ValidationEndpointConfiguration: &ValidationEndpointConfigurationProperty{
//   			RequestMethod: jsii.String("requestMethod"),
//   			RequestPath: jsii.String("requestPath"),
//   		},
//   	},
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
type CfnConnectionTypeMixinProps struct {
	// Configuration that defines the base URL and additional request parameters needed during connection creation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-connectiontype.html#cfn-glue-connectiontype-connectionproperties
	//
	ConnectionProperties interface{} `field:"optional" json:"connectionProperties" yaml:"connectionProperties"`
	// The name of the connection type.
	//
	// Must be prefixed with REST-.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-connectiontype.html#cfn-glue-connectiontype-connectiontype
	//
	ConnectionType *string `field:"optional" json:"connectionType" yaml:"connectionType"`
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
	// Configuration for HTTP request and response handling.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-connectiontype.html#cfn-glue-connectiontype-restconfiguration
	//
	RestConfiguration interface{} `field:"optional" json:"restConfiguration" yaml:"restConfiguration"`
	// Tags to assign to the connection type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-connectiontype.html#cfn-glue-connectiontype-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

