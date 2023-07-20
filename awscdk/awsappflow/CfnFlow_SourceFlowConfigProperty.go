package awsappflow


// Contains information about the configuration of the source connector used in the flow.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceFlowConfigProperty := &SourceFlowConfigProperty{
//   	ConnectorType: jsii.String("connectorType"),
//   	SourceConnectorProperties: &SourceConnectorPropertiesProperty{
//   		Amplitude: &AmplitudeSourcePropertiesProperty{
//   			Object: jsii.String("object"),
//   		},
//   		CustomConnector: &CustomConnectorSourcePropertiesProperty{
//   			EntityName: jsii.String("entityName"),
//
//   			// the properties below are optional
//   			CustomProperties: map[string]*string{
//   				"customPropertiesKey": jsii.String("customProperties"),
//   			},
//   			DataTransferApi: &DataTransferApiProperty{
//   				Name: jsii.String("name"),
//   				Type: jsii.String("type"),
//   			},
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
//
//   			// the properties below are optional
//   			S3InputFormatConfig: &S3InputFormatConfigProperty{
//   				S3InputFileType: jsii.String("s3InputFileType"),
//   			},
//   		},
//   		Salesforce: &SalesforceSourcePropertiesProperty{
//   			Object: jsii.String("object"),
//
//   			// the properties below are optional
//   			DataTransferApi: jsii.String("dataTransferApi"),
//   			EnableDynamicFieldUpdate: jsii.Boolean(false),
//   			IncludeDeletedRecords: jsii.Boolean(false),
//   		},
//   		SapoData: &SAPODataSourcePropertiesProperty{
//   			ObjectPath: jsii.String("objectPath"),
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
//   			Object: jsii.String("object"),
//
//   			// the properties below are optional
//   			DocumentType: jsii.String("documentType"),
//   			IncludeAllVersions: jsii.Boolean(false),
//   			IncludeRenditions: jsii.Boolean(false),
//   			IncludeSourceFiles: jsii.Boolean(false),
//   		},
//   		Zendesk: &ZendeskSourcePropertiesProperty{
//   			Object: jsii.String("object"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	ApiVersion: jsii.String("apiVersion"),
//   	ConnectorProfileName: jsii.String("connectorProfileName"),
//   	IncrementalPullConfig: &IncrementalPullConfigProperty{
//   		DatetimeTypeFieldName: jsii.String("datetimeTypeFieldName"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-sourceflowconfig.html
//
type CfnFlow_SourceFlowConfigProperty struct {
	// The type of connector, such as Salesforce, Amplitude, and so on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-sourceflowconfig.html#cfn-appflow-flow-sourceflowconfig-connectortype
	//
	ConnectorType *string `field:"required" json:"connectorType" yaml:"connectorType"`
	// Specifies the information that is required to query a particular source connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-sourceflowconfig.html#cfn-appflow-flow-sourceflowconfig-sourceconnectorproperties
	//
	SourceConnectorProperties interface{} `field:"required" json:"sourceConnectorProperties" yaml:"sourceConnectorProperties"`
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
	// Defines the configuration for a scheduled incremental data pull.
	//
	// If a valid configuration is provided, the fields specified in the configuration are used when querying for the incremental data pull.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-sourceflowconfig.html#cfn-appflow-flow-sourceflowconfig-incrementalpullconfig
	//
	IncrementalPullConfig interface{} `field:"optional" json:"incrementalPullConfig" yaml:"incrementalPullConfig"`
}

