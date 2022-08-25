package awselasticloadbalancingv2


// A Target Group and weight combination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var networkTargetGroup networkTargetGroup
//
//   networkWeightedTargetGroup := &networkWeightedTargetGroup{
//   	targetGroup: networkTargetGroup,
//
//   	// the properties below are optional
//   	weight: jsii.Number(123),
//   }
//
type NetworkWeightedTargetGroup struct {
	// The target group.
	TargetGroup INetworkTargetGroup `field:"required" json:"targetGroup" yaml:"targetGroup"`
	// The target group's weight.
	//
	// Range is [0..1000).
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

