package awsbedrock


// Day of the month on which a monthly refresh runs.
//
// Exactly one variant is set: an explicit day number, or the last calendar day of the month.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   var lastDayOfMonth interface{}
//
//   dayOfMonthProperty := &DayOfMonthProperty{
//   	DayNumber: jsii.Number(123),
//   	LastDayOfMonth: lastDayOfMonth,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-dayofmonth.html
//
type CfnDataSourcePropsMixin_DayOfMonthProperty struct {
	// Specific day of the month, 1 through 28 (capped at 28 to avoid month-length ambiguity).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-dayofmonth.html#cfn-bedrock-datasource-dayofmonth-daynumber
	//
	DayNumber *float64 `field:"optional" json:"dayNumber" yaml:"dayNumber"`
	// Run on the last calendar day of each month.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-dayofmonth.html#cfn-bedrock-datasource-dayofmonth-lastdayofmonth
	//
	LastDayOfMonth interface{} `field:"optional" json:"lastDayOfMonth" yaml:"lastDayOfMonth"`
}

