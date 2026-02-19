package awsdynamodb


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   globalReadProvisionedThroughputSettingsProperty := &GlobalReadProvisionedThroughputSettingsProperty{
//   	ReadCapacityUnits: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-globalreadprovisionedthroughputsettings.html
//
type CfnGlobalTable_GlobalReadProvisionedThroughputSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-globalreadprovisionedthroughputsettings.html#cfn-dynamodb-globaltable-globalreadprovisionedthroughputsettings-readcapacityunits
	//
	ReadCapacityUnits *float64 `field:"optional" json:"readCapacityUnits" yaml:"readCapacityUnits"`
}

