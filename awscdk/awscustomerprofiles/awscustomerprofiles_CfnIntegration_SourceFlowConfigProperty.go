package awscustomerprofiles


// The configuration that controls how Customer Profiles retrieves data from the source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceFlowConfigProperty := &sourceFlowConfigProperty{
//   	connectorType: jsii.String("connectorType"),
//   	sourceConnectorProperties: &sourceConnectorPropertiesProperty{
//   		marketo: &marketoSourcePropertiesProperty{
//   			object: jsii.String("object"),
//   		},
//   		s3: &s3SourcePropertiesProperty{
//   			bucketName: jsii.String("bucketName"),
//
//   			// the properties below are optional
//   			bucketPrefix: jsii.String("bucketPrefix"),
//   		},
//   		salesforce: &salesforceSourcePropertiesProperty{
//   			object: jsii.String("object"),
//
//   			// the properties below are optional
//   			enableDynamicFieldUpdate: jsii.Boolean(false),
//   			includeDeletedRecords: jsii.Boolean(false),
//   		},
//   		serviceNow: &serviceNowSourcePropertiesProperty{
//   			object: jsii.String("object"),
//   		},
//   		zendesk: &zendeskSourcePropertiesProperty{
//   			object: jsii.String("object"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	connectorProfileName: jsii.String("connectorProfileName"),
//   	incrementalPullConfig: &incrementalPullConfigProperty{
//   		datetimeTypeFieldName: jsii.String("datetimeTypeFieldName"),
//   	},
//   }
//
type CfnIntegration_SourceFlowConfigProperty struct {
	// The type of connector, such as Salesforce, Marketo, and so on.
	ConnectorType *string `field:"required" json:"connectorType" yaml:"connectorType"`
	// Specifies the information that is required to query a particular source connector.
	SourceConnectorProperties interface{} `field:"required" json:"sourceConnectorProperties" yaml:"sourceConnectorProperties"`
	// The name of the Amazon AppFlow connector profile.
	//
	// This name must be unique for each connector profile in the AWS account .
	ConnectorProfileName *string `field:"optional" json:"connectorProfileName" yaml:"connectorProfileName"`
	// Defines the configuration for a scheduled incremental data pull.
	//
	// If a valid configuration is provided, the fields specified in the configuration are used when querying for the incremental data pull.
	IncrementalPullConfig interface{} `field:"optional" json:"incrementalPullConfig" yaml:"incrementalPullConfig"`
}

