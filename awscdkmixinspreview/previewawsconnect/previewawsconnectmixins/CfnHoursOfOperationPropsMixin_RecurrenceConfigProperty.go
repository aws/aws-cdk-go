package previewawsconnectmixins


// Configuration for recurring hours of operation overrides.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   recurrenceConfigProperty := &RecurrenceConfigProperty{
//   	RecurrencePattern: &RecurrencePatternProperty{
//   		ByMonth: []interface{}{
//   			jsii.Number(123),
//   		},
//   		ByMonthDay: []interface{}{
//   			jsii.Number(123),
//   		},
//   		ByWeekdayOccurrence: []interface{}{
//   			jsii.Number(123),
//   		},
//   		Frequency: jsii.String("frequency"),
//   		Interval: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-hoursofoperation-recurrenceconfig.html
//
type CfnHoursOfOperationPropsMixin_RecurrenceConfigProperty struct {
	// Pattern for recurring hours of operation overrides.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-hoursofoperation-recurrenceconfig.html#cfn-connect-hoursofoperation-recurrenceconfig-recurrencepattern
	//
	RecurrencePattern interface{} `field:"optional" json:"recurrencePattern" yaml:"recurrencePattern"`
}

