package previewawscloudtrailmixins


// Specifies the frequency for a dashboard refresh schedule.
//
// For a custom dashboard, you can schedule a refresh for every 1, 6, 12, or 24 hours, or every day.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   frequencyProperty := &FrequencyProperty{
//   	Unit: jsii.String("unit"),
//   	Value: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-dashboard-frequency.html
//
type CfnDashboardPropsMixin_FrequencyProperty struct {
	// The unit to use for the refresh.
	//
	// For custom dashboards, the unit can be `HOURS` or `DAYS` .
	//
	// For the Highlights dashboard, the `Unit` must be `HOURS` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-dashboard-frequency.html#cfn-cloudtrail-dashboard-frequency-unit
	//
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
	// The value for the refresh schedule.
	//
	// For custom dashboards, the following values are valid when the unit is `HOURS` : `1` , `6` , `12` , `24`
	//
	// For custom dashboards, the only valid value when the unit is `DAYS` is `1` .
	//
	// For the Highlights dashboard, the `Value` must be `6` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-dashboard-frequency.html#cfn-cloudtrail-dashboard-frequency-value
	//
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

