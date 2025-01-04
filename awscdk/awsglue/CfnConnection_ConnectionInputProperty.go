package awsglue


// A structure that is used to specify a connection to create or update.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var connectionProperties interface{}
//
//   connectionInputProperty := &ConnectionInputProperty{
//   	ConnectionType: jsii.String("connectionType"),
//
//   	// the properties below are optional
//   	ConnectionProperties: connectionProperties,
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-connectioninput.html
//
type CfnConnection_ConnectionInputProperty struct {
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
	// - `SNAPCHATADS` - Designates a connection to Snapchat Ads.
	// - `STRIPE` - Designates a connection to Stripe.
	// - `ZENDESK` - Designates a connection to Zendesk.
	// - `ZOHOCRM` - Designates a connection to Zoho CRM.
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
	ConnectionType *string `field:"required" json:"connectionType" yaml:"connectionType"`
	// These key-value pairs define parameters for the connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-connectioninput.html#cfn-glue-connection-connectioninput-connectionproperties
	//
	ConnectionProperties interface{} `field:"optional" json:"connectionProperties" yaml:"connectionProperties"`
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
}

