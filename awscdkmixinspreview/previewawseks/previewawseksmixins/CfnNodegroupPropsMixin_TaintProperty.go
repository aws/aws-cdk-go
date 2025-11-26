package previewawseksmixins


// A property that allows a node to repel a `Pod` .
//
// For more information, see [Node taints on managed node groups](https://docs.aws.amazon.com/eks/latest/userguide/node-taints-managed-node-groups.html) in the *Amazon EKS User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   taintProperty := &TaintProperty{
//   	Effect: jsii.String("effect"),
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-nodegroup-taint.html
//
type CfnNodegroupPropsMixin_TaintProperty struct {
	// The effect of the taint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-nodegroup-taint.html#cfn-eks-nodegroup-taint-effect
	//
	Effect *string `field:"optional" json:"effect" yaml:"effect"`
	// The key of the taint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-nodegroup-taint.html#cfn-eks-nodegroup-taint-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value of the taint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-nodegroup-taint.html#cfn-eks-nodegroup-taint-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

