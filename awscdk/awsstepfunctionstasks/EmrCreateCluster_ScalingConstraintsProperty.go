package awsstepfunctionstasks


// The upper and lower EC2 instance limits for an automatic scaling policy.
//
// Automatic scaling activities triggered by automatic scaling
// rules will not cause an instance group to grow above or below these limits.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scalingConstraintsProperty := &ScalingConstraintsProperty{
//   	MaxCapacity: jsii.Number(123),
//   	MinCapacity: jsii.Number(123),
//   }
//
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_ScalingConstraints.html
//
type EmrCreateCluster_ScalingConstraintsProperty struct {
	// The upper boundary of EC2 instances in an instance group beyond which scaling activities are not allowed to grow.
	//
	// Scale-out
	// activities will not add instances beyond this boundary.
	MaxCapacity *float64 `field:"required" json:"maxCapacity" yaml:"maxCapacity"`
	// The lower boundary of EC2 instances in an instance group below which scaling activities are not allowed to shrink.
	//
	// Scale-in
	// activities will not terminate instances below this boundary.
	MinCapacity *float64 `field:"required" json:"minCapacity" yaml:"minCapacity"`
}

