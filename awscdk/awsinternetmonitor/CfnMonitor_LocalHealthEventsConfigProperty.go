package awsinternetmonitor


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   localHealthEventsConfigProperty := &LocalHealthEventsConfigProperty{
//   	HealthScoreThreshold: jsii.Number(123),
//   	MinTrafficImpact: jsii.Number(123),
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-internetmonitor-monitor-localhealtheventsconfig.html
//
type CfnMonitor_LocalHealthEventsConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-internetmonitor-monitor-localhealtheventsconfig.html#cfn-internetmonitor-monitor-localhealtheventsconfig-healthscorethreshold
	//
	HealthScoreThreshold *float64 `field:"optional" json:"healthScoreThreshold" yaml:"healthScoreThreshold"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-internetmonitor-monitor-localhealtheventsconfig.html#cfn-internetmonitor-monitor-localhealtheventsconfig-mintrafficimpact
	//
	MinTrafficImpact *float64 `field:"optional" json:"minTrafficImpact" yaml:"minTrafficImpact"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-internetmonitor-monitor-localhealtheventsconfig.html#cfn-internetmonitor-monitor-localhealtheventsconfig-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

