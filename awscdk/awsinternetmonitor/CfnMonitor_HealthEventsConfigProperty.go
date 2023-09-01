package awsinternetmonitor


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   healthEventsConfigProperty := &HealthEventsConfigProperty{
//   	AvailabilityLocalHealthEventsConfig: &LocalHealthEventsConfigProperty{
//   		HealthScoreThreshold: jsii.Number(123),
//   		MinTrafficImpact: jsii.Number(123),
//   		Status: jsii.String("status"),
//   	},
//   	AvailabilityScoreThreshold: jsii.Number(123),
//   	PerformanceLocalHealthEventsConfig: &LocalHealthEventsConfigProperty{
//   		HealthScoreThreshold: jsii.Number(123),
//   		MinTrafficImpact: jsii.Number(123),
//   		Status: jsii.String("status"),
//   	},
//   	PerformanceScoreThreshold: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-internetmonitor-monitor-healtheventsconfig.html
//
type CfnMonitor_HealthEventsConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-internetmonitor-monitor-healtheventsconfig.html#cfn-internetmonitor-monitor-healtheventsconfig-availabilitylocalhealtheventsconfig
	//
	AvailabilityLocalHealthEventsConfig interface{} `field:"optional" json:"availabilityLocalHealthEventsConfig" yaml:"availabilityLocalHealthEventsConfig"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-internetmonitor-monitor-healtheventsconfig.html#cfn-internetmonitor-monitor-healtheventsconfig-availabilityscorethreshold
	//
	AvailabilityScoreThreshold *float64 `field:"optional" json:"availabilityScoreThreshold" yaml:"availabilityScoreThreshold"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-internetmonitor-monitor-healtheventsconfig.html#cfn-internetmonitor-monitor-healtheventsconfig-performancelocalhealtheventsconfig
	//
	PerformanceLocalHealthEventsConfig interface{} `field:"optional" json:"performanceLocalHealthEventsConfig" yaml:"performanceLocalHealthEventsConfig"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-internetmonitor-monitor-healtheventsconfig.html#cfn-internetmonitor-monitor-healtheventsconfig-performancescorethreshold
	//
	PerformanceScoreThreshold *float64 `field:"optional" json:"performanceScoreThreshold" yaml:"performanceScoreThreshold"`
}

