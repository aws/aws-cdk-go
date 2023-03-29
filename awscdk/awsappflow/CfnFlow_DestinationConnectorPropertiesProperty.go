package awsappflow


// This stores the information that is required to query a particular connector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   destinationConnectorPropertiesProperty := &DestinationConnectorPropertiesProperty{
//   	CustomConnector: &CustomConnectorDestinationPropertiesProperty{
//   		EntityName: jsii.String("entityName"),
//
//   		// the properties below are optional
//   		CustomProperties: map[string]*string{
//   			"customPropertiesKey": jsii.String("customProperties"),
//   		},
//   		ErrorHandlingConfig: &ErrorHandlingConfigProperty{
//   			BucketName: jsii.String("bucketName"),
//   			BucketPrefix: jsii.String("bucketPrefix"),
//   			FailOnFirstError: jsii.Boolean(false),
//   		},
//   		IdFieldNames: []*string{
//   			jsii.String("idFieldNames"),
//   		},
//   		WriteOperationType: jsii.String("writeOperationType"),
//   	},
//   	EventBridge: &EventBridgeDestinationPropertiesProperty{
//   		Object: jsii.String("object"),
//
//   		// the properties below are optional
//   		ErrorHandlingConfig: &ErrorHandlingConfigProperty{
//   			BucketName: jsii.String("bucketName"),
//   			BucketPrefix: jsii.String("bucketPrefix"),
//   			FailOnFirstError: jsii.Boolean(false),
//   		},
//   	},
//   	LookoutMetrics: &LookoutMetricsDestinationPropertiesProperty{
//   		Object: jsii.String("object"),
//   	},
//   	Marketo: &MarketoDestinationPropertiesProperty{
//   		Object: jsii.String("object"),
//
//   		// the properties below are optional
//   		ErrorHandlingConfig: &ErrorHandlingConfigProperty{
//   			BucketName: jsii.String("bucketName"),
//   			BucketPrefix: jsii.String("bucketPrefix"),
//   			FailOnFirstError: jsii.Boolean(false),
//   		},
//   	},
//   	Redshift: &RedshiftDestinationPropertiesProperty{
//   		IntermediateBucketName: jsii.String("intermediateBucketName"),
//   		Object: jsii.String("object"),
//
//   		// the properties below are optional
//   		BucketPrefix: jsii.String("bucketPrefix"),
//   		ErrorHandlingConfig: &ErrorHandlingConfigProperty{
//   			BucketName: jsii.String("bucketName"),
//   			BucketPrefix: jsii.String("bucketPrefix"),
//   			FailOnFirstError: jsii.Boolean(false),
//   		},
//   	},
//   	S3: &S3DestinationPropertiesProperty{
//   		BucketName: jsii.String("bucketName"),
//
//   		// the properties below are optional
//   		BucketPrefix: jsii.String("bucketPrefix"),
//   		S3OutputFormatConfig: &S3OutputFormatConfigProperty{
//   			AggregationConfig: &AggregationConfigProperty{
//   				AggregationType: jsii.String("aggregationType"),
//   				TargetFileSize: jsii.Number(123),
//   			},
//   			FileType: jsii.String("fileType"),
//   			PrefixConfig: &PrefixConfigProperty{
//   				PathPrefixHierarchy: []*string{
//   					jsii.String("pathPrefixHierarchy"),
//   				},
//   				PrefixFormat: jsii.String("prefixFormat"),
//   				PrefixType: jsii.String("prefixType"),
//   			},
//   			PreserveSourceDataTyping: jsii.Boolean(false),
//   		},
//   	},
//   	Salesforce: &SalesforceDestinationPropertiesProperty{
//   		Object: jsii.String("object"),
//
//   		// the properties below are optional
//   		DataTransferApi: jsii.String("dataTransferApi"),
//   		ErrorHandlingConfig: &ErrorHandlingConfigProperty{
//   			BucketName: jsii.String("bucketName"),
//   			BucketPrefix: jsii.String("bucketPrefix"),
//   			FailOnFirstError: jsii.Boolean(false),
//   		},
//   		IdFieldNames: []*string{
//   			jsii.String("idFieldNames"),
//   		},
//   		WriteOperationType: jsii.String("writeOperationType"),
//   	},
//   	SapoData: &SAPODataDestinationPropertiesProperty{
//   		ObjectPath: jsii.String("objectPath"),
//
//   		// the properties below are optional
//   		ErrorHandlingConfig: &ErrorHandlingConfigProperty{
//   			BucketName: jsii.String("bucketName"),
//   			BucketPrefix: jsii.String("bucketPrefix"),
//   			FailOnFirstError: jsii.Boolean(false),
//   		},
//   		IdFieldNames: []*string{
//   			jsii.String("idFieldNames"),
//   		},
//   		SuccessResponseHandlingConfig: &SuccessResponseHandlingConfigProperty{
//   			BucketName: jsii.String("bucketName"),
//   			BucketPrefix: jsii.String("bucketPrefix"),
//   		},
//   		WriteOperationType: jsii.String("writeOperationType"),
//   	},
//   	Snowflake: &SnowflakeDestinationPropertiesProperty{
//   		IntermediateBucketName: jsii.String("intermediateBucketName"),
//   		Object: jsii.String("object"),
//
//   		// the properties below are optional
//   		BucketPrefix: jsii.String("bucketPrefix"),
//   		ErrorHandlingConfig: &ErrorHandlingConfigProperty{
//   			BucketName: jsii.String("bucketName"),
//   			BucketPrefix: jsii.String("bucketPrefix"),
//   			FailOnFirstError: jsii.Boolean(false),
//   		},
//   	},
//   	Upsolver: &UpsolverDestinationPropertiesProperty{
//   		BucketName: jsii.String("bucketName"),
//   		S3OutputFormatConfig: &UpsolverS3OutputFormatConfigProperty{
//   			PrefixConfig: &PrefixConfigProperty{
//   				PathPrefixHierarchy: []*string{
//   					jsii.String("pathPrefixHierarchy"),
//   				},
//   				PrefixFormat: jsii.String("prefixFormat"),
//   				PrefixType: jsii.String("prefixType"),
//   			},
//
//   			// the properties below are optional
//   			AggregationConfig: &AggregationConfigProperty{
//   				AggregationType: jsii.String("aggregationType"),
//   				TargetFileSize: jsii.Number(123),
//   			},
//   			FileType: jsii.String("fileType"),
//   		},
//
//   		// the properties below are optional
//   		BucketPrefix: jsii.String("bucketPrefix"),
//   	},
//   	Zendesk: &ZendeskDestinationPropertiesProperty{
//   		Object: jsii.String("object"),
//
//   		// the properties below are optional
//   		ErrorHandlingConfig: &ErrorHandlingConfigProperty{
//   			BucketName: jsii.String("bucketName"),
//   			BucketPrefix: jsii.String("bucketPrefix"),
//   			FailOnFirstError: jsii.Boolean(false),
//   		},
//   		IdFieldNames: []*string{
//   			jsii.String("idFieldNames"),
//   		},
//   		WriteOperationType: jsii.String("writeOperationType"),
//   	},
//   }
//
type CfnFlow_DestinationConnectorPropertiesProperty struct {
	// The properties that are required to query the custom Connector.
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

