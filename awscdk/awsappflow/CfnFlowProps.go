package awsappflow

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnFlow`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFlowProps := &CfnFlowProps{
//   	DestinationFlowConfigList: []interface{}{
//   		&DestinationFlowConfigProperty{
//   			ConnectorType: jsii.String("connectorType"),
//   			DestinationConnectorProperties: &DestinationConnectorPropertiesProperty{
//   				CustomConnector: &CustomConnectorDestinationPropertiesProperty{
//   					EntityName: jsii.String("entityName"),
//
//   					// the properties below are optional
//   					CustomProperties: map[string]*string{
//   						"customPropertiesKey": jsii.String("customProperties"),
//   					},
//   					ErrorHandlingConfig: &ErrorHandlingConfigProperty{
//   						BucketName: jsii.String("bucketName"),
//   						BucketPrefix: jsii.String("bucketPrefix"),
//   						FailOnFirstError: jsii.Boolean(false),
//   					},
//   					IdFieldNames: []*string{
//   						jsii.String("idFieldNames"),
//   					},
//   					WriteOperationType: jsii.String("writeOperationType"),
//   				},
//   				EventBridge: &EventBridgeDestinationPropertiesProperty{
//   					Object: jsii.String("object"),
//
//   					// the properties below are optional
//   					ErrorHandlingConfig: &ErrorHandlingConfigProperty{
//   						BucketName: jsii.String("bucketName"),
//   						BucketPrefix: jsii.String("bucketPrefix"),
//   						FailOnFirstError: jsii.Boolean(false),
//   					},
//   				},
//   				LookoutMetrics: &LookoutMetricsDestinationPropertiesProperty{
//   					Object: jsii.String("object"),
//   				},
//   				Marketo: &MarketoDestinationPropertiesProperty{
//   					Object: jsii.String("object"),
//
//   					// the properties below are optional
//   					ErrorHandlingConfig: &ErrorHandlingConfigProperty{
//   						BucketName: jsii.String("bucketName"),
//   						BucketPrefix: jsii.String("bucketPrefix"),
//   						FailOnFirstError: jsii.Boolean(false),
//   					},
//   				},
//   				Redshift: &RedshiftDestinationPropertiesProperty{
//   					IntermediateBucketName: jsii.String("intermediateBucketName"),
//   					Object: jsii.String("object"),
//
//   					// the properties below are optional
//   					BucketPrefix: jsii.String("bucketPrefix"),
//   					ErrorHandlingConfig: &ErrorHandlingConfigProperty{
//   						BucketName: jsii.String("bucketName"),
//   						BucketPrefix: jsii.String("bucketPrefix"),
//   						FailOnFirstError: jsii.Boolean(false),
//   					},
//   				},
//   				S3: &S3DestinationPropertiesProperty{
//   					BucketName: jsii.String("bucketName"),
//
//   					// the properties below are optional
//   					BucketPrefix: jsii.String("bucketPrefix"),
//   					S3OutputFormatConfig: &S3OutputFormatConfigProperty{
//   						AggregationConfig: &AggregationConfigProperty{
//   							AggregationType: jsii.String("aggregationType"),
//   							TargetFileSize: jsii.Number(123),
//   						},
//   						FileType: jsii.String("fileType"),
//   						PrefixConfig: &PrefixConfigProperty{
//   							PathPrefixHierarchy: []*string{
//   								jsii.String("pathPrefixHierarchy"),
//   							},
//   							PrefixFormat: jsii.String("prefixFormat"),
//   							PrefixType: jsii.String("prefixType"),
//   						},
//   						PreserveSourceDataTyping: jsii.Boolean(false),
//   					},
//   				},
//   				Salesforce: &SalesforceDestinationPropertiesProperty{
//   					Object: jsii.String("object"),
//
//   					// the properties below are optional
//   					DataTransferApi: jsii.String("dataTransferApi"),
//   					ErrorHandlingConfig: &ErrorHandlingConfigProperty{
//   						BucketName: jsii.String("bucketName"),
//   						BucketPrefix: jsii.String("bucketPrefix"),
//   						FailOnFirstError: jsii.Boolean(false),
//   					},
//   					IdFieldNames: []*string{
//   						jsii.String("idFieldNames"),
//   					},
//   					WriteOperationType: jsii.String("writeOperationType"),
//   				},
//   				SapoData: &SAPODataDestinationPropertiesProperty{
//   					ObjectPath: jsii.String("objectPath"),
//
//   					// the properties below are optional
//   					ErrorHandlingConfig: &ErrorHandlingConfigProperty{
//   						BucketName: jsii.String("bucketName"),
//   						BucketPrefix: jsii.String("bucketPrefix"),
//   						FailOnFirstError: jsii.Boolean(false),
//   					},
//   					IdFieldNames: []*string{
//   						jsii.String("idFieldNames"),
//   					},
//   					SuccessResponseHandlingConfig: &SuccessResponseHandlingConfigProperty{
//   						BucketName: jsii.String("bucketName"),
//   						BucketPrefix: jsii.String("bucketPrefix"),
//   					},
//   					WriteOperationType: jsii.String("writeOperationType"),
//   				},
//   				Snowflake: &SnowflakeDestinationPropertiesProperty{
//   					IntermediateBucketName: jsii.String("intermediateBucketName"),
//   					Object: jsii.String("object"),
//
//   					// the properties below are optional
//   					BucketPrefix: jsii.String("bucketPrefix"),
//   					ErrorHandlingConfig: &ErrorHandlingConfigProperty{
//   						BucketName: jsii.String("bucketName"),
//   						BucketPrefix: jsii.String("bucketPrefix"),
//   						FailOnFirstError: jsii.Boolean(false),
//   					},
//   				},
//   				Upsolver: &UpsolverDestinationPropertiesProperty{
//   					BucketName: jsii.String("bucketName"),
//   					S3OutputFormatConfig: &UpsolverS3OutputFormatConfigProperty{
//   						PrefixConfig: &PrefixConfigProperty{
//   							PathPrefixHierarchy: []*string{
//   								jsii.String("pathPrefixHierarchy"),
//   							},
//   							PrefixFormat: jsii.String("prefixFormat"),
//   							PrefixType: jsii.String("prefixType"),
//   						},
//
//   						// the properties below are optional
//   						AggregationConfig: &AggregationConfigProperty{
//   							AggregationType: jsii.String("aggregationType"),
//   							TargetFileSize: jsii.Number(123),
//   						},
//   						FileType: jsii.String("fileType"),
//   					},
//
//   					// the properties below are optional
//   					BucketPrefix: jsii.String("bucketPrefix"),
//   				},
//   				Zendesk: &ZendeskDestinationPropertiesProperty{
//   					Object: jsii.String("object"),
//
//   					// the properties below are optional
//   					ErrorHandlingConfig: &ErrorHandlingConfigProperty{
//   						BucketName: jsii.String("bucketName"),
//   						BucketPrefix: jsii.String("bucketPrefix"),
//   						FailOnFirstError: jsii.Boolean(false),
//   					},
//   					IdFieldNames: []*string{
//   						jsii.String("idFieldNames"),
//   					},
//   					WriteOperationType: jsii.String("writeOperationType"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			ApiVersion: jsii.String("apiVersion"),
//   			ConnectorProfileName: jsii.String("connectorProfileName"),
//   		},
//   	},
//   	FlowName: jsii.String("flowName"),
//   	SourceFlowConfig: &SourceFlowConfigProperty{
//   		ConnectorType: jsii.String("connectorType"),
//   		SourceConnectorProperties: &SourceConnectorPropertiesProperty{
//   			Amplitude: &AmplitudeSourcePropertiesProperty{
//   				Object: jsii.String("object"),
//   			},
//   			CustomConnector: &CustomConnectorSourcePropertiesProperty{
//   				EntityName: jsii.String("entityName"),
//
//   				// the properties below are optional
//   				CustomProperties: map[string]*string{
//   					"customPropertiesKey": jsii.String("customProperties"),
//   				},
//   			},
//   			Datadog: &DatadogSourcePropertiesProperty{
//   				Object: jsii.String("object"),
//   			},
//   			Dynatrace: &DynatraceSourcePropertiesProperty{
//   				Object: jsii.String("object"),
//   			},
//   			GoogleAnalytics: &GoogleAnalyticsSourcePropertiesProperty{
//   				Object: jsii.String("object"),
//   			},
//   			InforNexus: &InforNexusSourcePropertiesProperty{
//   				Object: jsii.String("object"),
//   			},
//   			Marketo: &MarketoSourcePropertiesProperty{
//   				Object: jsii.String("object"),
//   			},
//   			Pardot: &PardotSourcePropertiesProperty{
//   				Object: jsii.String("object"),
//   			},
//   			S3: &S3SourcePropertiesProperty{
//   				BucketName: jsii.String("bucketName"),
//   				BucketPrefix: jsii.String("bucketPrefix"),
//
//   				// the properties below are optional
//   				S3InputFormatConfig: &S3InputFormatConfigProperty{
//   					S3InputFileType: jsii.String("s3InputFileType"),
//   				},
//   			},
//   			Salesforce: &SalesforceSourcePropertiesProperty{
//   				Object: jsii.String("object"),
//
//   				// the properties below are optional
//   				DataTransferApi: jsii.String("dataTransferApi"),
//   				EnableDynamicFieldUpdate: jsii.Boolean(false),
//   				IncludeDeletedRecords: jsii.Boolean(false),
//   			},
//   			SapoData: &SAPODataSourcePropertiesProperty{
//   				ObjectPath: jsii.String("objectPath"),
//   			},
//   			ServiceNow: &ServiceNowSourcePropertiesProperty{
//   				Object: jsii.String("object"),
//   			},
//   			Singular: &SingularSourcePropertiesProperty{
//   				Object: jsii.String("object"),
//   			},
//   			Slack: &SlackSourcePropertiesProperty{
//   				Object: jsii.String("object"),
//   			},
//   			Trendmicro: &TrendmicroSourcePropertiesProperty{
//   				Object: jsii.String("object"),
//   			},
//   			Veeva: &VeevaSourcePropertiesProperty{
//   				Object: jsii.String("object"),
//
//   				// the properties below are optional
//   				DocumentType: jsii.String("documentType"),
//   				IncludeAllVersions: jsii.Boolean(false),
//   				IncludeRenditions: jsii.Boolean(false),
//   				IncludeSourceFiles: jsii.Boolean(false),
//   			},
//   			Zendesk: &ZendeskSourcePropertiesProperty{
//   				Object: jsii.String("object"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		ApiVersion: jsii.String("apiVersion"),
//   		ConnectorProfileName: jsii.String("connectorProfileName"),
//   		IncrementalPullConfig: &IncrementalPullConfigProperty{
//   			DatetimeTypeFieldName: jsii.String("datetimeTypeFieldName"),
//   		},
//   	},
//   	Tasks: []interface{}{
//   		&TaskProperty{
//   			SourceFields: []*string{
//   				jsii.String("sourceFields"),
//   			},
//   			TaskType: jsii.String("taskType"),
//
//   			// the properties below are optional
//   			ConnectorOperator: &ConnectorOperatorProperty{
//   				Amplitude: jsii.String("amplitude"),
//   				CustomConnector: jsii.String("customConnector"),
//   				Datadog: jsii.String("datadog"),
//   				Dynatrace: jsii.String("dynatrace"),
//   				GoogleAnalytics: jsii.String("googleAnalytics"),
//   				InforNexus: jsii.String("inforNexus"),
//   				Marketo: jsii.String("marketo"),
//   				Pardot: jsii.String("pardot"),
//   				S3: jsii.String("s3"),
//   				Salesforce: jsii.String("salesforce"),
//   				SapoData: jsii.String("sapoData"),
//   				ServiceNow: jsii.String("serviceNow"),
//   				Singular: jsii.String("singular"),
//   				Slack: jsii.String("slack"),
//   				Trendmicro: jsii.String("trendmicro"),
//   				Veeva: jsii.String("veeva"),
//   				Zendesk: jsii.String("zendesk"),
//   			},
//   			DestinationField: jsii.String("destinationField"),
//   			TaskProperties: []interface{}{
//   				&TaskPropertiesObjectProperty{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	TriggerConfig: &TriggerConfigProperty{
//   		TriggerType: jsii.String("triggerType"),
//
//   		// the properties below are optional
//   		ActivateFlowOnCreate: jsii.Boolean(false),
//   		TriggerProperties: &ScheduledTriggerPropertiesProperty{
//   			ScheduleExpression: jsii.String("scheduleExpression"),
//
//   			// the properties below are optional
//   			DataPullMode: jsii.String("dataPullMode"),
//   			FirstExecutionFrom: jsii.Number(123),
//   			FlowErrorDeactivationThreshold: jsii.Number(123),
//   			ScheduleEndTime: jsii.Number(123),
//   			ScheduleOffset: jsii.Number(123),
//   			ScheduleStartTime: jsii.Number(123),
//   			TimeZone: jsii.String("timeZone"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	KmsArn: jsii.String("kmsArn"),
//   	MetadataCatalogConfig: &MetadataCatalogConfigProperty{
//   		GlueDataCatalog: &GlueDataCatalogProperty{
//   			DatabaseName: jsii.String("databaseName"),
//   			RoleArn: jsii.String("roleArn"),
//   			TablePrefix: jsii.String("tablePrefix"),
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnFlowProps struct {
	// The configuration that controls how Amazon AppFlow places data in the destination connector.
	DestinationFlowConfigList interface{} `field:"required" json:"destinationFlowConfigList" yaml:"destinationFlowConfigList"`
	// The specified name of the flow.
	//
	// Spaces are not allowed. Use underscores (_) or hyphens (-) only.
	FlowName *string `field:"required" json:"flowName" yaml:"flowName"`
	// Contains information about the configuration of the source connector used in the flow.
	SourceFlowConfig interface{} `field:"required" json:"sourceFlowConfig" yaml:"sourceFlowConfig"`
	// A list of tasks that Amazon AppFlow performs while transferring the data in the flow run.
	Tasks interface{} `field:"required" json:"tasks" yaml:"tasks"`
	// The trigger settings that determine how and when Amazon AppFlow runs the specified flow.
	TriggerConfig interface{} `field:"required" json:"triggerConfig" yaml:"triggerConfig"`
	// A user-entered description of the flow.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The ARN (Amazon Resource Name) of the Key Management Service (KMS) key you provide for encryption.
	//
	// This is required if you do not want to use the Amazon AppFlow-managed KMS key. If you don't provide anything here, Amazon AppFlow uses the Amazon AppFlow-managed KMS key.
	KmsArn *string `field:"optional" json:"kmsArn" yaml:"kmsArn"`
	// `AWS::AppFlow::Flow.MetadataCatalogConfig`.
	MetadataCatalogConfig interface{} `field:"optional" json:"metadataCatalogConfig" yaml:"metadataCatalogConfig"`
	// The tags used to organize, track, or control access for your flow.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

