package mixinsawsglue


// A structure that is used to specify a connection to create or update.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var athenaProperties interface{}
//   var connectionProperties interface{}
//   var customAuthenticationCredentials interface{}
//   var pythonProperties interface{}
//   var sparkProperties interface{}
//   var tokenUrlParametersMap interface{}
//
//   connectionInputProperty := &ConnectionInputProperty{
//   	AthenaProperties: athenaProperties,
//   	AuthenticationConfiguration: &AuthenticationConfigurationInputProperty{
//   		AuthenticationType: jsii.String("authenticationType"),
//   		BasicAuthenticationCredentials: &BasicAuthenticationCredentialsProperty{
//   			Password: jsii.String("password"),
//   			Username: jsii.String("username"),
//   		},
//   		CustomAuthenticationCredentials: customAuthenticationCredentials,
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
//   		OAuth2Properties: &OAuth2PropertiesInputProperty{
//   			AuthorizationCodeProperties: &AuthorizationCodePropertiesProperty{
//   				AuthorizationCode: jsii.String("authorizationCode"),
//   				RedirectUri: jsii.String("redirectUri"),
//   			},
//   			OAuth2ClientApplication: &OAuth2ClientApplicationProperty{
//   				AwsManagedClientApplicationReference: jsii.String("awsManagedClientApplicationReference"),
//   				UserManagedClientApplicationClientId: jsii.String("userManagedClientApplicationClientId"),
//   			},
//   			OAuth2Credentials: &OAuth2CredentialsProperty{
//   				AccessToken: jsii.String("accessToken"),
//   				JwtToken: jsii.String("jwtToken"),
//   				RefreshToken: jsii.String("refreshToken"),
//   				UserManagedClientApplicationClientSecret: jsii.String("userManagedClientApplicationClientSecret"),
//   			},
//   			OAuth2GrantType: jsii.String("oAuth2GrantType"),
//   			TokenUrl: jsii.String("tokenUrl"),
//   			TokenUrlParametersMap: tokenUrlParametersMap,
//   		},
//   		SecretArn: jsii.String("secretArn"),
//   	},
//   	ConnectionProperties: connectionProperties,
//   	ConnectionType: jsii.String("connectionType"),
//   	Description: jsii.String("description"),
//   	MatchCriteria: []*string{
//   		jsii.String("matchCriteria"),
//   	},
//   	Name: jsii.String("name"),
//   	PhysicalConnectionRequirements: &PhysicalConnectionRequirementsProperty{
//   		AvailabilityZone: jsii.String("availabilityZone"),
//   		SecurityGroupIdList: []*string{
//   			jsii.String("securityGroupIdList"),
//   		},
//   		SubnetId: jsii.String("subnetId"),
//   	},
//   	PythonProperties: pythonProperties,
//   	SparkProperties: sparkProperties,
//   	ValidateCredentials: jsii.Boolean(false),
//   	ValidateForComputeEnvironments: []*string{
//   		jsii.String("validateForComputeEnvironments"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-connectioninput.html
//
type CfnConnectionPropsMixin_ConnectionInputProperty struct {
	// Connection properties specific to the Athena compute environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-connectioninput.html#cfn-glue-connection-connectioninput-athenaproperties
	//
	AthenaProperties interface{} `field:"optional" json:"athenaProperties" yaml:"athenaProperties"`
	// The authentication properties of the connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-connectioninput.html#cfn-glue-connection-connectioninput-authenticationconfiguration
	//
	AuthenticationConfiguration interface{} `field:"optional" json:"authenticationConfiguration" yaml:"authenticationConfiguration"`
	// These key-value pairs define parameters for the connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-connectioninput.html#cfn-glue-connection-connectioninput-connectionproperties
	//
	ConnectionProperties interface{} `field:"optional" json:"connectionProperties" yaml:"connectionProperties"`
	// The type of the connection. Currently, these types are supported:.
	//
	// - `JDBC` - Designates a connection to a database through Java Database Connectivity (JDBC).
	//
	// `JDBC` Connections use the following ConnectionParameters.
	//
	// - Required: All of ( `HOST` , `PORT` , `JDBC_ENGINE` ) or `JDBC_CONNECTION_URL` .
	// - Required: All of ( `USERNAME` , `PASSWORD` ) or `SECRET_ID` .
	// - Optional: `JDBC_ENFORCE_SSL` , `CUSTOM_JDBC_CERT` , `CUSTOM_JDBC_CERT_STRING` , `SKIP_CUSTOM_JDBC_CERT_VALIDATION` . These parameters are used to configure SSL with JDBC.
	// - `KAFKA` - Designates a connection to an Apache Kafka streaming platform.
	//
	// `KAFKA` Connections use the following ConnectionParameters.
	//
	// - Required: `KAFKA_BOOTSTRAP_SERVERS` .
	// - Optional: `KAFKA_SSL_ENABLED` , `KAFKA_CUSTOM_CERT` , `KAFKA_SKIP_CUSTOM_CERT_VALIDATION` . These parameters are used to configure SSL with `KAFKA` .
	// - Optional: `KAFKA_CLIENT_KEYSTORE` , `KAFKA_CLIENT_KEYSTORE_PASSWORD` , `KAFKA_CLIENT_KEY_PASSWORD` , `ENCRYPTED_KAFKA_CLIENT_KEYSTORE_PASSWORD` , `ENCRYPTED_KAFKA_CLIENT_KEY_PASSWORD` . These parameters are used to configure TLS client configuration with SSL in `KAFKA` .
	// - Optional: `KAFKA_SASL_MECHANISM` . Can be specified as `SCRAM-SHA-512` , `GSSAPI` , or `AWS_MSK_IAM` .
	// - Optional: `KAFKA_SASL_SCRAM_USERNAME` , `KAFKA_SASL_SCRAM_PASSWORD` , `ENCRYPTED_KAFKA_SASL_SCRAM_PASSWORD` . These parameters are used to configure SASL/SCRAM-SHA-512 authentication with `KAFKA` .
	// - Optional: `KAFKA_SASL_GSSAPI_KEYTAB` , `KAFKA_SASL_GSSAPI_KRB5_CONF` , `KAFKA_SASL_GSSAPI_SERVICE` , `KAFKA_SASL_GSSAPI_PRINCIPAL` . These parameters are used to configure SASL/GSSAPI authentication with `KAFKA` .
	// - `MONGODB` - Designates a connection to a MongoDB document database.
	//
	// `MONGODB` Connections use the following ConnectionParameters.
	//
	// - Required: `CONNECTION_URL` .
	// - Required: All of ( `USERNAME` , `PASSWORD` ) or `SECRET_ID` .
	// - `VIEW_VALIDATION_REDSHIFT` - Designates a connection used for view validation by Amazon Redshift.
	// - `VIEW_VALIDATION_ATHENA` - Designates a connection used for view validation by Amazon Athena.
	// - `NETWORK` - Designates a network connection to a data source within an Amazon Virtual Private Cloud environment (Amazon VPC).
	//
	// `NETWORK` Connections do not require ConnectionParameters. Instead, provide a PhysicalConnectionRequirements.
	// - `MARKETPLACE` - Uses configuration settings contained in a connector purchased from AWS Marketplace to read from and write to data stores that are not natively supported by AWS Glue .
	//
	// `MARKETPLACE` Connections use the following ConnectionParameters.
	//
	// - Required: `CONNECTOR_TYPE` , `CONNECTOR_URL` , `CONNECTOR_CLASS_NAME` , `CONNECTION_URL` .
	// - Required for `JDBC` `CONNECTOR_TYPE` connections: All of ( `USERNAME` , `PASSWORD` ) or `SECRET_ID` .
	// - `CUSTOM` - Uses configuration settings contained in a custom connector to read from and write to data stores that are not natively supported by AWS Glue .
	//
	// Additionally, a `ConnectionType` for the following SaaS connectors is supported:
	//
	// - `FACEBOOKADS` - Designates a connection to Facebook Ads.
	// - `GOOGLEADS` - Designates a connection to Google Ads.
	// - `GOOGLESHEETS` - Designates a connection to Google Sheets.
	// - `GOOGLEANALYTICS4` - Designates a connection to Google Analytics 4.
	// - `HUBSPOT` - Designates a connection to HubSpot.
	// - `INSTAGRAMADS` - Designates a connection to Instagram Ads.
	// - `INTERCOM` - Designates a connection to Intercom.
	// - `JIRACLOUD` - Designates a connection to Jira Cloud.
	// - `MARKETO` - Designates a connection to Adobe Marketo Engage.
	// - `NETSUITEERP` - Designates a connection to Oracle NetSuite.
	// - `SALESFORCE` - Designates a connection to Salesforce using OAuth authentication.
	// - `SALESFORCEMARKETINGCLOUD` - Designates a connection to Salesforce Marketing Cloud.
	// - `SALESFORCEPARDOT` - Designates a connection to Salesforce Marketing Cloud Account Engagement (MCAE).
	// - `SAPODATA` - Designates a connection to SAP OData.
	// - `SERVICENOW` - Designates a connection to ServiceNow.
	// - `SLACK` - Designates a connection to Slack.
	// - `SNOWFLAKE` - Designates a connection to Snowflake.
	// - `SNAPCHATADS` - Designates a connection to Snapchat Ads.
	// - `STRIPE` - Designates a connection to Stripe.
	// - `ZENDESK` - Designates a connection to Zendesk.
	// - `ZOHOCRM` - Designates a connection to Zoho CRM.
	// - `ADOBEANALYTICS` - Designates a connection to Adobe Analytics.
	// - `LINKEDIN` - Designates a connection to LinkedIn.
	// - `MIXPANEL` - Designates a connection to Mixpanel.
	// - `ASANA` - Designates a connection to Asana.
	// - `SMARTSHEET` - Designates a connection to Smartsheet.
	// - `DATADOG` - Designates a connection to Datadog.
	// - `WOOCOMMERCE` - Designates a connection to WooCommerce.
	// - `PAYPAL` - Designates a connection to PayPal.
	// - `QUICKBOOKS` - Designates a connection to QuickBooks.
	// - `FACEBOOKPAGEINSIGHTS` - Designates a connection to Facebook Page Insights.
	// - `FRESHDESK` - Designates a connection to Freshdesk.
	// - `TWILIO` - Designates a connection to Twilio.
	// - `DOCUSIGNMONITOR` - Designates a connection to DocuSign Monitor.
	// - `FRESHSALES` - Designates a connection to Freshsales.
	// - `ZOOM` - Designates a connection to Zoom.
	// - `GOOGLESEARCHCONSOLE` - Designates a connection to Google Search Console.
	// - `SALESFORCECOMMERCECLOUD` - Designates a connection to Salesforce Commerce Cloud.
	// - `SAPCONCUR` - Designates a connection to SAP Concur.
	// - `DYNATRACE` - Designates a connection to Dynatrace.
	// - `MICROSOFTDYNAMIC365FINANCEANDOPS` - Designates a connection to Microsoft Dynamics 365 Finance and Operations.
	// - `MICROSOFTTEAMS` - Designates a connection to Microsoft Teams.
	// - `BLACKBAUDRAISEREDGENXT` - Designates a connection to Blackbaud Raiser's Edge NXT.
	// - `MAILCHIMP` - Designates a connection to Mailchimp.
	// - `GITLAB` - Designates a connection to GitLab.
	// - `PENDO` - Designates a connection to Pendo.
	// - `PRODUCTBOARD` - Designates a connection to Productboard.
	// - `CIRCLECI` - Designates a connection to CircleCI.
	// - `PIPEDIVE` - Designates a connection to Pipedrive.
	// - `SENDGRID` - Designates a connection to SendGrid.
	//
	// For more information on the connection parameters needed for a particular connector, see the documentation for the connector in [Adding an AWS Glue connection](https://docs.aws.amazon.com/glue/latest/dg/console-connections.html) in the AWS Glue User Guide.
	//
	// `SFTP` is not supported.
	//
	// For more information about how optional ConnectionProperties are used to configure features in AWS Glue , consult [AWS Glue connection properties](https://docs.aws.amazon.com/glue/latest/dg/connection-defining.html) .
	//
	// For more information about how optional ConnectionProperties are used to configure features in AWS Glue Studio, consult [Using connectors and connections](https://docs.aws.amazon.com/glue/latest/ug/connectors-chapter.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-connectioninput.html#cfn-glue-connection-connectioninput-connectiontype
	//
	ConnectionType *string `field:"optional" json:"connectionType" yaml:"connectionType"`
	// The description of the connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-connectioninput.html#cfn-glue-connection-connectioninput-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A list of criteria that can be used in selecting this connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-connectioninput.html#cfn-glue-connection-connectioninput-matchcriteria
	//
	MatchCriteria *[]*string `field:"optional" json:"matchCriteria" yaml:"matchCriteria"`
	// The name of the connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-connectioninput.html#cfn-glue-connection-connectioninput-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The physical connection requirements, such as virtual private cloud (VPC) and `SecurityGroup` , that are needed to successfully make this connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-connectioninput.html#cfn-glue-connection-connectioninput-physicalconnectionrequirements
	//
	PhysicalConnectionRequirements interface{} `field:"optional" json:"physicalConnectionRequirements" yaml:"physicalConnectionRequirements"`
	// Connection properties specific to the Python compute environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-connectioninput.html#cfn-glue-connection-connectioninput-pythonproperties
	//
	PythonProperties interface{} `field:"optional" json:"pythonProperties" yaml:"pythonProperties"`
	// Connection properties specific to the Spark compute environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-connectioninput.html#cfn-glue-connection-connectioninput-sparkproperties
	//
	SparkProperties interface{} `field:"optional" json:"sparkProperties" yaml:"sparkProperties"`
	// A flag to validate the credentials during create connection.
	//
	// Default is true.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-connectioninput.html#cfn-glue-connection-connectioninput-validatecredentials
	//
	ValidateCredentials interface{} `field:"optional" json:"validateCredentials" yaml:"validateCredentials"`
	// The compute environments that the specified connection properties are validated against.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-connectioninput.html#cfn-glue-connection-connectioninput-validateforcomputeenvironments
	//
	ValidateForComputeEnvironments *[]*string `field:"optional" json:"validateForComputeEnvironments" yaml:"validateForComputeEnvironments"`
}

