package awsemr


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   spotResizingSpecificationProperty := &SpotResizingSpecificationProperty{
//   	AllocationStrategy: jsii.String("allocationStrategy"),
//   	TimeoutDurationMinutes: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-instancefleetconfig-spotresizingspecification.html
//
type CfnInstanceFleetConfig_SpotResizingSpecificationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-instancefleetconfig-spotresizingspecification.html#cfn-emr-instancefleetconfig-spotresizingspecification-allocationstrategy
	//
	AllocationStrategy *string `field:"optional" json:"allocationStrategy" yaml:"allocationStrategy"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-instancefleetconfig-spotresizingspecification.html#cfn-emr-instancefleetconfig-spotresizingspecification-timeoutdurationminutes
	//
	TimeoutDurationMinutes *float64 `field:"optional" json:"timeoutDurationMinutes" yaml:"timeoutDurationMinutes"`
}

