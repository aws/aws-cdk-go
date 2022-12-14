package awsappflow


// The connector-specific profile properties required by each connector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectorProfilePropertiesProperty := &connectorProfilePropertiesProperty{
//   	customConnector: &customConnectorProfilePropertiesProperty{
//   		oAuth2Properties: &oAuth2PropertiesProperty{
//   			oAuth2GrantType: jsii.String("oAuth2GrantType"),
//   			tokenUrl: jsii.String("tokenUrl"),
//   			tokenUrlCustomProperties: map[string]*string{
//   				"tokenUrlCustomPropertiesKey": jsii.String("tokenUrlCustomProperties"),
//   			},
//   		},
//   		profileProperties: map[string]*string{
//   			"profilePropertiesKey": jsii.String("profileProperties"),
//   		},
//   	},
//   	datadog: &datadogConnectorProfilePropertiesProperty{
//   		instanceUrl: jsii.String("instanceUrl"),
//   	},
//   	dynatrace: &dynatraceConnectorProfilePropertiesProperty{
//   		instanceUrl: jsii.String("instanceUrl"),
//   	},
//   	inforNexus: &inforNexusConnectorProfilePropertiesProperty{
//   		instanceUrl: jsii.String("instanceUrl"),
//   	},
//   	marketo: &marketoConnectorProfilePropertiesProperty{
//   		instanceUrl: jsii.String("instanceUrl"),
//   	},
//   	redshift: &redshiftConnectorProfilePropertiesProperty{
//   		bucketName: jsii.String("bucketName"),
//   		roleArn: jsii.String("roleArn"),
//
//   		// the properties below are optional
//   		bucketPrefix: jsii.String("bucketPrefix"),
//   		clusterIdentifier: jsii.String("clusterIdentifier"),
//   		dataApiRoleArn: jsii.String("dataApiRoleArn"),
//   		databaseName: jsii.String("databaseName"),
//   		databaseUrl: jsii.String("databaseUrl"),
//   		isRedshiftServerless: jsii.Boolean(false),
//   		workgroupName: jsii.String("workgroupName"),
//   	},
//   	salesforce: &salesforceConnectorProfilePropertiesProperty{
//   		instanceUrl: jsii.String("instanceUrl"),
//   		isSandboxEnvironment: jsii.Boolean(false),
//   	},
//   	sapoData: &sAPODataConnectorProfilePropertiesProperty{
//   		applicationHostUrl: jsii.String("applicationHostUrl"),
//   		applicationServicePath: jsii.String("applicationServicePath"),
//   		clientNumber: jsii.String("clientNumber"),
//   		logonLanguage: jsii.String("logonLanguage"),
//   		oAuthProperties: &oAuthPropertiesProperty{
//   			authCodeUrl: jsii.String("authCodeUrl"),
//   			oAuthScopes: []*string{
//   				jsii.String("oAuthScopes"),
//   			},
//   			tokenUrl: jsii.String("tokenUrl"),
//   		},
//   		portNumber: jsii.Number(123),
//   		privateLinkServiceName: jsii.String("privateLinkServiceName"),
//   	},
//   	serviceNow: &serviceNowConnectorProfilePropertiesProperty{
//   		instanceUrl: jsii.String("instanceUrl"),
//   	},
//   	slack: &slackConnectorProfilePropertiesProperty{
//   		instanceUrl: jsii.String("instanceUrl"),
//   	},
//   	snowflake: &snowflakeConnectorProfilePropertiesProperty{
//   		bucketName: jsii.String("bucketName"),
//   		stage: jsii.String("stage"),
//   		warehouse: jsii.String("warehouse"),
//
//   		// the properties below are optional
//   		accountName: jsii.String("accountName"),
//   		bucketPrefix: jsii.String("bucketPrefix"),
//   		privateLinkServiceName: jsii.String("privateLinkServiceName"),
//   		region: jsii.String("region"),
//   	},
//   	veeva: &veevaConnectorProfilePropertiesProperty{
//   		instanceUrl: jsii.String("instanceUrl"),
//   	},
//   	zendesk: &zendeskConnectorProfilePropertiesProperty{
//   		instanceUrl: jsii.String("instanceUrl"),
//   	},
//   }
//
type CfnConnectorProfile_ConnectorProfilePropertiesProperty struct {
	// `CfnConnectorProfile.ConnectorProfilePropertiesProperty.CustomConnector`.
	CustomConnector interface{} `field:"optional" json:"customConnector" yaml:"customConnector"`
	// The connector-specific properties required by Datadog.
	Datadog interface{} `field:"optional" json:"datadog" yaml:"datadog"`
	// The connector-specific properties required by Dynatrace.
	Dynatrace interface{} `field:"optional" json:"dynatrace" yaml:"dynatrace"`
	// The connector-specific properties required by Infor Nexus.
	InforNexus interface{} `field:"optional" json:"inforNexus" yaml:"inforNexus"`
	// The connector-specific properties required by Marketo.
	Marketo interface{} `field:"optional" json:"marketo" yaml:"marketo"`
	// The connector-specific properties required by Amazon Redshift.
	Redshift interface{} `field:"optional" json:"redshift" yaml:"redshift"`
	// The connector-specific properties required by Salesforce.
	Salesforce interface{} `field:"optional" json:"salesforce" yaml:"salesforce"`
	// The connector-specific profile properties required when using SAPOData.
	SapoData interface{} `field:"optional" json:"sapoData" yaml:"sapoData"`
	// The connector-specific properties required by serviceNow.
	ServiceNow interface{} `field:"optional" json:"serviceNow" yaml:"serviceNow"`
	// The connector-specific properties required by Slack.
	Slack interface{} `field:"optional" json:"slack" yaml:"slack"`
	// The connector-specific properties required by Snowflake.
	Snowflake interface{} `field:"optional" json:"snowflake" yaml:"snowflake"`
	// The connector-specific properties required by Veeva.
	Veeva interface{} `field:"optional" json:"veeva" yaml:"veeva"`
	// The connector-specific properties required by Zendesk.
	Zendesk interface{} `field:"optional" json:"zendesk" yaml:"zendesk"`
}

