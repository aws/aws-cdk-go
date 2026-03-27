package awsxray


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   samplingRateBoostProperty := &SamplingRateBoostProperty{
//   	CooldownWindowMinutes: jsii.Number(123),
//   	MaxRate: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-xray-samplingrule-samplingrateboost.html
//
type CfnSamplingRulePropsMixin_SamplingRateBoostProperty struct {
	// Time window (in minutes) in which only one sampling rate boost can be triggered.
	//
	// After a boost occurs, no further boosts are allowed until the next window.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-xray-samplingrule-samplingrateboost.html#cfn-xray-samplingrule-samplingrateboost-cooldownwindowminutes
	//
	CooldownWindowMinutes *float64 `field:"optional" json:"cooldownWindowMinutes" yaml:"cooldownWindowMinutes"`
	// The maximum sampling rate X-Ray will apply when it detects anomalies.
	//
	// X-Ray determines the appropriate rate between your baseline and the maximum, depending on anomaly activity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-xray-samplingrule-samplingrateboost.html#cfn-xray-samplingrule-samplingrateboost-maxrate
	//
	MaxRate *float64 `field:"optional" json:"maxRate" yaml:"maxRate"`
}

