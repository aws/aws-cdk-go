package awsnetworkflowmonitor


// A local resource is the host where the agent is installed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   monitorLocalResourceProperty := &MonitorLocalResourceProperty{
//   	Identifier: jsii.String("identifier"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkflowmonitor-monitor-monitorlocalresource.html
//
type CfnMonitorPropsMixin_MonitorLocalResourceProperty struct {
	// The identifier of the local resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkflowmonitor-monitor-monitorlocalresource.html#cfn-networkflowmonitor-monitor-monitorlocalresource-identifier
	//
	Identifier *string `field:"optional" json:"identifier" yaml:"identifier"`
	// The type of the local resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkflowmonitor-monitor-monitorlocalresource.html#cfn-networkflowmonitor-monitor-monitorlocalresource-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

