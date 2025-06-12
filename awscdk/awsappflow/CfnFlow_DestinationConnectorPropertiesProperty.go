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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-destinationconnectorproperties.html
//
type CfnFlow_DestinationConnectorPropertiesProperty struct {
	// The properties that are required to query the custom Connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-destinationconnectorproperties.html#cfn-appflow-flow-destinationconnectorproperties-customconnector
	//
	CustomConnector interface{} `field:"optional" json:"customConnector" yaml:"customConnector"`
	// The properties required to query Amazon EventBridge.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-destinationconnectorproperties.html#cfn-appflow-flow-destinationconnectorproperties-eventbridge
	//
	EventBridge interface{} `field:"optional" json:"eventBridge" yaml:"eventBridge"`
	// The properties required to query Amazon Lookout for Metrics.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-destinationconnectorproperties.html#cfn-appflow-flow-destinationconnectorproperties-lookoutmetrics
	//
	LookoutMetrics interface{} `field:"optional" json:"lookoutMetrics" yaml:"lookoutMetrics"`
	// The properties required to query Marketo.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-destinationconnectorproperties.html#cfn-appflow-flow-destinationconnectorproperties-marketo
	//
	Marketo interface{} `field:"optional" json:"marketo" yaml:"marketo"`
	// The properties required to query Amazon Redshift.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-destinationconnectorproperties.html#cfn-appflow-flow-destinationconnectorproperties-redshift
	//
	Redshift interface{} `field:"optional" json:"redshift" yaml:"redshift"`
	// The properties required to query Amazon S3.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-destinationconnectorproperties.html#cfn-appflow-flow-destinationconnectorproperties-s3
	//
	S3 interface{} `field:"optional" json:"s3" yaml:"s3"`
	// The properties required to query Salesforce.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-destinationconnectorproperties.html#cfn-appflow-flow-destinationconnectorproperties-salesforce
	//
	Salesforce interface{} `field:"optional" json:"salesforce" yaml:"salesforce"`
	// The properties required to query SAPOData.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-destinationconnectorproperties.html#cfn-appflow-flow-destinationconnectorproperties-sapodata
	//
	SapoData interface{} `field:"optional" json:"sapoData" yaml:"sapoData"`
	// The properties required to query Snowflake.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-destinationconnectorproperties.html#cfn-appflow-flow-destinationconnectorproperties-snowflake
	//
	Snowflake interface{} `field:"optional" json:"snowflake" yaml:"snowflake"`
	// The properties required to query Upsolver.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-destinationconnectorproperties.html#cfn-appflow-flow-destinationconnectorproperties-upsolver
	//
	Upsolver interface{} `field:"optional" json:"upsolver" yaml:"upsolver"`
	// The properties required to query Zendesk.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-destinationconnectorproperties.html#cfn-appflow-flow-destinationconnectorproperties-zendesk
	//
	Zendesk interface{} `field:"optional" json:"zendesk" yaml:"zendesk"`
}

