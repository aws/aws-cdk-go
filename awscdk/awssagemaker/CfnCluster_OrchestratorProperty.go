package awssagemaker


// Specifies parameter(s) specific to the orchestrator, e.g. specify the EKS cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   orchestratorProperty := &OrchestratorProperty{
//   	Eks: &ClusterOrchestratorEksConfigProperty{
//   		ClusterArn: jsii.String("clusterArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-orchestrator.html
//
type CfnCluster_OrchestratorProperty struct {
	// Specifies parameter(s) related to EKS as orchestrator, e.g. the EKS cluster nodes will attach to,.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-orchestrator.html#cfn-sagemaker-cluster-orchestrator-eks
	//
	Eks interface{} `field:"required" json:"eks" yaml:"eks"`
}

