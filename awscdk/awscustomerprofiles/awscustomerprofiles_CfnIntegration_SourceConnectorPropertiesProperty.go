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
//   sourceConnectorPropertiesProperty := &sourceConnectorPropertiesProperty{
//   	marketo: &marketoSourcePropertiesProperty{
//   		object: jsii.String("object"),
//   	},
//   	s3: &s3SourcePropertiesProperty{
//   		bucketName: jsii.String("bucketName"),
//
//   		// the properties below are optional
//   		bucketPrefix: jsii.String("bucketPrefix"),
//   	},
//   	salesforce: &salesforceSourcePropertiesProperty{
//   		object: jsii.String("object"),
//
//   		// the properties below are optional
//   		enableDynamicFieldUpdate: jsii.Boolean(false),
//   		includeDeletedRecords: jsii.Boolean(false),
//   	},
//   	serviceNow: &serviceNowSourcePropertiesProperty{
//   		object: jsii.String("object"),
//   	},
//   	zendesk: &zendeskSourcePropertiesProperty{
//   		object: jsii.String("object"),
//   	},
//   }
//
type CfnIntegration_SourceConnectorPropertiesProperty struct {
	// The properties that are applied when Marketo is being used as a source.
	Marketo interface{} `field:"optional" json:"marketo" yaml:"marketo"`
	// The properties that are applied when Amazon S3 is being used as the flow source.
	S3 interface{} `field:"optional" json:"s3" yaml:"s3"`
	// The properties that are applied when Salesforce is being used as a source.
	Salesforce interface{} `field:"optional" json:"salesforce" yaml:"salesforce"`
	// The properties that are applied when ServiceNow is being used as a source.
	ServiceNow interface{} `field:"optional" json:"serviceNow" yaml:"serviceNow"`
	// The properties that are applied when using Zendesk as a flow source.
	Zendesk interface{} `field:"optional" json:"zendesk" yaml:"zendesk"`
}

