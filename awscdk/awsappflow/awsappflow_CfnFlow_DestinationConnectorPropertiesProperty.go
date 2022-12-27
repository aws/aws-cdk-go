package awsappflow


// This stores the information that is required to query a particular connector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   destinationConnectorPropertiesProperty := &destinationConnectorPropertiesProperty{
//   	customConnector: &customConnectorDestinationPropertiesProperty{
//   		entityName: jsii.String("entityName"),
//
//   		// the properties below are optional
//   		customProperties: map[string]*string{
//   			"customPropertiesKey": jsii.String("customProperties"),
//   		},
//   		errorHandlingConfig: &errorHandlingConfigProperty{
//   			bucketName: jsii.String("bucketName"),
//   			bucketPrefix: jsii.String("bucketPrefix"),
//   			failOnFirstError: jsii.Boolean(false),
//   		},
//   		idFieldNames: []*string{
//   			jsii.String("idFieldNames"),
//   		},
//   		writeOperationType: jsii.String("writeOperationType"),
//   	},
//   	eventBridge: &eventBridgeDestinationPropertiesProperty{
//   		object: jsii.String("object"),
//
//   		// the properties below are optional
//   		errorHandlingConfig: &errorHandlingConfigProperty{
//   			bucketName: jsii.String("bucketName"),
//   			bucketPrefix: jsii.String("bucketPrefix"),
//   			failOnFirstError: jsii.Boolean(false),
//   		},
//   	},
//   	lookoutMetrics: &lookoutMetricsDestinationPropertiesProperty{
//   		object: jsii.String("object"),
//   	},
//   	marketo: &marketoDestinationPropertiesProperty{
//   		object: jsii.String("object"),
//
//   		// the properties below are optional
//   		errorHandlingConfig: &errorHandlingConfigProperty{
//   			bucketName: jsii.String("bucketName"),
//   			bucketPrefix: jsii.String("bucketPrefix"),
//   			failOnFirstError: jsii.Boolean(false),
//   		},
//   	},
//   	redshift: &redshiftDestinationPropertiesProperty{
//   		intermediateBucketName: jsii.String("intermediateBucketName"),
//   		object: jsii.String("object"),
//
//   		// the properties below are optional
//   		bucketPrefix: jsii.String("bucketPrefix"),
//   		errorHandlingConfig: &errorHandlingConfigProperty{
//   			bucketName: jsii.String("bucketName"),
//   			bucketPrefix: jsii.String("bucketPrefix"),
//   			failOnFirstError: jsii.Boolean(false),
//   		},
//   	},
//   	s3: &s3DestinationPropertiesProperty{
//   		bucketName: jsii.String("bucketName"),
//
//   		// the properties below are optional
//   		bucketPrefix: jsii.String("bucketPrefix"),
//   		s3OutputFormatConfig: &s3OutputFormatConfigProperty{
//   			aggregationConfig: &aggregationConfigProperty{
//   				aggregationType: jsii.String("aggregationType"),
//   				targetFileSize: jsii.Number(123),
//   			},
//   			fileType: jsii.String("fileType"),
//   			prefixConfig: &prefixConfigProperty{
//   				pathPrefixHierarchy: []*string{
//   					jsii.String("pathPrefixHierarchy"),
//   				},
//   				prefixFormat: jsii.String("prefixFormat"),
//   				prefixType: jsii.String("prefixType"),
//   			},
//   			preserveSourceDataTyping: jsii.Boolean(false),
//   		},
//   	},
//   	salesforce: &salesforceDestinationPropertiesProperty{
//   		object: jsii.String("object"),
//
//   		// the properties below are optional
//   		dataTransferApi: jsii.String("dataTransferApi"),
//   		errorHandlingConfig: &errorHandlingConfigProperty{
//   			bucketName: jsii.String("bucketName"),
//   			bucketPrefix: jsii.String("bucketPrefix"),
//   			failOnFirstError: jsii.Boolean(false),
//   		},
//   		idFieldNames: []*string{
//   			jsii.String("idFieldNames"),
//   		},
//   		writeOperationType: jsii.String("writeOperationType"),
//   	},
//   	sapoData: &sAPODataDestinationPropertiesProperty{
//   		objectPath: jsii.String("objectPath"),
//
//   		// the properties below are optional
//   		errorHandlingConfig: &errorHandlingConfigProperty{
//   			bucketName: jsii.String("bucketName"),
//   			bucketPrefix: jsii.String("bucketPrefix"),
//   			failOnFirstError: jsii.Boolean(false),
//   		},
//   		idFieldNames: []*string{
//   			jsii.String("idFieldNames"),
//   		},
//   		successResponseHandlingConfig: &successResponseHandlingConfigProperty{
//   			bucketName: jsii.String("bucketName"),
//   			bucketPrefix: jsii.String("bucketPrefix"),
//   		},
//   		writeOperationType: jsii.String("writeOperationType"),
//   	},
//   	snowflake: &snowflakeDestinationPropertiesProperty{
//   		intermediateBucketName: jsii.String("intermediateBucketName"),
//   		object: jsii.String("object"),
//
//   		// the properties below are optional
//   		bucketPrefix: jsii.String("bucketPrefix"),
//   		errorHandlingConfig: &errorHandlingConfigProperty{
//   			bucketName: jsii.String("bucketName"),
//   			bucketPrefix: jsii.String("bucketPrefix"),
//   			failOnFirstError: jsii.Boolean(false),
//   		},
//   	},
//   	upsolver: &upsolverDestinationPropertiesProperty{
//   		bucketName: jsii.String("bucketName"),
//   		s3OutputFormatConfig: &upsolverS3OutputFormatConfigProperty{
//   			prefixConfig: &prefixConfigProperty{
//   				pathPrefixHierarchy: []*string{
//   					jsii.String("pathPrefixHierarchy"),
//   				},
//   				prefixFormat: jsii.String("prefixFormat"),
//   				prefixType: jsii.String("prefixType"),
//   			},
//
//   			// the properties below are optional
//   			aggregationConfig: &aggregationConfigProperty{
//   				aggregationType: jsii.String("aggregationType"),
//   				targetFileSize: jsii.Number(123),
//   			},
//   			fileType: jsii.String("fileType"),
//   		},
//
//   		// the properties below are optional
//   		bucketPrefix: jsii.String("bucketPrefix"),
//   	},
//   	zendesk: &zendeskDestinationPropertiesProperty{
//   		object: jsii.String("object"),
//
//   		// the properties below are optional
//   		errorHandlingConfig: &errorHandlingConfigProperty{
//   			bucketName: jsii.String("bucketName"),
//   			bucketPrefix: jsii.String("bucketPrefix"),
//   			failOnFirstError: jsii.Boolean(false),
//   		},
//   		idFieldNames: []*string{
//   			jsii.String("idFieldNames"),
//   		},
//   		writeOperationType: jsii.String("writeOperationType"),
//   	},
//   }
//
type CfnFlow_DestinationConnectorPropertiesProperty struct {
	// `CfnFlow.DestinationConnectorPropertiesProperty.CustomConnector`.
	CustomConnector interface{} `field:"optional" json:"customConnector" yaml:"customConnector"`
	// The properties required to query Amazon EventBridge.
	EventBridge interface{} `field:"optional" json:"eventBridge" yaml:"eventBridge"`
	// The properties required to query Amazon Lookout for Metrics.
	LookoutMetrics interface{} `field:"optional" json:"lookoutMetrics" yaml:"lookoutMetrics"`
	// The properties required to query Marketo.
	Marketo interface{} `field:"optional" json:"marketo" yaml:"marketo"`
	// The properties required to query Amazon Redshift.
	Redshift interface{} `field:"optional" json:"redshift" yaml:"redshift"`
	// The properties required to query Amazon S3.
	S3 interface{} `field:"optional" json:"s3" yaml:"s3"`
	// The properties required to query Salesforce.
	Salesforce interface{} `field:"optional" json:"salesforce" yaml:"salesforce"`
	// The properties required to query SAPOData.
	SapoData interface{} `field:"optional" json:"sapoData" yaml:"sapoData"`
	// The properties required to query Snowflake.
	Snowflake interface{} `field:"optional" json:"snowflake" yaml:"snowflake"`
	// The properties required to query Upsolver.
	Upsolver interface{} `field:"optional" json:"upsolver" yaml:"upsolver"`
	// The properties required to query Zendesk.
	Zendesk interface{} `field:"optional" json:"zendesk" yaml:"zendesk"`
}

