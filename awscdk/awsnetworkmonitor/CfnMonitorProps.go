package awsnetworkmonitor

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnMonitor`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMonitorProps := &CfnMonitorProps{
//   	MonitorName: jsii.String("monitorName"),
//
//   	// the properties below are optional
//   	AggregationPeriod: jsii.Number(123),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmonitor-monitor.html
//
type CfnMonitorProps struct {
	// The name identifying the monitor.
	//
	// It can contain only letters, underscores (_), or dashes (-), and can be up to 200 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmonitor-monitor.html#cfn-networkmonitor-monitor-monitorname
	//
	MonitorName *string `field:"required" json:"monitorName" yaml:"monitorName"`
	// The time, in seconds, that metrics are aggregated and sent to Amazon CloudWatch.
	//
	// Valid values are either 30 or 60.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmonitor-monitor.html#cfn-networkmonitor-monitor-aggregationperiod
	//
	AggregationPeriod *float64 `field:"optional" json:"aggregationPeriod" yaml:"aggregationPeriod"`
	// The tags for the monitor.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmonitor-monitor.html#cfn-networkmonitor-monitor-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

