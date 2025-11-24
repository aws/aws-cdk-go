package mixinsawsinternetmonitor


// Configuration information that determines the threshold and other conditions for when Internet Monitor creates a health event for a local performance or availability issue, when scores cross a threshold for one or more city-networks.
//
// Defines the percentages, for performance scores or availability scores, that are the local thresholds for when Amazon CloudWatch Internet Monitor creates a health event. Also defines whether a local threshold is enabled or disabled, and the minimum percentage of overall traffic that must be impacted by an issue before Internet Monitor creates an event when a	threshold is crossed for a local health score.
//
// If you don't set a local health event threshold, the default value is 60%.
//
// For more information, see [Change health event thresholds](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-IM-overview.html#IMUpdateThresholdFromOverview) in the Internet Monitor section of the *Amazon CloudWatch User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   localHealthEventsConfigProperty := &LocalHealthEventsConfigProperty{
//   	HealthScoreThreshold: jsii.Number(123),
//   	MinTrafficImpact: jsii.Number(123),
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-internetmonitor-monitor-localhealtheventsconfig.html
//
type CfnMonitorPropsMixin_LocalHealthEventsConfigProperty struct {
	// The health event threshold percentage set for a local health score.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-internetmonitor-monitor-localhealtheventsconfig.html#cfn-internetmonitor-monitor-localhealtheventsconfig-healthscorethreshold
	//
	HealthScoreThreshold *float64 `field:"optional" json:"healthScoreThreshold" yaml:"healthScoreThreshold"`
	// The minimum percentage of overall traffic for an application that must be impacted by an issue before Internet Monitor creates an event when a threshold is crossed for a local health score.
	//
	// If you don't set a minimum traffic impact threshold, the default value is 0.01%.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-internetmonitor-monitor-localhealtheventsconfig.html#cfn-internetmonitor-monitor-localhealtheventsconfig-mintrafficimpact
	//
	MinTrafficImpact *float64 `field:"optional" json:"minTrafficImpact" yaml:"minTrafficImpact"`
	// The status of whether Internet Monitor creates a health event based on a threshold percentage set for a local health score.
	//
	// The status can be `ENABLED` or `DISABLED` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-internetmonitor-monitor-localhealtheventsconfig.html#cfn-internetmonitor-monitor-localhealtheventsconfig-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

