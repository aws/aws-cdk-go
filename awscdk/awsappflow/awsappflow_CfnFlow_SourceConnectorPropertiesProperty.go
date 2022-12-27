package awsappflow


// Specifies the information that is required to query a particular connector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceConnectorPropertiesProperty := &sourceConnectorPropertiesProperty{
//   	amplitude: &amplitudeSourcePropertiesProperty{
//   		object: jsii.String("object"),
//   	},
//   	customConnector: &customConnectorSourcePropertiesProperty{
//   		entityName: jsii.String("entityName"),
//
//   		// the properties below are optional
//   		customProperties: map[string]*string{
//   			"customPropertiesKey": jsii.String("customProperties"),
//   		},
//   	},
//   	datadog: &datadogSourcePropertiesProperty{
//   		object: jsii.String("object"),
//   	},
//   	dynatrace: &dynatraceSourcePropertiesProperty{
//   		object: jsii.String("object"),
//   	},
//   	googleAnalytics: &googleAnalyticsSourcePropertiesProperty{
//   		object: jsii.String("object"),
//   	},
//   	inforNexus: &inforNexusSourcePropertiesProperty{
//   		object: jsii.String("object"),
//   	},
//   	marketo: &marketoSourcePropertiesProperty{
//   		object: jsii.String("object"),
//   	},
//   	s3: &s3SourcePropertiesProperty{
//   		bucketName: jsii.String("bucketName"),
//   		bucketPrefix: jsii.String("bucketPrefix"),
//
//   		// the properties below are optional
//   		s3InputFormatConfig: &s3InputFormatConfigProperty{
//   			s3InputFileType: jsii.String("s3InputFileType"),
//   		},
//   	},
//   	salesforce: &salesforceSourcePropertiesProperty{
//   		object: jsii.String("object"),
//
//   		// the properties below are optional
//   		dataTransferApi: jsii.String("dataTransferApi"),
//   		enableDynamicFieldUpdate: jsii.Boolean(false),
//   		includeDeletedRecords: jsii.Boolean(false),
//   	},
//   	sapoData: &sAPODataSourcePropertiesProperty{
//   		objectPath: jsii.String("objectPath"),
//   	},
//   	serviceNow: &serviceNowSourcePropertiesProperty{
//   		object: jsii.String("object"),
//   	},
//   	singular: &singularSourcePropertiesProperty{
//   		object: jsii.String("object"),
//   	},
//   	slack: &slackSourcePropertiesProperty{
//   		object: jsii.String("object"),
//   	},
//   	trendmicro: &trendmicroSourcePropertiesProperty{
//   		object: jsii.String("object"),
//   	},
//   	veeva: &veevaSourcePropertiesProperty{
//   		object: jsii.String("object"),
//
//   		// the properties below are optional
//   		documentType: jsii.String("documentType"),
//   		includeAllVersions: jsii.Boolean(false),
//   		includeRenditions: jsii.Boolean(false),
//   		includeSourceFiles: jsii.Boolean(false),
//   	},
//   	zendesk: &zendeskSourcePropertiesProperty{
//   		object: jsii.String("object"),
//   	},
//   }
//
type CfnFlow_SourceConnectorPropertiesProperty struct {
	// Specifies the information that is required for querying Amplitude.
	Amplitude interface{} `field:"optional" json:"amplitude" yaml:"amplitude"`
	// `CfnFlow.SourceConnectorPropertiesProperty.CustomConnector`.
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

