package awseks


// A property that allows a node to repel a set of pods.
//
// For more information, see [Node taints on managed node groups](https://docs.aws.amazon.com/eks/latest/userguide/node-taints-managed-node-groups.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   taintProperty := &taintProperty{
//   	effect: jsii.String("effect"),
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//   }
//
type CfnNodegroup_TaintProperty struct {
	// The effect of the taint.
	Effect *string `field:"optional" json:"effect" yaml:"effect"`
	// The key of the taint.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value of the taint.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

