package awscustomerprofiles


// The properties that are applied when Salesforce is being used as a source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   salesforceSourcePropertiesProperty := &SalesforceSourcePropertiesProperty{
//   	Object: jsii.String("object"),
//
//   	// the properties below are optional
//   	EnableDynamicFieldUpdate: jsii.Boolean(false),
//   	IncludeDeletedRecords: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-salesforcesourceproperties.html
//
type CfnIntegration_SalesforceSourcePropertiesProperty struct {
	// The object specified in the Salesforce flow source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-salesforcesourceproperties.html#cfn-customerprofiles-integration-salesforcesourceproperties-object
	//
	Object *string `field:"required" json:"object" yaml:"object"`
	// The flag that enables dynamic fetching of new (recently added) fields in the Salesforce objects while running a flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-salesforcesourceproperties.html#cfn-customerprofiles-integration-salesforcesourceproperties-enabledynamicfieldupdate
	//
	EnableDynamicFieldUpdate interface{} `field:"optional" json:"enableDynamicFieldUpdate" yaml:"enableDynamicFieldUpdate"`
	// Indicates whether Amazon AppFlow includes deleted files in the flow run.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-salesforcesourceproperties.html#cfn-customerprofiles-integration-salesforcesourceproperties-includedeletedrecords
	//
	IncludeDeletedRecords interface{} `field:"optional" json:"includeDeletedRecords" yaml:"includeDeletedRecords"`
}

