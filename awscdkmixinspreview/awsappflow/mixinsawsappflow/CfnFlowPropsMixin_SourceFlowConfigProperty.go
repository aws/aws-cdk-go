package mixinsawsappflow


// Contains information about the configuration of the source connector used in the flow.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sourceFlowConfigProperty := &SourceFlowConfigProperty{
//   	ApiVersion: jsii.String("apiVersion"),
//   	ConnectorProfileName: jsii.String("connectorProfileName"),
//   	ConnectorType: jsii.String("connectorType"),
//   	IncrementalPullConfig: &IncrementalPullConfigProperty{
//   		DatetimeTypeFieldName: jsii.String("datetimeTypeFieldName"),
//   	},
//   	SourceConnectorProperties: &SourceConnectorPropertiesProperty{
//   		Amplitude: &AmplitudeSourcePropertiesProperty{
//   			Object: jsii.String("object"),
//   		},
//   		CustomConnector: &CustomConnectorSourcePropertiesProperty{
//   			CustomProperties: map[string]*string{
//   				"customPropertiesKey": jsii.String("customProperties"),
//   			},
//   			DataTransferApi: &DataTransferApiProperty{
//   				Name: jsii.String("name"),
//   				Type: jsii.String("type"),
//   			},
//   			EntityName: jsii.String("entityName"),
//   		},
//   		Datadog: &DatadogSourcePropertiesProperty{
//   			Object: jsii.String("object"),
//   		},
//   		Dynatrace: &DynatraceSourcePropertiesProperty{
//   			Object: jsii.String("object"),
//   		},
//   		GoogleAnalytics: &GoogleAnalyticsSourcePropertiesProperty{
//   			Object: jsii.String("object"),
//   		},
//   		InforNexus: &InforNexusSourcePropertiesProperty{
//   			Object: jsii.String("object"),
//   		},
//   		Marketo: &MarketoSourcePropertiesProperty{
//   			Object: jsii.String("object"),
//   		},
//   		Pardot: &PardotSourcePropertiesProperty{
//   			Object: jsii.String("object"),
//   		},
//   		S3: &S3SourcePropertiesProperty{
//   			BucketName: jsii.String("bucketName"),
//   			BucketPrefix: jsii.String("bucketPrefix"),
//   			S3InputFormatConfig: &S3InputFormatConfigProperty{
//   				S3InputFileType: jsii.String("s3InputFileType"),
//   			},
//   		},
//   		Salesforce: &SalesforceSourcePropertiesProperty{
//   			DataTransferApi: jsii.String("dataTransferApi"),
//   			EnableDynamicFieldUpdate: jsii.Boolean(false),
//   			IncludeDeletedRecords: jsii.Boolean(false),
//   			Object: jsii.String("object"),
//   		},
//   		SapoData: &SAPODataSourcePropertiesProperty{
//   			ObjectPath: jsii.String("objectPath"),
//   			PaginationConfig: &SAPODataPaginationConfigProperty{
//   				MaxPageSize: jsii.Number(123),
//   			},
//   			ParallelismConfig: &SAPODataParallelismConfigProperty{
//   				MaxParallelism: jsii.Number(123),
//   			},
//   		},
//   		ServiceNow: &ServiceNowSourcePropertiesProperty{
//   			Object: jsii.String("object"),
//   		},
//   		Singular: &SingularSourcePropertiesProperty{
//   			Object: jsii.String("object"),
//   		},
//   		Slack: &SlackSourcePropertiesProperty{
//   			Object: jsii.String("object"),
//   		},
//   		Trendmicro: &TrendmicroSourcePropertiesProperty{
//   			Object: jsii.String("object"),
//   		},
//   		Veeva: &VeevaSourcePropertiesProperty{
//   			DocumentType: jsii.String("documentType"),
//   			IncludeAllVersions: jsii.Boolean(false),
//   			IncludeRenditions: jsii.Boolean(false),
//   			IncludeSourceFiles: jsii.Boolean(false),
//   			Object: jsii.String("object"),
//   		},
//   		Zendesk: &ZendeskSourcePropertiesProperty{
//   			Object: jsii.String("object"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-sourceflowconfig.html
//
type CfnFlowPropsMixin_SourceFlowConfigProperty struct {
	// The API version of the connector when it's used as a source in the flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-sourceflowconfig.html#cfn-appflow-flow-sourceflowconfig-apiversion
	//
	ApiVersion *string `field:"optional" json:"apiVersion" yaml:"apiVersion"`
	// The name of the connector profile.
	//
	// This name must be unique for each connector profile in the AWS account .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-sourceflowconfig.html#cfn-appflow-flow-sourceflowconfig-connectorprofilename
	//
	ConnectorProfileName *string `field:"optional" json:"connectorProfileName" yaml:"connectorProfileName"`
	// The type of connector, such as Salesforce, Amplitude, and so on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-sourceflowconfig.html#cfn-appflow-flow-sourceflowconfig-connectortype
	//
	ConnectorType *string `field:"optional" json:"connectorType" yaml:"connectorType"`
	// Defines the configuration for a scheduled incremental data pull.
	//
	// If a valid configuration is provided, the fields specified in the configuration are used when querying for the incremental data pull.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-sourceflowconfig.html#cfn-appflow-flow-sourceflowconfig-incrementalpullconfig
	//
	IncrementalPullConfig interface{} `field:"optional" json:"incrementalPullConfig" yaml:"incrementalPullConfig"`
	// Specifies the information that is required to query a particular source connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-sourceflowconfig.html#cfn-appflow-flow-sourceflowconfig-sourceconnectorproperties
	//
	SourceConnectorProperties interface{} `field:"optional" json:"sourceConnectorProperties" yaml:"sourceConnectorProperties"`
}

