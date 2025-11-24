package mixinsawscustomerprofiles


// The configuration that controls how Customer Profiles retrieves data from the source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sourceFlowConfigProperty := &SourceFlowConfigProperty{
//   	ConnectorProfileName: jsii.String("connectorProfileName"),
//   	ConnectorType: jsii.String("connectorType"),
//   	IncrementalPullConfig: &IncrementalPullConfigProperty{
//   		DatetimeTypeFieldName: jsii.String("datetimeTypeFieldName"),
//   	},
//   	SourceConnectorProperties: &SourceConnectorPropertiesProperty{
//   		Marketo: &MarketoSourcePropertiesProperty{
//   			Object: jsii.String("object"),
//   		},
//   		S3: &S3SourcePropertiesProperty{
//   			BucketName: jsii.String("bucketName"),
//   			BucketPrefix: jsii.String("bucketPrefix"),
//   		},
//   		Salesforce: &SalesforceSourcePropertiesProperty{
//   			EnableDynamicFieldUpdate: jsii.Boolean(false),
//   			IncludeDeletedRecords: jsii.Boolean(false),
//   			Object: jsii.String("object"),
//   		},
//   		ServiceNow: &ServiceNowSourcePropertiesProperty{
//   			Object: jsii.String("object"),
//   		},
//   		Zendesk: &ZendeskSourcePropertiesProperty{
//   			Object: jsii.String("object"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-sourceflowconfig.html
//
type CfnIntegrationPropsMixin_SourceFlowConfigProperty struct {
	// The name of the Amazon AppFlow connector profile.
	//
	// This name must be unique for each connector profile in the AWS account .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-sourceflowconfig.html#cfn-customerprofiles-integration-sourceflowconfig-connectorprofilename
	//
	ConnectorProfileName *string `field:"optional" json:"connectorProfileName" yaml:"connectorProfileName"`
	// The type of connector, such as Salesforce, Marketo, and so on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-sourceflowconfig.html#cfn-customerprofiles-integration-sourceflowconfig-connectortype
	//
	ConnectorType *string `field:"optional" json:"connectorType" yaml:"connectorType"`
	// Defines the configuration for a scheduled incremental data pull.
	//
	// If a valid configuration is provided, the fields specified in the configuration are used when querying for the incremental data pull.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-sourceflowconfig.html#cfn-customerprofiles-integration-sourceflowconfig-incrementalpullconfig
	//
	IncrementalPullConfig interface{} `field:"optional" json:"incrementalPullConfig" yaml:"incrementalPullConfig"`
	// Specifies the information that is required to query a particular source connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-sourceflowconfig.html#cfn-customerprofiles-integration-sourceflowconfig-sourceconnectorproperties
	//
	SourceConnectorProperties interface{} `field:"optional" json:"sourceConnectorProperties" yaml:"sourceConnectorProperties"`
}

