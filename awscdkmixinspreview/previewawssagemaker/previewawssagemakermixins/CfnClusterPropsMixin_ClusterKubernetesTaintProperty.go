package previewawssagemakermixins


// A Kubernetes taint to apply to cluster nodes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   clusterKubernetesTaintProperty := &ClusterKubernetesTaintProperty{
//   	Effect: jsii.String("effect"),
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterkubernetestaint.html
//
type CfnClusterPropsMixin_ClusterKubernetesTaintProperty struct {
	// The effect of the taint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterkubernetestaint.html#cfn-sagemaker-cluster-clusterkubernetestaint-effect
	//
	Effect *string `field:"optional" json:"effect" yaml:"effect"`
	// The key of the taint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterkubernetestaint.html#cfn-sagemaker-cluster-clusterkubernetestaint-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value of the taint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterkubernetestaint.html#cfn-sagemaker-cluster-clusterkubernetestaint-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

