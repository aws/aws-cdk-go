package awscustomerprofiles


// Specifies the information that is required to query a particular Amazon AppFlow connector.
//
// Customer Profiles supports Salesforce, Zendesk, Marketo, ServiceNow and Amazon S3.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceConnectorPropertiesProperty := &SourceConnectorPropertiesProperty{
//   	Marketo: &MarketoSourcePropertiesProperty{
//   		Object: jsii.String("object"),
//   	},
//   	S3: &S3SourcePropertiesProperty{
//   		BucketName: jsii.String("bucketName"),
//
//   		// the properties below are optional
//   		BucketPrefix: jsii.String("bucketPrefix"),
//   	},
//   	Salesforce: &SalesforceSourcePropertiesProperty{
//   		Object: jsii.String("object"),
//
//   		// the properties below are optional
//   		EnableDynamicFieldUpdate: jsii.Boolean(false),
//   		IncludeDeletedRecords: jsii.Boolean(false),
//   	},
//   	ServiceNow: &ServiceNowSourcePropertiesProperty{
//   		Object: jsii.String("object"),
//   	},
//   	Zendesk: &ZendeskSourcePropertiesProperty{
//   		Object: jsii.String("object"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-sourceconnectorproperties.html
//
type CfnIntegration_SourceConnectorPropertiesProperty struct {
	// The properties that are applied when Marketo is being used as a source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-sourceconnectorproperties.html#cfn-customerprofiles-integration-sourceconnectorproperties-marketo
	//
	Marketo interface{} `field:"optional" json:"marketo" yaml:"marketo"`
	// The properties that are applied when Amazon S3 is being used as the flow source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-sourceconnectorproperties.html#cfn-customerprofiles-integration-sourceconnectorproperties-s3
	//
	S3 interface{} `field:"optional" json:"s3" yaml:"s3"`
	// The properties that are applied when Salesforce is being used as a source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-sourceconnectorproperties.html#cfn-customerprofiles-integration-sourceconnectorproperties-salesforce
	//
	Salesforce interface{} `field:"optional" json:"salesforce" yaml:"salesforce"`
	// The properties that are applied when ServiceNow is being used as a source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-sourceconnectorproperties.html#cfn-customerprofiles-integration-sourceconnectorproperties-servicenow
	//
	ServiceNow interface{} `field:"optional" json:"serviceNow" yaml:"serviceNow"`
	// The properties that are applied when using Zendesk as a flow source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-sourceconnectorproperties.html#cfn-customerprofiles-integration-sourceconnectorproperties-zendesk
	//
	Zendesk interface{} `field:"optional" json:"zendesk" yaml:"zendesk"`
}

