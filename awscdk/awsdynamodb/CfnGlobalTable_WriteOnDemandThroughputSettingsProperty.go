package awsdynamodb


// Sets the write request settings for a global table or a global secondary index.
//
// You must specify this setting if you set the `BillingMode` to `PAY_PER_REQUEST` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   writeOnDemandThroughputSettingsProperty := &WriteOnDemandThroughputSettingsProperty{
//   	MaxWriteRequestUnits: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-writeondemandthroughputsettings.html
//
type CfnGlobalTable_WriteOnDemandThroughputSettingsProperty struct {
	// Maximum number of write request settings for the specified replica of a global table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-writeondemandthroughputsettings.html#cfn-dynamodb-globaltable-writeondemandthroughputsettings-maxwriterequestunits
	//
	MaxWriteRequestUnits *float64 `field:"optional" json:"maxWriteRequestUnits" yaml:"maxWriteRequestUnits"`
}

