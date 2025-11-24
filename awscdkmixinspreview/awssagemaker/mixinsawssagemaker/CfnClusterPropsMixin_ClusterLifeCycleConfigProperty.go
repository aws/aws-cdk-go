package mixinsawssagemaker


// The lifecycle configuration for a SageMaker HyperPod cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   clusterLifeCycleConfigProperty := &ClusterLifeCycleConfigProperty{
//   	OnCreate: jsii.String("onCreate"),
//   	SourceS3Uri: jsii.String("sourceS3Uri"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterlifecycleconfig.html
//
type CfnClusterPropsMixin_ClusterLifeCycleConfigProperty struct {
	// The file name of the entrypoint script of lifecycle scripts under `SourceS3Uri` .
	//
	// This entrypoint script runs during cluster creation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterlifecycleconfig.html#cfn-sagemaker-cluster-clusterlifecycleconfig-oncreate
	//
	OnCreate *string `field:"optional" json:"onCreate" yaml:"onCreate"`
	// An Amazon S3 bucket path where your lifecycle scripts are stored.
	//
	// > Make sure that the S3 bucket path starts with `s3://sagemaker-` . The [IAM role for SageMaker HyperPod](https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-hyperpod-prerequisites.html#sagemaker-hyperpod-prerequisites-iam-role-for-hyperpod) has the managed [`AmazonSageMakerClusterInstanceRolePolicy`](https://docs.aws.amazon.com/sagemaker/latest/dg/security-iam-awsmanpol-cluster.html) attached, which allows access to S3 buckets with the specific prefix `sagemaker-` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterlifecycleconfig.html#cfn-sagemaker-cluster-clusterlifecycleconfig-sources3uri
	//
	SourceS3Uri *string `field:"optional" json:"sourceS3Uri" yaml:"sourceS3Uri"`
}

