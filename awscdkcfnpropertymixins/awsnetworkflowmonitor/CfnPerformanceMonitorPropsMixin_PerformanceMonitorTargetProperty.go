package awsnetworkflowmonitor


// A target that is monitored by a performance monitor.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   performanceMonitorTargetProperty := &PerformanceMonitorTargetProperty{
//   	Identifier: jsii.String("identifier"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkflowmonitor-performancemonitor-performancemonitortarget.html
//
type CfnPerformanceMonitorPropsMixin_PerformanceMonitorTargetProperty struct {
	// The identifier of the target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkflowmonitor-performancemonitor-performancemonitortarget.html#cfn-networkflowmonitor-performancemonitor-performancemonitortarget-identifier
	//
	Identifier *string `field:"optional" json:"identifier" yaml:"identifier"`
	// The type of the target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkflowmonitor-performancemonitor-performancemonitortarget.html#cfn-networkflowmonitor-performancemonitor-performancemonitortarget-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

