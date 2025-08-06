package awsinspectorv2


// The time.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   timeProperty := &TimeProperty{
//   	TimeOfDay: jsii.String("timeOfDay"),
//   	TimeZone: jsii.String("timeZone"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-cisscanconfiguration-time.html
//
type CfnCisScanConfiguration_TimeProperty struct {
	// The time of day in 24-hour format (00:00).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-cisscanconfiguration-time.html#cfn-inspectorv2-cisscanconfiguration-time-timeofday
	//
	TimeOfDay *string `field:"required" json:"timeOfDay" yaml:"timeOfDay"`
	// The timezone.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-cisscanconfiguration-time.html#cfn-inspectorv2-cisscanconfiguration-time-timezone
	//
	TimeZone *string `field:"required" json:"timeZone" yaml:"timeZone"`
}

