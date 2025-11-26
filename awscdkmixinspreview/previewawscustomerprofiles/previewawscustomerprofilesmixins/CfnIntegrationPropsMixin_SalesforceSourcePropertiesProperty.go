package previewawscustomerprofilesmixins


// The properties that are applied when Salesforce is being used as a source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   salesforceSourcePropertiesProperty := &SalesforceSourcePropertiesProperty{
//   	EnableDynamicFieldUpdate: jsii.Boolean(false),
//   	IncludeDeletedRecords: jsii.Boolean(false),
//   	Object: jsii.String("object"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-salesforcesourceproperties.html
//
type CfnIntegrationPropsMixin_SalesforceSourcePropertiesProperty struct {
	// The flag that enables dynamic fetching of new (recently added) fields in the Salesforce objects while running a flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-salesforcesourceproperties.html#cfn-customerprofiles-integration-salesforcesourceproperties-enabledynamicfieldupdate
	//
	EnableDynamicFieldUpdate interface{} `field:"optional" json:"enableDynamicFieldUpdate" yaml:"enableDynamicFieldUpdate"`
	// Indicates whether Amazon AppFlow includes deleted files in the flow run.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-salesforcesourceproperties.html#cfn-customerprofiles-integration-salesforcesourceproperties-includedeletedrecords
	//
	IncludeDeletedRecords interface{} `field:"optional" json:"includeDeletedRecords" yaml:"includeDeletedRecords"`
	// The object specified in the Salesforce flow source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-salesforcesourceproperties.html#cfn-customerprofiles-integration-salesforcesourceproperties-object
	//
	Object *string `field:"optional" json:"object" yaml:"object"`
}

