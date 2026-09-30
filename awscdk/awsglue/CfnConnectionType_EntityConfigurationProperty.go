package awsglue


// Configuration for interacting with a specific data entity.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   entityConfigurationProperty := &EntityConfigurationProperty{
//   	Schema: map[string]interface{}{
//   		"schemaKey": &FieldDefinitionProperty{
//   			"fieldDataType": jsii.String("fieldDataType"),
//   			"name": jsii.String("name"),
//
//   			// the properties below are optional
//   			"filterOverrides": &FilterOverridesProperty{
//   				"betweenConfiguration": &BetweenConfigurationProperty{
//   					"highBoundKey": jsii.String("highBoundKey"),
//   					"lowBoundKey": jsii.String("lowBoundKey"),
//   					"template": jsii.String("template"),
//   				},
//   				"dateTimeFormat": jsii.String("dateTimeFormat"),
//   				"fieldName": jsii.String("fieldName"),
//   				"operatorMappings": map[string]*string{
//   					"operatorMappingsKey": jsii.String("operatorMappings"),
//   				},
//   			},
//   			"isNullable": jsii.Boolean(false),
//   			"isOrderable": jsii.Boolean(false),
//   			"isPartitionable": jsii.Boolean(false),
//   			"isQueryable": jsii.Boolean(false),
//   			"responseDateFormat": jsii.String("responseDateFormat"),
//   		},
//   	},
//   	SourceConfiguration: &SourceConfigurationProperty{
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-entityconfiguration.html
//
type CfnConnectionType_EntityConfigurationProperty struct {
	// The schema definition for this entity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-entityconfiguration.html#cfn-glue-connectiontype-entityconfiguration-schema
	//
	Schema interface{} `field:"optional" json:"schema" yaml:"schema"`
	// Configuration that defines how to make requests to endpoints.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-entityconfiguration.html#cfn-glue-connectiontype-entityconfiguration-sourceconfiguration
	//
	SourceConfiguration interface{} `field:"optional" json:"sourceConfiguration" yaml:"sourceConfiguration"`
}

