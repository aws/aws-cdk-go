package awsbcm


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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-dashboard-datetimerange.html
//
type CfnDashboardPropsMixin_DateTimeRangeProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-dashboard-datetimerange.html#cfn-bcm-dashboard-datetimerange-endtime
	//
	EndTime interface{} `field:"optional" json:"endTime" yaml:"endTime"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-dashboard-datetimerange.html#cfn-bcm-dashboard-datetimerange-starttime
	//
	StartTime interface{} `field:"optional" json:"startTime" yaml:"startTime"`
}

