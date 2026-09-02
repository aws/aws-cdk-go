package awscloudwatch


// The warm-up configuration for an alarm.
//
// A warm-up period delays alarm evaluation after you create or update the alarm. This reduces alarm noise from missing data while a new resource or service starts up. During the warm-up period, the alarm stays in INSUFFICIENT_DATA and doesn't perform alarm actions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   warmUpConfigurationProperty := &WarmUpConfigurationProperty{
//   	OnlyStartEvaluatingAfterWarmUpPeriodEnds: jsii.Boolean(false),
//   	WarmUpPeriodDurationInMinutes: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-logalarm-warmupconfiguration.html
//
type CfnLogAlarmPropsMixin_WarmUpConfigurationProperty struct {
	// Specifies whether the alarm waits for the full warm-up period before it starts evaluating.
	//
	// If true, the alarm waits the entire WarmUpPeriodDurationInMinutes before it starts evaluating, even if metric data arrives earlier. If false, the alarm ends the warm-up period early and starts evaluating as soon as it has enough metric data to fill its evaluation window. This is the default behavior.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-logalarm-warmupconfiguration.html#cfn-cloudwatch-logalarm-warmupconfiguration-onlystartevaluatingafterwarmupperiodends
	//
	OnlyStartEvaluatingAfterWarmUpPeriodEnds interface{} `field:"optional" json:"onlyStartEvaluatingAfterWarmUpPeriodEnds" yaml:"onlyStartEvaluatingAfterWarmUpPeriodEnds"`
	// The length of the warm-up period, in minutes.
	//
	// For this duration after you create or update the alarm, the alarm stays in INSUFFICIENT_DATA and doesn't perform alarm actions. Valid values range from 1 to 2,880 minutes (2 days). You can change this value while the alarm is still in its warm-up period. Changes have no effect after the warm-up period ends.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-logalarm-warmupconfiguration.html#cfn-cloudwatch-logalarm-warmupconfiguration-warmupperioddurationinminutes
	//
	WarmUpPeriodDurationInMinutes *float64 `field:"optional" json:"warmUpPeriodDurationInMinutes" yaml:"warmUpPeriodDurationInMinutes"`
}

