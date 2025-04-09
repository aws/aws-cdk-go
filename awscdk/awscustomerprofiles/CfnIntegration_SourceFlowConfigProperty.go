package awscustomerprofiles


// The configuration that controls how Customer Profiles retrieves data from the source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceFlowConfigProperty := &SourceFlowConfigProperty{
//   	ConnectorType: jsii.String("connectorType"),
//   	SourceConnectorProperties: &SourceConnectorPropertiesProperty{
//   		Marketo: &MarketoSourcePropertiesProperty{
//   			Object: jsii.String("object"),
//   		},
//   		S3: &S3SourcePropertiesProperty{
//   			BucketName: jsii.String("bucketName"),
//
//   			// the properties below are optional
//   			BucketPrefix: jsii.String("bucketPrefix"),
//   		},
//   		Salesforce: &SalesforceSourcePropertiesProperty{
//   			Object: jsii.String("object"),
//
//   			// the properties below are optional
//   			EnableDynamicFieldUpdate: jsii.Boolean(false),
//   			IncludeDeletedRecords: jsii.Boolean(false),
//   		},
//   		ServiceNow: &ServiceNowSourcePropertiesProperty{
//   			Object: jsii.String("object"),
//   		},
//   		Zendesk: &ZendeskSourcePropertiesProperty{
//   			Object: jsii.String("object"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	ConnectorProfileName: jsii.String("connectorProfileName"),
//   	IncrementalPullConfig: &IncrementalPullConfigProperty{
//   		DatetimeTypeFieldName: jsii.String("datetimeTypeFieldName"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-sourceflowconfig.html
//
type CfnIntegration_SourceFlowConfigProperty struct {
	// The type of connector, such as Salesforce, Marketo, and so on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-sourceflowconfig.html#cfn-customerprofiles-integration-sourceflowconfig-connectortype
	//
	ConnectorType *string `field:"required" json:"connectorType" yaml:"connectorType"`
	// Specifies the information that is required to query a particular source connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-sourceflowconfig.html#cfn-customerprofiles-integration-sourceflowconfig-sourceconnectorproperties
	//
	SourceConnectorProperties interface{} `field:"required" json:"sourceConnectorProperties" yaml:"sourceConnectorProperties"`
	// The name of the Amazon AppFlow connector profile.
	//
	// This name must be unique for each connector profile in the AWS account .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-sourceflowconfig.html#cfn-customerprofiles-integration-sourceflowconfig-connectorprofilename
	//
	ConnectorProfileName *string `field:"optional" json:"connectorProfileName" yaml:"connectorProfileName"`
	// Defines the configuration for a scheduled incremental data pull.
	//
	// If a valid configuration is provided, the fields specified in the configuration are used when querying for the incremental data pull.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-sourceflowconfig.html#cfn-customerprofiles-integration-sourceflowconfig-incrementalpullconfig
	//
	IncrementalPullConfig interface{} `field:"optional" json:"incrementalPullConfig" yaml:"incrementalPullConfig"`
}

