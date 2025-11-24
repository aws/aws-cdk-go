package mixinsawssagemaker


// The configuration for the Amazon EKS cluster that is used as the orchestrator for the SageMaker HyperPod cluster.
//
// This includes the Amazon Resource Name (ARN) of the EKS cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   clusterOrchestratorEksConfigProperty := &ClusterOrchestratorEksConfigProperty{
//   	ClusterArn: jsii.String("clusterArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterorchestratoreksconfig.html
//
type CfnClusterPropsMixin_ClusterOrchestratorEksConfigProperty struct {
	// The Amazon Resource Name (ARN) of the SageMaker HyperPod cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterorchestratoreksconfig.html#cfn-sagemaker-cluster-clusterorchestratoreksconfig-clusterarn
	//
	ClusterArn *string `field:"optional" json:"clusterArn" yaml:"clusterArn"`
}

