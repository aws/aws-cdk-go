package awsec2


// The CPU performance to consider, using an instance family as the baseline reference.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cpuPerformanceFactorRequestProperty := &CpuPerformanceFactorRequestProperty{
//   	References: []interface{}{
//   		&PerformanceFactorReferenceRequestProperty{
//   			InstanceFamily: jsii.String("instanceFamily"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-cpuperformancefactorrequest.html
//
type CfnEC2Fleet_CpuPerformanceFactorRequestProperty struct {
	// Specify an instance family to use as the baseline reference for CPU performance.
	//
	// All instance types that match your specified attributes will be compared against the CPU performance of the referenced instance family, regardless of CPU manufacturer or architecture differences.
	//
	// > Currently, only one instance family can be specified in the list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-cpuperformancefactorrequest.html#cfn-ec2-ec2fleet-cpuperformancefactorrequest-references
	//
	References interface{} `field:"optional" json:"references" yaml:"references"`
}

