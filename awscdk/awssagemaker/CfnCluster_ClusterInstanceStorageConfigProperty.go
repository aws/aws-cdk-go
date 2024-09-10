package awssagemaker


// Defines the configuration for attaching additional storage to the instances in the SageMaker HyperPod cluster instance group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clusterInstanceStorageConfigProperty := &ClusterInstanceStorageConfigProperty{
//   	EbsVolumeConfig: &ClusterEbsVolumeConfigProperty{
//   		VolumeSizeInGb: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterinstancestorageconfig.html
//
type CfnCluster_ClusterInstanceStorageConfigProperty struct {
	// Defines the configuration for attaching additional Amazon Elastic Block Store (EBS) volumes to the instances in the SageMaker HyperPod cluster instance group.
	//
	// The additional EBS volume is attached to each instance within the SageMaker HyperPod cluster instance group and mounted to /opt/sagemaker.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterinstancestorageconfig.html#cfn-sagemaker-cluster-clusterinstancestorageconfig-ebsvolumeconfig
	//
	EbsVolumeConfig interface{} `field:"optional" json:"ebsVolumeConfig" yaml:"ebsVolumeConfig"`
}

