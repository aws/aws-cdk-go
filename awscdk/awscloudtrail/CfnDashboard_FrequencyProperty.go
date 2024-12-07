package awscloudtrail


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   frequencyProperty := &FrequencyProperty{
//   	Unit: jsii.String("unit"),
//   	Value: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-dashboard-frequency.html
//
type CfnDashboard_FrequencyProperty struct {
	// The frequency unit.
	//
	// Supported values are HOURS and DAYS.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-dashboard-frequency.html#cfn-cloudtrail-dashboard-frequency-unit
	//
	Unit *string `field:"required" json:"unit" yaml:"unit"`
	// The frequency value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-dashboard-frequency.html#cfn-cloudtrail-dashboard-frequency-value
	//
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

