package previewawsautoscalingmixins


// Specify an instance family to use as the baseline reference for CPU performance.
//
// All instance types that All instance types that match your specified attributes will be compared against the CPU performance of the referenced instance family, regardless of CPU manufacturer or architecture differences.
//
// > Currently only one instance family can be specified in the list.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   performanceFactorReferenceRequestProperty := &PerformanceFactorReferenceRequestProperty{
//   	InstanceFamily: jsii.String("instanceFamily"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-performancefactorreferencerequest.html
//
type CfnAutoScalingGroupPropsMixin_PerformanceFactorReferenceRequestProperty struct {
	// The instance family to use as a baseline reference.
	//
	// > Make sure that you specify the correct value for the instance family. The instance family is everything before the period (.) in the instance type name. For example, in the instance `c6i.large` , the instance family is `c6i` , not `c6` . For more information, see [Amazon EC2 instance type naming conventions](https://docs.aws.amazon.com/ec2/latest/instancetypes/instance-type-names.html) in *Amazon EC2 Instance Types* .
	//
	// The following instance types are *not supported* for performance protection.
	//
	// - `c1`
	// - `g3| g3s`
	// - `hpc7g`
	// - `m1| m2`
	// - `mac1 | mac2 | mac2-m1ultra | mac2-m2 | mac2-m2pro`
	// - `p3dn | p4d | p5`
	// - `t1`
	// - `u-12tb1 | u-18tb1 | u-24tb1 | u-3tb1 | u-6tb1 | u-9tb1 | u7i-12tb | u7in-16tb | u7in-24tb | u7in-32tb`
	//
	// If you performance protection by specifying a supported instance family, the returned instance types will exclude the preceding unsupported instance families.
	//
	// If you specify an unsupported instance family as a value for baseline performance, the API returns an empty response.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-performancefactorreferencerequest.html#cfn-autoscaling-autoscalinggroup-performancefactorreferencerequest-instancefamily
	//
	InstanceFamily *string `field:"optional" json:"instanceFamily" yaml:"instanceFamily"`
}

