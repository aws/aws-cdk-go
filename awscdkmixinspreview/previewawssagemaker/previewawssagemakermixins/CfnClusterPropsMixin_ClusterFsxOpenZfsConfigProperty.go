package previewawssagemakermixins


// Configuration for mounting an Amazon FSx OpenZFS file system to the instances in the SageMaker HyperPod cluster instance group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   clusterFsxOpenZfsConfigProperty := &ClusterFsxOpenZfsConfigProperty{
//   	DnsName: jsii.String("dnsName"),
//   	MountPath: jsii.String("mountPath"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterfsxopenzfsconfig.html
//
type CfnClusterPropsMixin_ClusterFsxOpenZfsConfigProperty struct {
	// The DNS name of the FSx for OpenZFS file system.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterfsxopenzfsconfig.html#cfn-sagemaker-cluster-clusterfsxopenzfsconfig-dnsname
	//
	DnsName *string `field:"optional" json:"dnsName" yaml:"dnsName"`
	// The mount path for the FSx for OpenZFS file system.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterfsxopenzfsconfig.html#cfn-sagemaker-cluster-clusterfsxopenzfsconfig-mountpath
	//
	MountPath *string `field:"optional" json:"mountPath" yaml:"mountPath"`
}

