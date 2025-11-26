package awssagemaker


// Kubernetes configuration for cluster nodes including labels and taints.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clusterKubernetesConfigProperty := &ClusterKubernetesConfigProperty{
//   	Labels: map[string]*string{
//   		"labelsKey": jsii.String("labels"),
//   	},
//   	Taints: []interface{}{
//   		&ClusterKubernetesTaintProperty{
//   			Effect: jsii.String("effect"),
//   			Key: jsii.String("key"),
//
//   			// the properties below are optional
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterkubernetesconfig.html
//
type CfnCluster_ClusterKubernetesConfigProperty struct {
	// A map of Kubernetes labels to apply to cluster nodes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterkubernetesconfig.html#cfn-sagemaker-cluster-clusterkubernetesconfig-labels
	//
	Labels interface{} `field:"optional" json:"labels" yaml:"labels"`
	// A list of Kubernetes taints to apply to cluster nodes.
	//
	// Maximum of 50 taints.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterkubernetesconfig.html#cfn-sagemaker-cluster-clusterkubernetesconfig-taints
	//
	Taints interface{} `field:"optional" json:"taints" yaml:"taints"`
}

