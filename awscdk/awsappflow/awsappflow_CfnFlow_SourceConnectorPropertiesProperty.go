package awsappflow


// Specifies the information that is required to query a particular connector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceConnectorPropertiesProperty := &SourceConnectorPropertiesProperty{
//   	Amplitude: &AmplitudeSourcePropertiesProperty{
//   		Object: jsii.String("object"),
//   	},
//   	CustomConnector: &CustomConnectorSourcePropertiesProperty{
//   		EntityName: jsii.String("entityName"),
//
//   		// the properties below are optional
//   		CustomProperties: map[string]*string{
//   			"customPropertiesKey": jsii.String("customProperties"),
//   		},
//   	},
//   	Datadog: &DatadogSourcePropertiesProperty{
//   		Object: jsii.String("object"),
//   	},
//   	Dynatrace: &DynatraceSourcePropertiesProperty{
//   		Object: jsii.String("object"),
//   	},
//   	GoogleAnalytics: &GoogleAnalyticsSourcePropertiesProperty{
//   		Object: jsii.String("object"),
//   	},
//   	InforNexus: &InforNexusSourcePropertiesProperty{
//   		Object: jsii.String("object"),
//   	},
//   	Marketo: &MarketoSourcePropertiesProperty{
//   		Object: jsii.String("object"),
//   	},
//   	Pardot: &PardotSourcePropertiesProperty{
//   		Object: jsii.String("object"),
//   	},
//   	S3: &S3SourcePropertiesProperty{
//   		BucketName: jsii.String("bucketName"),
//   		BucketPrefix: jsii.String("bucketPrefix"),
//
//   		// the properties below are optional
//   		S3InputFormatConfig: &S3InputFormatConfigProperty{
//   			S3InputFileType: jsii.String("s3InputFileType"),
//   		},
//   	},
//   	Salesforce: &SalesforceSourcePropertiesProperty{
//   		Object: jsii.String("object"),
//
//   		// the properties below are optional
//   		DataTransferApi: jsii.String("dataTransferApi"),
//   		EnableDynamicFieldUpdate: jsii.Boolean(false),
//   		IncludeDeletedRecords: jsii.Boolean(false),
//   	},
//   	SapoData: &SAPODataSourcePropertiesProperty{
//   		ObjectPath: jsii.String("objectPath"),
//   	},
//   	ServiceNow: &ServiceNowSourcePropertiesProperty{
//   		Object: jsii.String("object"),
//   	},
//   	Singular: &SingularSourcePropertiesProperty{
//   		Object: jsii.String("object"),
//   	},
//   	Slack: &SlackSourcePropertiesProperty{
//   		Object: jsii.String("object"),
//   	},
//   	Trendmicro: &TrendmicroSourcePropertiesProperty{
//   		Object: jsii.String("object"),
//   	},
//   	Veeva: &VeevaSourcePropertiesProperty{
//   		Object: jsii.String("object"),
//
//   		// the properties below are optional
//   		DocumentType: jsii.String("documentType"),
//   		IncludeAllVersions: jsii.Boolean(false),
//   		IncludeRenditions: jsii.Boolean(false),
//   		IncludeSourceFiles: jsii.Boolean(false),
//   	},
//   	Zendesk: &ZendeskSourcePropertiesProperty{
//   		Object: jsii.String("object"),
//   	},
//   }
//
type CfnFlow_SourceConnectorPropertiesProperty struct {
	// Specifies the information that is required for querying Amplitude.
	Amplitude interface{} `field:"optional" json:"amplitude" yaml:"amplitude"`
	// The properties that are applied when the custom connector is being used as a source.
	CustomConnector interface{} `field:"optional" json:"customConnector" yaml:"customConnector"`
	// Specifies the information that is required for querying Datadog.
	Datadog interface{} `field:"optional" json:"datadog" yaml:"datadog"`
	// Specifies the information that is required for querying Dynatrace.
	Dynatrace interface{} `field:"optional" json:"dynatrace" yaml:"dynatrace"`
	// Specifies the information that is required for querying Google Analytics.
	GoogleAnalytics interface{} `field:"optional" json:"googleAnalytics" yaml:"googleAnalytics"`
	// Specifies the information that is required for querying Infor Nexus.
	InforNexus interface{} `field:"optional" json:"inforNexus" yaml:"inforNexus"`
	// Specifies the information that is required for querying Marketo.
	Marketo interface{} `field:"optional" json:"marketo" yaml:"marketo"`
	// `CfnFlow.SourceConnectorPropertiesProperty.Pardot`.
	Pardot interface{} `field:"optional" json:"pardot" yaml:"pardot"`
	// Specifies the information that is required for querying Amazon S3.
	S3 interface{} `field:"optional" json:"s3" yaml:"s3"`
	// Specifies the information that is required for querying Salesforce.
	Salesforce interface{} `field:"optional" json:"salesforce" yaml:"salesforce"`
	// The properties that are applied when using SAPOData as a flow source.
	SapoData interface{} `field:"optional" json:"sapoData" yaml:"sapoData"`
	// Specifies the information that is required for querying ServiceNow.
	ServiceNow interface{} `field:"optional" json:"serviceNow" yaml:"serviceNow"`
	// Specifies the information that is required for querying Singular.
	Singular interface{} `field:"optional" json:"singular" yaml:"singular"`
	// Specifies the information that is required for querying Slack.
	Slack interface{} `field:"optional" json:"slack" yaml:"slack"`
	// Specifies the information that is required for querying Trend Micro.
	Trendmicro interface{} `field:"optional" json:"trendmicro" yaml:"trendmicro"`
	// Specifies the information that is required for querying Veeva.
	Veeva interface{} `field:"optional" json:"veeva" yaml:"veeva"`
	// Specifies the information that is required for querying Zendesk.
	Zendesk interface{} `field:"optional" json:"zendesk" yaml:"zendesk"`
}

