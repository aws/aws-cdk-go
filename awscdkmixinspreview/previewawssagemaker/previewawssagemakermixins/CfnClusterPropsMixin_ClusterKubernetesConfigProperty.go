package previewawssagemakermixins


// Kubernetes configuration that specifies labels and taints to be applied to cluster nodes in an instance group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   clusterKubernetesConfigProperty := &ClusterKubernetesConfigProperty{
//   	Labels: map[string]*string{
//   		"labelsKey": jsii.String("labels"),
//   	},
//   	Taints: []interface{}{
//   		&ClusterKubernetesTaintProperty{
//   			Effect: jsii.String("effect"),
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterkubernetesconfig.html
//
type CfnClusterPropsMixin_ClusterKubernetesConfigProperty struct {
	// Key-value pairs of labels to be applied to cluster nodes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterkubernetesconfig.html#cfn-sagemaker-cluster-clusterkubernetesconfig-labels
	//
	Labels interface{} `field:"optional" json:"labels" yaml:"labels"`
	// List of taints to be applied to cluster nodes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterkubernetesconfig.html#cfn-sagemaker-cluster-clusterkubernetesconfig-taints
	//
	Taints interface{} `field:"optional" json:"taints" yaml:"taints"`
}

