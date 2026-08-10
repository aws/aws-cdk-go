package awsnetworkflowmonitor


// A remote resource is the other endpoint in a network flow.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   monitorRemoteResourceProperty := &MonitorRemoteResourceProperty{
//   	Identifier: jsii.String("identifier"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkflowmonitor-monitor-monitorremoteresource.html
//
type CfnMonitorPropsMixin_MonitorRemoteResourceProperty struct {
	// The identifier of the remote resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkflowmonitor-monitor-monitorremoteresource.html#cfn-networkflowmonitor-monitor-monitorremoteresource-identifier
	//
	Identifier *string `field:"optional" json:"identifier" yaml:"identifier"`
	// The type of the remote resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkflowmonitor-monitor-monitorremoteresource.html#cfn-networkflowmonitor-monitor-monitorremoteresource-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

