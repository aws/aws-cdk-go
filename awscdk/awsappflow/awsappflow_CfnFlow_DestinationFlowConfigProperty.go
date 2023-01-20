package awsappflow


// Contains information about the configuration of destination connectors present in the flow.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   destinationFlowConfigProperty := &destinationFlowConfigProperty{
//   	connectorType: jsii.String("connectorType"),
//   	destinationConnectorProperties: &destinationConnectorPropertiesProperty{
//   		customConnector: &customConnectorDestinationPropertiesProperty{
//   			entityName: jsii.String("entityName"),
//
//   			// the properties below are optional
//   			customProperties: map[string]*string{
//   				"customPropertiesKey": jsii.String("customProperties"),
//   			},
//   			errorHandlingConfig: &errorHandlingConfigProperty{
//   				bucketName: jsii.String("bucketName"),
//   				bucketPrefix: jsii.String("bucketPrefix"),
//   				failOnFirstError: jsii.Boolean(false),
//   			},
//   			idFieldNames: []*string{
//   				jsii.String("idFieldNames"),
//   			},
//   			writeOperationType: jsii.String("writeOperationType"),
//   		},
//   		eventBridge: &eventBridgeDestinationPropertiesProperty{
//   			object: jsii.String("object"),
//
//   			// the properties below are optional
//   			errorHandlingConfig: &errorHandlingConfigProperty{
//   				bucketName: jsii.String("bucketName"),
//   				bucketPrefix: jsii.String("bucketPrefix"),
//   				failOnFirstError: jsii.Boolean(false),
//   			},
//   		},
//   		lookoutMetrics: &lookoutMetricsDestinationPropertiesProperty{
//   			object: jsii.String("object"),
//   		},
//   		marketo: &marketoDestinationPropertiesProperty{
//   			object: jsii.String("object"),
//
//   			// the properties below are optional
//   			errorHandlingConfig: &errorHandlingConfigProperty{
//   				bucketName: jsii.String("bucketName"),
//   				bucketPrefix: jsii.String("bucketPrefix"),
//   				failOnFirstError: jsii.Boolean(false),
//   			},
//   		},
//   		redshift: &redshiftDestinationPropertiesProperty{
//   			intermediateBucketName: jsii.String("intermediateBucketName"),
//   			object: jsii.String("object"),
//
//   			// the properties below are optional
//   			bucketPrefix: jsii.String("bucketPrefix"),
//   			errorHandlingConfig: &errorHandlingConfigProperty{
//   				bucketName: jsii.String("bucketName"),
//   				bucketPrefix: jsii.String("bucketPrefix"),
//   				failOnFirstError: jsii.Boolean(false),
//   			},
//   		},
//   		s3: &s3DestinationPropertiesProperty{
//   			bucketName: jsii.String("bucketName"),
//
//   			// the properties below are optional
//   			bucketPrefix: jsii.String("bucketPrefix"),
//   			s3OutputFormatConfig: &s3OutputFormatConfigProperty{
//   				aggregationConfig: &aggregationConfigProperty{
//   					aggregationType: jsii.String("aggregationType"),
//   					targetFileSize: jsii.Number(123),
//   				},
//   				fileType: jsii.String("fileType"),
//   				prefixConfig: &prefixConfigProperty{
//   					pathPrefixHierarchy: []*string{
//   						jsii.String("pathPrefixHierarchy"),
//   					},
//   					prefixFormat: jsii.String("prefixFormat"),
//   					prefixType: jsii.String("prefixType"),
//   				},
//   				preserveSourceDataTyping: jsii.Boolean(false),
//   			},
//   		},
//   		salesforce: &salesforceDestinationPropertiesProperty{
//   			object: jsii.String("object"),
//
//   			// the properties below are optional
//   			dataTransferApi: jsii.String("dataTransferApi"),
//   			errorHandlingConfig: &errorHandlingConfigProperty{
//   				bucketName: jsii.String("bucketName"),
//   				bucketPrefix: jsii.String("bucketPrefix"),
//   				failOnFirstError: jsii.Boolean(false),
//   			},
//   			idFieldNames: []*string{
//   				jsii.String("idFieldNames"),
//   			},
//   			writeOperationType: jsii.String("writeOperationType"),
//   		},
//   		sapoData: &sAPODataDestinationPropertiesProperty{
//   			objectPath: jsii.String("objectPath"),
//
//   			// the properties below are optional
//   			errorHandlingConfig: &errorHandlingConfigProperty{
//   				bucketName: jsii.String("bucketName"),
//   				bucketPrefix: jsii.String("bucketPrefix"),
//   				failOnFirstError: jsii.Boolean(false),
//   			},
//   			idFieldNames: []*string{
//   				jsii.String("idFieldNames"),
//   			},
//   			successResponseHandlingConfig: &successResponseHandlingConfigProperty{
//   				bucketName: jsii.String("bucketName"),
//   				bucketPrefix: jsii.String("bucketPrefix"),
//   			},
//   			writeOperationType: jsii.String("writeOperationType"),
//   		},
//   		snowflake: &snowflakeDestinationPropertiesProperty{
//   			intermediateBucketName: jsii.String("intermediateBucketName"),
//   			object: jsii.String("object"),
//
//   			// the properties below are optional
//   			bucketPrefix: jsii.String("bucketPrefix"),
//   			errorHandlingConfig: &errorHandlingConfigProperty{
//   				bucketName: jsii.String("bucketName"),
//   				bucketPrefix: jsii.String("bucketPrefix"),
//   				failOnFirstError: jsii.Boolean(false),
//   			},
//   		},
//   		upsolver: &upsolverDestinationPropertiesProperty{
//   			bucketName: jsii.String("bucketName"),
//   			s3OutputFormatConfig: &upsolverS3OutputFormatConfigProperty{
//   				prefixConfig: &prefixConfigProperty{
//   					pathPrefixHierarchy: []*string{
//   						jsii.String("pathPrefixHierarchy"),
//   					},
//   					prefixFormat: jsii.String("prefixFormat"),
//   					prefixType: jsii.String("prefixType"),
//   				},
//
//   				// the properties below are optional
//   				aggregationConfig: &aggregationConfigProperty{
//   					aggregationType: jsii.String("aggregationType"),
//   					targetFileSize: jsii.Number(123),
//   				},
//   				fileType: jsii.String("fileType"),
//   			},
//
//   			// the properties below are optional
//   			bucketPrefix: jsii.String("bucketPrefix"),
//   		},
//   		zendesk: &zendeskDestinationPropertiesProperty{
//   			object: jsii.String("object"),
//
//   			// the properties below are optional
//   			errorHandlingConfig: &errorHandlingConfigProperty{
//   				bucketName: jsii.String("bucketName"),
//   				bucketPrefix: jsii.String("bucketPrefix"),
//   				failOnFirstError: jsii.Boolean(false),
//   			},
//   			idFieldNames: []*string{
//   				jsii.String("idFieldNames"),
//   			},
//   			writeOperationType: jsii.String("writeOperationType"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	apiVersion: jsii.String("apiVersion"),
//   	connectorProfileName: jsii.String("connectorProfileName"),
//   }
//
type CfnFlow_DestinationFlowConfigProperty struct {
	// The type of destination connector, such as Sales force, Amazon S3, and so on.
	//
	// *Allowed Values* : `EventBridge | Redshift | S3 | Salesforce | Snowflake`.
	ConnectorType *string `field:"required" json:"connectorType" yaml:"connectorType"`
	// This stores the information that is required to query a particular connector.
	DestinationConnectorProperties interface{} `field:"required" json:"destinationConnectorProperties" yaml:"destinationConnectorProperties"`
	// `CfnFlow.DestinationFlowConfigProperty.ApiVersion`.
	ApiVersion *string `field:"optional" json:"apiVersion" yaml:"apiVersion"`
	// The name of the connector profile.
	//
	// This name must be unique for each connector profile in the AWS account .
	ConnectorProfileName *string `field:"optional" json:"connectorProfileName" yaml:"connectorProfileName"`
}

