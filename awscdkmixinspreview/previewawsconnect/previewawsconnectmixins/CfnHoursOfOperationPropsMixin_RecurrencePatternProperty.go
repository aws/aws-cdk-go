package previewawsconnectmixins


// Pattern for recurring hours of operation overrides.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   recurrencePatternProperty := &RecurrencePatternProperty{
//   	ByMonth: []interface{}{
//   		jsii.Number(123),
//   	},
//   	ByMonthDay: []interface{}{
//   		jsii.Number(123),
//   	},
//   	ByWeekdayOccurrence: []interface{}{
//   		jsii.Number(123),
//   	},
//   	Frequency: jsii.String("frequency"),
//   	Interval: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-hoursofoperation-recurrencepattern.html
//
type CfnHoursOfOperationPropsMixin_RecurrencePatternProperty struct {
	// List of months (1-12) for recurrence pattern.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-hoursofoperation-recurrencepattern.html#cfn-connect-hoursofoperation-recurrencepattern-bymonth
	//
	ByMonth interface{} `field:"optional" json:"byMonth" yaml:"byMonth"`
	// List of month days (-1 to 31) for recurrence pattern.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-hoursofoperation-recurrencepattern.html#cfn-connect-hoursofoperation-recurrencepattern-bymonthday
	//
	ByMonthDay interface{} `field:"optional" json:"byMonthDay" yaml:"byMonthDay"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-hoursofoperation-recurrencepattern.html#cfn-connect-hoursofoperation-recurrencepattern-byweekdayoccurrence
	//
	ByWeekdayOccurrence interface{} `field:"optional" json:"byWeekdayOccurrence" yaml:"byWeekdayOccurrence"`
	// The frequency of recurrence for hours of operation overrides.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-hoursofoperation-recurrencepattern.html#cfn-connect-hoursofoperation-recurrencepattern-frequency
	//
	Frequency *string `field:"optional" json:"frequency" yaml:"frequency"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-hoursofoperation-recurrencepattern.html#cfn-connect-hoursofoperation-recurrencepattern-interval
	//
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
}

