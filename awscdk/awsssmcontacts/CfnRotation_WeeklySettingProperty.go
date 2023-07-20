package awsssmcontacts


// Information about rotations that recur weekly.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   weeklySettingProperty := &WeeklySettingProperty{
//   	DayOfWeek: jsii.String("dayOfWeek"),
//   	HandOffTime: jsii.String("handOffTime"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmcontacts-rotation-weeklysetting.html
//
type CfnRotation_WeeklySettingProperty struct {
	// The day of the week when weekly recurring on-call shift rotations begins.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmcontacts-rotation-weeklysetting.html#cfn-ssmcontacts-rotation-weeklysetting-dayofweek
	//
	DayOfWeek *string `field:"required" json:"dayOfWeek" yaml:"dayOfWeek"`
	// The time of day when a weekly recurring on-call shift rotation begins.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmcontacts-rotation-weeklysetting.html#cfn-ssmcontacts-rotation-weeklysetting-handofftime
	//
	HandOffTime *string `field:"required" json:"handOffTime" yaml:"handOffTime"`
}

