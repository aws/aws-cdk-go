package awsbraket


// Defines a time range for spending limits, specifying when the limit is active.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   timePeriodProperty := &TimePeriodProperty{
//   	EndAt: jsii.String("endAt"),
//   	StartAt: jsii.String("startAt"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-braket-spendinglimit-timeperiod.html
//
type CfnSpendingLimitPropsMixin_TimePeriodProperty struct {
	// The end date and time for the spending limit period, in ISO 8601 format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-braket-spendinglimit-timeperiod.html#cfn-braket-spendinglimit-timeperiod-endat
	//
	EndAt *string `field:"optional" json:"endAt" yaml:"endAt"`
	// The start date and time for the spending limit period, in ISO 8601 format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-braket-spendinglimit-timeperiod.html#cfn-braket-spendinglimit-timeperiod-startat
	//
	StartAt *string `field:"optional" json:"startAt" yaml:"startAt"`
}

