package awsinternetmonitor


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   healthEventsConfigProperty := &HealthEventsConfigProperty{
//   	AvailabilityScoreThreshold: jsii.Number(123),
//   	PerformanceScoreThreshold: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-internetmonitor-monitor-healtheventsconfig.html
//
type CfnMonitor_HealthEventsConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-internetmonitor-monitor-healtheventsconfig.html#cfn-internetmonitor-monitor-healtheventsconfig-availabilityscorethreshold
	//
	AvailabilityScoreThreshold *float64 `field:"optional" json:"availabilityScoreThreshold" yaml:"availabilityScoreThreshold"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-internetmonitor-monitor-healtheventsconfig.html#cfn-internetmonitor-monitor-healtheventsconfig-performancescorethreshold
	//
	PerformanceScoreThreshold *float64 `field:"optional" json:"performanceScoreThreshold" yaml:"performanceScoreThreshold"`
}

