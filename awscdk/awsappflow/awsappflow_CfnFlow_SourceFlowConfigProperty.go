package awsappflow


// Contains information about the configuration of the source connector used in the flow.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceFlowConfigProperty := &sourceFlowConfigProperty{
//   	connectorType: jsii.String("connectorType"),
//   	sourceConnectorProperties: &sourceConnectorPropertiesProperty{
//   		amplitude: &amplitudeSourcePropertiesProperty{
//   			object: jsii.String("object"),
//   		},
//   		customConnector: &customConnectorSourcePropertiesProperty{
//   			entityName: jsii.String("entityName"),
//
//   			// the properties below are optional
//   			customProperties: map[string]*string{
//   				"customPropertiesKey": jsii.String("customProperties"),
//   			},
//   		},
//   		datadog: &datadogSourcePropertiesProperty{
//   			object: jsii.String("object"),
//   		},
//   		dynatrace: &dynatraceSourcePropertiesProperty{
//   			object: jsii.String("object"),
//   		},
//   		googleAnalytics: &googleAnalyticsSourcePropertiesProperty{
//   			object: jsii.String("object"),
//   		},
//   		inforNexus: &inforNexusSourcePropertiesProperty{
//   			object: jsii.String("object"),
//   		},
//   		marketo: &marketoSourcePropertiesProperty{
//   			object: jsii.String("object"),
//   		},
//   		s3: &s3SourcePropertiesProperty{
//   			bucketName: jsii.String("bucketName"),
//   			bucketPrefix: jsii.String("bucketPrefix"),
//
//   			// the properties below are optional
//   			s3InputFormatConfig: &s3InputFormatConfigProperty{
//   				s3InputFileType: jsii.String("s3InputFileType"),
//   			},
//   		},
//   		salesforce: &salesforceSourcePropertiesProperty{
//   			object: jsii.String("object"),
//
//   			// the properties below are optional
//   			dataTransferApi: jsii.String("dataTransferApi"),
//   			enableDynamicFieldUpdate: jsii.Boolean(false),
//   			includeDeletedRecords: jsii.Boolean(false),
//   		},
//   		sapoData: &sAPODataSourcePropertiesProperty{
//   			objectPath: jsii.String("objectPath"),
//   		},
//   		serviceNow: &serviceNowSourcePropertiesProperty{
//   			object: jsii.String("object"),
//   		},
//   		singular: &singularSourcePropertiesProperty{
//   			object: jsii.String("object"),
//   		},
//   		slack: &slackSourcePropertiesProperty{
//   			object: jsii.String("object"),
//   		},
//   		trendmicro: &trendmicroSourcePropertiesProperty{
//   			object: jsii.String("object"),
//   		},
//   		veeva: &veevaSourcePropertiesProperty{
//   			object: jsii.String("object"),
//
//   			// the properties below are optional
//   			documentType: jsii.String("documentType"),
//   			includeAllVersions: jsii.Boolean(false),
//   			includeRenditions: jsii.Boolean(false),
//   			includeSourceFiles: jsii.Boolean(false),
//   		},
//   		zendesk: &zendeskSourcePropertiesProperty{
//   			object: jsii.String("object"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	apiVersion: jsii.String("apiVersion"),
//   	connectorProfileName: jsii.String("connectorProfileName"),
//   	incrementalPullConfig: &incrementalPullConfigProperty{
//   		datetimeTypeFieldName: jsii.String("datetimeTypeFieldName"),
//   	},
//   }
//
type CfnFlow_SourceFlowConfigProperty struct {
	// The type of source connector, such as Salesforce, Amplitude, and so on.
	//
	// *Allowed Values* : S3 | Amplitude | Datadog | Dynatrace | Googleanalytics | Infornexus | Salesforce | Servicenow | Singular | Slack | Trendmicro | Veeva | Zendesk.
	ConnectorType *string `field:"required" json:"connectorType" yaml:"connectorType"`
	// Specifies the information that is required to query a particular source connector.
	SourceConnectorProperties interface{} `field:"required" json:"sourceConnectorProperties" yaml:"sourceConnectorProperties"`
	// `CfnFlow.SourceFlowConfigProperty.ApiVersion`.
	ApiVersion *string `field:"optional" json:"apiVersion" yaml:"apiVersion"`
	// The name of the connector profile.
	//
	// This name must be unique for each connector profile in the AWS account .
	ConnectorProfileName *string `field:"optional" json:"connectorProfileName" yaml:"connectorProfileName"`
	// Defines the configuration for a scheduled incremental data pull.
	//
	// If a valid configuration is provided, the fields specified in the configuration are used when querying for the incremental data pull.
	IncrementalPullConfig interface{} `field:"optional" json:"incrementalPullConfig" yaml:"incrementalPullConfig"`
}

