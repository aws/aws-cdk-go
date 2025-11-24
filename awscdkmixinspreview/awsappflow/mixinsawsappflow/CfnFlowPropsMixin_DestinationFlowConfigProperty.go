package mixinsawsappflow


// Contains information about the configuration of destination connectors present in the flow.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   destinationFlowConfigProperty := &DestinationFlowConfigProperty{
//   	ApiVersion: jsii.String("apiVersion"),
//   	ConnectorProfileName: jsii.String("connectorProfileName"),
//   	ConnectorType: jsii.String("connectorType"),
//   	DestinationConnectorProperties: &DestinationConnectorPropertiesProperty{
//   		CustomConnector: &CustomConnectorDestinationPropertiesProperty{
//   			CustomProperties: map[string]*string{
//   				"customPropertiesKey": jsii.String("customProperties"),
//   			},
//   			EntityName: jsii.String("entityName"),
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
//   			ErrorHandlingConfig: &ErrorHandlingConfigProperty{
//   				BucketName: jsii.String("bucketName"),
//   				BucketPrefix: jsii.String("bucketPrefix"),
//   				FailOnFirstError: jsii.Boolean(false),
//   			},
//   			Object: jsii.String("object"),
//   		},
//   		LookoutMetrics: &LookoutMetricsDestinationPropertiesProperty{
//   			Object: jsii.String("object"),
//   		},
//   		Marketo: &MarketoDestinationPropertiesProperty{
//   			ErrorHandlingConfig: &ErrorHandlingConfigProperty{
//   				BucketName: jsii.String("bucketName"),
//   				BucketPrefix: jsii.String("bucketPrefix"),
//   				FailOnFirstError: jsii.Boolean(false),
//   			},
//   			Object: jsii.String("object"),
//   		},
//   		Redshift: &RedshiftDestinationPropertiesProperty{
//   			BucketPrefix: jsii.String("bucketPrefix"),
//   			ErrorHandlingConfig: &ErrorHandlingConfigProperty{
//   				BucketName: jsii.String("bucketName"),
//   				BucketPrefix: jsii.String("bucketPrefix"),
//   				FailOnFirstError: jsii.Boolean(false),
//   			},
//   			IntermediateBucketName: jsii.String("intermediateBucketName"),
//   			Object: jsii.String("object"),
//   		},
//   		S3: &S3DestinationPropertiesProperty{
//   			BucketName: jsii.String("bucketName"),
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
//   			DataTransferApi: jsii.String("dataTransferApi"),
//   			ErrorHandlingConfig: &ErrorHandlingConfigProperty{
//   				BucketName: jsii.String("bucketName"),
//   				BucketPrefix: jsii.String("bucketPrefix"),
//   				FailOnFirstError: jsii.Boolean(false),
//   			},
//   			IdFieldNames: []*string{
//   				jsii.String("idFieldNames"),
//   			},
//   			Object: jsii.String("object"),
//   			WriteOperationType: jsii.String("writeOperationType"),
//   		},
//   		SapoData: &SAPODataDestinationPropertiesProperty{
//   			ErrorHandlingConfig: &ErrorHandlingConfigProperty{
//   				BucketName: jsii.String("bucketName"),
//   				BucketPrefix: jsii.String("bucketPrefix"),
//   				FailOnFirstError: jsii.Boolean(false),
//   			},
//   			IdFieldNames: []*string{
//   				jsii.String("idFieldNames"),
//   			},
//   			ObjectPath: jsii.String("objectPath"),
//   			SuccessResponseHandlingConfig: &SuccessResponseHandlingConfigProperty{
//   				BucketName: jsii.String("bucketName"),
//   				BucketPrefix: jsii.String("bucketPrefix"),
//   			},
//   			WriteOperationType: jsii.String("writeOperationType"),
//   		},
//   		Snowflake: &SnowflakeDestinationPropertiesProperty{
//   			BucketPrefix: jsii.String("bucketPrefix"),
//   			ErrorHandlingConfig: &ErrorHandlingConfigProperty{
//   				BucketName: jsii.String("bucketName"),
//   				BucketPrefix: jsii.String("bucketPrefix"),
//   				FailOnFirstError: jsii.Boolean(false),
//   			},
//   			IntermediateBucketName: jsii.String("intermediateBucketName"),
//   			Object: jsii.String("object"),
//   		},
//   		Upsolver: &UpsolverDestinationPropertiesProperty{
//   			BucketName: jsii.String("bucketName"),
//   			BucketPrefix: jsii.String("bucketPrefix"),
//   			S3OutputFormatConfig: &UpsolverS3OutputFormatConfigProperty{
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
//   			},
//   		},
//   		Zendesk: &ZendeskDestinationPropertiesProperty{
//   			ErrorHandlingConfig: &ErrorHandlingConfigProperty{
//   				BucketName: jsii.String("bucketName"),
//   				BucketPrefix: jsii.String("bucketPrefix"),
//   				FailOnFirstError: jsii.Boolean(false),
//   			},
//   			IdFieldNames: []*string{
//   				jsii.String("idFieldNames"),
//   			},
//   			Object: jsii.String("object"),
//   			WriteOperationType: jsii.String("writeOperationType"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-destinationflowconfig.html
//
type CfnFlowPropsMixin_DestinationFlowConfigProperty struct {
	// The API version that the destination connector uses.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-destinationflowconfig.html#cfn-appflow-flow-destinationflowconfig-apiversion
	//
	ApiVersion *string `field:"optional" json:"apiVersion" yaml:"apiVersion"`
	// The name of the connector profile.
	//
	// This name must be unique for each connector profile in the AWS account .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-destinationflowconfig.html#cfn-appflow-flow-destinationflowconfig-connectorprofilename
	//
	ConnectorProfileName *string `field:"optional" json:"connectorProfileName" yaml:"connectorProfileName"`
	// The type of destination connector, such as Sales force, Amazon S3, and so on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-destinationflowconfig.html#cfn-appflow-flow-destinationflowconfig-connectortype
	//
	ConnectorType *string `field:"optional" json:"connectorType" yaml:"connectorType"`
	// This stores the information that is required to query a particular connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-destinationflowconfig.html#cfn-appflow-flow-destinationflowconfig-destinationconnectorproperties
	//
	DestinationConnectorProperties interface{} `field:"optional" json:"destinationConnectorProperties" yaml:"destinationConnectorProperties"`
}

