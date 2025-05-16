package awsec2


// The baseline performance to consider, using an instance family as a baseline reference.
//
// The instance family establishes the lowest acceptable level of performance. Amazon EC2 uses this baseline to guide instance type selection, but there is no guarantee that the selected instance types will always exceed the baseline for every application.
//
// Currently, this parameter only supports CPU performance as a baseline performance factor. For example, specifying `c6i` would use the CPU performance of the `c6i` family as the baseline reference.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   baselinePerformanceFactorsRequestProperty := &BaselinePerformanceFactorsRequestProperty{
//   	Cpu: &CpuPerformanceFactorRequestProperty{
//   		References: []interface{}{
//   			&PerformanceFactorReferenceRequestProperty{
//   				InstanceFamily: jsii.String("instanceFamily"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-baselineperformancefactorsrequest.html
//
type CfnEC2Fleet_BaselinePerformanceFactorsRequestProperty struct {
	// The CPU performance to consider, using an instance family as the baseline reference.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-baselineperformancefactorsrequest.html#cfn-ec2-ec2fleet-baselineperformancefactorsrequest-cpu
	//
	Cpu interface{} `field:"optional" json:"cpu" yaml:"cpu"`
}

