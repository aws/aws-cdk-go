package mixinsawsappflow

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnFlowPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnFlowMixinProps := &CfnFlowMixinProps{
//   	Description: jsii.String("description"),
//   	DestinationFlowConfigList: []interface{}{
//   		&DestinationFlowConfigProperty{
//   			ApiVersion: jsii.String("apiVersion"),
//   			ConnectorProfileName: jsii.String("connectorProfileName"),
//   			ConnectorType: jsii.String("connectorType"),
//   			DestinationConnectorProperties: &DestinationConnectorPropertiesProperty{
//   				CustomConnector: &CustomConnectorDestinationPropertiesProperty{
//   					CustomProperties: map[string]*string{
//   						"customPropertiesKey": jsii.String("customProperties"),
//   					},
//   					EntityName: jsii.String("entityName"),
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
//   					ErrorHandlingConfig: &ErrorHandlingConfigProperty{
//   						BucketName: jsii.String("bucketName"),
//   						BucketPrefix: jsii.String("bucketPrefix"),
//   						FailOnFirstError: jsii.Boolean(false),
//   					},
//   					Object: jsii.String("object"),
//   				},
//   				LookoutMetrics: &LookoutMetricsDestinationPropertiesProperty{
//   					Object: jsii.String("object"),
//   				},
//   				Marketo: &MarketoDestinationPropertiesProperty{
//   					ErrorHandlingConfig: &ErrorHandlingConfigProperty{
//   						BucketName: jsii.String("bucketName"),
//   						BucketPrefix: jsii.String("bucketPrefix"),
//   						FailOnFirstError: jsii.Boolean(false),
//   					},
//   					Object: jsii.String("object"),
//   				},
//   				Redshift: &RedshiftDestinationPropertiesProperty{
//   					BucketPrefix: jsii.String("bucketPrefix"),
//   					ErrorHandlingConfig: &ErrorHandlingConfigProperty{
//   						BucketName: jsii.String("bucketName"),
//   						BucketPrefix: jsii.String("bucketPrefix"),
//   						FailOnFirstError: jsii.Boolean(false),
//   					},
//   					IntermediateBucketName: jsii.String("intermediateBucketName"),
//   					Object: jsii.String("object"),
//   				},
//   				S3: &S3DestinationPropertiesProperty{
//   					BucketName: jsii.String("bucketName"),
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
//   					DataTransferApi: jsii.String("dataTransferApi"),
//   					ErrorHandlingConfig: &ErrorHandlingConfigProperty{
//   						BucketName: jsii.String("bucketName"),
//   						BucketPrefix: jsii.String("bucketPrefix"),
//   						FailOnFirstError: jsii.Boolean(false),
//   					},
//   					IdFieldNames: []*string{
//   						jsii.String("idFieldNames"),
//   					},
//   					Object: jsii.String("object"),
//   					WriteOperationType: jsii.String("writeOperationType"),
//   				},
//   				SapoData: &SAPODataDestinationPropertiesProperty{
//   					ErrorHandlingConfig: &ErrorHandlingConfigProperty{
//   						BucketName: jsii.String("bucketName"),
//   						BucketPrefix: jsii.String("bucketPrefix"),
//   						FailOnFirstError: jsii.Boolean(false),
//   					},
//   					IdFieldNames: []*string{
//   						jsii.String("idFieldNames"),
//   					},
//   					ObjectPath: jsii.String("objectPath"),
//   					SuccessResponseHandlingConfig: &SuccessResponseHandlingConfigProperty{
//   						BucketName: jsii.String("bucketName"),
//   						BucketPrefix: jsii.String("bucketPrefix"),
//   					},
//   					WriteOperationType: jsii.String("writeOperationType"),
//   				},
//   				Snowflake: &SnowflakeDestinationPropertiesProperty{
//   					BucketPrefix: jsii.String("bucketPrefix"),
//   					ErrorHandlingConfig: &ErrorHandlingConfigProperty{
//   						BucketName: jsii.String("bucketName"),
//   						BucketPrefix: jsii.String("bucketPrefix"),
//   						FailOnFirstError: jsii.Boolean(false),
//   					},
//   					IntermediateBucketName: jsii.String("intermediateBucketName"),
//   					Object: jsii.String("object"),
//   				},
//   				Upsolver: &UpsolverDestinationPropertiesProperty{
//   					BucketName: jsii.String("bucketName"),
//   					BucketPrefix: jsii.String("bucketPrefix"),
//   					S3OutputFormatConfig: &UpsolverS3OutputFormatConfigProperty{
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
//   					},
//   				},
//   				Zendesk: &ZendeskDestinationPropertiesProperty{
//   					ErrorHandlingConfig: &ErrorHandlingConfigProperty{
//   						BucketName: jsii.String("bucketName"),
//   						BucketPrefix: jsii.String("bucketPrefix"),
//   						FailOnFirstError: jsii.Boolean(false),
//   					},
//   					IdFieldNames: []*string{
//   						jsii.String("idFieldNames"),
//   					},
//   					Object: jsii.String("object"),
//   					WriteOperationType: jsii.String("writeOperationType"),
//   				},
//   			},
//   		},
//   	},
//   	FlowName: jsii.String("flowName"),
//   	FlowStatus: jsii.String("flowStatus"),
//   	KmsArn: jsii.String("kmsArn"),
//   	MetadataCatalogConfig: &MetadataCatalogConfigProperty{
//   		GlueDataCatalog: &GlueDataCatalogProperty{
//   			DatabaseName: jsii.String("databaseName"),
//   			RoleArn: jsii.String("roleArn"),
//   			TablePrefix: jsii.String("tablePrefix"),
//   		},
//   	},
//   	SourceFlowConfig: &SourceFlowConfigProperty{
//   		ApiVersion: jsii.String("apiVersion"),
//   		ConnectorProfileName: jsii.String("connectorProfileName"),
//   		ConnectorType: jsii.String("connectorType"),
//   		IncrementalPullConfig: &IncrementalPullConfigProperty{
//   			DatetimeTypeFieldName: jsii.String("datetimeTypeFieldName"),
//   		},
//   		SourceConnectorProperties: &SourceConnectorPropertiesProperty{
//   			Amplitude: &AmplitudeSourcePropertiesProperty{
//   				Object: jsii.String("object"),
//   			},
//   			CustomConnector: &CustomConnectorSourcePropertiesProperty{
//   				CustomProperties: map[string]*string{
//   					"customPropertiesKey": jsii.String("customProperties"),
//   				},
//   				DataTransferApi: &DataTransferApiProperty{
//   					Name: jsii.String("name"),
//   					Type: jsii.String("type"),
//   				},
//   				EntityName: jsii.String("entityName"),
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
//   				S3InputFormatConfig: &S3InputFormatConfigProperty{
//   					S3InputFileType: jsii.String("s3InputFileType"),
//   				},
//   			},
//   			Salesforce: &SalesforceSourcePropertiesProperty{
//   				DataTransferApi: jsii.String("dataTransferApi"),
//   				EnableDynamicFieldUpdate: jsii.Boolean(false),
//   				IncludeDeletedRecords: jsii.Boolean(false),
//   				Object: jsii.String("object"),
//   			},
//   			SapoData: &SAPODataSourcePropertiesProperty{
//   				ObjectPath: jsii.String("objectPath"),
//   				PaginationConfig: &SAPODataPaginationConfigProperty{
//   					MaxPageSize: jsii.Number(123),
//   				},
//   				ParallelismConfig: &SAPODataParallelismConfigProperty{
//   					MaxParallelism: jsii.Number(123),
//   				},
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
//   				DocumentType: jsii.String("documentType"),
//   				IncludeAllVersions: jsii.Boolean(false),
//   				IncludeRenditions: jsii.Boolean(false),
//   				IncludeSourceFiles: jsii.Boolean(false),
//   				Object: jsii.String("object"),
//   			},
//   			Zendesk: &ZendeskSourcePropertiesProperty{
//   				Object: jsii.String("object"),
//   			},
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Tasks: []interface{}{
//   		&TaskProperty{
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
//   			SourceFields: []*string{
//   				jsii.String("sourceFields"),
//   			},
//   			TaskProperties: []interface{}{
//   				&TaskPropertiesObjectProperty{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			TaskType: jsii.String("taskType"),
//   		},
//   	},
//   	TriggerConfig: &TriggerConfigProperty{
//   		TriggerProperties: &ScheduledTriggerPropertiesProperty{
//   			DataPullMode: jsii.String("dataPullMode"),
//   			FirstExecutionFrom: jsii.Number(123),
//   			FlowErrorDeactivationThreshold: jsii.Number(123),
//   			ScheduleEndTime: jsii.Number(123),
//   			ScheduleExpression: jsii.String("scheduleExpression"),
//   			ScheduleOffset: jsii.Number(123),
//   			ScheduleStartTime: jsii.Number(123),
//   			TimeZone: jsii.String("timeZone"),
//   		},
//   		TriggerType: jsii.String("triggerType"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appflow-flow.html
//
type CfnFlowMixinProps struct {
	// A user-entered description of the flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appflow-flow.html#cfn-appflow-flow-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The configuration that controls how Amazon AppFlow places data in the destination connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appflow-flow.html#cfn-appflow-flow-destinationflowconfiglist
	//
	DestinationFlowConfigList interface{} `field:"optional" json:"destinationFlowConfigList" yaml:"destinationFlowConfigList"`
	// The specified name of the flow.
	//
	// Spaces are not allowed. Use underscores (_) or hyphens (-) only.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appflow-flow.html#cfn-appflow-flow-flowname
	//
	FlowName *string `field:"optional" json:"flowName" yaml:"flowName"`
	// Sets the status of the flow. You can specify one of the following values:.
	//
	// - **Active** - The flow runs based on the trigger settings that you defined. Active scheduled flows run as scheduled, and active event-triggered flows run when the specified change event occurs. However, active on-demand flows run only when you manually start them by using Amazon AppFlow.
	// - **Suspended** - You can use this option to deactivate an active flow. Scheduled and event-triggered flows will cease to run until you reactive them. This value only affects scheduled and event-triggered flows. It has no effect for on-demand flows.
	//
	// If you omit the FlowStatus parameter, Amazon AppFlow creates the flow with a default status. The default status for on-demand flows is Active. The default status for scheduled and event-triggered flows is Draft, which means they’re not yet active.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appflow-flow.html#cfn-appflow-flow-flowstatus
	//
	FlowStatus *string `field:"optional" json:"flowStatus" yaml:"flowStatus"`
	// The ARN (Amazon Resource Name) of the Key Management Service (KMS) key you provide for encryption.
	//
	// This is required if you do not want to use the Amazon AppFlow-managed KMS key. If you don't provide anything here, Amazon AppFlow uses the Amazon AppFlow-managed KMS key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appflow-flow.html#cfn-appflow-flow-kmsarn
	//
	KmsArn *string `field:"optional" json:"kmsArn" yaml:"kmsArn"`
	// Specifies the configuration that Amazon AppFlow uses when it catalogs your data.
	//
	// When Amazon AppFlow catalogs your data, it stores metadata in a data catalog.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appflow-flow.html#cfn-appflow-flow-metadatacatalogconfig
	//
	MetadataCatalogConfig interface{} `field:"optional" json:"metadataCatalogConfig" yaml:"metadataCatalogConfig"`
	// Contains information about the configuration of the source connector used in the flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appflow-flow.html#cfn-appflow-flow-sourceflowconfig
	//
	SourceFlowConfig interface{} `field:"optional" json:"sourceFlowConfig" yaml:"sourceFlowConfig"`
	// The tags used to organize, track, or control access for your flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appflow-flow.html#cfn-appflow-flow-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// A list of tasks that Amazon AppFlow performs while transferring the data in the flow run.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appflow-flow.html#cfn-appflow-flow-tasks
	//
	Tasks interface{} `field:"optional" json:"tasks" yaml:"tasks"`
	// The trigger settings that determine how and when Amazon AppFlow runs the specified flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appflow-flow.html#cfn-appflow-flow-triggerconfig
	//
	TriggerConfig interface{} `field:"optional" json:"triggerConfig" yaml:"triggerConfig"`
}

