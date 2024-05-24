package awsdynamodb


// Sets the read request settings for a replica table or a replica global secondary index.
//
// You must specify this setting if you set the `BillingMode` to `PAY_PER_REQUEST` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   readOnDemandThroughputSettingsProperty := &ReadOnDemandThroughputSettingsProperty{
//   	MaxReadRequestUnits: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-readondemandthroughputsettings.html
//
type CfnGlobalTable_ReadOnDemandThroughputSettingsProperty struct {
	// Maximum number of read request units for the specified replica of a global table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-readondemandthroughputsettings.html#cfn-dynamodb-globaltable-readondemandthroughputsettings-maxreadrequestunits
	//
	MaxReadRequestUnits *float64 `field:"optional" json:"maxReadRequestUnits" yaml:"maxReadRequestUnits"`
}

