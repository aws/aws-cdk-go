package mixinsawsssmcontacts


// Information about when an on-call rotation is in effect and how long the rotation period lasts.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   recurrenceSettingsProperty := &RecurrenceSettingsProperty{
//   	DailySettings: []*string{
//   		jsii.String("dailySettings"),
//   	},
//   	MonthlySettings: []interface{}{
//   		&MonthlySettingProperty{
//   			DayOfMonth: jsii.Number(123),
//   			HandOffTime: jsii.String("handOffTime"),
//   		},
//   	},
//   	NumberOfOnCalls: jsii.Number(123),
//   	RecurrenceMultiplier: jsii.Number(123),
//   	ShiftCoverages: []interface{}{
//   		&ShiftCoverageProperty{
//   			CoverageTimes: []interface{}{
//   				&CoverageTimeProperty{
//   					EndTime: jsii.String("endTime"),
//   					StartTime: jsii.String("startTime"),
//   				},
//   			},
//   			DayOfWeek: jsii.String("dayOfWeek"),
//   		},
//   	},
//   	WeeklySettings: []interface{}{
//   		&WeeklySettingProperty{
//   			DayOfWeek: jsii.String("dayOfWeek"),
//   			HandOffTime: jsii.String("handOffTime"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmcontacts-rotation-recurrencesettings.html
//
type CfnRotationPropsMixin_RecurrenceSettingsProperty struct {
	// Information about on-call rotations that recur daily.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmcontacts-rotation-recurrencesettings.html#cfn-ssmcontacts-rotation-recurrencesettings-dailysettings
	//
	DailySettings *[]*string `field:"optional" json:"dailySettings" yaml:"dailySettings"`
	// Information about on-call rotations that recur monthly.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmcontacts-rotation-recurrencesettings.html#cfn-ssmcontacts-rotation-recurrencesettings-monthlysettings
	//
	MonthlySettings interface{} `field:"optional" json:"monthlySettings" yaml:"monthlySettings"`
	// The number of contacts, or shift team members designated to be on call concurrently during a shift.
	//
	// For example, in an on-call schedule that contains ten contacts, a value of `2` designates that two of them are on call at any given time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmcontacts-rotation-recurrencesettings.html#cfn-ssmcontacts-rotation-recurrencesettings-numberofoncalls
	//
	NumberOfOnCalls *float64 `field:"optional" json:"numberOfOnCalls" yaml:"numberOfOnCalls"`
	// The number of days, weeks, or months a single rotation lasts.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmcontacts-rotation-recurrencesettings.html#cfn-ssmcontacts-rotation-recurrencesettings-recurrencemultiplier
	//
	RecurrenceMultiplier *float64 `field:"optional" json:"recurrenceMultiplier" yaml:"recurrenceMultiplier"`
	// Information about the days of the week included in on-call rotation coverage.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmcontacts-rotation-recurrencesettings.html#cfn-ssmcontacts-rotation-recurrencesettings-shiftcoverages
	//
	ShiftCoverages interface{} `field:"optional" json:"shiftCoverages" yaml:"shiftCoverages"`
	// Information about on-call rotations that recur weekly.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmcontacts-rotation-recurrencesettings.html#cfn-ssmcontacts-rotation-recurrencesettings-weeklysettings
	//
	WeeklySettings interface{} `field:"optional" json:"weeklySettings" yaml:"weeklySettings"`
}

