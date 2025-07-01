package awssagemaker


// Defines the configuration for attaching an additional Amazon Elastic Block Store (EBS) volume to each instance of the SageMaker HyperPod cluster instance group.
//
// To learn more, see [SageMaker HyperPod release notes: June 20, 2024](https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-hyperpod-release-notes.html#sagemaker-hyperpod-release-notes-20240620) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clusterEbsVolumeConfigProperty := &ClusterEbsVolumeConfigProperty{
//   	VolumeSizeInGb: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterebsvolumeconfig.html
//
type CfnCluster_ClusterEbsVolumeConfigProperty struct {
	// The size in gigabytes (GB) of the additional EBS volume to be attached to the instances in the SageMaker HyperPod cluster instance group.
	//
	// The additional EBS volume is attached to each instance within the SageMaker HyperPod cluster instance group and mounted to `/opt/sagemaker` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterebsvolumeconfig.html#cfn-sagemaker-cluster-clusterebsvolumeconfig-volumesizeingb
	//
	VolumeSizeInGb *float64 `field:"optional" json:"volumeSizeInGb" yaml:"volumeSizeInGb"`
}

