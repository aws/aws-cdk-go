package awsecs


// The minimum and maximum number of vCPUs for instance type selection.
//
// This allows you to specify a range of vCPU counts that meet your workload requirements.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vCpuCountRangeRequestProperty := &VCpuCountRangeRequestProperty{
//   	Min: jsii.Number(123),
//
//   	// the properties below are optional
//   	Max: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-vcpucountrangerequest.html
//
type CfnCapacityProvider_VCpuCountRangeRequestProperty struct {
	// The minimum number of vCPUs.
	//
	// Instance types with fewer vCPUs than this value are excluded from selection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-vcpucountrangerequest.html#cfn-ecs-capacityprovider-vcpucountrangerequest-min
	//
	Min *float64 `field:"required" json:"min" yaml:"min"`
	// The maximum number of vCPUs.
	//
	// Instance types with more vCPUs than this value are excluded from selection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-vcpucountrangerequest.html#cfn-ecs-capacityprovider-vcpucountrangerequest-max
	//
	Max *float64 `field:"optional" json:"max" yaml:"max"`
}

