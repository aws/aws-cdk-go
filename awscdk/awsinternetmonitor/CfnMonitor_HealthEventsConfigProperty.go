package awsinternetmonitor


// Define the health event threshold percentages for the performance score and availability score for your application's monitor.
//
// Amazon CloudWatch Internet Monitor creates a health event when there's an internet issue that affects your application end users where a health score percentage is at or below a set threshold.
//
// If you don't set a health event threshold, the default value is 95%.
//
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
	// The health event threshold percentage set for availability scores.
	//
	// When the global availability score is at or below this percentage, Internet Monitor creates a health event.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-internetmonitor-monitor-healtheventsconfig.html#cfn-internetmonitor-monitor-healtheventsconfig-availabilityscorethreshold
	//
	AvailabilityScoreThreshold *float64 `field:"optional" json:"availabilityScoreThreshold" yaml:"availabilityScoreThreshold"`
	// The health event threshold percentage set for performance scores.
	//
	// When the global performance score is at or below this percentage, Internet Monitor creates a health event.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-internetmonitor-monitor-healtheventsconfig.html#cfn-internetmonitor-monitor-healtheventsconfig-performancescorethreshold
	//
	PerformanceScoreThreshold *float64 `field:"optional" json:"performanceScoreThreshold" yaml:"performanceScoreThreshold"`
}

