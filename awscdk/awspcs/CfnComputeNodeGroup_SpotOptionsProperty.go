package awspcs


// Additional configuration when you specify `SPOT` as the `purchaseOption` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   spotOptionsProperty := &SpotOptionsProperty{
//   	AllocationStrategy: jsii.String("allocationStrategy"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-computenodegroup-spotoptions.html
//
type CfnComputeNodeGroup_SpotOptionsProperty struct {
	// The Amazon EC2 allocation strategy AWS PCS uses to provision EC2 instances.
	//
	// AWS PCS supports lowest price, capacity optimized, and price capacity optimized. If you don't provide this option, it defaults to price capacity optimized.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-computenodegroup-spotoptions.html#cfn-pcs-computenodegroup-spotoptions-allocationstrategy
	//
	AllocationStrategy *string `field:"optional" json:"allocationStrategy" yaml:"allocationStrategy"`
}

