package awsglue


// Configuration that defines how to make requests to endpoints.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceConfigurationProperty := &SourceConfigurationProperty{
//   	FilterConfiguration: &FilterConfigurationProperty{
//   		FilterMode: jsii.String("filterMode"),
//
//   		// the properties below are optional
//   		BetweenConfiguration: &BetweenConfigurationProperty{
//   			HighBoundKey: jsii.String("highBoundKey"),
//   			LowBoundKey: jsii.String("lowBoundKey"),
//   			Template: jsii.String("template"),
//   		},
//   		DateTimeFormat: jsii.String("dateTimeFormat"),
//   		FilterStringConfiguration: &FilterStringConfigurationProperty{
//   			QueryParameterName: jsii.String("queryParameterName"),
//
//   			// the properties below are optional
//   			QuoteCharacter: jsii.String("quoteCharacter"),
//   			QuoteStringValues: jsii.Boolean(false),
//   		},
//   		OperatorMappings: map[string]*string{
//   			"operatorMappingsKey": jsii.String("operatorMappings"),
//   		},
//   		StripQuotes: jsii.Boolean(false),
//   	},
//   	PaginationConfiguration: &PaginationConfigurationProperty{
//   		CursorConfiguration: &CursorConfigurationProperty{
//   			NextPage: &ExtractedParameterProperty{
//   				DefaultValue: jsii.String("defaultValue"),
//   				Key: jsii.String("key"),
//   				PropertyLocation: jsii.String("propertyLocation"),
//   				Value: &ResponseExtractionMappingProperty{
//   					ContentPath: jsii.String("contentPath"),
//   					HeaderKey: jsii.String("headerKey"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			LimitParameter: &ExtractedParameterProperty{
//   				DefaultValue: jsii.String("defaultValue"),
//   				Key: jsii.String("key"),
//   				PropertyLocation: jsii.String("propertyLocation"),
//   				Value: &ResponseExtractionMappingProperty{
//   					ContentPath: jsii.String("contentPath"),
//   					HeaderKey: jsii.String("headerKey"),
//   				},
//   			},
//   		},
//   		OffsetConfiguration: &OffsetConfigurationProperty{
//   			LimitParameter: &ExtractedParameterProperty{
//   				DefaultValue: jsii.String("defaultValue"),
//   				Key: jsii.String("key"),
//   				PropertyLocation: jsii.String("propertyLocation"),
//   				Value: &ResponseExtractionMappingProperty{
//   					ContentPath: jsii.String("contentPath"),
//   					HeaderKey: jsii.String("headerKey"),
//   				},
//   			},
//   			OffsetParameter: &ExtractedParameterProperty{
//   				DefaultValue: jsii.String("defaultValue"),
//   				Key: jsii.String("key"),
//   				PropertyLocation: jsii.String("propertyLocation"),
//   				Value: &ResponseExtractionMappingProperty{
//   					ContentPath: jsii.String("contentPath"),
//   					HeaderKey: jsii.String("headerKey"),
//   				},
//   			},
//   		},
//   	},
//   	RequestMethod: jsii.String("requestMethod"),
//   	RequestParameters: []interface{}{
//   		&ConnectorPropertyProperty{
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
//   	RequestPath: jsii.String("requestPath"),
//   	ResponseConfiguration: &ResponseConfigurationProperty{
//   		ResultPath: jsii.String("resultPath"),
//
//   		// the properties below are optional
//   		ErrorPath: jsii.String("errorPath"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-sourceconfiguration.html
//
type CfnConnectionType_SourceConfigurationProperty struct {
	// Configuration that defines how filter predicates are applied to REST API requests, supporting both query parameter and filter string strategies.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-sourceconfiguration.html#cfn-glue-connectiontype-sourceconfiguration-filterconfiguration
	//
	FilterConfiguration interface{} `field:"optional" json:"filterConfiguration" yaml:"filterConfiguration"`
	// Configuration for handling paginated responses.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-sourceconfiguration.html#cfn-glue-connectiontype-sourceconfiguration-paginationconfiguration
	//
	PaginationConfiguration interface{} `field:"optional" json:"paginationConfiguration" yaml:"paginationConfiguration"`
	// The HTTP method to use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-sourceconfiguration.html#cfn-glue-connectiontype-sourceconfiguration-requestmethod
	//
	RequestMethod *string `field:"optional" json:"requestMethod" yaml:"requestMethod"`
	// Request parameters configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-sourceconfiguration.html#cfn-glue-connectiontype-sourceconfiguration-requestparameters
	//
	RequestParameters interface{} `field:"optional" json:"requestParameters" yaml:"requestParameters"`
	// The URL path for the REST endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-sourceconfiguration.html#cfn-glue-connectiontype-sourceconfiguration-requestpath
	//
	RequestPath *string `field:"optional" json:"requestPath" yaml:"requestPath"`
	// Configuration for parsing JSON responses from REST API calls.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-sourceconfiguration.html#cfn-glue-connectiontype-sourceconfiguration-responseconfiguration
	//
	ResponseConfiguration interface{} `field:"optional" json:"responseConfiguration" yaml:"responseConfiguration"`
}

