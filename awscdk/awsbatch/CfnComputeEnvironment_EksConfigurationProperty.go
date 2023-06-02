package awsbatch


// Configuration for the Amazon EKS cluster that supports the AWS Batch compute environment.
//
// The cluster must exist before the compute environment can be created.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eksConfigurationProperty := &EksConfigurationProperty{
//   	EksClusterArn: jsii.String("eksClusterArn"),
//   	KubernetesNamespace: jsii.String("kubernetesNamespace"),
//   }
//
type CfnComputeEnvironment_EksConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of the Amazon EKS cluster.
	//
	// An example is `arn: *aws* :eks: *us-east-1* : *123456789012* :cluster/ *ClusterForBatch*` .
	EksClusterArn *string `field:"required" json:"eksClusterArn" yaml:"eksClusterArn"`
	// The namespace of the Amazon EKS cluster.
	//
	// AWS Batch manages pods in this namespace. The value can't left empty or null. It must be fewer than 64 characters long, can't be set to `default` , can't start with " `kube-` ," and must match this regular expression: `^[a-z0-9]([-a-z0-9]*[a-z0-9])?$` . For more information, see [Namespaces](https://docs.aws.amazon.com/https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/) in the Kubernetes documentation.
	KubernetesNamespace *string `field:"required" json:"kubernetesNamespace" yaml:"kubernetesNamespace"`
}

