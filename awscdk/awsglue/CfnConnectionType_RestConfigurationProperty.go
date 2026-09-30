package awsglue


// Configuration for HTTP request and response handling.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   restConfigurationProperty := &RestConfigurationProperty{
//   	EntityConfigurations: map[string]interface{}{
//   		"entityConfigurationsKey": &EntityConfigurationProperty{
//   			"schema": map[string]interface{}{
//   				"schemaKey": &FieldDefinitionProperty{
//   					"fieldDataType": jsii.String("fieldDataType"),
//   					"name": jsii.String("name"),
//
//   					// the properties below are optional
//   					"filterOverrides": &FilterOverridesProperty{
//   						"betweenConfiguration": &BetweenConfigurationProperty{
//   							"highBoundKey": jsii.String("highBoundKey"),
//   							"lowBoundKey": jsii.String("lowBoundKey"),
//   							"template": jsii.String("template"),
//   						},
//   						"dateTimeFormat": jsii.String("dateTimeFormat"),
//   						"fieldName": jsii.String("fieldName"),
//   						"operatorMappings": map[string]*string{
//   							"operatorMappingsKey": jsii.String("operatorMappings"),
//   						},
//   					},
//   					"isNullable": jsii.Boolean(false),
//   					"isOrderable": jsii.Boolean(false),
//   					"isPartitionable": jsii.Boolean(false),
//   					"isQueryable": jsii.Boolean(false),
//   					"responseDateFormat": jsii.String("responseDateFormat"),
//   				},
//   			},
//   			"sourceConfiguration": &SourceConfigurationProperty{
//   				"filterConfiguration": &FilterConfigurationProperty{
//   					"filterMode": jsii.String("filterMode"),
//
//   					// the properties below are optional
//   					"betweenConfiguration": &BetweenConfigurationProperty{
//   						"highBoundKey": jsii.String("highBoundKey"),
//   						"lowBoundKey": jsii.String("lowBoundKey"),
//   						"template": jsii.String("template"),
//   					},
//   					"dateTimeFormat": jsii.String("dateTimeFormat"),
//   					"filterStringConfiguration": &FilterStringConfigurationProperty{
//   						"queryParameterName": jsii.String("queryParameterName"),
//
//   						// the properties below are optional
//   						"quoteCharacter": jsii.String("quoteCharacter"),
//   						"quoteStringValues": jsii.Boolean(false),
//   					},
//   					"operatorMappings": map[string]*string{
//   						"operatorMappingsKey": jsii.String("operatorMappings"),
//   					},
//   					"stripQuotes": jsii.Boolean(false),
//   				},
//   				"paginationConfiguration": &PaginationConfigurationProperty{
//   					"cursorConfiguration": &CursorConfigurationProperty{
//   						"nextPage": &ExtractedParameterProperty{
//   							"defaultValue": jsii.String("defaultValue"),
//   							"key": jsii.String("key"),
//   							"propertyLocation": jsii.String("propertyLocation"),
//   							"value": &ResponseExtractionMappingProperty{
//   								"contentPath": jsii.String("contentPath"),
//   								"headerKey": jsii.String("headerKey"),
//   							},
//   						},
//
//   						// the properties below are optional
//   						"limitParameter": &ExtractedParameterProperty{
//   							"defaultValue": jsii.String("defaultValue"),
//   							"key": jsii.String("key"),
//   							"propertyLocation": jsii.String("propertyLocation"),
//   							"value": &ResponseExtractionMappingProperty{
//   								"contentPath": jsii.String("contentPath"),
//   								"headerKey": jsii.String("headerKey"),
//   							},
//   						},
//   					},
//   					"offsetConfiguration": &OffsetConfigurationProperty{
//   						"limitParameter": &ExtractedParameterProperty{
//   							"defaultValue": jsii.String("defaultValue"),
//   							"key": jsii.String("key"),
//   							"propertyLocation": jsii.String("propertyLocation"),
//   							"value": &ResponseExtractionMappingProperty{
//   								"contentPath": jsii.String("contentPath"),
//   								"headerKey": jsii.String("headerKey"),
//   							},
//   						},
//   						"offsetParameter": &ExtractedParameterProperty{
//   							"defaultValue": jsii.String("defaultValue"),
//   							"key": jsii.String("key"),
//   							"propertyLocation": jsii.String("propertyLocation"),
//   							"value": &ResponseExtractionMappingProperty{
//   								"contentPath": jsii.String("contentPath"),
//   								"headerKey": jsii.String("headerKey"),
//   							},
//   						},
//   					},
//   				},
//   				"requestMethod": jsii.String("requestMethod"),
//   				"requestParameters": []interface{}{
//   					&ConnectorPropertyProperty{
//   						"name": jsii.String("name"),
//   						"propertyType": jsii.String("propertyType"),
//   						"required": jsii.Boolean(false),
//
//   						// the properties below are optional
//   						"allowedValues": []*string{
//   							jsii.String("allowedValues"),
//   						},
//   						"defaultValue": jsii.String("defaultValue"),
//   						"keyOverride": jsii.String("keyOverride"),
//   						"propertyLocation": jsii.String("propertyLocation"),
//   					},
//   				},
//   				"requestPath": jsii.String("requestPath"),
//   				"responseConfiguration": &ResponseConfigurationProperty{
//   					"resultPath": jsii.String("resultPath"),
//
//   					// the properties below are optional
//   					"errorPath": jsii.String("errorPath"),
//   				},
//   			},
//   		},
//   	},
//   	GlobalSourceConfiguration: &SourceConfigurationProperty{
//   		FilterConfiguration: &FilterConfigurationProperty{
//   			FilterMode: jsii.String("filterMode"),
//
//   			// the properties below are optional
//   			BetweenConfiguration: &BetweenConfigurationProperty{
//   				HighBoundKey: jsii.String("highBoundKey"),
//   				LowBoundKey: jsii.String("lowBoundKey"),
//   				Template: jsii.String("template"),
//   			},
//   			DateTimeFormat: jsii.String("dateTimeFormat"),
//   			FilterStringConfiguration: &FilterStringConfigurationProperty{
//   				QueryParameterName: jsii.String("queryParameterName"),
//
//   				// the properties below are optional
//   				QuoteCharacter: jsii.String("quoteCharacter"),
//   				QuoteStringValues: jsii.Boolean(false),
//   			},
//   			OperatorMappings: map[string]*string{
//   				"operatorMappingsKey": jsii.String("operatorMappings"),
//   			},
//   			StripQuotes: jsii.Boolean(false),
//   		},
//   		PaginationConfiguration: &PaginationConfigurationProperty{
//   			CursorConfiguration: &CursorConfigurationProperty{
//   				NextPage: &ExtractedParameterProperty{
//   					DefaultValue: jsii.String("defaultValue"),
//   					Key: jsii.String("key"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   					Value: &ResponseExtractionMappingProperty{
//   						ContentPath: jsii.String("contentPath"),
//   						HeaderKey: jsii.String("headerKey"),
//   					},
//   				},
//
//   				// the properties below are optional
//   				LimitParameter: &ExtractedParameterProperty{
//   					DefaultValue: jsii.String("defaultValue"),
//   					Key: jsii.String("key"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   					Value: &ResponseExtractionMappingProperty{
//   						ContentPath: jsii.String("contentPath"),
//   						HeaderKey: jsii.String("headerKey"),
//   					},
//   				},
//   			},
//   			OffsetConfiguration: &OffsetConfigurationProperty{
//   				LimitParameter: &ExtractedParameterProperty{
//   					DefaultValue: jsii.String("defaultValue"),
//   					Key: jsii.String("key"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   					Value: &ResponseExtractionMappingProperty{
//   						ContentPath: jsii.String("contentPath"),
//   						HeaderKey: jsii.String("headerKey"),
//   					},
//   				},
//   				OffsetParameter: &ExtractedParameterProperty{
//   					DefaultValue: jsii.String("defaultValue"),
//   					Key: jsii.String("key"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   					Value: &ResponseExtractionMappingProperty{
//   						ContentPath: jsii.String("contentPath"),
//   						HeaderKey: jsii.String("headerKey"),
//   					},
//   				},
//   			},
//   		},
//   		RequestMethod: jsii.String("requestMethod"),
//   		RequestParameters: []interface{}{
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
//   		RequestPath: jsii.String("requestPath"),
//   		ResponseConfiguration: &ResponseConfigurationProperty{
//   			ResultPath: jsii.String("resultPath"),
//
//   			// the properties below are optional
//   			ErrorPath: jsii.String("errorPath"),
//   		},
//   	},
//   	ValidationEndpointConfiguration: &ValidationEndpointConfigurationProperty{
//   		RequestMethod: jsii.String("requestMethod"),
//   		RequestPath: jsii.String("requestPath"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-restconfiguration.html
//
type CfnConnectionType_RestConfigurationProperty struct {
	// A map of entity configurations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-restconfiguration.html#cfn-glue-connectiontype-restconfiguration-entityconfigurations
	//
	EntityConfigurations interface{} `field:"optional" json:"entityConfigurations" yaml:"entityConfigurations"`
	// Configuration that defines how to make requests to endpoints.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-restconfiguration.html#cfn-glue-connectiontype-restconfiguration-globalsourceconfiguration
	//
	GlobalSourceConfiguration interface{} `field:"optional" json:"globalSourceConfiguration" yaml:"globalSourceConfiguration"`
	// Configuration for the validation endpoint.
	//
	// Only supports RequestMethod and RequestPath.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-restconfiguration.html#cfn-glue-connectiontype-restconfiguration-validationendpointconfiguration
	//
	ValidationEndpointConfiguration interface{} `field:"optional" json:"validationEndpointConfiguration" yaml:"validationEndpointConfiguration"`
}

