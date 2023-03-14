package awscustomerprofiles


// The properties that are applied when Salesforce is being used as a source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   salesforceSourcePropertiesProperty := &salesforceSourcePropertiesProperty{
//   	object: jsii.String("object"),
//
//   	// the properties below are optional
//   	enableDynamicFieldUpdate: jsii.Boolean(false),
//   	includeDeletedRecords: jsii.Boolean(false),
//   }
//
type CfnIntegration_SalesforceSourcePropertiesProperty struct {
	// The object specified in the Salesforce flow source.
	Object *string `field:"required" json:"object" yaml:"object"`
	// The flag that enables dynamic fetching of new (recently added) fields in the Salesforce objects while running a flow.
	EnableDynamicFieldUpdate interface{} `field:"optional" json:"enableDynamicFieldUpdate" yaml:"enableDynamicFieldUpdate"`
	// Indicates whether Amazon AppFlow includes deleted files in the flow run.
	IncludeDeletedRecords interface{} `field:"optional" json:"includeDeletedRecords" yaml:"includeDeletedRecords"`
}

