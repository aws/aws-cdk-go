package awsnetworkflowmonitor

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnPerformanceMonitor`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPerformanceMonitorProps := &CfnPerformanceMonitorProps{
//   	MonitorName: jsii.String("monitorName"),
//   	Targets: []interface{}{
//   		&PerformanceMonitorTargetProperty{
//   			Identifier: jsii.String("identifier"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	ScopeArn: jsii.String("scopeArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkflowmonitor-performancemonitor.html
//
type CfnPerformanceMonitorProps struct {
	// The name of the performance monitor.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkflowmonitor-performancemonitor.html#cfn-networkflowmonitor-performancemonitor-monitorname
	//
	MonitorName *string `field:"required" json:"monitorName" yaml:"monitorName"`
	// The targets that are monitored by the performance monitor.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkflowmonitor-performancemonitor.html#cfn-networkflowmonitor-performancemonitor-targets
	//
	Targets interface{} `field:"required" json:"targets" yaml:"targets"`
	// The Amazon Resource Name (ARN) of the scope for the performance monitor.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkflowmonitor-performancemonitor.html#cfn-networkflowmonitor-performancemonitor-scopearn
	//
	ScopeArn *string `field:"optional" json:"scopeArn" yaml:"scopeArn"`
	// The tags for the performance monitor.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkflowmonitor-performancemonitor.html#cfn-networkflowmonitor-performancemonitor-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

