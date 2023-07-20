package awsssmcontacts


// Information about on-call rotations that recur monthly.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   monthlySettingProperty := &MonthlySettingProperty{
//   	DayOfMonth: jsii.Number(123),
//   	HandOffTime: jsii.String("handOffTime"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmcontacts-rotation-monthlysetting.html
//
type CfnRotation_MonthlySettingProperty struct {
	// The day of the month when monthly recurring on-call rotations begin.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmcontacts-rotation-monthlysetting.html#cfn-ssmcontacts-rotation-monthlysetting-dayofmonth
	//
	DayOfMonth *float64 `field:"required" json:"dayOfMonth" yaml:"dayOfMonth"`
	// The time of day when a monthly recurring on-call shift rotation begins.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmcontacts-rotation-monthlysetting.html#cfn-ssmcontacts-rotation-monthlysetting-handofftime
	//
	HandOffTime *string `field:"required" json:"handOffTime" yaml:"handOffTime"`
}

