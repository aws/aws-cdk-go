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
//   		DataTransferApi: &DataTransferApiProperty{
//   			Name: jsii.String("name"),
//   			Type: jsii.String("type"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-sourceconnectorproperties.html
//
type CfnFlow_SourceConnectorPropertiesProperty struct {
	// Specifies the information that is required for querying Amplitude.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-sourceconnectorproperties.html#cfn-appflow-flow-sourceconnectorproperties-amplitude
	//
	Amplitude interface{} `field:"optional" json:"amplitude" yaml:"amplitude"`
	// The properties that are applied when the custom connector is being used as a source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-sourceconnectorproperties.html#cfn-appflow-flow-sourceconnectorproperties-customconnector
	//
	CustomConnector interface{} `field:"optional" json:"customConnector" yaml:"customConnector"`
	// Specifies the information that is required for querying Datadog.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-sourceconnectorproperties.html#cfn-appflow-flow-sourceconnectorproperties-datadog
	//
	Datadog interface{} `field:"optional" json:"datadog" yaml:"datadog"`
	// Specifies the information that is required for querying Dynatrace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-sourceconnectorproperties.html#cfn-appflow-flow-sourceconnectorproperties-dynatrace
	//
	Dynatrace interface{} `field:"optional" json:"dynatrace" yaml:"dynatrace"`
	// Specifies the information that is required for querying Google Analytics.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-sourceconnectorproperties.html#cfn-appflow-flow-sourceconnectorproperties-googleanalytics
	//
	GoogleAnalytics interface{} `field:"optional" json:"googleAnalytics" yaml:"googleAnalytics"`
	// Specifies the information that is required for querying Infor Nexus.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-sourceconnectorproperties.html#cfn-appflow-flow-sourceconnectorproperties-infornexus
	//
	InforNexus interface{} `field:"optional" json:"inforNexus" yaml:"inforNexus"`
	// Specifies the information that is required for querying Marketo.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-sourceconnectorproperties.html#cfn-appflow-flow-sourceconnectorproperties-marketo
	//
	Marketo interface{} `field:"optional" json:"marketo" yaml:"marketo"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-sourceconnectorproperties.html#cfn-appflow-flow-sourceconnectorproperties-pardot
	//
	Pardot interface{} `field:"optional" json:"pardot" yaml:"pardot"`
	// Specifies the information that is required for querying Amazon S3.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-sourceconnectorproperties.html#cfn-appflow-flow-sourceconnectorproperties-s3
	//
	S3 interface{} `field:"optional" json:"s3" yaml:"s3"`
	// Specifies the information that is required for querying Salesforce.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-sourceconnectorproperties.html#cfn-appflow-flow-sourceconnectorproperties-salesforce
	//
	Salesforce interface{} `field:"optional" json:"salesforce" yaml:"salesforce"`
	// The properties that are applied when using SAPOData as a flow source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-sourceconnectorproperties.html#cfn-appflow-flow-sourceconnectorproperties-sapodata
	//
	SapoData interface{} `field:"optional" json:"sapoData" yaml:"sapoData"`
	// Specifies the information that is required for querying ServiceNow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-sourceconnectorproperties.html#cfn-appflow-flow-sourceconnectorproperties-servicenow
	//
	ServiceNow interface{} `field:"optional" json:"serviceNow" yaml:"serviceNow"`
	// Specifies the information that is required for querying Singular.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-sourceconnectorproperties.html#cfn-appflow-flow-sourceconnectorproperties-singular
	//
	Singular interface{} `field:"optional" json:"singular" yaml:"singular"`
	// Specifies the information that is required for querying Slack.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-sourceconnectorproperties.html#cfn-appflow-flow-sourceconnectorproperties-slack
	//
	Slack interface{} `field:"optional" json:"slack" yaml:"slack"`
	// Specifies the information that is required for querying Trend Micro.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-sourceconnectorproperties.html#cfn-appflow-flow-sourceconnectorproperties-trendmicro
	//
	Trendmicro interface{} `field:"optional" json:"trendmicro" yaml:"trendmicro"`
	// Specifies the information that is required for querying Veeva.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-sourceconnectorproperties.html#cfn-appflow-flow-sourceconnectorproperties-veeva
	//
	Veeva interface{} `field:"optional" json:"veeva" yaml:"veeva"`
	// Specifies the information that is required for querying Zendesk.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-sourceconnectorproperties.html#cfn-appflow-flow-sourceconnectorproperties-zendesk
	//
	Zendesk interface{} `field:"optional" json:"zendesk" yaml:"zendesk"`
}

