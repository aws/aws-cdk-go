package mixinsawsemr


// The resize specification for Spot Instances in the instance fleet, which contains the resize timeout period.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   spotResizingSpecificationProperty := &SpotResizingSpecificationProperty{
//   	AllocationStrategy: jsii.String("allocationStrategy"),
//   	TimeoutDurationMinutes: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-instancefleetconfig-spotresizingspecification.html
//
type CfnInstanceFleetConfigPropsMixin_SpotResizingSpecificationProperty struct {
	// Specifies the allocation strategy to use to launch Spot instances during a resize.
	//
	// If you run Amazon EMR releases 6.9.0 or higher, the default is `price-capacity-optimized` . If you run Amazon EMR releases 6.8.0 or lower, the default is `capacity-optimized` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-instancefleetconfig-spotresizingspecification.html#cfn-emr-instancefleetconfig-spotresizingspecification-allocationstrategy
	//
	AllocationStrategy *string `field:"optional" json:"allocationStrategy" yaml:"allocationStrategy"`
	// Spot resize timeout in minutes.
	//
	// If Spot Instances are not provisioned within this time, the resize workflow will stop provisioning of Spot instances. Minimum value is 5 minutes and maximum value is 10,080 minutes (7 days). The timeout applies to all resize workflows on the Instance Fleet. The resize could be triggered by Amazon EMR Managed Scaling or by the customer (via Amazon EMR Console, Amazon EMR CLI modify-instance-fleet or Amazon EMR SDK ModifyInstanceFleet API) or by Amazon EMR due to Amazon EC2 Spot Reclamation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-instancefleetconfig-spotresizingspecification.html#cfn-emr-instancefleetconfig-spotresizingspecification-timeoutdurationminutes
	//
	TimeoutDurationMinutes *float64 `field:"optional" json:"timeoutDurationMinutes" yaml:"timeoutDurationMinutes"`
}

