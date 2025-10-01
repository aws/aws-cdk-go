package awscodeguruprofiler


// A reference to a ProfilingGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   profilingGroupReference := &ProfilingGroupReference{
//   	ProfilingGroupArn: jsii.String("profilingGroupArn"),
//   	ProfilingGroupName: jsii.String("profilingGroupName"),
//   }
//
type ProfilingGroupReference struct {
	// The ARN of the ProfilingGroup resource.
	ProfilingGroupArn *string `field:"required" json:"profilingGroupArn" yaml:"profilingGroupArn"`
	// The ProfilingGroupName of the ProfilingGroup resource.
	ProfilingGroupName *string `field:"required" json:"profilingGroupName" yaml:"profilingGroupName"`
}

