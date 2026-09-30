package awsglue


// Configuration for HTTP request and response handling.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   restConfigurationProperty := &RestConfigurationProperty{
//   	EntityConfigurations: map[string]interface{}{
//   		"entityConfigurationsKey": &EntityConfigurationProperty{
//   			"schema": map[string]interface{}{
//   				"schemaKey": &FieldDefinitionProperty{
//   					"fieldDataType": jsii.String("fieldDataType"),
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
//   					"name": jsii.String("name"),
//   					"responseDateFormat": jsii.String("responseDateFormat"),
//   				},
//   			},
//   			"sourceConfiguration": &SourceConfigurationProperty{
//   				"filterConfiguration": &FilterConfigurationProperty{
//   					"betweenConfiguration": &BetweenConfigurationProperty{
//   						"highBoundKey": jsii.String("highBoundKey"),
//   						"lowBoundKey": jsii.String("lowBoundKey"),
//   						"template": jsii.String("template"),
//   					},
//   					"dateTimeFormat": jsii.String("dateTimeFormat"),
//   					"filterMode": jsii.String("filterMode"),
//   					"filterStringConfiguration": &FilterStringConfigurationProperty{
//   						"queryParameterName": jsii.String("queryParameterName"),
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
//   						"limitParameter": &ExtractedParameterProperty{
//   							"defaultValue": jsii.String("defaultValue"),
//   							"key": jsii.String("key"),
//   							"propertyLocation": jsii.String("propertyLocation"),
//   							"value": &ResponseExtractionMappingProperty{
//   								"contentPath": jsii.String("contentPath"),
//   								"headerKey": jsii.String("headerKey"),
//   							},
//   						},
//   						"nextPage": &ExtractedParameterProperty{
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
//   						"allowedValues": []*string{
//   							jsii.String("allowedValues"),
//   						},
//   						"defaultValue": jsii.String("defaultValue"),
//   						"keyOverride": jsii.String("keyOverride"),
//   						"name": jsii.String("name"),
//   						"propertyLocation": jsii.String("propertyLocation"),
//   						"propertyType": jsii.String("propertyType"),
//   						"required": jsii.Boolean(false),
//   					},
//   				},
//   				"requestPath": jsii.String("requestPath"),
//   				"responseConfiguration": &ResponseConfigurationProperty{
//   					"errorPath": jsii.String("errorPath"),
//   					"resultPath": jsii.String("resultPath"),
//   				},
//   			},
//   		},
//   	},
//   	GlobalSourceConfiguration: &SourceConfigurationProperty{
//   		FilterConfiguration: &FilterConfigurationProperty{
//   			BetweenConfiguration: &BetweenConfigurationProperty{
//   				HighBoundKey: jsii.String("highBoundKey"),
//   				LowBoundKey: jsii.String("lowBoundKey"),
//   				Template: jsii.String("template"),
//   			},
//   			DateTimeFormat: jsii.String("dateTimeFormat"),
//   			FilterMode: jsii.String("filterMode"),
//   			FilterStringConfiguration: &FilterStringConfigurationProperty{
//   				QueryParameterName: jsii.String("queryParameterName"),
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
//   				LimitParameter: &ExtractedParameterProperty{
//   					DefaultValue: jsii.String("defaultValue"),
//   					Key: jsii.String("key"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   					Value: &ResponseExtractionMappingProperty{
//   						ContentPath: jsii.String("contentPath"),
//   						HeaderKey: jsii.String("headerKey"),
//   					},
//   				},
//   				NextPage: &ExtractedParameterProperty{
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
//   		RequestPath: jsii.String("requestPath"),
//   		ResponseConfiguration: &ResponseConfigurationProperty{
//   			ErrorPath: jsii.String("errorPath"),
//   			ResultPath: jsii.String("resultPath"),
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
type CfnConnectionTypePropsMixin_RestConfigurationProperty struct {
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

