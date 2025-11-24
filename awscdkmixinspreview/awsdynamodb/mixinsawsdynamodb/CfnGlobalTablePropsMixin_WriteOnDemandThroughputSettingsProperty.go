package mixinsawsdynamodb


// Sets the write request settings for a global table or a global secondary index.
//
// You can only specify this setting if your resource uses the `PAY_PER_REQUEST` `BillingMode` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   writeOnDemandThroughputSettingsProperty := &WriteOnDemandThroughputSettingsProperty{
//   	MaxWriteRequestUnits: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-writeondemandthroughputsettings.html
//
type CfnGlobalTablePropsMixin_WriteOnDemandThroughputSettingsProperty struct {
	// Maximum number of write request settings for the specified replica of a global table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-writeondemandthroughputsettings.html#cfn-dynamodb-globaltable-writeondemandthroughputsettings-maxwriterequestunits
	//
	MaxWriteRequestUnits *float64 `field:"optional" json:"maxWriteRequestUnits" yaml:"maxWriteRequestUnits"`
}

