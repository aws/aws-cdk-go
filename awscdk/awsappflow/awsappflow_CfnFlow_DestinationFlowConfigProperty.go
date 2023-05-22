package awsappflow


// Contains information about the configuration of destination connectors present in the flow.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   destinationFlowConfigProperty := &DestinationFlowConfigProperty{
//   	ConnectorType: jsii.String("connectorType"),
//   	DestinationConnectorProperties: &DestinationConnectorPropertiesProperty{
//   		CustomConnector: &CustomConnectorDestinationPropertiesProperty{
//   			EntityName: jsii.String("entityName"),
//
//   			// the properties below are optional
//   			CustomProperties: map[string]*string{
//   				"customPropertiesKey": jsii.String("customProperties"),
//   			},
//   			ErrorHandlingConfig: &ErrorHandlingConfigProperty{
//   				BucketName: jsii.String("bucketName"),
//   				BucketPrefix: jsii.String("bucketPrefix"),
//   				FailOnFirstError: jsii.Boolean(false),
//   			},
//   			IdFieldNames: []*string{
//   				jsii.String("idFieldNames"),
//   			},
//   			WriteOperationType: jsii.String("writeOperationType"),
//   		},
//   		EventBridge: &EventBridgeDestinationPropertiesProperty{
//   			Object: jsii.String("object"),
//
//   			// the properties below are optional
//   			ErrorHandlingConfig: &ErrorHandlingConfigProperty{
//   				BucketName: jsii.String("bucketName"),
//   				BucketPrefix: jsii.String("bucketPrefix"),
//   				FailOnFirstError: jsii.Boolean(false),
//   			},
//   		},
//   		LookoutMetrics: &LookoutMetricsDestinationPropertiesProperty{
//   			Object: jsii.String("object"),
//   		},
//   		Marketo: &MarketoDestinationPropertiesProperty{
//   			Object: jsii.String("object"),
//
//   			// the properties below are optional
//   			ErrorHandlingConfig: &ErrorHandlingConfigProperty{
//   				BucketName: jsii.String("bucketName"),
//   				BucketPrefix: jsii.String("bucketPrefix"),
//   				FailOnFirstError: jsii.Boolean(false),
//   			},
//   		},
//   		Redshift: &RedshiftDestinationPropertiesProperty{
//   			IntermediateBucketName: jsii.String("intermediateBucketName"),
//   			Object: jsii.String("object"),
//
//   			// the properties below are optional
//   			BucketPrefix: jsii.String("bucketPrefix"),
//   			ErrorHandlingConfig: &ErrorHandlingConfigProperty{
//   				BucketName: jsii.String("bucketName"),
//   				BucketPrefix: jsii.String("bucketPrefix"),
//   				FailOnFirstError: jsii.Boolean(false),
//   			},
//   		},
//   		S3: &S3DestinationPropertiesProperty{
//   			BucketName: jsii.String("bucketName"),
//
//   			// the properties below are optional
//   			BucketPrefix: jsii.String("bucketPrefix"),
//   			S3OutputFormatConfig: &S3OutputFormatConfigProperty{
//   				AggregationConfig: &AggregationConfigProperty{
//   					AggregationType: jsii.String("aggregationType"),
//   					TargetFileSize: jsii.Number(123),
//   				},
//   				FileType: jsii.String("fileType"),
//   				PrefixConfig: &PrefixConfigProperty{
//   					PathPrefixHierarchy: []*string{
//   						jsii.String("pathPrefixHierarchy"),
//   					},
//   					PrefixFormat: jsii.String("prefixFormat"),
//   					PrefixType: jsii.String("prefixType"),
//   				},
//   				PreserveSourceDataTyping: jsii.Boolean(false),
//   			},
//   		},
//   		Salesforce: &SalesforceDestinationPropertiesProperty{
//   			Object: jsii.String("object"),
//
//   			// the properties below are optional
//   			DataTransferApi: jsii.String("dataTransferApi"),
//   			ErrorHandlingConfig: &ErrorHandlingConfigProperty{
//   				BucketName: jsii.String("bucketName"),
//   				BucketPrefix: jsii.String("bucketPrefix"),
//   				FailOnFirstError: jsii.Boolean(false),
//   			},
//   			IdFieldNames: []*string{
//   				jsii.String("idFieldNames"),
//   			},
//   			WriteOperationType: jsii.String("writeOperationType"),
//   		},
//   		SapoData: &SAPODataDestinationPropertiesProperty{
//   			ObjectPath: jsii.String("objectPath"),
//
//   			// the properties below are optional
//   			ErrorHandlingConfig: &ErrorHandlingConfigProperty{
//   				BucketName: jsii.String("bucketName"),
//   				BucketPrefix: jsii.String("bucketPrefix"),
//   				FailOnFirstError: jsii.Boolean(false),
//   			},
//   			IdFieldNames: []*string{
//   				jsii.String("idFieldNames"),
//   			},
//   			SuccessResponseHandlingConfig: &SuccessResponseHandlingConfigProperty{
//   				BucketName: jsii.String("bucketName"),
//   				BucketPrefix: jsii.String("bucketPrefix"),
//   			},
//   			WriteOperationType: jsii.String("writeOperationType"),
//   		},
//   		Snowflake: &SnowflakeDestinationPropertiesProperty{
//   			IntermediateBucketName: jsii.String("intermediateBucketName"),
//   			Object: jsii.String("object"),
//
//   			// the properties below are optional
//   			BucketPrefix: jsii.String("bucketPrefix"),
//   			ErrorHandlingConfig: &ErrorHandlingConfigProperty{
//   				BucketName: jsii.String("bucketName"),
//   				BucketPrefix: jsii.String("bucketPrefix"),
//   				FailOnFirstError: jsii.Boolean(false),
//   			},
//   		},
//   		Upsolver: &UpsolverDestinationPropertiesProperty{
//   			BucketName: jsii.String("bucketName"),
//   			S3OutputFormatConfig: &UpsolverS3OutputFormatConfigProperty{
//   				PrefixConfig: &PrefixConfigProperty{
//   					PathPrefixHierarchy: []*string{
//   						jsii.String("pathPrefixHierarchy"),
//   					},
//   					PrefixFormat: jsii.String("prefixFormat"),
//   					PrefixType: jsii.String("prefixType"),
//   				},
//
//   				// the properties below are optional
//   				AggregationConfig: &AggregationConfigProperty{
//   					AggregationType: jsii.String("aggregationType"),
//   					TargetFileSize: jsii.Number(123),
//   				},
//   				FileType: jsii.String("fileType"),
//   			},
//
//   			// the properties below are optional
//   			BucketPrefix: jsii.String("bucketPrefix"),
//   		},
//   		Zendesk: &ZendeskDestinationPropertiesProperty{
//   			Object: jsii.String("object"),
//
//   			// the properties below are optional
//   			ErrorHandlingConfig: &ErrorHandlingConfigProperty{
//   				BucketName: jsii.String("bucketName"),
//   				BucketPrefix: jsii.String("bucketPrefix"),
//   				FailOnFirstError: jsii.Boolean(false),
//   			},
//   			IdFieldNames: []*string{
//   				jsii.String("idFieldNames"),
//   			},
//   			WriteOperationType: jsii.String("writeOperationType"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	ApiVersion: jsii.String("apiVersion"),
//   	ConnectorProfileName: jsii.String("connectorProfileName"),
//   }
//
type CfnFlow_DestinationFlowConfigProperty struct {
	// The type of destination connector, such as Sales force, Amazon S3, and so on.
	//
	// *Allowed Values* : `EventBridge | Redshift | S3 | Salesforce | Snowflake`.
	ConnectorType *string `field:"required" json:"connectorType" yaml:"connectorType"`
	// This stores the information that is required to query a particular connector.
	DestinationConnectorProperties interface{} `field:"required" json:"destinationConnectorProperties" yaml:"destinationConnectorProperties"`
	// The API version that the destination connector uses.
	ApiVersion *string `field:"optional" json:"apiVersion" yaml:"apiVersion"`
	// The name of the connector profile.
	//
	// This name must be unique for each connector profile in the AWS account .
	ConnectorProfileName *string `field:"optional" json:"connectorProfileName" yaml:"connectorProfileName"`
}

