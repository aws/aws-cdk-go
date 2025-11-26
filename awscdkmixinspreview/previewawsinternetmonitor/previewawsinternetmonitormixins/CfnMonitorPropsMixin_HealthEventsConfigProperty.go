package previewawsinternetmonitormixins


// Define the health event threshold percentages for the performance score and availability score for your application's monitor.
//
// Amazon CloudWatch Internet Monitor creates a health event when there's an internet issue that affects your application end users where a health score percentage is at or below a set threshold.
//
// If you don't set a health event threshold, the default value is 95%.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
type CfnMonitorPropsMixin_HealthEventsConfigProperty struct {
	// The configuration that determines the threshold and other conditions for when Internet Monitor creates a health event for a local availability issue.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-internetmonitor-monitor-healtheventsconfig.html#cfn-internetmonitor-monitor-healtheventsconfig-availabilitylocalhealtheventsconfig
	//
	AvailabilityLocalHealthEventsConfig interface{} `field:"optional" json:"availabilityLocalHealthEventsConfig" yaml:"availabilityLocalHealthEventsConfig"`
	// The health event threshold percentage set for availability scores.
	//
	// When the overall availability score is at or below this percentage, Internet Monitor creates a health event.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-internetmonitor-monitor-healtheventsconfig.html#cfn-internetmonitor-monitor-healtheventsconfig-availabilityscorethreshold
	//
	AvailabilityScoreThreshold *float64 `field:"optional" json:"availabilityScoreThreshold" yaml:"availabilityScoreThreshold"`
	// The configuration that determines the threshold and other conditions for when Internet Monitor creates a health event for a local performance issue.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-internetmonitor-monitor-healtheventsconfig.html#cfn-internetmonitor-monitor-healtheventsconfig-performancelocalhealtheventsconfig
	//
	PerformanceLocalHealthEventsConfig interface{} `field:"optional" json:"performanceLocalHealthEventsConfig" yaml:"performanceLocalHealthEventsConfig"`
	// The health event threshold percentage set for performance scores.
	//
	// When the overall performance score is at or below this percentage, Internet Monitor creates a health event.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-internetmonitor-monitor-healtheventsconfig.html#cfn-internetmonitor-monitor-healtheventsconfig-performancescorethreshold
	//
	PerformanceScoreThreshold *float64 `field:"optional" json:"performanceScoreThreshold" yaml:"performanceScoreThreshold"`
}

