package awsnetworkflowmonitor

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnMonitorPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnMonitorMixinProps := &CfnMonitorMixinProps{
//   	LocalResources: []interface{}{
//   		&MonitorLocalResourceProperty{
//   			Identifier: jsii.String("identifier"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	MonitorName: jsii.String("monitorName"),
//   	RemoteResources: []interface{}{
//   		&MonitorRemoteResourceProperty{
//   			Identifier: jsii.String("identifier"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	ScopeArn: jsii.String("scopeArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkflowmonitor-monitor.html
//
type CfnMonitorMixinProps struct {
	// The local resources to monitor.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkflowmonitor-monitor.html#cfn-networkflowmonitor-monitor-localresources
	//
	LocalResources interface{} `field:"optional" json:"localResources" yaml:"localResources"`
	// The name of the monitor.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkflowmonitor-monitor.html#cfn-networkflowmonitor-monitor-monitorname
	//
	MonitorName *string `field:"optional" json:"monitorName" yaml:"monitorName"`
	// The remote resources to monitor.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkflowmonitor-monitor.html#cfn-networkflowmonitor-monitor-remoteresources
	//
	RemoteResources interface{} `field:"optional" json:"remoteResources" yaml:"remoteResources"`
	// The Amazon Resource Name (ARN) of the scope for the monitor.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkflowmonitor-monitor.html#cfn-networkflowmonitor-monitor-scopearn
	//
	ScopeArn *string `field:"optional" json:"scopeArn" yaml:"scopeArn"`
	// The tags for the monitor.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkflowmonitor-monitor.html#cfn-networkflowmonitor-monitor-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

