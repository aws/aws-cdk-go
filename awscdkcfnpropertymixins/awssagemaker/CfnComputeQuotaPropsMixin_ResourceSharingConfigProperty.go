package awssagemaker


// Idle-resource-sharing strategy for this team's quota.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   resourceSharingConfigProperty := &ResourceSharingConfigProperty{
//   	AbsoluteBorrowLimits: []interface{}{
//   		&ComputeQuotaResourceConfigProperty{
//   			AcceleratorPartition: &AcceleratorPartitionConfigProperty{
//   				Count: jsii.Number(123),
//   				Type: jsii.String("type"),
//   			},
//   			Accelerators: jsii.Number(123),
//   			Count: jsii.Number(123),
//   			InstanceType: jsii.String("instanceType"),
//   			MemoryInGiB: jsii.Number(123),
//   			VCpu: jsii.Number(123),
//   		},
//   	},
//   	BorrowLimit: jsii.Number(123),
//   	Strategy: jsii.String("strategy"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-computequota-resourcesharingconfig.html
//
type CfnComputeQuotaPropsMixin_ResourceSharingConfigProperty struct {
	// Per-instance-type absolute caps on resources this team can borrow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-computequota-resourcesharingconfig.html#cfn-sagemaker-computequota-resourcesharingconfig-absoluteborrowlimits
	//
	AbsoluteBorrowLimits interface{} `field:"optional" json:"absoluteBorrowLimits" yaml:"absoluteBorrowLimits"`
	// Maximum percentage of the team's idle quota that can be borrowed by other teams.
	//
	// Smithy admits 0..10000; in practice the service-documented operational range is 1..500. Defaults service-side to 50 when omitted on Create.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-computequota-resourcesharingconfig.html#cfn-sagemaker-computequota-resourcesharingconfig-borrowlimit
	//
	BorrowLimit *float64 `field:"optional" json:"borrowLimit" yaml:"borrowLimit"`
	// How this team participates in idle-resource sharing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-computequota-resourcesharingconfig.html#cfn-sagemaker-computequota-resourcesharingconfig-strategy
	//
	Strategy *string `field:"optional" json:"strategy" yaml:"strategy"`
}

