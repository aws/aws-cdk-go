package awsappflow


// The connector-specific profile properties required by each connector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectorProfilePropertiesProperty := &ConnectorProfilePropertiesProperty{
//   	CustomConnector: &CustomConnectorProfilePropertiesProperty{
//   		OAuth2Properties: &OAuth2PropertiesProperty{
//   			OAuth2GrantType: jsii.String("oAuth2GrantType"),
//   			TokenUrl: jsii.String("tokenUrl"),
//   			TokenUrlCustomProperties: map[string]*string{
//   				"tokenUrlCustomPropertiesKey": jsii.String("tokenUrlCustomProperties"),
//   			},
//   		},
//   		ProfileProperties: map[string]*string{
//   			"profilePropertiesKey": jsii.String("profileProperties"),
//   		},
//   	},
//   	Datadog: &DatadogConnectorProfilePropertiesProperty{
//   		InstanceUrl: jsii.String("instanceUrl"),
//   	},
//   	Dynatrace: &DynatraceConnectorProfilePropertiesProperty{
//   		InstanceUrl: jsii.String("instanceUrl"),
//   	},
//   	InforNexus: &InforNexusConnectorProfilePropertiesProperty{
//   		InstanceUrl: jsii.String("instanceUrl"),
//   	},
//   	Marketo: &MarketoConnectorProfilePropertiesProperty{
//   		InstanceUrl: jsii.String("instanceUrl"),
//   	},
//   	Pardot: &PardotConnectorProfilePropertiesProperty{
//   		BusinessUnitId: jsii.String("businessUnitId"),
//
//   		// the properties below are optional
//   		InstanceUrl: jsii.String("instanceUrl"),
//   		IsSandboxEnvironment: jsii.Boolean(false),
//   	},
//   	Redshift: &RedshiftConnectorProfilePropertiesProperty{
//   		BucketName: jsii.String("bucketName"),
//   		RoleArn: jsii.String("roleArn"),
//
//   		// the properties below are optional
//   		BucketPrefix: jsii.String("bucketPrefix"),
//   		ClusterIdentifier: jsii.String("clusterIdentifier"),
//   		DataApiRoleArn: jsii.String("dataApiRoleArn"),
//   		DatabaseName: jsii.String("databaseName"),
//   		DatabaseUrl: jsii.String("databaseUrl"),
//   		IsRedshiftServerless: jsii.Boolean(false),
//   		WorkgroupName: jsii.String("workgroupName"),
//   	},
//   	Salesforce: &SalesforceConnectorProfilePropertiesProperty{
//   		InstanceUrl: jsii.String("instanceUrl"),
//   		IsSandboxEnvironment: jsii.Boolean(false),
//   		UsePrivateLinkForMetadataAndAuthorization: jsii.Boolean(false),
//   	},
//   	SapoData: &SAPODataConnectorProfilePropertiesProperty{
//   		ApplicationHostUrl: jsii.String("applicationHostUrl"),
//   		ApplicationServicePath: jsii.String("applicationServicePath"),
//   		ClientNumber: jsii.String("clientNumber"),
//   		DisableSso: jsii.Boolean(false),
//   		LogonLanguage: jsii.String("logonLanguage"),
//   		OAuthProperties: &OAuthPropertiesProperty{
//   			AuthCodeUrl: jsii.String("authCodeUrl"),
//   			OAuthScopes: []*string{
//   				jsii.String("oAuthScopes"),
//   			},
//   			TokenUrl: jsii.String("tokenUrl"),
//   		},
//   		PortNumber: jsii.Number(123),
//   		PrivateLinkServiceName: jsii.String("privateLinkServiceName"),
//   	},
//   	ServiceNow: &ServiceNowConnectorProfilePropertiesProperty{
//   		InstanceUrl: jsii.String("instanceUrl"),
//   	},
//   	Slack: &SlackConnectorProfilePropertiesProperty{
//   		InstanceUrl: jsii.String("instanceUrl"),
//   	},
//   	Snowflake: &SnowflakeConnectorProfilePropertiesProperty{
//   		BucketName: jsii.String("bucketName"),
//   		Stage: jsii.String("stage"),
//   		Warehouse: jsii.String("warehouse"),
//
//   		// the properties below are optional
//   		AccountName: jsii.String("accountName"),
//   		BucketPrefix: jsii.String("bucketPrefix"),
//   		PrivateLinkServiceName: jsii.String("privateLinkServiceName"),
//   		Region: jsii.String("region"),
//   	},
//   	Veeva: &VeevaConnectorProfilePropertiesProperty{
//   		InstanceUrl: jsii.String("instanceUrl"),
//   	},
//   	Zendesk: &ZendeskConnectorProfilePropertiesProperty{
//   		InstanceUrl: jsii.String("instanceUrl"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-connectorprofileproperties.html
//
type CfnConnectorProfile_ConnectorProfilePropertiesProperty struct {
	// The properties required by the custom connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-connectorprofileproperties.html#cfn-appflow-connectorprofile-connectorprofileproperties-customconnector
	//
	CustomConnector interface{} `field:"optional" json:"customConnector" yaml:"customConnector"`
	// The connector-specific properties required by Datadog.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-connectorprofileproperties.html#cfn-appflow-connectorprofile-connectorprofileproperties-datadog
	//
	Datadog interface{} `field:"optional" json:"datadog" yaml:"datadog"`
	// The connector-specific properties required by Dynatrace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-connectorprofileproperties.html#cfn-appflow-connectorprofile-connectorprofileproperties-dynatrace
	//
	Dynatrace interface{} `field:"optional" json:"dynatrace" yaml:"dynatrace"`
	// The connector-specific properties required by Infor Nexus.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-connectorprofileproperties.html#cfn-appflow-connectorprofile-connectorprofileproperties-infornexus
	//
	InforNexus interface{} `field:"optional" json:"inforNexus" yaml:"inforNexus"`
	// The connector-specific properties required by Marketo.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-connectorprofileproperties.html#cfn-appflow-connectorprofile-connectorprofileproperties-marketo
	//
	Marketo interface{} `field:"optional" json:"marketo" yaml:"marketo"`
	// The connector-specific properties required by Salesforce Pardot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-connectorprofileproperties.html#cfn-appflow-connectorprofile-connectorprofileproperties-pardot
	//
	Pardot interface{} `field:"optional" json:"pardot" yaml:"pardot"`
	// The connector-specific properties required by Amazon Redshift.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-connectorprofileproperties.html#cfn-appflow-connectorprofile-connectorprofileproperties-redshift
	//
	Redshift interface{} `field:"optional" json:"redshift" yaml:"redshift"`
	// The connector-specific properties required by Salesforce.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-connectorprofileproperties.html#cfn-appflow-connectorprofile-connectorprofileproperties-salesforce
	//
	Salesforce interface{} `field:"optional" json:"salesforce" yaml:"salesforce"`
	// The connector-specific profile properties required when using SAPOData.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-connectorprofileproperties.html#cfn-appflow-connectorprofile-connectorprofileproperties-sapodata
	//
	SapoData interface{} `field:"optional" json:"sapoData" yaml:"sapoData"`
	// The connector-specific properties required by serviceNow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-connectorprofileproperties.html#cfn-appflow-connectorprofile-connectorprofileproperties-servicenow
	//
	ServiceNow interface{} `field:"optional" json:"serviceNow" yaml:"serviceNow"`
	// The connector-specific properties required by Slack.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-connectorprofileproperties.html#cfn-appflow-connectorprofile-connectorprofileproperties-slack
	//
	Slack interface{} `field:"optional" json:"slack" yaml:"slack"`
	// The connector-specific properties required by Snowflake.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-connectorprofileproperties.html#cfn-appflow-connectorprofile-connectorprofileproperties-snowflake
	//
	Snowflake interface{} `field:"optional" json:"snowflake" yaml:"snowflake"`
	// The connector-specific properties required by Veeva.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-connectorprofileproperties.html#cfn-appflow-connectorprofile-connectorprofileproperties-veeva
	//
	Veeva interface{} `field:"optional" json:"veeva" yaml:"veeva"`
	// The connector-specific properties required by Zendesk.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-connectorprofileproperties.html#cfn-appflow-connectorprofile-connectorprofileproperties-zendesk
	//
	Zendesk interface{} `field:"optional" json:"zendesk" yaml:"zendesk"`
}

