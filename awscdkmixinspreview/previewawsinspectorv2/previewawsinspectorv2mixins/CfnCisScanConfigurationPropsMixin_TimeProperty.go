package previewawsinspectorv2mixins


// The time.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   timeProperty := &TimeProperty{
//   	TimeOfDay: jsii.String("timeOfDay"),
//   	TimeZone: jsii.String("timeZone"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-cisscanconfiguration-time.html
//
type CfnCisScanConfigurationPropsMixin_TimeProperty struct {
	// The time of day in 24-hour format (00:00).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-cisscanconfiguration-time.html#cfn-inspectorv2-cisscanconfiguration-time-timeofday
	//
	TimeOfDay *string `field:"optional" json:"timeOfDay" yaml:"timeOfDay"`
	// The timezone.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-cisscanconfiguration-time.html#cfn-inspectorv2-cisscanconfiguration-time-timezone
	//
	TimeZone *string `field:"optional" json:"timeZone" yaml:"timeZone"`
}

