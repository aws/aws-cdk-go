package awsbcm


// A date expressed either as an absolute instant or as an offset from now.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   dateTimeValueProperty := &DateTimeValueProperty{
//   	Type: jsii.String("type"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-scheduledreport-datetimevalue.html
//
type CfnScheduledReportPropsMixin_DateTimeValueProperty struct {
	// Whether a date value is absolute or relative.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-scheduledreport-datetimevalue.html#cfn-bcm-scheduledreport-datetimevalue-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The date, or an ISO 8601 duration when Type is RELATIVE.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-scheduledreport-datetimevalue.html#cfn-bcm-scheduledreport-datetimevalue-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

