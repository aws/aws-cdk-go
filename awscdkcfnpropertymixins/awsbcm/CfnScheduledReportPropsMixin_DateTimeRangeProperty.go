package awsbcm


// An absolute or relative date range.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   dateTimeRangeProperty := &DateTimeRangeProperty{
//   	EndTime: &DateTimeValueProperty{
//   		Type: jsii.String("type"),
//   		Value: jsii.String("value"),
//   	},
//   	StartTime: &DateTimeValueProperty{
//   		Type: jsii.String("type"),
//   		Value: jsii.String("value"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-scheduledreport-datetimerange.html
//
type CfnScheduledReportPropsMixin_DateTimeRangeProperty struct {
	// A date expressed either as an absolute instant or as an offset from now.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-scheduledreport-datetimerange.html#cfn-bcm-scheduledreport-datetimerange-endtime
	//
	EndTime interface{} `field:"optional" json:"endTime" yaml:"endTime"`
	// A date expressed either as an absolute instant or as an offset from now.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-scheduledreport-datetimerange.html#cfn-bcm-scheduledreport-datetimerange-starttime
	//
	StartTime interface{} `field:"optional" json:"startTime" yaml:"startTime"`
}

