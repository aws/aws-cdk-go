package awssagemaker


// The lifecycle configuration for a SageMaker HyperPod cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clusterLifeCycleConfigProperty := &ClusterLifeCycleConfigProperty{
//   	OnCreate: jsii.String("onCreate"),
//   	SourceS3Uri: jsii.String("sourceS3Uri"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterlifecycleconfig.html
//
type CfnCluster_ClusterLifeCycleConfigProperty struct {
	// The file name of the entrypoint script of lifecycle scripts under SourceS3Uri.
	//
	// This entrypoint script runs during cluster creation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterlifecycleconfig.html#cfn-sagemaker-cluster-clusterlifecycleconfig-oncreate
	//
	OnCreate *string `field:"required" json:"onCreate" yaml:"onCreate"`
	// An Amazon S3 bucket path where your lifecycle scripts are stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterlifecycleconfig.html#cfn-sagemaker-cluster-clusterlifecycleconfig-sources3uri
	//
	SourceS3Uri *string `field:"required" json:"sourceS3Uri" yaml:"sourceS3Uri"`
}

